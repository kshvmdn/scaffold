package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
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

type Path struct {
	FilePath    string
	IsDirectory bool
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

func recursivelyFormPathList(dirPtr *Folder, parent string) []Path {
	if parent != "" {
		parent = fmt.Sprintf("%s/%s", parent, dirPtr.Name)
	} else {
		parent = dirPtr.Name
	}

	parentPath := Path{parent, true}

	paths := []Path{parentPath}

	for _, file := range dirPtr.Files {
		paths = append(paths, Path{
			fmt.Sprintf("%s/%s", parent, file.Name),
			false,
		})
	}

	for _, folder := range dirPtr.Folders {
		paths = append(paths, recursivelyFormPathList(&folder, parent)...)
	}

	return paths
}

func main() {
	templatePtr := flag.String("template", "", "Template file (directory structure)")
	directoryPtr := flag.String("directory", ".", "Root directory")
	flag.Parse()

	if *templatePtr == "" {
		log.Fatal("Missing -template argument")
	}

	if _, err := os.Stat(*templatePtr); os.IsNotExist(err) {
		log.Fatal("Template file doesn't exist")
	}

	usr, _ := user.Current()
	contents := parseFileContents(*templatePtr)
	directory := *directoryPtr

	fmt.Println(usr.HomeDir)

	// Trim trailing `/` on directory name, this is added when forming the mkdir command later
	directory = strings.TrimRight(directory, "/")
	// Replace ~/ with absolute path to home directory
	directory = strings.Replace(directory, "~/", fmt.Sprintf("%s/", usr.HomeDir), 1)

	rootPtr := &Folder{directory, []File{}, []Folder{}}

	traverseDirectory(rootPtr, contents)

	paths := recursivelyFormPathList(rootPtr, "")

	for _, path := range paths {
		fmt.Printf("Creating %s...", path.FilePath)

		if path.IsDirectory {
			err := os.MkdirAll(path.FilePath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err := os.Create(path.FilePath)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("done.")
	}
}
