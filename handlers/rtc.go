package handlers

import (
    "net/http"
    "os"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    rtctokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
    "agora-token-server/utils"
)

func GetRtcToken(c *gin.Context) {
    appID := os.Getenv("APP_ID")
    appCertificate := os.Getenv("APP_CERTIFICATE")

    channelName := c.Param("channelName")
    roleStr := c.Param("role")
    tokentype := c.Param("tokentype")
    uidStr := c.Param("uid")

    if channelName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Channel name is required"})
        return
    }

    var role rtctokenbuilder.Role
    if roleStr == "publisher" {
        role = rtctokenbuilder.Role(1) // RolePublisher
    } else {
        role = rtctokenbuilder.Role(2) // RoleSubscriber
    }

    expireTime := c.DefaultQuery("expiry", "3600")
    expireTime64, err := strconv.ParseUint(expireTime, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid expiry time", "error": err.Error()})
        return
    }
    expireTimestamp := uint32(time.Now().UTC().Unix()) + uint32(expireTime64)

    token, err := utils.GenerateRtcToken(appID, appCertificate, channelName, uidStr, tokentype, role, expireTimestamp)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate RTC token", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "rtcToken": token,
        "channel": channelName,
    })
}