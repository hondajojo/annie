package utils

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"path"

	"github.com/hondajojo/annie/request"
)

// Match1 return result of first match
func Match1(pattern, text string) []string {
	re := regexp.MustCompile(pattern)
	value := re.FindStringSubmatch(text)
	return value
}

// MatchAll return all matching results
func MatchAll(pattern, text string) [][]string {
	re := regexp.MustCompile(pattern)
	value := re.FindAllStringSubmatch(text, -1)
	return value
}

// FileSize return the file size of the specified path file
func FileSize(filePath string) int64 {
	file, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return 0
	}
	return file.Size()
}

// Domain get the domain of given URL
func Domain(url string) string {
	domainPattern := `([a-z0-9][-a-z0-9]{0,62})\.` +
		`(com\.cn|com\.hk|` +
		`cn|com|net|edu|gov|biz|org|info|pro|name|xxx|xyz|` +
		`me|top|cc|tv|tt)`
	domain := Match1(domainPattern, url)[1]
	return domain
}

// FileName Converts a string to a valid filename
func FileName(name string) string {
	// FIXME(hondajojo) file name can't have /
	newName := strings.Replace(name, "/", " ", -1)
	newName = strings.Replace(newName, ":", "：", -1)
	return newName
}

// FilePath gen valid filename
func FilePath(base, name, ext string, escape bool) string {
	fileName := fmt.Sprintf("%s.%s", name, ext)
	if escape {
		fileName = FileName(fileName)
	}
	return path.Join(base, fileName)
}

// StringInSlice if a string is in the list
func StringInSlice(str string, list []string) bool {
	for _, a := range list {
		if a == str {
			return true
		}
	}
	return false
}

// GetNameAndExt return the name and ext of the URL
// https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg ->
// 1f5a87801a0711e898b12b640777720f, jpg
func GetNameAndExt(uri string) (string, string) {
	u, _ := url.ParseRequestURI(uri)
	s := strings.Split(u.Path, "/")
	filename := strings.Split(s[len(s)-1], ".")
	if len(filename) > 1 {
		return filename[0], filename[1]
	} else {
		// Image url like this
		// https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg/w650
		// has no suffix
		contentType := request.ContentType(uri, uri)
		return filename[0], strings.Split(contentType, "/")[1]
	}
}
