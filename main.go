package main

import (
	"./vk"
	"log"
	"os"
	"sort"
	"strconv"
)

type ByLikes []vk.Photo

func (a ByLikes) Len() int           { return len(a) }
func (a ByLikes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLikes) Less(i, j int) bool { return a[i].LikesCount() > a[j].LikesCount() }

func main() {
	api := vk.NewApi(os.Args[1], os.Args[2], "photos")

	aid, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln(err)
	}

	albums, err := api.PhotosGetAlbums(aid)
	if err != nil {
		log.Fatalln(err)
	}

	for _, album := range albums {
		photos, err := api.GetPhotos(album)

		if err != nil {
			log.Fatalln(err)
		}
		
		if len(photos) == 0 {continue }

		log.Printf("% 4d %s", len(photos), album.Title)

		hasLikes := make([]vk.Photo, 0, len(photos))

		for _, photo := range photos {
			if photo.HasLikes() {
				hasLikes = append(hasLikes, photo)
			}
		}

		log.Printf("  Likes: % 4d", len(hasLikes))

		sort.Sort(ByLikes(hasLikes))

		for _, photo := range hasLikes[:5] {
			log.Printf("    % 4d https://vk.com/photo%d_%d", photo.LikesCount(), album.OwnerId, photo.Id)
		}
	}

	log.Println("OK!")
}