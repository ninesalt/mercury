package user

import (
	"fmt"
	"testing"
)

var u User

func init() {
	u = CreateUser("paul")
}

func TestAddToBalance(t *testing.T) {
	fmt.Println("|ee")
	stock := "AAPL"
	toAdd := float32(10)
	oldBalance := u.Balances()[stock]
	u.AddToBalance(stock, toAdd)
	newBalance := u.Balances()[stock]
	expected := oldBalance + toAdd

	// test adding a postive number
	if expected != newBalance {
		t.Errorf("Expected balance to increase. Got %v, expected %v", newBalance, expected)
	}

	// test adding a negative number
	oldBalance = newBalance
	expected = newBalance
	u.AddToBalance(stock, float32(-20))
	newBalance = u.Balances()[stock]

	if newBalance != expected {
		t.Errorf("Should not add negative numbers to balance. Got %v, expected %v", newBalance, expected)
	}

}

func TestSubtractFromBalance(t *testing.T) {
	stock := "UBER"
	toSub := float32(10)
	oldBalance := u.Balances()[stock]
	u.SubtractFromBalance(stock, toSub)
	newBalance := u.Balances()[stock]
	expected := oldBalance - toSub

	// test subtracting a postive number
	// for now balance is allowed to go negative as long as it's properly subtracted
	if expected != newBalance {
		t.Errorf("Expected balance to increase. Got %v, expected %v", newBalance, expected)
	}

	// test subtracting a negative number
	oldBalance = newBalance
	expected = newBalance
	u.AddToBalance(stock, float32(-20))
	newBalance = u.Balances()[stock]

	if newBalance != expected {
		t.Errorf("Should not add negative numbers to balance. Got %v, expected %v", newBalance, expected)
	}
}

func TestBalances(t *testing.T) {
	u = CreateUser("robert")
	if u.Balances() == nil {
		t.Errorf("Expected default balances map to not be nil")
	}
}
