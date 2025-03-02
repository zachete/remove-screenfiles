package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getPreparedDirList(dirList []string) []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Can't get home dir")
	}

	finalList := make([]string, 0, 2)
	for _, dir := range dirList {
		finalDirPath := filepath.Join(homeDir, dir)
		finalList = append(finalList, finalDirPath)
	}
	
	return finalList
}

func getTargetFilesPath(targetDirPath string) []string {
	files, err := os.ReadDir(targetDirPath)
	if err != nil {
		log.Fatalln("Can't read dir");
	}

	targetFiles := make([]string, 0, 5)
	for _, file := range files {
		fileName := file.Name()
		if !strings.HasPrefix(fileName, "Screen") {
			continue
		} 

		fullRelPath := filepath.Join(targetDirPath, file.Name())
		fullAbsPath, err := filepath.Abs(fullRelPath)
		if err != nil {
			log.Fatalln("Can't get absolute path")
		}
		targetFiles = append(targetFiles, fullAbsPath)
	}

	return targetFiles
}

func main() {
	targetDirList := []string{"/Documents", "/Desktop"}
	preparedDirList := getPreparedDirList(targetDirList)
	for _, dirPath := range preparedDirList {
		targetFiles := getTargetFilesPath(dirPath)
		total := 0
		for _, path := range(targetFiles) {
			os.Remove(path)
			total++
			log.Printf("Removed %v\n", path)
		}
		log.Printf("Total %v files removed inside %v", total, dirPath)
	}

}