package index

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func readFiles(fileName []string) []string {
	docList := make([]string, 0)
	for _, sourceName := range fileName {
		docs := readFile(sourceName)
		if docs != nil && len(docs) > 0 {
			docList = append(docList, docs...)
		}
	}
	return docList
}

// 可改用流读取
func readFile(fileName string) []string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	docList := strings.Split(string(content), "\n")
	if len(docList) == 0 {
		log.Infof("readFile err: %v", "docList is empty\n")
		return nil
	}
	return docList
}
