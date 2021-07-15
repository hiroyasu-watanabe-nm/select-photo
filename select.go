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

// このプログラムはDeskTop下にあるフォルダの.pngファイルをARWに変換し、移動するプログラム
// NOTE: go run main.go DeskTop下にあるフォルダ名 DeskTop下にあるフォルダ名
// これだと SelectPhotoっていう振る舞いではなくって、moveFolderInDesktopになる。
// よってこれは夏目くんのやりたいことなの？
func main() {
	flag.Parse()

	// 移動元ディレクトリを引数から取得 存在を確認 ないのであればエラー
	source_dir := flag.Arg(0)
	
	// 移動先ディレクトリを引数から取得 存在を確認　ないのであればエラー
	destination_dir := flag.Arg(1)
	
	// 移動したい拡張子を引数から取得 ないのであればエラー
	extension := flag.Arg(2)
	if len(extension) <= 0 {
		panic("拡張子が指定されていません。")
	}
	
	// 移動元のディレクトリに配置されている全ファイルを取得 エラーがあれば処理終了
	files, err = getFiles(source_dir)
	if err != nil{		
		log.Fatal(err)
		panic(err)
	}	
	// 移動元のディレクトリに配置されているARW拡張子をもつファイルと同じ名前のJPG拡張子のファイルだけを抽出
	targetFiles := getTargetFiles(files)
	
	// 移動元のディレクトリにあるファイルを全部移動先のディレクトリに移動する
	moveFiles(targetFiles, source_dir, destination_dir)
}

// 移動元ディレクトリの中にあるファイルを全て取得
func getFiles(dir string) ([]os.FileInfo, error) {
	// 引数のディレクトリ配下に配置されているファイルを全取得して返す
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// 拡張子が".JPG"のファイルを取り出し、拡張子部分を".ARW"に置き換えてファイル名をスライスに格納
func getSelectedFileName(files []os.FileInfo) ([]os.FileInfo) {
	var ret []string
	var ARWFileMap map[string]os.FileInfo
	var JPGFileMap map[string]os.FileInfo
	// 引数のファイル一覧をループして、ファイル名の拡張子がARWのものを抜き出す
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ARW" {
			ARWFileMap[file.Name()] = file
		}else if filepath.Ext(file.Name()) == ".JPG"{
			JPGFileMap[file.Name()] = file
		}
	}
	for k,v := range ARWFileMap{
		_, isOK := JPGFileMap[k]
		if !isOK {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}

// 移動元ディレクトリの全ファイルの中で、ファイル名が、スライス内の要素と一致するファイルを移動先ディレクトリに移動
func moveFiles(files []os.FileInfo, source_dir string, destination_dir string) {
	for _, file := range files {
		if err := os.Rename(source_dir+"/"+file.Name(), destination_dir+"/"+file.Name()); err != nil {
			fmt.Println(err)
		}
	}
}
