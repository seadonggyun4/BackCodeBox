package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct { //은행계좌
	balance int //잔액 (메모리 보호대상) Mutex
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) { //인출
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) { //입금
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account    //slice 배열 (메모리 보호대상) Mutex
var globalLock *sync.Mutex //글로벌 락을 만들어 정교하게 보호한다.

func Transfer(sender, receiver int, money int) { //송금자, 받는자, 송금액
	globalLock.Lock()
	accounts[sender].Widthdraw(money) // 어카운트 배열에서 돈을 빼고
	accounts[receiver].Deposit(money) //리시버에게 송금
	globalLock.Unlock()
}

func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()
	return total
}

func RandomTransfer() { //랜덤함수 레퍼런스
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance() //샌더 잔액여부 확인
		if balance > 0 {
			break //샌더 잔액이 있으면 멈춘다
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ { //20개 어카운트 배열 1000 지정
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
		globalLock = &sync.Mutex{}
	}

	PrintTotalBalance()

	for i := 0; i < 10; i++ { //쓰레드개수를 (1개로 하면 일정한 카운트를 유지한다)
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
