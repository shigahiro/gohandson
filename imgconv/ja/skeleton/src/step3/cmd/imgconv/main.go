package main

import (
	"bufio"
	"fmt"
	"os"
)

func run() error {

	if len(os.Args) < 3 {
		return fmt.Errorf("引数が足りません。")
	}

	src, dst := os.Args[1], os.Args[2]

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした。%s", src)
	}
	// TODO: 関数終了時にファイルを閉じる

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ファイルを書き出せませんでした。%s", dst)
	}
	// TODO: 関数終了時にファイルを閉じる

	scanner := bufio.NewScanner(sf)
	// TODO: sfから1行ずつ読み込み、"行数:"を前に付けてdfに書き出す。
	for i := 1; scanner.Scan(); i++ {
		fmt.Fprintf(df, "%d:%s\n", i, scanner.Text())
	}

	// TODO: scannerから得られたエラーを返す
}

func main() {
	if err := run(); err != nil {
		// TODO: 標準エラー出力（os.Stderr）にエラーを出力する
		os.Exit(1)
	}
}
