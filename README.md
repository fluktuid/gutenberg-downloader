
# Gutenberg Downloader

A brief description of what this project does and who it's for

Add badges from somewhere like: [shields.io](https://shields.io/)

[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)


## Usage

#### download books

Download all english books as epubs with images from gutenberg-downloader.

```bash
gutenberg-downloader download -l en -f epub.image
```


| Flag              | default       | Description                                   | Example       |
| :---------------- | :------------ | :-------------------------------------------- | :------------ |
| `-l, --languages` | `en`          | 2-char the language code of downloaded books  | de,en         |
| `-f, --filetypes` | `txt`         | filetypes for downloaded books<br>(txt, html, epub.image, epub.noimages, kindle.images, kindle.noimages, mp3)| epub.images |
| `-o, --outfolder` | `./downloads` | Folder for storing downloaded files           | "~/Downloads" |
| `-p, --parallel`  | `10`          | max amount of parallel downloads              | 15            |

#### Search Books

```bash
gutenberg-downloader search [arg]
```

| Parameter | Type     | Description                             |
| :-------- | :------- | :-------------------------------------- |
| `arg`     | `string` | **Required**. String to search in books |

## Installation


### custom build
```bash
# build
git clone <repo>
go build -o gutenberg-downloader .
chmod +x gutenberg-downloader
# use
./gutenberg-downloader --help
```

### Linux
```bash
# download release
SYS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=(UNAME -m)
curl -o gutenberg-downloader "https://github.com/fluktuid/gutenberg-downloader/releases/download/v0.2.0/gutenberg-downloader_$(SYS)_$(ARCH)"
chmod +x gutenberg-downloader
sudo mv gutenberg-downloader /usr/local/bin/
# use
./gutenberg-downloader --help
```


### MacOS

```bash
brew tap fluktuid/tap
brew install gutenberg-downloader
```

## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## Feedback

If you have any feedback, please reach out to me and create an issue.

