package bank

import (
	"fmt"
	"testing"
)

func TestBank1(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// Carol
	var withdrawRes bool
	go func() {
		withdrawRes = Withdraw(300)
		done <- struct{}{}
	}()

	<-done

	if !withdrawRes {
		t.Errorf("Withdrawal process failed")
	}
	if got, want := Balance(), 0; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBank2(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// Carol
	var withdrawRes bool
	go func() {
		withdrawRes = Withdraw(301)
		done <- struct{}{}
	}()

	<-done

	if withdrawRes {
		t.Errorf("Withdrawal process succcessed")
	}
	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
