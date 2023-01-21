package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
	"github.com/alexpts/go-next/next/layer"
)

type Layer = layer.Layer
type HandlerCtx = layer.HandlerCtx

func main() {
	app := next.ProvideMicroApp(nil, nil)
	app.Use(Layer{}, func(ctx *HandlerCtx) {
		ctx.Response.AppendBodyString(`Hello`)
	})

	server := &fasthttp.Server{
		Handler: app.FastHttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
