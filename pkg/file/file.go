package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	set "github.com/deckarep/golang-set"
)

func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		/*
			if file.IsDir() {
				paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
				continue
			}
		*/
		name := file.Name()
		if name == "output" && file.IsDir() {
			continue
		}
		paths = append(paths, filepath.Join(dir, name))
	}

	return paths
}

func GetExtList(paths []string) []string {
	var extList []interface{}
	extList = append(extList, "other")
	for _, path := range paths {
		ext := GetExt(path)
		if ext == "" {
			continue
		}
		extList = append(extList, ext)
	}
	set := set.NewSetFromSlice(extList).String()
	return strings.Split(set[4:len(set)-1], ", ")
}

func GetExt(path string) string {
	ext := filepath.Ext(path)
	fInfo, _ := os.Stat(path)
	if ext == "" || fInfo.IsDir() {
		return ""
	}
	return ext[1:]
}

func Sort(paths []string, rootDir string, extList []string) {
	for _, e := range extList {
		if err := os.MkdirAll(filepath.Join(rootDir, "output", e), 0755); err != nil {
			fmt.Println(err)
		}
	}

	for _, path := range paths {
		//fmt.Println(path)
		ext := GetExt(path)
		if ext == "" {
			ext = "other"
		}
		filename := filepath.Base(path)
		move2Path := filepath.Join(rootDir, "output", ext, filename)
		//fmt.Println("=> " + move2Path)
		if err := os.Rename(path, move2Path); err != nil {
			fmt.Println(err)
		}
	}
}
