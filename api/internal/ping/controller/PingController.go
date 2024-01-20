package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (_ *PingController) Index(context *gin.Context) {
	viper.Debug()
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
