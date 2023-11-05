package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	clientIDGoogle     = "861003596694-oiq644rrdk1bnts64a55k9t48ljffnuf.apps.googleusercontent.com"
	clientSecretGoogle = "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"
	redirectURIGoogle  = "http://localhost:8000/google/callback"
)

type googleRespBody struct {
	Acces_token string `json:"access_token"`
}
type googleUsers struct {
	Username string `json:"name"`
}

func (h *Handler) googleLogin(w http.ResponseWriter, r *http.Request) {
	authUrl := "https://accounts.google.com/o/oauth2/v2/auth"
	params := url.Values{}
	params.Add("client_id", clientIDGoogle)
	params.Add("redirect_uri", redirectURIGoogle)
	params.Add("scope", "https://www.googleapis.com/auth/userinfo.profile")
	params.Add("response_type", "code")
	redirectURL := fmt.Sprintf("%s?%s", authUrl, params.Encode())

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (h *Handler) googleLoginCallBack(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	fmt.Println(code)
	tokenURL := "https://accounts.google.com/o/oauth2/token"
	data := url.Values{}
	data.Add("code", code)
	data.Add("client_id", clientIDGoogle)
	data.Add("client_secret", clientSecretGoogle)
	data.Add("redirect_uri", redirectURIGoogle)
	data.Add("grant_type", "authorization_code")

	accessUrl := fmt.Sprintf("%s?%s", tokenURL, data.Encode())
	fmt.Println("url", accessUrl)
	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		fmt.Println("HEre")
		h.logError(w, r, err, 500)
		return
	}
	defer resp.Body.Close()
	var body googleRespBody
	err = json.NewDecoder(resp.Body).Decode(&body)

	fmt.Println("TOKEN:", body.Acces_token)

	apiURL := "https://www.googleapis.com/oauth2/v3/userinfo"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+body.Acces_token)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Client err:", err)
		return
	}
	defer response.Body.Close()
	info, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading response ", err)
		return
	}
	var googleUser googleUsers
	err = json.Unmarshal(info, &googleUser)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}

	exsits, err := h.Service.Authorization.GetUserByName(googleUser.Username)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	if !exsits {
		fmt.Println("HERE")
		cookie, err := h.Service.Authorization.CreateUserOauth(googleUser.Username)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		fmt.Println("COOKIE:", cookie)
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	} else {
		cookie, err := h.Service.Authorization.UpdateSession(googleUser.Username)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		fmt.Println("COOKIE:", cookie)
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	}
}
