//1 月 1 日から 2 の 8 乗の 256 日目にあたる 9 月 12 日(うるう年では 9 月 13 日)はロシアではプログラマーの日とされています。
//あなたはプログラマーの日まで 2 の n 乗日目の区切りを調べたくなりました。
//
//1 月 1 日から何日目か x が与えられるので 2 の n 乗日目であれば「OK」そうでない場合は「NG」と出力してください。
//
//例えば
//
//64
//と入力されたとき 64 は 2 の 6 乗なので
//OK
//と出力してください。

package main

import (
	"bufio"
	"fmt"
	"strconv"
)
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin)//一行ずつ標準入力
	s, _ := reader.ReadString('\n')//string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)//先頭末尾のスペース改行を削る
	i, _ := strconv.Atoi(s)
	for {
		if
		i%2 != 0{
			fmt.Println("NG")
			break
		}else if i == 2 {
			fmt.Println("OK")
			break
		}else{
			i = i/2
			continue
		}

	}



	}
