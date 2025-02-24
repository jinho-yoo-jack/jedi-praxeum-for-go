package bank

type Sender struct {
	name      string
	accountId string
}

type Receiver struct {
	name      string
	accountId string
}

type Transaction struct {
	sender   Sender
	receiver Receiver
	amount   int
	date     string
}

type Account struct {
	id      string
	owner   string
	balance int
}

// NewAccount
// Constructor Of Account
func NewAccount(id string, owner string) *Account {
	return &Account{id: id, owner: owner, balance: 0}
}

func (a *Account) Id() string {
	return a.id
}

func WithDrawFunc(a *Account, amount int) {
	a.balance -= amount
}

func (a *Account) WithDraw(amount int) {
	a.balance -= amount
}

func (a Account) WithDrawByValue(amount int) int {
	a.balance -= amount
	return a.balance
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) GetBalance() int {
	return a.balance
}
