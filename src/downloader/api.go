package downloader

type Filetype string

type Book string

const (
	FT_HTML          = Filetype("html")
	FT_TXT           = Filetype("txt")
	FT_EPUB          = FT_EPUB_IMG
	FT_EPUB_IMG      = Filetype("epub.images")
	FT_EPUB_NO_IMG   = Filetype("epub.noimages")
	FT_KINDLE        = FT_KINDLE_IMG
	FT_KINDLE_IMG    = Filetype("kindle.images")
	FT_KINDLE_NO_IMG = Filetype("kindle.noimages")
	FT_MP3           = Filetype("mp3")

	BASE_URL     = "https://www.gutenberg.org"
	BASE_PATH    = BASE_ROBOT + BASE_HARVEST
	BASE_ROBOT   = "/robot"
	BASE_HARVEST = "/harvest"
)
