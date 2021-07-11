//あなたは友達とシャボン玉を飛ばして遊んでいました。ふと、あなたは、自分が作ったシャボン玉がどこまで飛んで行くのかが気になりました。
//
//あなたは、時刻 0 のときに座標 (x, y) にシャボン玉を作ります。このシャボン玉は、時刻 i - 1 と時刻 i の間 (1 ≦ i ≦ T) に風を受けます。この風により、まずシャボン玉の x 座標が a_i 増えます。その次に、シャボン玉の y 座標が b_i 増えます。
//
//シャボン玉は、y 座標が 0 以下になると、地面に接触し割れてなくなってしまいます。
//
//時刻 T までに吹く風の情報が与えられるので、シャボン玉が通った x 座標の最大値を求めるプログラムを作成してください。
//
//例)
//入力例 1 の状況を考えます。シャボン玉が最初 (1, 1) にあるとします。時刻 4 までに次のように風が吹くとします。
//
//・時刻 0 から時刻 1: x 座標 +4, y 座標 +2
//・時刻 1 から時刻 2: x 座標 -5, y 座標 -4
//・時刻 2 から時刻 3: x 座標 +3, y 座標 +3
//・時刻 3 から時刻 4: x 座標 +3, y 座標 +3

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
	s = strings.TrimSpace(s)            //先頭末尾のスペース改行を削る
	t := strings.Split(s, " ") //中のスペースで分割
	T, _ := strconv.Atoi(t[0])
	x, _ := strconv.Atoi(t[1])
	y, _ := strconv.Atoi(t[2])
	max := x

	for i := 0; i < T; i++ {
		s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)        //先頭末尾のスペース改行を削る
		u := strings.Split(s, " ") //中のスペースで分割
		wx, _ := strconv.Atoi(u[0])
		wy, _ := strconv.Atoi(u[1])
		x += wx
		y += wy
		if  max < x{
			max = x
		}
		if y <= 0{
			break
		}

	}
	fmt.Println(max)
}