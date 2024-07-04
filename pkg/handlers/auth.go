package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"jilt.com/m/pkg/models"
)

var (
	googleOauthConfig *oauth2.Config
	facebookOauthConfig *oauth2.Config
	githubOauthConfig *oauth2.Config
	discordOauthConfig *oauth2.Config
)

func init() {
	// Replace these with your Google OAuth credentials
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google",
		ClientID:     "760820018039-vhlp72gm70411lf6sla8125cg49t9jdd.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/facebook",
		ClientID:     "390126090201300",
		ClientSecret: "61138312baa063576def63bf49af0b40",
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
	githubOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/github",
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func GenerateStateOauthCookie(w http.ResponseWriter) string {
    var expiration = 365 * 24 * 60 * 60
    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: time.Now().Add(time.Second * time.Duration(expiration))}
    http.SetCookie(w, &cookie)
    return state
}

// GET /login
func Login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "navbar", "login")
}

// GET /signup
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "navbar", "signup")
}

func Profile(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if user.IsAdmin {
			generateHTML(writer, user, "layout", "admin.navbar", "auth.profile")
		} else {
			generateHTML(writer, user, "layout", "auth.navbar", "auth.profile")
		}
	}
}

func EditProfile(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		generateHTML(writer, user, "layout", "auth.navbar", "auth.edit_profile")
	}
}

// POST /signup
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	success("User created: ", user.Name)
	http.Redirect(writer, request, "/login", http.StatusFound)
}

// POST /authenticate
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find email")
	}
	if user.CheckPassword(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}

		success("Authenticated user: ", user.Email)
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", http.StatusFound)
	} else {
		http.Redirect(writer, request, "/login", http.StatusFound)
	}
}

// GET /logout
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", http.StatusFound)
}

const (
	DiscordclientID     = "1258028813489668126"
	DiscordclientSecret = "_dDeeGjWrt2UCqU-boMS4GBiVZ9JIJgd"
	DiscordredirectURI  = "http://10.31.35.143:8080/auth/discord"
	Discordredirect_URI = "http%3A%2F%2F10.31.35.143%3A8080%2Fauth%2Fdiscord"
)

