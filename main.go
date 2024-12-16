package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "agora-token-server/handlers"
    "agora-token-server/middleware"
)

var appID, appCertificate string

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    appID = os.Getenv("APP_ID")
    appCertificate = os.Getenv("APP_CERTIFICATE")

    if appID == "" || appCertificate == "" {
        log.Fatal("FATAL ERROR: APP_ID or APP_CERTIFICATE is missing")
    }
}

func main() {
    r := gin.Default()
    r.Use(middleware.Nocache())
    r.Use(middleware.Cors())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.GET("rtc/:channelName/:role/:tokentype/:uid/", handlers.GetRtcToken)
    r.GET("rtm/:uid/", handlers.GetRtmToken)
    r.GET("rte/:channelName/:role/:tokentype/:uid/", handlers.GetBothTokens)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
