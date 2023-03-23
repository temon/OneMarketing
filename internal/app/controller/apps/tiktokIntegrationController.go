package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"os"
)

//api := &tiktokAdsV1_3.TiktokAdsAPIv1_3{
//	AccessToken: "my-access-token",
//	BaseURL:     "https://ads.tiktok.com/api",
//}
//adapter := &tiktokAds.TiktokAdsAPIAdapter{Api: api}
//
//campaigns, err := adapter.GetCampaigns()
//if err != nil {
//	fmt.Println("error")
//}
//fmt.Println("testing")
//fmt.Println(campaigns)

func TiktokAuthorize(c *gin.Context) {
	appId := os.Getenv("TIKTOK_ADS_APP_ID")
	redirectUri := os.Getenv("TIKTOK_ADS_REDIRECT_URI")
	state := os.Getenv("TIKTOK_ADS_STATE")
	scope := os.Getenv("TIKTOK_ADS_STATE")

	authURL := fmt.Sprintf("https://ads.tiktok.com/open_api/authorize?app_id=%s&redirect_uri=%s&response_type=code&state=%s&scope=%s",
		appId,
		redirectUri,
		state,
		scope,
	)

	c.Redirect(http.StatusFound, authURL)
}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func TiktokCallbackHandler(c *gin.Context) {

	appId := os.Getenv("TIKTOK_ADS_APP_ID")
	redirectUri := os.Getenv("TIKTOK_ADS_REDIRECT_URI")
	secret := os.Getenv("TIKTOK_ADS_APP_SECRET")

	code := c.Query("code")

	// todo: move this into a service
	// get the token
	tokenURL := "https://ads.tiktok.com/open_api/oauth2/access_token"
	req, _ := http.NewRequest("POST", tokenURL, nil)
	q := req.URL.Query()
	q.Add("app_id", appId)
	q.Add("secret", secret)
	q.Add("code", code)
	q.Add("grant_type", "authorization_code")
	q.Add("redirect_uri", redirectUri)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer dclose(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	// todo: save the token and refresh token
	var tokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	err := json.Unmarshal(body, &tokens)

	if err != nil {
		log.Fatal("Failed to parse json")
	}

	c.Status(200)
}
