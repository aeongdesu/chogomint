package main

import (
	"flag"
	"fmt"
	"github.com/aeongdesu/chogomint/utils"
)

func main() {
	// get flag
	id := flag.Int("id", 0, "osu!collector id")
	// osdb := flag.Bool("g", false, "generates osdb file")
	mirror := flag.String("m", "https://api.nerinyan.moe/d/", "beatmapset download url")
	flag.Parse()

	if flag.Args() == nil || *id == 0 {
		flag.PrintDefaults()
		return
	}

	collectionData := utils.GetCollectionData(*id)

	fmt.Println(collectionData.Name)
	
	// if *osdb != nil { utils.GenerateOsdb(collectionData) }

	utils.DownloadMaps(collectionData, *mirror)
}