package utils

import (
	"fmt"
	"github.com/alitto/pond"
	"io"
	"net/http"
	"os"
)

func DownloadMaps(collectionData Collection, mirror string) {
	pool := pond.New(8, 0, pond.MinWorkers(8))

	if err := os.Mkdir(fmt.Sprint(collectionData.Id), 0755); err != nil {
		panic(err)
	}

	for _, beatmap := range collectionData.Beatmapsets {
		beatmap := beatmap
		pool.Submit(func() {
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
		})
	}
	pool.StopAndWait()
	fmt.Println("Finished downloading collection:", collectionData.Name)
}
