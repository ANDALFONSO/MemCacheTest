package main

import (
	"memcache/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handlers.MemCache(server)
	server.Run(":9090")
}
