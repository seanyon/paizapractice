/*
ある正の整数aとbがスペース区切りで入力されます。


aとbを比較し大きい方の値を出力して下さい。aとbが同じ場合は「eq」と出力して下さい。
*/
package main

import (
	"bufio"
	"fmt"
	"math"
)
import "os"
import "strings"
import "strconv"

func main() {
	reader := bufio.NewReader(os.Stdin)//一行ずつ標準入力
	s, _ := reader.ReadString('\n')//string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)//先頭末尾のスペース削る
	t := strings.Split(s, " ")//スペースで区切る
	u := make([]float64, len(t))//float64のスライス
	for i := 0; i < 2; i++{
		u[i], _  = strconv.ParseFloat(t[i],64)//string→float64に
	}
	if u[0]==u[1]{
		fmt.Println("eq")
	} else {
		fmt.Println(math.Max(u[0], u[1]))//math.Maxはfloat64にしか使えない
	}



}
