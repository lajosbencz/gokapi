package server

import (
	"context"
	"fmt"
	wampClient "github.com/gammazero/nexus/v3/client"
	wampRouter "github.com/gammazero/nexus/v3/router"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/markbates/pkger"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	WampRealm = "gokapi"
)

type Server struct {
	host            string
	port            int
	webuiPath       string
	HTTPServer      *http.Server
	HTTPRouter      *http.ServeMux
	WebSocketRouter wampRouter.Router
	WebSocketServer *wampRouter.WebsocketServer
	WebSocketClient *wampClient.Client
}

func NewServer(host string, port int) (*Server, error) {
	s := &Server{
		host: host,
		port: port,
	}
	// WebSocketRouter
	routerConfig := &wampRouter.Config{
		RealmConfigs: []*wampRouter.RealmConfig{
			{
				URI:           wamp.URI(WampRealm),
				AnonymousAuth: true,
				AllowDisclose: true,
			},
		},
	}
	var err error
	s.WebSocketRouter, err = wampRouter.NewRouter(routerConfig, nil)
	if err != nil {
		return nil, err
	}

	// WebSocketServer
	s.WebSocketServer = wampRouter.NewWebsocketServer(s.WebSocketRouter)
	s.WebSocketServer.Upgrader.EnableCompression = true
	s.WebSocketServer.EnableTrackingCookie = true
	s.WebSocketServer.KeepAlive = 30 * time.Second
	s.WebSocketServer.AllowOrigins([]string{"*"})

	// HTTPRouter
	s.HTTPRouter = http.NewServeMux()
	s.HTTPRouter.Handle("/ws/", s.WebSocketServer)
	s.HTTPRouter.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(pkger.Dir("/webui/dist"))))
	s.HTTPRouter.Handle("/", http.FileServer(pkger.Dir("/web")))

	s.HTTPServer = &http.Server{
		Addr:    s.GetListenAddress(),
		Handler: s.HTTPRouter,
	}

	return s, nil
}

func (s *Server) GetListenAddress() string {
	return s.host + ":" + strconv.Itoa(s.port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.WebSocketRouter.Close()
	return s.HTTPServer.Shutdown(ctx)
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.GetListenAddress(), s.HTTPRouter)
}

func (s *Server) GetWebSocketClient() (*wampClient.Client, error) {
	if s.WebSocketClient == nil {
		logger := log.New(os.Stdout, "CALLEE> ", log.LstdFlags)
		cfg := wampClient.Config{
			Realm:  WampRealm,
			Logger: logger,
		}
		var err error
		s.WebSocketClient, err = wampClient.ConnectLocal(s.WebSocketRouter, cfg)
		if err != nil {
			return nil, err
		}
	}
	return s.WebSocketClient, nil
}

func (s *Server) Register(topic string, callee wampClient.InvocationHandler) (*wampClient.Client, error) {
	client, err := s.GetWebSocketClient()
	if err != nil {
		return nil, err
	}
	if err = client.Register(topic, callee, nil); err != nil {
		return nil, fmt.Errorf("failed to register %q: %s", topic, err)
	}
	fmt.Println("Topic registered:", topic)
	return client, nil
}

func (s *Server) Publish(topic string, args wamp.List, kwargs wamp.Dict, options wamp.Dict) error {
	client, err := s.GetWebSocketClient()
	if err != nil {
		return err
	}
	return client.Publish(topic, options, args, kwargs)
}
