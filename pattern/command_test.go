package pattern

import "testing"

func TestCommand(t *testing.T) {
	s := StockTrade{}
	b := NewBuyStock(s)
	se := NewSellStock(s)
	a := new(Agent)
	err := a.placeOrder(b)
	if err != nil {
		t.Fatal(err)
	}
	err = a.placeOrder(se)
	if err != nil {
		t.Fatal(err)
	}
}
