package handlers

import (
	"github.com/gin-gonic/gin"
)

func ConsentHandler(ctx *gin.Context) {
	consentRequest := ConsentRequest{}
	challenge, err := GetInfoFromHydraAdmin(ctx, "consent", &consentRequest)
	if err != nil {
		return
	}

	//ignore consent page here, accept directly.
	acceptConsentRequest := AcceptConsentRequest{
		GrantScope:               consentRequest.RequestedScope,
		GrantAccessTokenAudience: consentRequest.RequestedAccessTokenAudience,
	}
	AcceptChallenge(ctx, "consent", challenge, acceptConsentRequest)
}
