package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var HydraAdmin = ""

func GetInfoFromHydraAdmin(ctx *gin.Context, flow string, result interface{}) (string, error) {
	challenge := ctx.Request.URL.Query().Get(fmt.Sprintf("%s_challenge", flow))
	log.Printf("challenge: %s", challenge)

	url := fmt.Sprintf("%s/oauth2/auth/requests/%s?challenge=%s", HydraAdmin, flow, challenge)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if result != nil {
		json.Unmarshal(body, result)
	}
	return challenge, nil
}

func AcceptChallenge(ctx *gin.Context, flow string, challenge string, acceptRequest interface{}) error {
	acceptUrl := fmt.Sprintf("%s/oauth2/auth/requests/%s/accept?challenge=%s", HydraAdmin, flow, challenge)
	acceptRequestBody, err := json.Marshal(acceptRequest)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, acceptUrl, bytes.NewBuffer(acceptRequestBody))
	hydraAcceptRes, err := client.Do(req)
	if err != nil {
		return err
	}

	hydraAcceptBody, err := ioutil.ReadAll(hydraAcceptRes.Body)
	if err != nil {
		return err
	}

	completedRequest := CompletedRequest{}
	if err := json.Unmarshal(hydraAcceptBody, &completedRequest); err != nil {
		return err
	}
	return render.Redirect{
		Code:     http.StatusFound,
		Request:  ctx.Request,
		Location: completedRequest.RedirectTo,
	}.Render(ctx.Writer)
}

func GetEnvStrValue(key string, defaultValue string) string {
	var envValue = os.Getenv(key)
	if envValue == "" {
		log.Printf(fmt.Sprintf("Get ENV [%s] -> [%s] (by default)", key, defaultValue))
		return defaultValue
	} else {
		log.Printf(fmt.Sprintf("Get ENV [%s] -> [%s]", key, envValue))
		return envValue
	}
}
