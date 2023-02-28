package vk

import (
	"fmt"
	"io"
	"log"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func (api *VkApi) login() error {
	urlLogin, err := url.Parse(URL_LOGIN)
	if err != nil {
		return newError("Can't parse URL")
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return newError("Can't create cookies")
	}
	api._client.Jar = jar

	var values = url.Values{}

	values.Set("act", "login")
	values.Set("to", "")
	values.Set("expire", "")
	values.Set("_origin", "https://vk.com")
	values.Set("ip_h", "https://vk.com")
	values.Set("lg_h", "https://vk.com")
	values.Set("email", api.Login)
	values.Set("pass", api.password)

	response, err := api._client.PostForm(urlLogin.String(), values)
	if err != nil {
		return newError(err.Error())
	}

	endUrl := response.Request.URL.String()
	log.Printf("endUrl: %s", endUrl)
	defer response.Body.Close()
	out, err := os.Create("filename.ext")
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(out, response.Body)

	var cookies = api._client.CookieMap(urlLogin)

	for k, v := range cookies {
		log.Printf("%s => %s\n", k, v)
	}

	if remixSid, ok := cookies["remixsid"]; ok {
		api._sid = remixSid
		api._login_p = cookies["p"]
		api._login_l = cookies["l"]

		if strings.Contains(endUrl, "act=blocked") {
			return newAuthError("Account is blocked")
		} else if strings.Contains(endUrl, "security_check") {
			return newAuthError("Security check is required")
		}

		return nil
	} else if strings.Contains(endUrl, "sid=") {
		return newAuthError("Authorization error (capcha)")
	} else if strings.Contains(endUrl, "m=") {
		return newAuthError("Bad password")
	}

	return newAuthError(fmt.Sprintf("Unknown error (%s)", endUrl))
}
