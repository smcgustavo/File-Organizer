package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func path(argument string) string {
	if argument != "" {
		return argument
	}
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return currentPath
}

func allFilesList(directory string) []fs.DirEntry {
	dir, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	files, err := dir.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}
	dir.Close()
	return files
}

func fileNamesList(directory string) []fs.DirEntry {
	list := allFilesList(directory)
	var newList []fs.DirEntry
	for _, item := range list {
		if !item.IsDir() {
			newList = append(newList, item)
		}

	}
	return newList
}

func extractStrings(files []fs.DirEntry) []string {
	var list []string
	for _, file := range files {
		list = append(list, file.Name())
	}
	return list
}

func removeDelimeters(list []string) []string {
	delimeters := []string{"_", "-", "."}
	for _, delimeter := range delimeters {
		for idx, item := range list {
			list[idx] = strings.ReplaceAll(item, delimeter, " ")
		}
	}
	return list
}

func extractTags(list []string) []string {
	var tags []string
	for _, item := range list {
		tags = append(tags, strings.Split(item, " ")...)
	}
	for _, tag := range tags {
		if tag == "" {

		}
	}
	return tags
}

func countTags(list []string) map[string]int {
	tagsCount := make(map[string]int)
	for _, word := range list {
		tagsCount[word] += 1
	}
	return tagsCount
}

func getTag(fileName string, tags map[string]int) string {
	highestValue := 0
	selectedTag := ""
	for key := range tags {
		if strings.Contains(fileName, key) && tags[key] > highestValue {
			selectedTag = key
			highestValue = tags[key]
		}
	}
	return selectedTag
}

func existFolder(folderName string, path string) bool {
	info, err := os.Stat(path + "/" + folderName)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

func makeFolder(folderName string, path string) {
	if existFolder(folderName, path) {
		return
	}
	err := os.Mkdir(path+"/"+folderName, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func clusterFiles(files []fs.DirEntry, tagsMap map[string]int, path string) {
	for _, item := range files {
		choosedTag := getTag(item.Name(), tagsMap)
		makeFolder(choosedTag, path)
		folderPath := path + "/" + choosedTag + "/" + item.Name()
		os.Rename(path+"/"+item.Name(), folderPath)
	}
}

func main() {
	directory := path("/home/gustavosmc/Downloads")
	files := fileNamesList(directory)
	fmt.Println(files)
	list := extractStrings(files)
	fmt.Println(list)
	list_2 := removeDelimeters(list)
	fmt.Println(list_2)
	list_3 := extractTags(list_2)
	tagMap := countTags(list_3)
	fmt.Println(tagMap)
	clusterFiles(files, tagMap, "/home/gustavosmc/Downloads")
}

/*
	func printAllFileNames(files []fs.DirEntry) {
		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
	    }
		}
	}
*/
