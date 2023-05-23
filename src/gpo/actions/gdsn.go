package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
)

var GDSNBaseUrl = "https://healthcare-test.firstbase.ch/api/v1/"
var GDSNAuthPath = "/login"
var GDSNAuthToken string

type GDSNCredentials struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type GDSNAuth struct {
	Url         string
	Credentials GDSNCredentials
}

func (a GDSNAuth) getAuthToken() (*http.Cookie, error) {
	jsonValue, err := json.Marshal(a.Credentials)
	if err != nil {
		return nil, fmt.Errorf("composing credentials: %s", err)
	}
	response, err := http.Post(a.Url, "application/json;charset=UTF-8", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("login with GDSN: %s", err)
	}
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing response body: %s", err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status code: %d (%s)", response.StatusCode, string(bodyBytes))
	}

	var authCookie *http.Cookie
	for _, cookie := range response.Cookies() {
		if cookie.Name == "AUTHENTICATION_TOKEN" {
			authCookie = cookie
		}
	}
	return authCookie, nil
}

func NewGDSNAuth() GDSNAuth {
	return GDSNAuth{
		Url: GDSNBaseUrl + GDSNAuthPath,
		Credentials: GDSNCredentials{
			Username: os.Getenv("GDSN_USER"),
			Password: os.Getenv("GDSN_PASSWORD"),
		},
	}
}

// https://healthcare-test.firstbase.ch/api/v1/items/60884522006724:0673978000008:756
func (q GDSNQuery) getItemByGTIN(gtin string, gln string, countryCode string) (string, error) {
	url := fmt.Sprintf("%s/items/%s:%s:%s", GDSNBaseUrl, gtin, gln, countryCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("setting up new get request: %s", err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", fmt.Errorf("setting up cookie jar: %s", err)
	}
	jar.SetCookies(req.URL, []*http.Cookie{&q.authCookie})
	client := &http.Client{Jar: jar}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("querying GDSN database: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("querying GDSN database: unexpected HTTP status code: %d", response.StatusCode)
	}

	// var content map[string]interface{}
	// err = json.NewDecoder(response.Body).Decode(&content)
	// if err != nil {
	//     return "", fmt.Errorf("decoding json reponse: %s", err)
	// }
	// return content, err
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("querying GDSN database: reading body: %s", err)
	}
	return string(bodyBytes), nil
}

//https://healthcare-test.firstbase.ch/api/v1/items?format=json&keyword=manufacturerName:Med*&count=100&fields=eclassCode&fields=brandName

// func (q GDSNQuery) getItemsByKeyword(keyword string, allowNonCompliantItems bool) (string, error) {

// }

type GDSNQuery struct {
	domain     string `default:"https://healthcare-test.firstbase.ch/"`
	format     string `default:"json"`
	authCookie http.Cookie
}

// gdsnAuthCookie, err := NewGDSNAuth().getAuthToken()

// if err != nil {
// 	return fmt.Errorf("GDSN authentication: %s", err)
// }
// q := GDSNQuery{
// 	authCookie: *gdsnAuthCookie,
// }
// reponse, err := q.getItemByGTIN("60884522006724", "0673978000008", "756")
// if err != nil {
// 	log.Fatal(err)
// }
// c.Logger().Info(reponse)
