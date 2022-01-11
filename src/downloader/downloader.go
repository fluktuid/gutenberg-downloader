package downloader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(path string, uurl string) error {
	u, err := url.Parse(uurl)
	if err != nil {
		panic(err)
	}
	// Get the data
	resp, err := http.Get(uurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	err = os.MkdirAll(path+filepath.Dir(u.Path), 0755)
	if err != nil {
		return err
	}
	out, err := os.Create(path + u.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
