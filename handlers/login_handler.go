package handlers

import (
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	challenge, err := GetInfoFromHydraAdmin(ctx, "login", nil)
	if err != nil {
		return
	}

	//ignore login page here, accept directly.
	acceptRequest := AcceptLoginRequest{Subject: "123"}

	AcceptChallenge(ctx, "login", challenge, acceptRequest)
}
