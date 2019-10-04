package fund

import (
	"testing"
	"log"
	"strconv"
)

func Test_fund(t *testing.T) {
	fund := NewFund(100)

	if fund.Balance() == 100 {
		log.Println("Withdrawing 100 works")
	} else {
		t.Errorf("Expectd 100 but got " + strconv.FormatFloat(fund.Balance(), 'f', 2, 64))
	}
}

func Test_Withdraw(t *testing.T) {
	fund := NewFund(100)
	fund.Withdraw(5)

	if fund.Balance() == 95 {
		log.Println("Withdrawing 5 works")
	} else {
		t.Errorf("Expectd 5 but got " + strconv.FormatFloat(fund.Balance(), 'f', 2, 64))
	}
}

func Bench_Fund(b *testing.B) {
	fund := NewFund(float64 (b.N))

	for i:=0; i<b.N; i++ {
		fund.Withdraw(1)
	}
	if fund.Balance() == 0 {
		log.Println("All funds withdrawn successfully")
	} else {
		b.Errorf("Something went wrong! :(")
	}
}