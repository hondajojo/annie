package extractors

import (
	"github.com/hondajojo/annie/downloader"
	"github.com/hondajojo/annie/parser"
	"github.com/hondajojo/annie/request"
	"github.com/hondajojo/annie/utils"
)

// Pixivision download function
func Pixivision(url string) downloader.VideoData {
	html := request.Get(url)
	title, urls := parser.GetImages(url, html, "am__work__illust  ", nil)
	data := downloader.VideoData{
		Site:  "pixivision pixivision.net",
		Title: utils.FileName(title),
		Type:  "image",
		URLs:  urls,
		Size:  0,
	}
	data.Download(url)
	return data
}
