package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AcceptToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

const (
	clientID          = "8d4b1ea8b6f092ffe262"
	clientSecret      = "1166c105718802c643c6e9c237478321af1b9381"
	gitHubredirectURI = "http://localhost:8000/github/callback"
)

type githubInfo struct {
	Username string `json:"login"`
	Email    string `json:"email"`
}

func (h *Handler) githubLogin(w http.ResponseWriter, r *http.Request) {
	authUrl := "https://github.com/login/oauth/authorize"
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("redirect_uri", gitHubredirectURI)
	params.Add("scope", "user")
	params.Add("response_type", "code")
	redirectURL := fmt.Sprintf("%s?%s", authUrl, params.Encode())

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func (h *Handler) githubLoginCallBack(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	fmt.Println(code)
	tokenURL := "https://github.com/login/oauth/access_token"
	data := url.Values{}
	data.Add("code", code)
	data.Add("client_id", clientID)
	data.Add("client_secret", clientSecret)
	data.Add("redirect_uri", gitHubredirectURI)

	accessUrl := fmt.Sprintf("%s?%s", tokenURL, data.Encode())
	fmt.Println("url", accessUrl)
	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		fmt.Println("HEre")
		h.logError(w, r, err, 500)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	elements := strings.Split(string(body), "&")

	fmt.Println("token", elements[0][13:])

	// Use the access token to make a request to the GitHub user API
	apiURL := "https://api.github.com/user"
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+elements[0][13:])

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
	var githubUser githubInfo
	err = json.Unmarshal(info, &githubUser)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}
	fmt.Println("USER", githubUser)

	exsits, err := h.Service.Authorization.GetUserByName(githubUser.Username)
	if err != nil {
		h.logError(w, r, err, http.StatusBadRequest)
		return
	}

	if !exsits {
		fmt.Println("HERE")
		cookie, err := h.Service.Authorization.CreateUserOauth(githubUser.Username)
		if err != nil {
			h.logError(w, r, err, http.StatusBadRequest)
			return
		}
		fmt.Println("COOKIE:", cookie)
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "/", http.StatusSeeOther)

		return
	} else {
		cookie, err := h.Service.Authorization.UpdateSession(githubUser.Username)
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
