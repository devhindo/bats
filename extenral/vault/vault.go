package vault

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
)



func GetEnvVariable(name s) {
	env := ""
	return env
}


/*
HCP_API_TOKEN=$(curl --location "https://auth.idp.hashicorp.com/oauth2/token" \
--header "Content-Type: application/x-www-form-urlencoded" \
--data-urlencode "client_id=$HCP_CLIENT_ID" \
--data-urlencode "client_secret=$HCP_CLIENT_SECRET" \
--data-urlencode "grant_type=client_credentials" \
--data-urlencode "audience=https://api.hashicorp.cloud" | jq -r .access_token)
*/

func generate_HCP_API_TOKEN(HCP_CLIENT_ID string, HCP_CLIENT_SECRET string) (string, error) {
	data := url.Values{}
	data.Set("client_id", "HCP_CLIENT_ID")
	data.Set("client_secret", HCP_CLIENT_SECRET)
	data.Set("grant_type", "client_credentials")
	data.Set("audience", "https://api.hashicorp.cloud")

	req, err := http.NewRequest("POST", "https://auth.idp.hashicorp.com/oauth2/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	a:=http.StatusAccepted

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set()


}

/*
-- read secrets
curl \
--location  "env.SECRET_LOC"\
--request GET \
--header "Authorization: Bearer $HCP_API_TOKEN" | jq
*/