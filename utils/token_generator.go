package utils

import (
    "fmt"
    "strconv"
    "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
)

// GenerateRtcToken generates a token for Agora RTC service
func GenerateRtcToken(appID, appCertificate, channelName, uidStr, tokentype string, role rtctokenbuilder2.Role, expireTimestamp uint32) (string, error) {
    if appID == "" || appCertificate == "" {
        return "", fmt.Errorf("APP_ID or APP_CERTIFICATE is missing")
    }

    if tokentype == "userAccount" {
        token, err := rtctokenbuilder2.BuildTokenWithAccount(appID, appCertificate, channelName, uidStr, role, expireTimestamp)
        if err != nil {
            return "", fmt.Errorf("failed to build token with account: %v", err)
        }
        return token, nil
    }

    uid64, err := strconv.ParseUint(uidStr, 10, 64)
    if err != nil {
        return "", fmt.Errorf("failed to parse uid: %v", err)
    }

    token, err := rtctokenbuilder2.BuildTokenWithUid(appID, appCertificate, channelName, uint32(uid64), role, expireTimestamp)
    if err != nil {
        return "", fmt.Errorf("failed to build token with UID: %v", err)
    }
    return token, nil
}