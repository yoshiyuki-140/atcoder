package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s [file]", os.Args[0])
	}

	fname := os.Args[1]
	fset := token.NewFileSet()
	pf, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)

	if err != nil {
		log.Fatalf("could not parse %s: %v", fname, err)
	}

	pf.Comments = nil

	cmap := ast.NewCommentMap(fset, pf, pf.Comments)

	for node, _ := range cmap {
		cmap[node] = nil
	}

	if err := printer.Fprint(os.Stdout, fset, pf); err != nil {
		log.Fatalf("could not print %s: %v", fname, err)
	}
}

// このコードを`remove_comments.go`として保存し、以下のように実行します。

// go run remove_comments.go [your file.go]

// これで`[your file.go]`からコメントが削除されたコードが出力されます。ただし、`[your file.go]`は実際のコードファイルのパスに置き換えてください。このコードは全てのコメントを削除します。これにより、`/* ... */`スタイルと`//`スタイルの両方のコメントが削除されます。

// なお、この方法はすべてのコメントを削除しますが一部のコメント（`go:generate`などの指示、Cgoの指示など）は重要な役割を果たしていますので、注意して使用してください。
