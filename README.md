# select-photo

## about
- ".JPG"で写真を選別し、残ったJPEGファイルと同じファイル名で拡張子が".RAW"のファイルを別ディレクトリに移動するプログラム
- 実行ファイルの階層 ~/○○/○○/○○/select.go
- 移動元ディレクトリの階層 ~/Desktop/移動元ディレクトリ名
- 移動先ディレクトリの階層 ~/Desktop/移動先ディレクトリ名

## run
- go run select.go 移動元ディレクトリ名 移動先ディレクトリ名

## example
【移動元】C0001.JPG, C0001.ARW, C0002.ARW, C0003.JPG, C0003.ARW, C0004.JPG, C0004.ARW, C0005.ARW

→

【移動元】C0001.JPG, C0002.ARW, C0003.JPG, C0004.JPG, C0005.ARW
【移動先】C0001.ARW, C0003.ARW, C0004.ARW
