package main

import (
	"net/http"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/alexpts/go-next/next"
	"github.com/alexpts/go-next/next/layer"
)

func FromMux(handler http.HandlerFunc) layer.Handler {
	fasthttpHandler := fasthttpadaptor.NewFastHTTPHandlerFunc(handler)
	return func(cxt *layer.HandlerCtx) {
		fasthttpHandler(cxt.RequestCtx)
	}
}

func netHttpHandlerFunc(w http.ResponseWriter, request *http.Request) {
	_, _ = w.Write([]byte(request.RequestURI))
}

func main() {
	app := next.ProvideMicroApp(nil, nil)

	wrapHandler := FromMux(netHttpHandlerFunc)

	wrapHandlerWithNext := func() layer.Handler {
		return func(ctx *layer.HandlerCtx) {
			fasthttpHandler := fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, request *http.Request) {
				_, _ = w.Write([]byte(request.Method + " "))
				ctx.Next()
			})

			fasthttpHandler(ctx.RequestCtx)
		}
	}()

	app.Use(layer.Layer{}, wrapHandlerWithNext, wrapHandler)

	server := &fasthttp.Server{
		Handler: app.FastHttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
