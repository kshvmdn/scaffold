package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "path"
	"regexp"
	"strings"
	_ "sync"
)

type File struct {
	Name string
}

type Folder struct {
	Name    string
	Files   []File
	Folders []Folder
}

func parseFileContents(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	contents := string(b)
	return strings.Split(contents, "\n")
}

func walkDirectoryStructure(rootPtr *Folder, contents []string) {
	for _, element := range contents {
		if strings.TrimSpace(element) == "" {
			continue
		}

		isDirectory, err := regexp.MatchString("/$", element)
		if err != nil {
			log.Fatal(err)
		}

		level := len(regexp.MustCompile("(\\s)(\\s)").FindAllString(element, -1))

		current := rootPtr
		element = strings.TrimSpace(element)

		for i := 0; i < level; i++ {
			current = &current.Folders[len(current.Folders)-1]
		}

		if isDirectory {
			current.Folders = append(current.Folders, Folder{element, []File{}, []Folder{}})
		} else {
			current.Files = append(current.Files, File{element})
		}
	}
}

func main() {
	configPtr := flag.String("config", "", "Config file")
	directoryPtr := flag.String("directory", ".", "Root directory")
	flag.Parse()

	if *configPtr == "" {
		log.Fatal("Missing --config argument")
	}

	if _, err := os.Stat(*configPtr); os.IsNotExist(err) {
		log.Fatal("Config file doesn't exist")
	}

	contents := parseFileContents(*configPtr)

	rootPtr := &Folder{*directoryPtr, []File{}, []Folder{}}
	walkDirectoryStructure(rootPtr, contents)

	fmt.Println(*rootPtr)
}
