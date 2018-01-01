package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func main() {
	//fmt.Println("Enter Folder Path:")
	//var searchDir string
	//fmt.Scanln(&searchDir)
	fmt.Println("Enter Image Width:")
	var imgWidth int
	fmt.Scanln(&imgWidth)
	fmt.Println("Enter Image Height:")
	var imgHeight int
	fmt.Scanln(&imgHeight)
	var fileInfo os.FileInfo
	searchDir, _ := os.Getwd()
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		fmt.Print(err)
	} else {
		for _, file := range fileList {
			fileInfo, err = os.Stat(file)
			if fileInfo.IsDir() {
				continue
			}
			fmt.Println(file)
			src, err := imaging.Open(file)
			if err != nil {
				continue
			}
			dstImage := imaging.Resize(src, imgWidth, imgHeight, imaging.Lanczos)
			var extension = filepath.Ext(file)
			var name = file[0 : len(file)-len(extension)]
			err = imaging.Save(dstImage, name+"_."+extension)
			if err != nil {
				fmt.Print("Save failed: %v", err)
			}
		}
	}
}
