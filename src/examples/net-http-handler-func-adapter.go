package main

import (
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
)

func netHttpHandlerFunc(w http.ResponseWriter, request *http.Request) {
	_, _ = w.Write([]byte(request.RequestURI))
}

func main() {
	app := next.NewApp()

	wrapHandler := next.FromHttpHandlerFunc(netHttpHandlerFunc)
	app.Use(next.Config{}, wrapHandler)

	server := &fasthttp.Server{
		Handler: app.FasthttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
