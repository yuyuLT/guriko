package enemyhand

import (
	"math/rand"
	"time"
)
func CreateEnemyHand()int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}