package main

import (
    "go-cyber/api"
    "go-cyber/db"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    db := db.InitDB()
    defer db.Close()

    r.POST("/user/create", api.CreateUser)
    r.GET("/user/get", api.GetUser)
    r.PUT("/user/update", api.UpdateUser)
    r.POST("/gacha/draw", api.DrawGacha)
    r.GET("/character/list", api.ListCharacters)

    r.Run(":8080")
}
