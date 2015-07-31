package main

import (
	"./vk"
	"log"
	"os"
)

func main() {
	api := vk.NewApi(os.Args[0], os.Args[1], "photos")

	albums, err := api.PhotosGetAlbums(-94709215)
	if err != nil {
		log.Fatalln(err)
	}

	for _, album := range albums {
		photos, err := api.GetPhotos(album)

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("% 4d %s", len(photos), album.Title)

		hasLikes := make([]vk.Photo, 0, len(photos))

		for _, photo := range photos {
			if photo.HasLikes() {
				hasLikes = append(hasLikes, photo)
			}
		}

		log.Printf("  Likes: % 4d", len(hasLikes))
	}

	log.Println("OK!")
}