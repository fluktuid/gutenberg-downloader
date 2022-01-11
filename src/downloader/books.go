package downloader

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

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
		getLinks(resp.Body, ch)
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
