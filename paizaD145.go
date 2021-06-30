/*
あなたはあるゲームをプレイしています。

プレイヤーはスタミナのポイント N を所持しており1プレイあたり M ポイントずつ消費されます。
プレイヤーのスタミナと1プレイあたりの消費ポイントがスペース区切りで与えられるので何回プレイ出来るかを出力してください。

例えば入力例 1 の場合

30 7
30 / 7 = 4 あまり 2　となり 4 回プレイすることが可能なので以下のように出力してください
4
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	t := strings.Split(s, " ")
	u := make([]int, len(t))
	for i := 0; i < 2; i++ {
		u[i], _ = strconv.Atoi(t[i])
	}
	x := float64(u[0] / u[1])
	fmt.Println(math.Floor(x))

}
