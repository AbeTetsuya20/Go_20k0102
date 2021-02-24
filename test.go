package main

import "fmt"

func main(){
	fmt.Println("テストの開始します")

	defer func(){
		fmt.Println("遅延実行の開始")

		if err := recover();err != nil{
			fmt.Println("パニックが起こっています err:",err)
		}

		fmt.Println("遅延実行の終了")
	}()

	arr := [...]int{1,2,3}

	index := 10

	fmt.Println(arr[index])
	fmt.Println("パニックが起きました。")

}