package main

import (
	"context"
	"flag"
	"fmt"
	wampClient "github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/lajosbencz/gokapi/internal/server"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	var (
		addr  string
		port  int
		open  bool
		debug bool
	)
	flag.StringVar(&addr, "addr", "localhost", "Address to listen on")
	flag.IntVar(&port, "port", 4000, "Port to listen on")
	flag.BoolVar(&open, "open", false, "Should browser be opened")
	flag.BoolVar(&debug, "debug", false, "Verbose output")
	flag.Parse()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-shutdown
		log.Printf("system call:%+v", osCall)
		cancel()
	}()

	srv, err := server.NewServer(addr, port)
	if err != nil {
		panic(err)
	}
	if _, err := srv.Register("worldtime", worldTime); err != nil {
		log.Println(err)
	}
	if _, err := srv.Register("slowtime", worldTimeSlow); err != nil {
		log.Println(err)
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	url := "http://" + srv.GetListenAddress()
	log.Println("Listening on " + url)
	if open {
		openBrowser(url)
	}

	<-ctx.Done()
	log.Println("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func worldTime(ctx context.Context, inv *wamp.Invocation) wampClient.InvokeResult {
	now := time.Now()
	results := wamp.List{fmt.Sprintf("UTC: %s", now.UTC())}
	for _, arg := range inv.Arguments {
		locName, ok := wamp.AsString(arg)
		if !ok {
			continue
		}
		loc, err := time.LoadLocation(locName)
		if err != nil {
			results = append(results, fmt.Sprintf("%s: %s", locName, err))
			continue
		}
		results = append(results, fmt.Sprintf("%s: %s", locName, now.In(loc)))
	}

	return wampClient.InvokeResult{Args: results}
}

func worldTimeSlow(ctx context.Context, inv *wamp.Invocation) wampClient.InvokeResult {
	time.Sleep(time.Second * 2)
	return worldTime(ctx, inv)
}
