package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	Version  string
	Released bool
}

type Stream struct {
	Name          string
	EngineVersion string
}

type Devspace struct {
	Version string
}

/*
all engines
/v1/engines

released engines
/v1/engines?released=true

used engines
/v1/streams

devspace info
/v1/devspace

*/

func main() {
	engines := make([]Engine, 0, 32)
	// streams := make([]Stream, 0, 8)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "DevSpace API Server is Running")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/v1/engines", func(c *gin.Context) {
		c.JSON(http.StatusOK, engines)
	})

	router.POST("/v1/engines", func(c *gin.Context) {
		engine := Engine{}

		if c.ShouldBindJSON(&engine) == nil {
			c.String(http.StatusOK, "Failed to Bind Engine")
			return
		}

		engines = append(engines, engine)

		c.JSON(http.StatusOK, engine)
	})

	// router.GET("/v1/streams", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, streams)
	// })

	// router.POST("/v1/streams", func(c *gin.Context) {
	// 	stream := Stream{}

	// 	if c.ShouldBindJSON(&stream) == nil {
	// 		c.String(http.StatusOK, "Failed to Bind Stream")
	// 		return
	// 	}

	// 	streams = append(streams, stream)

	// 	c.JSON(http.StatusOK, stream)
	// })

	// router.POST("/user/:name/*action", func(c *gin.Context) {
	// 	if c.FullPath() == "/user/:name/*action" {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "name : " + c.Param("name"),
	// 		})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "not matched",
	// 		})
	// 	}
	// })

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
