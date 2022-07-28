package enemyhand

import (
	"math/rand"
	"time"
)

//自分の手が連続3回同じかどうか判定
func JudgeConsecutive(hands []int) bool{
	if len(hands) >= 4{
		if hands[len(hands)-4] == hands[len(hands)-3] && hands[len(hands)-4] == hands[len(hands)-2]{
			return true
		}
	}
	
	return false
}

//自分の手が3回連続で同じの場合、それに勝てる手を選択
func CreateWinHand(hands[]int) int{
	var hand int
	switch hands[len(hands)-2]{
		case 0:
			hand = 2;
		case 1:
			hand = 0;
		case 2:
			hand = 1;
	} 

	return hand
}

func CreateEnemyHand() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

func CreateOptimizedEnemyHand() int {

	var hand int
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(5) {
		case 0,1:
			hand = 0
		case 2,3:
			hand = 1	
		case 4:
			hand = 2
	}

	return hand
}