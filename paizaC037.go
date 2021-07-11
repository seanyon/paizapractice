//夜アニメの放送開始時刻は、しばしば "時" を表す部分の数字が24以上になることを許容し次のように表記することがあります。
//
//
//01/27 24:30
//
//この日付を "時" の部分が0以上23以下になるような表記で、
//
//
//01/28 00:30
//
//と表すことができます。
//
//
//前者の表記方法は時として便利ではありますが、実際に放送される "日" が知りたい場合は後者のほうが便利です。
//
//
//そこで、日時が与えられるので、"時" の部分が 0 以上 23 以下になるように修正した日時を出力するプログラムを作成してください。
//
//ただし、"時" を修正した結果、日がその月に存在しない日になった場合でも月を調整することなくそのまま出力してください。もしかすると、このプログラムが使われる世界ではそのような日付が存在するかもしれないからです。
package main

import (
	"bufio"
	"fmt"
	"strconv"
)
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin) //一行ずつ標準入力
	s, _ := reader.ReadString('\n')     //string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)//先頭末尾のスペース改行を削る
	t := strings.Split(s, " ")//スペースで分割
	date := strings.Split(t[0], "/")
	month, _ := strconv.Atoi(date[0])
	day, _ := strconv.Atoi(date[1])
	time := strings.Split(t[1], ":")
	hour, _ := strconv.Atoi(time[0])
	minute, _ := strconv.Atoi(time[1])
	if hour >= 24{
		var plusdate int
		plusdate = hour/24
		hour = hour%24
		day += plusdate
	}

	fmt.Printf("%02d/%02d %02d:%02d", month, day, hour, minute)//0埋め

	}

