# hydra-login-and-consent-services

This is a reference implementation for the User Login and Consent flow of ORY Hydra version 1.0.x in Golang. 
This POC ignores the login consent pages, accept the challenge directly.

Use below hydra command to create the client_id and client_secret
```
hydra clients create --endpoint http://127.0.0.1:4445/ --id code-client --secret secret --grant-types authorization_code,refresh_token --response-types code,id_token  --callbacks http://127.0.0.1:9092/callback --token-endpoint-auth-method client_secret_post
```