func secureRandomString(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Discord(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("access_token") == "" {
		url := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=identify+email", DiscordclientID, Discordredirect_URI)
		http.Redirect(writer, request, url, http.StatusFound)
		return
	} else {
		accessToken := request.URL.Query().Get("access_token")
		tokenType := request.URL.Query().Get("token_type")
		url := "https://discord.com/api/users/@me"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(writer, "Failed to create request", http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, accessToken))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(writer, "Failed to send request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(writer, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		var discordUser map[string]interface{}
		if err := json.Unmarshal(body, &discordUser); err != nil {
			http.Error(writer, "Failed to parse JSON", http.StatusInternalServerError)
			return
		}

		user, err := models.UserByDiscordID(discordUser["id"].(string))
		if err != nil {
			secureRandomResult, err := secureRandomString(16)
			if err != nil {
				http.Error(writer, "Failed to generate secure random string", http.StatusInternalServerError)
				return
			}
			user = models.User{
				Name:      discordUser["global_name"].(string),
				Email:     discordUser["email"].(string),
				Password:  secureRandomResult,
				DiscordID: discordUser["id"].(string),
			}
			if err := user.Create(); err != nil {
				http.Error(writer, "Cannot create user", http.StatusInternalServerError)
				return
			}
		}

		session, err := user.CreateSession()
		if err != nil {
			http.Error(writer, "Cannot create session", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", http.StatusFound)
	}
}

func AuthDiscord(writer http.ResponseWriter, request *http.Request) {
	code := request.URL.Query().Get("code")
	if code == "" {
		http.Redirect(writer, request, "/login", http.StatusFound)
		return
	}

	tokenURL := "https://discord.com/api/oauth2/token"
	data := url.Values{
		"client_id":     {DiscordclientID},
		"client_secret": {DiscordclientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {DiscordredirectURI},
	}

	resp, err := http.PostForm(tokenURL, data)
	if err != nil {
		http.Error(writer, "Failed to exchange token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(writer, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		http.Error(writer, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		http.Error(writer, "Invalid access token", http.StatusInternalServerError)
		return
	}

	// Redirect to frontend with access token in query parameters
	redirectURL := fmt.Sprintf("/discord?access_token=%s&token_type=%s", accessToken, tokenResponse["token_type"].(string))
	http.Redirect(writer, request, redirectURL, http.StatusFound)
}
func Google(writer http.ResponseWriter, request *http.Request) {
	state := GenerateStateOauthCookie(writer)
	url := googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func AuthGoogle(w http.ResponseWriter, r *http.Request) {
    state := r.FormValue("state")
	oauthState, _ := r.Cookie("oauthstate")
	fmt.Println("state: ", state)
	fmt.Println("oauthState: ", oauthState)
    if state != oauthState.Value {
        fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthState.Value, state)
        generateHTML(w, nil, "layout", "navbar", "error")
        return
    }

    code := r.FormValue("code")
    token, err := googleOauthConfig.Exchange(r.Context(), code)
    if err != nil {
        fmt.Printf("Code exchange failed: %v\n", err)
        generateHTML(w, nil, "layout", "navbar", "error")
        return
    }

    // Exemple de récupération des informations de l'utilisateur avec le token d'accès
    client := googleOauthConfig.Client(r.Context(), token)
    response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        fmt.Printf("Failed to get userinfo: %v\n", err)
        generateHTML(w, nil, "layout", "navbar", "error")
        return
    }
    defer response.Body.Close()

    var userInfo map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&userInfo)
    if err != nil {
        fmt.Printf("Failed to parse userinfo response: %v\n", err)
        generateHTML(w, nil, "layout", "navbar", "error")
        return
    }

	emptyUser := models.User{}
	lookuser, _ := models.UserByEmail(userInfo["email"].(string))
	if lookuser != emptyUser {
		http.Redirect(w, r, "/err?msg=Email already used", http.StatusTemporaryRedirect)
		return
	}

	user, err := models.UserByGoogleID(userInfo["id"].(string))
	if err != nil {
		secureRandomResult, err := secureRandomString(16)
		if err != nil {
			http.Error(w, "Failed to generate secure random string", http.StatusInternalServerError)
			return
		}
		user = models.User{
			Name:     userInfo["name"].(string),
			Email:    userInfo["email"].(string),
			Password: secureRandomResult,
			GoogleID: userInfo["id"].(string),
		}
		if err := user.Create(); err != nil {
			http.Error(w, "Cannot create user", http.StatusInternalServerError)
			return
		}
	}

	session, err := user.CreateSession()
	if err != nil {
		http.Error(w, "Cannot create session", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}


func Facebook(writer http.ResponseWriter, request *http.Request) {
	state := GenerateStateOauthCookie(writer)
	url := facebookOauthConfig.AuthCodeURL(state)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func AuthFacebook(writer http.ResponseWriter, request *http.Request) {
	state := request.FormValue("state")
	oauthState, _ := request.Cookie("oauthstate")
	if state != oauthState.Value {
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthState.Value, state)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	code := request.FormValue("code")
	token, err := facebookOauthConfig.Exchange(request.Context(), code)
	if err != nil {
		fmt.Printf("Code exchange failed: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	// Exemple de récupération des informations de l'utilisateur avec le token d'accès
	client := facebookOauthConfig.Client(request.Context(), token)
	response, err := client.Get("https://graph.facebook.com/me?fields=id,name,email")
	if err != nil {
		fmt.Printf("Failed to get userinfo: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		fmt.Printf("Failed to parse userinfo response: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	emptyUser := models.User{}
	lookuser, _ := models.UserByEmail(userInfo["email"].(string))
	if lookuser != emptyUser {
		http.Redirect(writer, request, "/err?msg=Email already used", http.StatusTemporaryRedirect)
		return
	}

	user, err := models.UserByFacebookID(userInfo["id"].(string))
	if err != nil {
		secureRandomResult, err := secureRandomString(16)
		if err != nil {
			http.Error(writer, "Failed to generate secure random string", http.StatusInternalServerError)
			return
		}
		user = models.User{
			Name:        userInfo["name"].(string),
			Email:       userInfo["email"].(string),
			Password:    secureRandomResult,
			FacebookID:  userInfo["id"].(string),
		}
		if err := user.Create(); err != nil {
			http.Error(writer, "Cannot create user", http.StatusInternalServerError)
			return
		}
	}

	session, err := user.CreateSession()
	if err != nil {
		http.Error(writer, "Cannot create session", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", http.StatusFound)
}

func Github(writer http.ResponseWriter, request *http.Request) {
	state := GenerateStateOauthCookie(writer)
	url := githubOauthConfig.AuthCodeURL(state)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func AuthGithub(writer http.ResponseWriter, request *http.Request) {
	state := request.FormValue("state")
	oauthState, _ := request.Cookie("oauthstate")
	if state != oauthState.Value {
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthState.Value, state)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	code := request.FormValue("code")
	token, err := githubOauthConfig.Exchange(request.Context(), code)
	if err != nil {
		fmt.Printf("Code exchange failed: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	// Exemple de récupération des informations de l'utilisateur avec le token d'accès
	client := githubOauthConfig.Client(request.Context(), token)
	response, err := client.Get("https://api.github.com/user")
	if err != nil {
		fmt.Printf("Failed to get userinfo: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		fmt.Printf("Failed to parse userinfo response: %v\n", err)
		generateHTML(writer, nil, "layout", "navbar", "error")
		return
	}

	emptyUser := models.User{}
	lookuser, _ := models.UserByEmail(userInfo["email"].(string))
	if lookuser != emptyUser {
		http.Redirect(writer, request, "/err?msg=Email already used", http.StatusTemporaryRedirect)
		return
	}

	user, err := models.UserByGithubID(userInfo["id"].(string))
	if err != nil {
		secureRandomResult, err := secureRandomString(16)
		if err != nil {
			http.Error(writer, "Failed to generate secure random string", http.StatusInternalServerError)
			return
		}
		user = models.User{
			Name:      userInfo["name"].(string),
			Email:     userInfo["email"].(string),
			Password:  secureRandomResult,
			GithubID:  userInfo["id"].(string),
		}
		if err := user.Create(); err != nil {
			http.Error(writer, "Cannot create user", http.StatusInternalServerError)
			return
		}
	}

	session, err := user.CreateSession()
	if err != nil {
		http.Error(writer, "Cannot create session", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	http.Redirect(writer, request, "/", http.StatusFound)
}