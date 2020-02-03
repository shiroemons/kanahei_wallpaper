package main

import (
	"strconv"
	"time"

	"github.com/reujab/wallpaper"
)

const (
	stdZeroMonth = "01"
)

func main() {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	month := t.Format(stdZeroMonth)
	url := "http://www.kanahei.com/upload/" + year + "/" + month + "/01_16x9.jpg"

	wallpaper.SetFromURL(url)
}
