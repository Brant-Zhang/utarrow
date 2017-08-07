package pattern

import (
	"container/list"
	"fmt"
)

//Command
type Order interface {
	execute() error
}

//Reveiver
type StockTrade struct{}

func (s *StockTrade) buy() {
	fmt.Println("i will buy some stocks")
}

func (s *StockTrade) sell() {
	fmt.Println("i will sell some stocks")
}

//Ivoker
type Agent struct {
	OrdersQueue *list.List
}

//调用者较简单，执行命令接收和执行
func (a *Agent) placeOrder(o Order) error {
	if a.OrdersQueue == nil {
		a.OrdersQueue = list.New()
	}
	a.OrdersQueue.PushBack(o)
	e := a.OrdersQueue.Back()
	err := e.Value.(Order).execute()
	a.OrdersQueue.Remove(e)
	return err
}

//ConcreteCommand
//我们可以根据实际需要扩展命令类
type BuyStockOrder struct {
	stock StockTrade
}

//在concreteCommand中通过构造函数定义该命令针对哪个Recevier
//定义命令接受的主体
func NewBuyStock(st StockTrade) *BuyStockOrder {
	b := new(BuyStockOrder)
	b.stock = st
	return b
}

func (b *BuyStockOrder) execute() error {
	b.stock.buy()
	return nil
}

//ConcreteCommand
type SellStockOrder struct {
	stock StockTrade
}

func NewSellStock(st StockTrade) *SellStockOrder {
	s := new(SellStockOrder)
	s.stock = st
	return s
}

func (s *SellStockOrder) execute() error {
	s.stock.sell()
	return nil
}
