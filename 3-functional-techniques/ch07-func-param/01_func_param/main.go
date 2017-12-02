package main

import (
	"server"
	. "utils"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
	"fmt"
)

func init() {
	GetOptions()
	InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	Info.Printf("Config %+v", Config)
	newServer, err := server.New(
		server.MaxConcurrentConnections(4),
		server.MaxNumber(256), // Config.MaxNumber
		server.UseNumberHandler(true),
		server.FormatNumber(func(x int) (string, error) { return fmt.Sprintf("%x", x), nil }),  // anonymous fcn
		//server.FormatNumber(func(x int) (string, error) { return fmt.Sprintf("%b", x), nil }),  // anonymous fcn
		//server.FormatNumber(func(x int) (string, error) { return "", errors.New("FormatNumber error") }),  // anonymous fcn
	)
	if err != nil {
		Error.Printf("unable to initialize server: %v", err)
		os.Exit(1)
	}
	srv := &http.Server{
		Addr:    ":"+ Config.Port,
		Handler: newServer,
	}

	go func() {
		<-quit
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
		defer cancel()
		Info.Println("shutting down server...")
		if err := srv.Shutdown( ctx ); err != nil {
			Error.Printf("unable to shutdown server: %v", err)
		}
	}()
	Error.Println("server started at localhost:"+ Config.Port)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		Error.Printf("ListenAndServe error: %v", err)
	}
	Info.Println("server shutdown gracefully")
}
