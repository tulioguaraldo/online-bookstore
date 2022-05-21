package main

import (
	env "github.com/tulioguaraldo/online-bookstore/config/environment"
	"github.com/tulioguaraldo/online-bookstore/server"
)

func main() {
	env.GetEnvironmentConfig()

	server := server.NewServer()
	server.RunServer()
}
