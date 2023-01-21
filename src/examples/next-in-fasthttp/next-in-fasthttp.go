// main
// fasthttp -> fasthttp -> next -> fasthttp
package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
	"github.com/alexpts/go-next/next/layer"
)

// nextAppToFasthttpMd - convert next app to fasthttp middleware
func nextAppToFasthttpMd(app *next.MicroApp, handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	app.Use(layer.Layer{}, func(ctx *layer.HandlerCtx) {
		handler(ctx.RequestCtx)
	})

	return app.FastHttpHandler
}

func fasthttpMd1(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.AppendBodyString(`fasthttpMd-1 | `)
		next(ctx)
	}
}

func fasthttpMd2(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.AppendBodyString(`fasthttpMd-2 | `)
		next(ctx)
	}
}

func fasthttpHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.AppendBodyString(`fasthttp-3 | `)
}

func main() {
	app := next.ProvideMicroApp(nil, nil)

	app.Use(layer.Layer{}, func(ctx *layer.HandlerCtx) {
		ctx.Response.AppendBodyString(`next | `)
		ctx.Next()
	})

	handler := fasthttpMd2(fasthttpHandler)
	appAsFasthttpMd := nextAppToFasthttpMd(&app, handler) // app as middleware of fasthttp + delegate to next handler
	allHandler := fasthttpMd1(appAsFasthttpMd)

	server := &fasthttp.Server{
		Handler: allHandler,
	}

	_ = server.ListenAndServe(":3000")
}
