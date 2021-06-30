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

	// 移動元ディレクトリを指定
	source_dir := "../../../Desktop/" + flag.Arg(0)

	// 移動先ディレクトリを指定
	destination_dir := "../../../Desktop/" + flag.Arg(1)

	var files []os.FileInfo = getFiles(source_dir)
	moveRawFiles(files, getSelectedFileName(files), source_dir, destination_dir)
}

// 移動元ディレクトリの中にあるファイルを全て取得
func getFiles(source_dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(source_dir)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

// 拡張子が".JPG"のファイル名を取り出し、拡張子部分を".ARW"に置き換えてファイル名をスライスに格納
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

// 移動元ディレクトリのファイルの中で、ファイル名が、スライス内の要素と一致するファイルを移動先ディレクトリに移動
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
