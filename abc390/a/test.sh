#!/bin/bash

function testCase1() {
    if diff <(cat ./tests/sample-1.in | go run main.go) ./tests/sample-1.out; then
        echo "✅ 出力は sample-1.out と一致しています。"
    else
        echo "❌ 出力が sample-1.out と異なります。"
    fi
}

function testCase2() {
    if diff <(cat ./tests/sample-2.in | go run main.go) ./tests/sample-2.out; then
        echo "✅ 出力は sample-2.out と一致しています。"
    else
        echo "❌ 出力が sample-2.out と異なります。"
    fi
}

function testCase3() {
    if diff <(cat ./tests/sample-3.in | go run main.go) ./tests/sample-3.out; then
        echo "✅ 出力は sample-3.out と一致しています。"
    else
        echo "❌ 出力が sample-3.out と異なります。"
    fi
}

function testCase4() {
    if diff <(cat ./tests/sample-4.in | go run main.go) ./tests/sample-4.out; then
        echo "✅ 出力は sample-4.out と一致しています。"
    else
        echo "❌ 出力が sample-4.out と異なります。"
    fi
}

testCase1
testCase2
testCase3
testCase4
exit 0