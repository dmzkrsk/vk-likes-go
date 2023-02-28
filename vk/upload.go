package vk

import (
	"errors"
	"net/url"
	"strconv"
)

// PhotosGetUploadServer return an URL of upload server
func (api *VkApi) PhotosGetUploadServer(owner, album int) (string, error) {
	var upload struct {
		UploadURL string `json:"upload_url"`
		OwnerID   int    `json:"owner_id"`
		AlbumID   int    `json:"album_id"`
	}

	var parameters = url.Values{}

	parameters.Set("album_id", strconv.Itoa(album))
	parameters.Set("group_id", strconv.Itoa(owner))

	err := api.method(&upload, "photos.getUploadServer", parameters, true)
	if err != nil {
		return "", err
	}

	if upload.OwnerID != owner {
		return "", errors.New("Invalid owner")
	}
	if upload.AlbumID != album {
		return "", errors.New("Invalid album")
	}

	return upload.UploadURL, nil
}
