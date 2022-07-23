package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
)

func main() {
	app := next.NewApp()
	app.Use(next.Config{}, func(ctx *next.HandlerCxt) {
		ctx.Response.AppendBodyString(`Hello`)
	})

	server := &fasthttp.Server{
		Handler: app.FasthttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
