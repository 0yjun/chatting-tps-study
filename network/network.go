package network

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
type Network struct{
	Engine *gin.Engine
} 

func NewServer()*Network  {
	n := &Network{Engine: gin.New()}
	n.Engine.Use(gin.Logger())
	n.Engine.Use(gin.Recovery())
	n.Engine.Use(cors.New(cors.Config{
		AllowWebSockets: true,
		AllowOrigins: []string{"*"},
		AllowMethods:  []string{"GET","POST","PUT"},
		AllowHeaders:  []string{"*"},
		AllowCredentials: true,
	}))
	return n;
	
}

func (n*Network) StartServer() error {
	log.Println("server is run");
	return n.Engine.Run(":8080");
}