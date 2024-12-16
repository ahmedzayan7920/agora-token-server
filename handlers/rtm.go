package handlers

import (
    "net/http"
    "os"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    rtmtokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtmtokenbuilder"
)

func GetRtmToken(c *gin.Context) {
    appID := os.Getenv("APP_ID")
    appCertificate := os.Getenv("APP_CERTIFICATE")
    
    uidStr := c.Param("uid")
    
    expireTime := c.DefaultQuery("expiry", "3600")
    expireTime64, err := strconv.ParseUint(expireTime, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid expiry time", "error": err.Error()})
        return
    }
    expireTimestamp := uint32(time.Now().UTC().Unix()) + uint32(expireTime64)

    rtmToken, err := rtmtokenbuilder.BuildToken(appID, appCertificate, uidStr, expireTimestamp, "")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate RTM token", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "rtmToken": rtmToken,
    })
}