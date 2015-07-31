package vk
import (
	"net/http"
	"encoding/json"
	"net/url"
	"io/ioutil"
)

type Client struct {
	*http.Client
}

func (this *Client) GetJson(url string, result interface{}) error {
	json_response, err := this.Get(url)
	if err != nil {return newError(err.Error())}

	var vkErrWrapper struct {
		Error VkApiError
	}

	response_bytes, err := ioutil.ReadAll(json_response.Body)
	if err != nil {return newError(err.Error())}

	json.Unmarshal(response_bytes, &vkErrWrapper)

	if vkErrWrapper.Error.Code > 0 {
		return vkErrWrapper.Error
	}

	var successResponse struct {
		Response json.RawMessage
	}

	err = json.Unmarshal(response_bytes, &successResponse)
	if err != nil {return newError(err.Error())}
	err = json.Unmarshal(successResponse.Response, result)
	if err != nil {return newError(err.Error())}

	return nil
}

func (this *Client) CookieMap(u *url.URL) map[string]string {
	cookies := make(map[string]string)

	for _, cookie := range this.Jar.Cookies(u) {
		cookies[cookie.Name] = cookie.Value
	}

	return cookies
}