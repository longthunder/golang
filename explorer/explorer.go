// explorer
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Folder struct {
	Name string
	Path string
	Time string
	Size string
}

var data []Folder
var dirs = []string{"D:\\Quake III Arena\\"}

func main() {

}

func scanDir() {
	data = make([]Folder, 0, 0)
	for _, dir := range dirs {
		infos, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println("Read dir error", dir)
			continue
		}
		for _, info := range infos {
			if info.IsDir() {
				data = append(data, scan(info.Name(), dir+info.Name()))
			} else {
				data = append(data, Folder{Name: info.Name()})
			}
		}
	}
}

func scan(name, path string) Folder {
	var time = time.Now()
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Println(filepath.Dir(path), path, "---", info.Name(), info.Size())
			info.ModTime()
		}
		return nil
	})
	return Folder{name}
}
