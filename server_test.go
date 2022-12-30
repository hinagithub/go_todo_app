package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {

	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	eg.Go(func() error {
		s := NewServer(l, mux)
		return s.Run(ctx)
	})

	// run関数に終了通知を送信する
	cancel()
	// run関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
