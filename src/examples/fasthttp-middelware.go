package main

import (
	"github.com/valyala/fasthttp"

	next2 "github.com/alexpts/go-next/next"
)

func appToFasthttpMiddleware(app next2.App, handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	app.Use(next2.Config{}, func(ctx *next2.HandlerCxt) {
		handler(ctx.RequestCtx)
	})

	return app.FasthttpHandler
}

func handler1(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.AppendBodyString(`fasthttp-1-`)
		requestHandler(ctx)
	}
}

func handler2(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.AppendBodyString(`fasthttp-2-`)
		requestHandler(ctx)
	}
}

func handler3(ctx *fasthttp.RequestCtx) {
	ctx.Response.AppendBodyString(`fasthttp-3-`)
}

func main() {
	app := next2.NewApp()

	app.Use(next2.Config{}, func(ctx *next2.HandlerCxt) {
		ctx.Response.AppendBodyString(`next-1-`)
		ctx.Next()
	})

	afterHandler := handler2(handler3)
	appHandler := appToFasthttpMiddleware(app, afterHandler) // app as middleware of fasthttp + delegate to next handler
	allHandler := handler1(appHandler)

	server := &fasthttp.Server{
		Handler: allHandler,
	}

	_ = server.ListenAndServe(":3000")
}
