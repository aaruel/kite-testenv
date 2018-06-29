package main

import (
	"github.com/aaruel/kite-testenv/gateway/repo"
	"github.com/aaruel/kite-testenv/gateway/router"
)

func main() {
	repoInst := repo.New()
	router.NewRouter(repoInst)
	select {}
}
