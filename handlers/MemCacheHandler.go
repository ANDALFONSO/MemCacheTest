package handlers

import (
	"memcache/entity"
	"memcache/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var memCacheService service.IMemCacheService = service.NewMemCache()

func MemCache(server *gin.Engine) {
	server.GET("/ping", ping)
	server.POST("/mem-cache", createKey)
	server.DELETE("/mem-cache", deleteKey)
	server.GET("/mem-cache", getKey)
}

func createKey(c *gin.Context) {
	key, ok := c.GetQuery("key")
	if !ok {
		c.JSON(404, "key not found")
		return
	}

	var personalData entity.PersonalData
	if err := c.BindJSON(&personalData); err != nil {
		c.JSON(http.StatusInternalServerError, "format invalid")
		return
	}

	c.JSON(http.StatusOK, memCacheService.CreateKey(key, &personalData))
}

func deleteKey(c *gin.Context) {
	key, ok := c.GetQuery("key")
	if !ok {
		c.JSON(http.StatusNotFound, "key not found")
		return
	}
	c.JSON(http.StatusOK, memCacheService.DeleteKey(key))
}

func getKey(c *gin.Context) {
	key, ok := c.GetQuery("key")
	if !ok {
		c.JSON(http.StatusNotFound, "key not found")
		return
	}
	c.JSON(http.StatusOK, memCacheService.GetKey(key))
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
