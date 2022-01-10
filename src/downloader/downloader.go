package downloader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/net/html"
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

func GetBooksLinks(filetypes []string, languages []string, responseCh chan<- string) error {
	link := BASE_URL + BASE_PATH
	offset := 0
	for offset >= 0 {
		_offset, err := getBooksPage(link, offset, filetypes, languages, responseCh)
		if err != nil || offset == -1 {
			return err
		}
		offset = _offset
	}
	return nil
}

func getBooksPage(link string, offset int, filetypes []string, languages []string, responseCh chan<- string) (int, error) {
	base, err := url.Parse(link)
	if err != nil {
		return -1, err
	}

	params := url.Values{}
	ft_param := strings.Join(filetypes, ",")
	params.Add("filetypes[]", ft_param)
	lang_param := strings.Join(languages, ",")
	params.Add("langs[]", lang_param)
	params.Add("offset", strconv.Itoa(offset))
	base.RawQuery = params.Encode()
	resp, err := http.Get(base.String())
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	ch := make(chan string)
	go func() {
		parse(resp.Body, ch)
		close(ch)
	}()

	for elem := range ch {
		if strings.Contains(elem, "harvest") {

			u, _ := url.Parse(elem)
			offset, _ := strconv.Atoi(u.Query().Get("offset"))
			return offset, nil
		}
		responseCh <- elem
	}

	return -1, nil
}

func parse(r io.Reader, ch chan string) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}
			ch <- url
		}
	}
}

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}
