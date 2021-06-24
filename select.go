package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	source_dir := "../../../Desktop/" + flag.Arg(0)
	destination_dir := "../../../Desktop/" + flag.Arg(1)

	var files []os.FileInfo = getFiles(source_dir)
	moveRawFiles(files, getSelectedFileName(files), source_dir, destination_dir)
}

func getFiles(source_dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(source_dir)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func getSelectedFileName(files []os.FileInfo) []string {
	var selected_file_names []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".JPG" {
			selected_file_name := strings.Replace(file.Name(), ".JPG", ".ARW", -1)
			selected_file_names = append(selected_file_names, selected_file_name)
		}
	}
	return selected_file_names
}

func moveRawFiles(files []os.FileInfo, selected_file_names []string, source_dir string, destination_dir string) {
	for _, file := range files {
		for _, selected_file_name := range selected_file_names {
			if file.Name() == selected_file_name {
				if err := os.Rename(source_dir+"/"+file.Name(), destination_dir+"/"+file.Name()); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
