package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 値の格納
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行読み込み
	N, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Printf("エラー%v\n", err)
		return
	}

	// matrix初期化
	adjacencyMatrix := make([][]bool, N)
	for i := 0; i < N; i++ {
		adjacencyMatrix[i] = make([]bool, N)
	}

	for i := 0; i < N; i++ {
		scanner.Scan()                           // 一行分読み込み
		tokens := strings.Fields(scanner.Text()) // 空白を削除して配列に格納
		for j, token := range tokens {
			if token == "1" {
				adjacencyMatrix[i][j] = true
			}
		}
	}

	// 中身確認
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if adjacencyMatrix[i][j] {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
			if j == N-1 {
				fmt.Println()
			}
		}
	}

	// メインの処理
	// 番号を格納する2重スライスを初期化
	data := make([][]int, N)
	for i := 0; i < N; i++ {
		data[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if adjacencyMatrix[i][j] {
				data[i] = append(data[i], j+1)
			}
		}
	}
	printMatrix(data, N)
}

func printMatrix(matrix [][]int, N int) {
	var i, j int
	// 行列の中身確認
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			if j == (N - 1) {
				fmt.Println(matrix[i][j])
			} else {
				fmt.Print(matrix[i][j])
			}
		}
	}
}
