package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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
	return strings.Split(string(b), "\n")
}

func traverseDirectory(rootPtr *Folder, contents []string) {
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
			current.Folders = append(current.Folders, Folder{strings.TrimRight(element, "/"), []File{}, []Folder{}})
		} else {
			current.Files = append(current.Files, File{element})
		}
	}
}

func formMkdirString(rootPtr *Folder) string {
	if len(rootPtr.Folders) == 0 {
		return rootPtr.Name
	}

	// In the case that there is one subdir, we want to avoid adding braces since
	// mkdir treats these as part of the directory name.
	if len(rootPtr.Folders) == 1 {
		return fmt.Sprintf("%s/%s", rootPtr.Name, formMkdirString(&rootPtr.Folders[0]))
	}

	mkdirString := fmt.Sprintf("%s/{", rootPtr.Name)

	for _, folder := range rootPtr.Folders {
		mkdirString += fmt.Sprintf("%s,", formMkdirString(&folder))
	}

	return fmt.Sprintf("%s}", strings.TrimRight(mkdirString, ","))
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

	// Trim trailing `/` on directory name
	rootPtr := &Folder{strings.TrimRight(*directoryPtr, "/"), []File{}, []Folder{}}

	traverseDirectory(rootPtr, contents)

	mkdirString := formMkdirString(rootPtr)
	fmt.Println(mkdirString)
}
