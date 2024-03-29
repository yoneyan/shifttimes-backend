package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func ControllerAPI() {
	router := gin.Default()
	router.Use(cors)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/schedule")

			//
			// Storage
			//
			//v1.POST("/storage", storage.Add)
			//v1.GET("/storage", storage.GetAll)
			//v1.GET("/storage/uuid/:uuid", storage.GetUUID)
			//v1.GET("/storage/id/:id", storage.Get)
			//v1.GET("/storage/name/:name", storage.GetName)
			//v1.PUT("/storage/:id", storage.Update)
			//Download
		}
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(8080), router))
}

func cors(c *gin.Context) {

	//c.Header("Access-Control-Allow-Headers", "Accept, Content-ID, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-ID", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
