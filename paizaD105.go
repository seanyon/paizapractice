/*
あなたは2つの文字列の長さを目で比較していましたが長い文字では目視で確認するのが辛くなってきました。
2つの文字列が入力されるので文字列の長さが一致していれば "Yes" 、違う場合は "No" と出力するプログラムを作成してください。

入力例 1 の場合

Paiza
Pizza
それぞれの文字列の長さは 5 文字なので以下のように出力して下さい。

Yes
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	l := len(s)

	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	m := len(s)

	if l == m {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}

}
