//N を 2 以上の整数とし、N の約数のうち N 自身を除いたものの和を S とします。 このとき
//
//・N = S となるような N を完全数
//・|N-S| = 1 となるような N をほぼ完全数
//
//と言うことにしましょう。ここで、|N-S| は N-S の絶対値を意味します。
//
//たとえば、N = 28 のとき、28 の約数は 1, 2, 4, 7, 14, 28 なので、S = 1+2+4+7+14 = 28 となります。従って、28 は完全数です。
//また、N = 16 のときには S = 1+2+4+8 = 15 となるので、16 はほぼ完全数です。
//
//入力された整数が完全数かほぼ完全数かそのいずれでもないかを判定するプログラムを作成してください。

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
	n, _ := strconv.Atoi(s)//string→int
	var ns []int
	for i := 0; i<n; i++ {
		s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)        //先頭末尾のスペース改行を削る
		v, _ := strconv.Atoi(s) //string→int
		ns = append(ns, v)

	}
	for i := 0; i<n; i++ {
		var vs []int //約数を保持するslice
		for a := 1; a< ns[i]; a++ {
			if ns[i]%a == 0 {
				vs = append(vs, a) //割り切れるなら追加
			} else {
				continue
			}
		} //全約数を足し合わせる
		var ts int
		for _, value := range vs {
			ts += value
		}
		if ts == ns[i] {
			fmt.Println("perfect")
		} else if ts == ns[i]-1 || ts == ns[i]+1 {
			fmt.Println("nearly")
		} else {
			fmt.Println("neither")
		}

	}


}