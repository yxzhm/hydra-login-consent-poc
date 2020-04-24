package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

func LoginHandler(ctx *gin.Context) {
	challenge, err := GetInfoFromHydraAdmin(ctx, "login", nil)
	if err != nil {
		return
	}

	render.Redirect{
		Code:     http.StatusFound,
		Request:  ctx.Request,
		Location: fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&response_type=code&scope=user:email+read:org&state=%s", OemClientID, challenge),
	}.Render(ctx.Writer)

}

func AcceptLogin(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("state")
	code := c.Request.URL.Query().Get("code")
	SetChallengeAndOEMCode(challenge, code)
	if code != "" && challenge != "" {
		acceptRequest := AcceptLoginRequest{Subject: "123", Context: OEMCode{
			Code: code,
		}}

		AcceptChallenge(c, "login", challenge, acceptRequest)
	}
}
