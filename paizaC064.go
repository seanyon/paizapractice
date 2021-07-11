//入力は以下のフォーマットで与えられます。
//
//M N
//c_1
//c_2
//...
//c_M
//a_{1,1} a_{1,2}  ... a_{1,M}
//a_{2,1} a_{2,2}  ... a_{2,M}
//...
//a_{N,1} a_{N,2}  ... a_{N,M}
//・1 行目にそれぞれ料理の品数、就活生の人数を表す整数 M , N がこの順で半角スペース区切りで与えられます。
//・続く M 行のうちの i 行目 (1 ≦ i ≦ M) には、100 g あたりのカロリーを表す整数 c_i がこの順で半角スペース区切りで与えられます。
//・さらに続く N 行のうちの i 行目 (1 ≦ i ≦ N) には M 個の整数が半角スペース区切りで与えられます。i 行目の j 番目 (1 ≦ j ≦ M) の整数 a_{i, j} は i 番目の就活生が食べた、j 番目の料理の量を表します。
//・入力は合計で M + N + 1 行となり、入力値最終行の末尾に改行が 1 つ入ります。
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
	M, _ := strconv.Atoi(t[0])
	N, _ := strconv.Atoi(t[1])

	var cs []float64
	for i := 0; i < M; i++ {//料理ごとのカロリー
		s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)        //先頭末尾のスペース改行を削る
		c_i, _ := strconv.Atoi(s)
		cs = append(cs, float64(c_i))
	}

	for i := 0; i < N; i++ {//人ごと
		var tc_i int
		s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)        //先頭末尾のスペース改行を削る
		t := strings.Split(s, " ")      //中のスペースで分割
		for i := 0; i < M; i++ {
			gd_i, _ := strconv.Atoi(t[i])
			gdi := float64(gd_i)//この時点でfloat型にしなければならない。
			cal:= int((gdi/100)*cs[i])
			tc_i += cal
		}

		fmt.Println(tc_i)
	}

}
