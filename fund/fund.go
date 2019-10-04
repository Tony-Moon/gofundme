package fund

type Fund struct {
	balance float64
}

func NewFund(i float64) *Fund{
	return &Fund {
		balance: i,
	}
}

func (f *Fund) Balance() float64 {
	return f.balance
}

func (f *Fund) Withdraw(i float64) {
	f.balance -= i
}