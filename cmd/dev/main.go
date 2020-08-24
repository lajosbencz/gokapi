package main

import (
	"context"
	"fmt"
	wampClient "github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/lajosbencz/gokapi/internal/server"
	"time"
)

func main() {
	srv, err := server.NewServer("localhost", 4000)
	if err != nil {
		panic(err)
	}
	if _, err := srv.Register("worldtime", worldTime); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Listening on http://", srv.GetListenAddress())
	if err = srv.Listen(); err != nil {
		fmt.Println(err)
	}
}

func worldTime(ctx context.Context, inv *wamp.Invocation) wampClient.InvokeResult {
	time.Sleep(2 * time.Second)
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
