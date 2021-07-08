/*
11 月 11 日が迫るある日、あなたはゾロ目の日付を判定するプログラムを作成することにしました。
ゾロ目の日とはある日付に対して 月 と 日 に含まれる全ての桁の数字が同じものを指します。
スペース区切りで月と日が与えられるのでゾロ目の日であれば "Yes" そうでなければ "No" と出力してください。

例えば入力例 1 では 11 月 1 日が以下のように与えられます。

11 1
この場合、月と日ともに 1 のゾロ目なので

Yes
と出力して下さい。
*/
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
	s = strings.TrimSpace(s)//先頭末尾のスペース削る1
	s = strings.Replace(s, " ", "", -1)
	for i := 0; i < len(s)-1; i++{
		if s[i] != s[i+1]{
			fmt.Println("No")
			break
		}else if i == len(s)-2{
			fmt.Println("Yes")
			break
		}else{
			continue
		}

		}





}