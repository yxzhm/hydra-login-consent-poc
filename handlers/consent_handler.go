package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ConsentHandler(ctx *gin.Context) {
	consentRequest := ConsentRequest{}
	challenge, err := GetInfoFromHydraAdmin(ctx, "consent", &consentRequest)
	if err != nil {
		return
	}
	oemCode := GetChallengeAndOEMCode(consentRequest.LoginChallenge)
	//Get OEM authorization code here, shall get the OEM token here or somewhere else
	//
	log.Printf(oemCode)

	//ignore consent page here, accept directly.
	acceptConsentRequest := AcceptConsentRequest{
		GrantScope:               consentRequest.RequestedScope,
		GrantAccessTokenAudience: consentRequest.RequestedAccessTokenAudience,
	}
	AcceptChallenge(ctx, "consent", challenge, acceptConsentRequest)
}
