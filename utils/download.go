package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadMaps(collectionData Collection, mirror string) {
	var wg sync.WaitGroup

	if err := os.Mkdir(fmt.Sprint(collectionData.Id), 0755); err != nil {
		panic(err)
	}

	for _, beatmap := range collectionData.Beatmapsets {
		wg.Add(1)
		go func(beatmap Beatmaps) {
			defer wg.Done()
			fmt.Println("Downloading beatmapset:", beatmap.Id)
			url := fmt.Sprintf("%s%d", mirror, beatmap.Id)

			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			file, err := os.Create(fmt.Sprintf("%d/%d.osz", collectionData.Id, beatmap.Id))
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if _, err = io.Copy(file, res.Body); err != nil {
				panic(err)
			}
			fmt.Println("Downloaded beatmapset:", beatmap.Id)
		}(beatmap)
	}
	wg.Wait()
}
