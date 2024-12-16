package middleware

import "github.com/gin-gonic/gin"

func Nocache() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Cache-Control", "private, no-cache, no-store, must-revalidate")
        c.Header("Expires", "-1")
        c.Header("Pragma", "no-cache")
        c.Header("Access-Control-Allow-Origin", "*")
    }
}
