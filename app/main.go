package main

import (
	"fmt"
	"strconv"
	"example.com/module/enemyhand"
)

func main() {
	var steps int
	input_flg := 0

	//段数を入力
	for input_flg == 0 {

		fmt.Print("ゴールまでの段数(3の倍数)を入力してください > ")
		fmt.Scan(&steps)

		if steps % 3 != 0 	{
			fmt.Println("段数は3の倍数で入力してください")
		}

		if steps % 3 == 0 {
			input_flg = 1;
		} 
	}

	my_hands := []int{}

	p_steps := steps
	c_steps := steps

	end_flg := 0

	for end_flg == 0{

		//ジャンケンの手を入力
		rpc_flg := 0
		var my_hand int

		for rpc_flg == 0 {	
			fmt.Print("グーリーコ！(グー:0 チョキ:1 パー:2 を入力) > ")
			fmt.Scan(&my_hand)

			if my_hand != 0 && my_hand != 1 && my_hand != 2 {
				fmt.Println("グー:0 チョキ:1 パー:2 のいずれかを入力してください")
			}else{
				rpc_flg = 1;
			}
		}

		//自分が出した手を記録
		my_hands = append(my_hands, my_hand)

		var enemy_hand int

		//自分の手が連続3回同じ手の場合、それに勝てる手を出す
		if enemyhand.JudgeConsecutive(my_hands){
			enemy_hand = enemyhand.CreateWinHand(my_hands)
		}else{
			//相手の手を乱数で生成（自分がゴールに近づいたとき出し方に変化を加える）
			if p_steps <= 15 {
				enemy_hand = enemyhand.CreateEnemyHand()
			}else{
				enemy_hand = enemyhand.CreateOptimizedEnemyHand()
			}
		}
		
		var enemy_rpc string
		switch enemy_hand {
		case 0:
			enemy_rpc = "グー"
		case 1:
			enemy_rpc = "チョキ"
		case 2:
			enemy_rpc = "パー"
		}

		//ジャンケンの結果を判定し、移動を行う

		judge_value := (my_hand - enemy_hand + 3) % 3
		var rpc_result string
		switch judge_value {
		case 0:
			rpc_result = "引き分けです"
		case 1:
			rpc_result = "あなたの負けです"
			c_steps -= advanceCount(enemy_hand)
		case 2: 
			rpc_result = "あなたの勝ちです"
			p_steps -= advanceCount(my_hand)
		}

		fmt.Println("相手の手は" + enemy_rpc + "でした " + rpc_result) 

		//終了チェックを行う
		if p_steps <= 0 {
			fmt.Println("あなたが先にゴールしました　ゲームを終了します")
			end_flg = 1
		} 

		if c_steps <= 0 {
			fmt.Println("相手が先にゴールしました　ゲームを終了します")
			end_flg = 1
		} 

		if end_flg == 0{
			fmt.Println("あなたは残り" + strconv.Itoa(p_steps) + "段です")
			fmt.Println("相手は残り" + strconv.Itoa(c_steps) + "段です")
		}	
	}
}

func advanceCount(hand int) int{
	var count int
	switch hand {
	case 0:
		count = 3
	case 1:
		count = 6
	case 2: 
		count = 6
	}
	return count
}