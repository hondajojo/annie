package extractors

import (
	"github.com/hondajojo/annie/downloader"
	"github.com/hondajojo/annie/parser"
	"github.com/hondajojo/annie/request"
	"github.com/hondajojo/annie/utils"
)

// Bcy download function
func Bcy(url string) downloader.VideoData {
	html := request.Get(url)
	title, urls := parser.GetImages(
		url, html, "detail_std detail_clickable", func(u string) string {
			// https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg/w650
			return u[:len(u)-5]
		},
	)

	data := downloader.VideoData{
		Site:  "半次元 bcy.net",
		Title: utils.FileName(title),
		Type:  "image",
		URLs:  urls,
		Size:  0,
	}
	data.Download(url)
	return data
}
