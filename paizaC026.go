//Rabbit社の社長はウサギが大好きで、 社内で1匹のウサギを飼育することにしました。
//ウサギが人参を好むことは広く知られています。 また、ウサギに詳しい社員によると、 人参は糖分が多いため与えすぎるのは良くないようです。
//
//そこで、人参は1日1本だけ与えることにし、 特に糖分が許容範囲内で質量が大きい人参を選ぶことにしました。
//
//プログラマーであるあなたはRabbit社の依頼を受け、 1 から N で番号付けられた N 個の人参のデータを入力として、 糖分が許容範囲内で質量の最も大きい人参を見つけるプログラムを作成することになりました。
//
//糖分には基準値Sと許容誤差 p があり、 糖分が S - p 以上 S + p 以下ならば許容範囲内となります。
//また、糖分が許容範囲内で質量の最も大きい人参 が複数ある場合は、それらのうち番号の最も小さいものを知ることができれば十分です。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //一行ずつ標準入力
	s, _ := reader.ReadString('\n')     //string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)            //先頭末尾のスペース改行を削る
	t := strings.Split(s, " ") //中のスペースで分割
	N, _ := strconv.Atoi(t[0])  //N個の人参
	S, _ := strconv.Atoi(t[1])
	P, _ := strconv.Atoi(t[2])
	var carotts []map[string]int
	var index = 0

	for i := 0; i < N; i++ {
		s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)
		t := strings.Split(s, " ") //中のスペースで分割
		carott_i := make(map[string]int)
		carott_i["name"] = i + 1
		carott_i["mass"], _ = strconv.Atoi(t[0])
		carott_i["sugar"], _ = strconv.Atoi(t[1])
		carotts = append(carotts, carott_i)
		if carott_i["sugar"] > S+P || carott_i["sugar"] < S-P {
			carott_i["mass"] = 0
		}
	}
		max := carotts[0]["mass"]
		for i := 0; i < N; i++{
			if max < carotts[i]["mass"]{
				max = carotts[i]["mass"]
				index = i
			}
		}
		if max == 0{
			fmt.Println("not found")
		}else{
			fmt.Println(index+1)
		}



}