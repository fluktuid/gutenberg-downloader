package downloader

import (
	"io"
	"net/http"
	"os"
)

// DownloadCatalog will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadCatalog() (string, error) {
	catalogDir, _ := os.MkdirTemp("", "catalog")
	// Get the data
	csv, err := downloadTar(catalogDir)

	return csv, err
}
func downloadTar(dir string) (string, error) {
	// Get the data
	resp, err := http.Get(BASE_URL + CATALOG_PATH + CATALOG_FILE)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	out, err := os.CreateTemp(dir, "catalog*.csv")
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return out.Name(), err
}
