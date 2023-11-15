package api

import (
    "net/http"
    "go-cyber/schema"
    "github.com/gin-gonic/gin"
)

func DrawGacha(c *gin.Context) {
    // token := c.GetHeader("x-token")
    var req schema.GachaDrawRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ガチャ実行のロジックを実装

    c.JSON(http.StatusOK, gin.H{"message": "Gacha results"})
}

// ListCharacters - ユーザ所持キャラクター一覧取得API
func ListCharacters(c *gin.Context) {
    // token := c.GetHeader("x-token")

    // キャラクター一覧取得のロジックを実装

    c.JSON(http.StatusOK, gin.H{"message": "Character list"})
}
