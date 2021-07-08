//あなたは半角アルファベットの苗字と名前からそれぞれ 1文字目を取りイニシャルを作ることにしました。
//半角スペース区切りで苗字と名前が入力されるので、1文字目を取り "." (半角ドット)で区切った文字列を出力してください。
//
//例えば以下の様な入力の場合
//
//Paiza Tarou
//以下の様に出力してください
//P.T
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
	t := strings.Split(s, " ")//スペースで区切る
	ln := t[0]
	fn:= t[1]
	fmt.Printf("%s.%s", string(ln[0]), string(fn[0]))





}
