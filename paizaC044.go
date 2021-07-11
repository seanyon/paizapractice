//じゃんけんは、2 人以上の参加者により行われるゲームです。 各参加者は、3 種類の手(グー・チョキ・パー) のいずれかを出します。 この手 の組み合わせにより勝者が決まります。
//
//勝敗の決定は以下の通りです。
//
//・グーは、チョキに勝ち、パーに敗れる
//・チョキは、パーに勝ち、グーに敗れる
//・パーは、グーに勝ち、チョキに敗れる
//・グー・チョキ・パーの 3 種類が出されている場合は引き分け
//・グー・チョキ・パーのいずれか 1 種類のみが出されている場合も引き分け
//
//3 人でじゃんけんした場合の例を下図に示しました。
//

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
	n, _ := strconv.Atoi(s)             //string→int, n人

	var rs int
	var ps int
	var ss int
	for i := 0; i < n; i++{
	s, _ := reader.ReadString('\n') //string型にする。返り値に改行含まれる
	s = strings.TrimSpace(s)        //先頭末尾のスペース改行を削る
	switch s {
	case "rock":
		rs += 1
	case "paper":
		ps += 1
	case "scissors":
		ss += 1
	}
}
	if rs == n ||  ps == n || ss == n || (rs>0 && ps >0 && ss >0){
		fmt.Println("draw")
	}else if rs == 0{
		fmt.Println("scissors")
	}else if ps == 0{
		fmt.Println("rock")
	}else{
		fmt.Println("paper")
	}



}

