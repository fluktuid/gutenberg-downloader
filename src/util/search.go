package util

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func SearchInFile(path string, search string, outChan chan<- string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, search) {
			outChan <- text
		}
	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)
	}
	return err
}
