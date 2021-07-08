//半角小文字アルファベットで構成された文字列 S が入力されます。
//
//文字列 S の中に含まれる "at" という文字列を全て "@" (アットマークに)置換して出力して下さい。
//
//例えば
//
//paizaatcodeattest
//と入力された場合
//
//paiza@code@test
//と出力してください。
package main

import (
	"bufio"
	"fmt"
)
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin)//一行ずつ標準入力
	s, _ := reader.ReadString('\n')//string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)//先頭末尾のスペース削る
	fmt.Println(strings.Replace(s, "at", "@", -1))



}