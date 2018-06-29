package main

import (
	"net/url"
	"strconv"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
)

func square(r *kite.Request) (interface{}, error) {
	a := r.Args.One().MustFloat64()
	return a * a, nil
}

func registerService(service *kite.Kite, port int) {
	err := service.RegisterForever(&url.URL{
		Scheme: "http",
		Host:   "localhost:" + strconv.Itoa(port),
		Path:   "kite",
	})
	if err != nil {
		panic(err)
	}
}

func start() {
	var port int = 5999

	service := kite.New("math", "1.0.0")
	service.Config = config.MustGet()
	service.Config.Port = port
	service.Config.KontrolURL = "http://localhost:6789/kite"

	service.HandleFunc("square", square)

	registerService(service, port)

	service.Run()
}

func main() {
	start()
}
