package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "os/exec"
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

func mkdir(rootPtr *Folder) string {
	if len(rootPtr.Folders) == 0 {
		return rootPtr.Name
	}

	mkdirString := fmt.Sprintf("%s/{", rootPtr.Name)

	for _, folder := range rootPtr.Folders {
		mkdirString += fmt.Sprintf("%s,", mkdir(&folder))
	}

	return fmt.Sprintf("%s}", mkdirString)
}

func touch(rootPtr *Folder) string {
	if len(rootPtr.Files) == 0 && len(rootPtr.Folders) == 0 {
		return ""
	}

	touchString := fmt.Sprintf("%s/{", rootPtr.Name)

	for _, file := range rootPtr.Files {
		touchString += fmt.Sprintf("%s,", file.Name)
	}

	for _, folder := range rootPtr.Folders {
		touchString += fmt.Sprintf("%s,", touch(&folder))
	}

	return fmt.Sprintf("%s}", touchString)
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

	// Trim trailing `/` on directory name, this is added when forming the mkdir command later
	directory := strings.TrimRight(*directoryPtr, "/")
	rootPtr := &Folder{directory, []File{}, []Folder{}}

	traverseDirectory(rootPtr, contents)

	mkdirCmd := mkdir(rootPtr)
	touchCmd := touch(rootPtr)

	fmt.Printf("mkdir -p %s; touch %s\n", mkdirCmd, touchCmd)

	// err := exec.Command("mkdir", "-p", mkdirCmd, "; ", "touch", touchCmd).Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
