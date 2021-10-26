package bank

var deposits = make(chan int) // send amount to deposit
type WithdrawEntry struct {
	amount int         // send amount to withdraw
	result chan<- bool // true: success withdraw
}

var withdraw = make(chan WithdrawEntry)
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraw <- WithdrawEntry{
		amount: amount,
		result: result,
	}
	return <-result
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case e := <-withdraw:
			if balance-e.amount < 0 {
				e.result <- false
				break
			}
			balance -= e.amount
			e.result <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
