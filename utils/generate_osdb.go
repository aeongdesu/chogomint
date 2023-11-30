package utils

import (
	"fmt"
	"os"
)

func GenerateOsdb(collectionData Collection) {
	file, err := os.Create(fmt.Sprint(collectionData.Id, "/collection.osdb"))
	if err != nil { panic(err) }
    defer file.Close()
}