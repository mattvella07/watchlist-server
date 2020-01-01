package main

import (
	"github.com/mattvella07/watchlist-server/server"
	"github.com/mattvella07/watchlist-server/server/conn"
)

func main() {
	conn.InitDB()
	defer conn.DB.Close()

	server.Start()
}
