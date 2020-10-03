package cleaner

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileType bool

func (fileType FileType) isDirectory() bool {
	return fileType == Directory
}

func (fileType FileType) matches(b bool) bool {
	if b {
		return fileType == Directory
	}
	return fileType == File
}

const (
	Directory FileType = true
	File               = false
)

type Cleaner struct {
	Root     string
	Confirm  bool
	Names    []string
	FileType FileType
}

func (config *Cleaner) Start() {
	confirm := config.Confirm

	filepath.Walk(config.Root, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path)
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}
		//fmt.Println(path)
		if config.FileType.matches(info.IsDir()) && config.hasName(info.Name()) {
			delete := !confirm

			if confirm {
				fmt.Printf("Delete folder %s? [y/n]: ", path)

				reader := bufio.NewReader(os.Stdin)
				response, err := reader.ReadString('\n')
				if err != nil {
					log.Fatal(err)
				}
				response = strings.ToLower(strings.TrimSpace(response))

				if response == "y" || response == "yes" {
					delete = true
				}
			}

			if delete {
				//err := os.RemoveAll(path)
				//if err != nil {
				//	fmt.Println(fmt.Sprintf("Failed to delete %s", path), err)
				//}
				fmt.Println(fmt.Sprintf("Deleted %s", path))
				return filepath.SkipDir
			}

		}
		return nil
	})
}

func (config *Cleaner) hasName(basename string) bool {
	for _, name := range config.Names {
		if name == basename {
			return true
		}
	}
	return false
}
