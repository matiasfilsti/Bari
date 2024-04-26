package readfilehelper

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func ReadLineOfFile(urlfilesites string) ([]string, string) {
	var fileLines []string
	file, err := os.Open(urlfilesites)
	if err != nil {

		//log.Fatal("Error opening Url's file")
		return fileLines, "error opening Url's file"
	}

	filescan := bufio.NewScanner(file)
	filescan.Split(bufio.ScanLines)
	for filescan.Scan() {
		urlValid, err := url.ParseRequestURI(filescan.Text())
		if err != nil {
			continue
		} else {
			fmt.Printf("Url valida: %v \n", urlValid)
			fileLines = append(fileLines, filescan.Text())
		}
	}
	file.Close()
	return fileLines, "nil"

}
