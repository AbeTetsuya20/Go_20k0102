package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("err.txt")
	if err != nil {
		fmt.Println("ファイルのオープンに失敗しました。")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	file.Close()

	fmt.Println("ファイルのオープンに成功しました。")

}
