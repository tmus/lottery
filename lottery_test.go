package lottery_test

import (
	"testing"

	"github.com/tmus/lottery"
)

func TestItCanWin(t *testing.T) {
	won := false

	lottery.Odds(1, 1).Success(func() {
		won = true
	}).Choose()

	if won != true {
		t.Error("Expected `won` to be true.")
	}
}

func TestItCanLose(t *testing.T) {
	lost := false

	// 0, 1 odds means it will always lose.
	lottery.Odds(0, 1).Failure(func() {
		lost = true
	}).Choose()

	if lost != true {
		t.Error("Expected `lost` to be true.")
	}
}

func TestItCanRunMultiple(t *testing.T) {
	count := 0

	lottery.Odds(1, 1).Success(func() {
		count++
	}).ChooseMany(10)

	if count != 10 {
		t.Error("Expected `count` to be `2`.")
	}
}
