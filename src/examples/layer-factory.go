package main

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/alexpts/go-next/next/layer"
)

type RouterConfig struct {
	Path         string
	Methods      string
	Controller   string
	Name         string
	Priority     int
	Handler      layer.Handler
	Restrictions layer.Restrictions
}

var HandlerMap = map[string]layer.Handler{
	"otherwise": MainPageAppHandler2,
	"mainPage":  MainPageAppHandler2,
	"mainPage2": MainPageAppHandler2,
	"hello":     MainPageAppHandler2,
}

func CreateLayers(projectDir string) []*layer.Layer {
	rawData, err := os.ReadFile(projectDir + "/config/router.yml")

	var routes map[string]RouterConfig
	err = yaml.Unmarshal(rawData, &routes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return factoryLayers(routes)
}

func factoryLayers(routes map[string]RouterConfig) []*layer.Layer {
	var layers []*layer.Layer
	var methods []string

	for name, route := range routes {
		handler := HandlerMap[route.Controller]
		if route.Name == `` {
			route.Name = name
		}

		if route.Methods != `` {
			methods = strings.Split(route.Methods, `|`) // GET|POST|PUT
		}

		l := layer.Layer{
			Handlers:     []layer.Handler{handler},
			Priority:     route.Priority,
			Name:         route.Name,
			Path:         route.Path,
			Restrictions: route.Restrictions,
			Methods:      methods,
		}

		layers = append(layers, &l)
	}

	return layers
}

func MainPageAppHandler2(ctx *layer.HandlerCtx) {
	ctx.Response.AppendBodyString(`MainPageAppHandler2`)
}
