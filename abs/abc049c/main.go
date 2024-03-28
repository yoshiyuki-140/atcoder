package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	words := [4]string{"dream", "dreamer", "erase", "eraser"}

	// fmt.Scan(&S)
	S = "dreamerase"

	for s := S; len(s) != 0; {
		for i, word := range words {
			// words内にある文字列が含まれているか否かを判定する
			if s := cutWord(s, word); s != "" {
				// 含まれていた場合
				S = s
				fmt.Println("S :", S, "i :", i)
				break
			}
			if i == len(words)-1 {
				break
			}
		}
		// 含まれていなかった場合
		fmt.Println("NO")
		break
	}
	fmt.Println(cutWord(S, "dreamere"))
	fmt.Println("YES")
}

func cutWord(str, targetWord string) string {
	// 単語を検索して切り出す
	index := strings.Index(str, targetWord)
	if index == -1 {
		// fmt.Println("単語が見つかりませんでした。")
		return ""
	}

	wordLength := len(targetWord)
	// 切り出し
	return str[:index] + str[index+wordLength:]
}
