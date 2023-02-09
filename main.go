package main

import (
	"ticketBooking/utils"

	"github.com/gin-gonic/gin"
)

func init(){
	utils.LoadEnvVariables()
	utils.ConnectingToDatabase()
}

func main(){
	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"Hello Men",
		})
	})

	r.Run()
}



