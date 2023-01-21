package main

import (
	"github.com/valyala/fasthttp"

	"github.com/alexpts/go-next/next"
	"github.com/alexpts/go-next/next/layer"
)

func Action(ctx *layer.HandlerCtx) {
	name, _ := ctx.UriParams["name"]
	ctx.Response.AppendBodyString(`Hello ` + name)
}

func Fallback404(ctx *layer.HandlerCtx) {
	ctx.Response.SetStatusCode(404)
	ctx.SetContentType("application/json")
	ctx.Response.AppendBody([]byte(`{"error": "not found handler"}`))
}

func main() {
	app := next.ProvideMicroApp(nil, nil)

	app.
		Use(layer.Layer{
			Path:     `/hello/{name}/`,
			Name:     `HelloAction`,
			Methods:  []string{`GET`, `POST`},
			Priority: 100,
			Restrictions: layer.Restrictions{
				`name`: `[a-z]+`,
			},
		}, Action).
		Use(layer.Layer{
			Name:     "Fallback 404 runner layer",
			Priority: -9999,
		}, Fallback404)

	server := &fasthttp.Server{
		Handler: app.FastHttpHandler,
	}

	_ = server.ListenAndServe(":3000")
}
