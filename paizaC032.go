//paiza カードは毎日のお買い物でポイントがたまるお得なポイントカードです。
//ポイントの付与基準は商品の種類によって以下のように異なります。
//ポイントは以下の手順で計算されます。
//1. 商品の種類ごとの合計購入金額を計算する。
//2. それぞれのポイント付与基準に基づき商品の種類ごとの付与ポイントを計算する。
//3. 商品の種類ごとの付与ポイントを合計する。(100円未満は切り捨て)
//
//買い物内訳と合計ポイントの例を見てみましょう。
//
//商品 1: 食料品 200 円
//商品 2: 書籍 240 円
//商品 3: 食料品 120 円
//商品 4: その他 460 円
//商品 5: 書籍 240 円
//商品 6: 衣類 3200 円
//
//食料品合計 320 円 → 15 ポイント
//書籍合計 480 円 → 12 ポイント
//衣類合計 3200 円 → 64 ポイント
//その他合計 460 円 → 4 ポイント => 合計ポイント → 95 ポイント
//このように購入した商品の種類と金額のリストが与えられるので、その買い物で付与されるポイントの合計を求めてください。
//・入力の 1 行目に購入した商品の総数を表す N が与えられます。
//・続く N 行目のうち i 行目 (1 ≦ i ≦ N) に商品の種類を表す整数 v_i および商品の金額を表す整数 p_i がこの順に半角スペース区切りで与えられます。
//　　・v_i について、0 は食料品、1 は書籍、2 は衣類、3 はその他を表します。
//・入力は合計で N + 1 行からなり、入力値最終行の末尾に改行が１つ入ります。
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

	var psum0 int //ポイント合計
	var psum1 int //ポイント合計
	var psum2 int //ポイント合計
	var psum3 int //ポイント合計

	for i := 0; i<n; i++ {
		s, _ := reader.ReadString('\n')     //string型にする。返り値に改行含まれる
		s = strings.TrimSpace(s)            //先頭末尾のスペース改行を削る
		t := strings.Split(s, " ") //中のスペースで分割
		v, _ := strconv.Atoi(t[0])  //v:種類
		p, _ := strconv.Atoi(t[1])//p：値段
		switch v {
		case 0:
			psum0 += p
		case 1:
			psum1 += p
		case 2:
			psum2 += p
		case 3:
			psum3 += p}
	}
	vsum :=  int(psum0/100)*5+int(psum1/100)*3+int(psum2/100)*2+int(psum3/100)*1
	fmt.Println(vsum)

}
