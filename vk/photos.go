package vk
import (
	"net/url"
	"strconv"
	"fmt"
)

type Album struct {
	Id int
	OwnerId int `json:"owner_id"`
	Title string
	Description string
	Size int
}

type Photo struct {
	Id int
	OwnerId int `json:"owner_id"`
	AlbumId int `json:"album_id"`
	Date Timestamp

	Likes Count
	Comments Count
}

func (photo Photo) HasLikes() bool {
	return photo.Likes.Count > 0
}

func (api *VkApi) GetPhotos(album Album) ([]Photo, error) {
	var photos []Photo

	const pageSize = DEFAULT_PHOTO_COUNT

	var parameters = url.Values{}

	parameters.Set("owner_id", strconv.Itoa(album.OwnerId))
	parameters.Set("album_id", strconv.Itoa(album.Id))
	parameters.Set("extended", "1")

	parameters.Set("count", strconv.Itoa(pageSize))

	for {
		var vkArray struct {
			Count int
			Items []Photo
		};

		parameters.Set("offset", strconv.Itoa(len(photos)))

		err := api.method(&vkArray, "photos.get", parameters, false)
		if err != nil {return nil, err}

		photos = append(photos, vkArray.Items...)

		if len(vkArray.Items) < pageSize {
			if(vkArray.Count != len(photos)) {
				return nil, newError(fmt.Sprintf("Results merge error (expected %d, total %d)", vkArray.Count, len(photos)))
			}

			return photos, nil
		}
	}

	panic("This should not be accessed")
}

func (api *VkApi) PhotosGetAlbums(owner int) ([]Album, error) {
	var albums struct {
		Count int
		Items []Album
	}

	var parameters = url.Values{}

	parameters.Set("owner_id", strconv.Itoa(owner))

	err := api.method(&albums, "photos.getAlbums", parameters, false)
	if err != nil {return nil, err}

	return albums.Items, nil
}
