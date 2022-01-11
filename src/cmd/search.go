package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fluktuid/gutenberg-downloader/src/downloader"
	"github.com/fluktuid/gutenberg-downloader/src/util"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "search",
	Short: "search some books",
	Long:  `Search some books`,
	Run:   runList,
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func runList(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Printf("%s\n", "need a search string")
		return
	}
	catalogFile, err := downloader.DownloadCatalog()
	if err != nil {
		panic(err)
	}
	elemChan := make(chan []string)
	go func() {
		file, err := os.Open(catalogFile)
		if err != nil {
			close(elemChan)
			return
		}
		defer file.Close()
		util.GetCSVRecords(file, elemChan)

		close(elemChan)
		defer file.Close()
	}()

	for elem := range elemChan {
		for _, v := range elem {
			if strings.Contains(v, args[0]) {
				fmt.Printf("-----\nID: %s\n%s (%s)\n  - %s\n", elem[0], elem[3], elem[4], elem[5])
			}
		}
	}

	// todo: delete catalog file
}
