package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCollectionData(id int) (collectionJson Collection) {
	// get collection data
	url := fmt.Sprintf("https://osucollector.com/api/collections/%d", id)
	resp, err := http.Get(url)
	if err != nil { panic(err) }
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(resp.StatusCode)
	}

	// json parsing
    err = json.NewDecoder(resp.Body).Decode(&collectionJson)
	if err != nil { panic(err) }
	return
}