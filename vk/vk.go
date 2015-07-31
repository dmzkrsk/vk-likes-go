package vk
import (
	"net/http"
	"net/url"
	"path"
	"time"
)

func (e VkError) String() string {
	return e.message
}

const VERSION string = "5.33"
const URL_LOGIN string = `https://login.vk.com/`
const URL_API string = `https://api.vk.com/method/`
const API_VERSION string = `5.33`
const DEFAULT_PHOTO_COUNT int = 500

type VkApi struct {
	Login string
	password string

	scope string

	token string

	_client *Client

	_throttle <-chan time.Time

	_sid string
	_login_p string
	_login_l string
}

func NewApi(login string, password string, scope string) *VkApi {
	rate := time.Second / 3
	throttle := time.Tick(rate)

	vkApi := &VkApi{Login: login, password: password, scope: scope, _client: &Client{new(http.Client)}, _throttle: throttle}

	return vkApi
}

func (api *VkApi) Auth() error {
	if api.Login == "" {return newError("No login provided")}
	if api.password == "" {return newError("No password provided")}

	return api.login()
}

func (api *VkApi) method(result interface{}, method string, parameters url.Values, auth bool) error {
	url, err := url.Parse(URL_API)
	if err != nil {return newError("Can't parse URL")}

	url.Path = path.Join(url.Path, method)
	parameters.Set("v", API_VERSION)
	parameters.Set("lang", "ru")
	parameters.Set("https", "0")
	url.RawQuery = parameters.Encode()

//	log.Println(url.String())

	<- api._throttle
	err = api._client.GetJson(url.String(), &result)
	if err != nil {return newError(err.Error())}

	return nil
}