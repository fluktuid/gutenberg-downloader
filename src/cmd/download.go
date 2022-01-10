package cmd

import (
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fluktuid/gutenberg-downloader/src/downloader"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download some books",
	Long:  `Download some books`,
	Run:   run,
}

var languages []string
var filetypes []string
var parallel int
var folder string

func init() {
	downloadCmd.PersistentFlags().StringSliceVarP(&languages, "languages", "l", []string{"en"}, "Languages in which books are searched for")
	downloadCmd.PersistentFlags().StringSliceVarP(&filetypes, "filetypes", "f", []string{"txt"}, "filetypes in which books are downloaded")
	downloadCmd.PersistentFlags().IntVarP(&parallel, "parallel", "p", 10, "max amount of go-routines used while downloading (max=100)")
	downloadCmd.PersistentFlags().StringVarP(&folder, "outfolder", "o", "./downloads", "folder used for storing files")
	logrus.SetLevel(logrus.DebugLevel)
}

func run(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	var cntDownloading int32
	downloadCh := make(chan string, 10)
	for i := 0; i < parallel && i < 100; i++ {
		go func(path string, ch <-chan string, wg sync.WaitGroup) {
			defer wg.Done()
			for elem := range ch {
				err := downloader.DownloadFile(path, elem)
				if err != nil {
					logrus.Error(err)
					continue
				}
				atomic.AddInt32(&cntDownloading, 1)
			}
		}(folder, downloadCh, wg)
	}
	err := downloader.GetBooksLinks(filetypes, languages, downloadCh)
	close(downloadCh)
	if err != nil {
		panic(err)
	}
	wg.Wait()
}
