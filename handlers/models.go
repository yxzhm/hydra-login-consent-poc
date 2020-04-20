package handlers

type AcceptLoginRequest struct {
	Subject string `json:"subject"`
}

type CompletedRequest struct {
	RedirectTo string `json:"redirect_to"`
}

type ConsentRequest struct {
	RequestedScope               []string `json:"requested_scope"`
	RequestedAccessTokenAudience []string `json:"requested_access_token_audience"`
}

type AcceptConsentRequest struct {
	GrantScope               []string `json:"grant_scope"`
	GrantAccessTokenAudience []string `json:"grant_access_token_audience"`
}
