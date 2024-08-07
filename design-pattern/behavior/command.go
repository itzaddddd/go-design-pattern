package behavior

import "fmt"

// // command interface
// type Command interface {
// 	execute()
// }

// // receiver
// type Light struct{}

// func NewLight() *Light {
// 	return &Light{}
// }

// func (l *Light) TurnOn() {
// 	fmt.Println("The light is on")
// }

// func (l *Light) TurnOff() {
// 	fmt.Println("The light is off")
// }

// // concerete command

// type TurnOnCommand struct {
// 	light Light
// }

// func NewTurnOnCommand(light Light) *TurnOnCommand {
// 	return &TurnOnCommand{
// 		light: light,
// 	}
// }

// func (t *TurnOnCommand) execute() {
// 	t.light.TurnOn()
// }

// type TurnOffCommand struct {
// 	light Light
// }

// func NewTurnOffCommand(light Light) *TurnOffCommand {
// 	return &TurnOffCommand{
// 		light: light,
// 	}
// }

// func (t *TurnOffCommand) execute() {
// 	t.light.TurnOff()
// }

// // invoker

// type RemoteControl struct {
// 	command Command
// }

// func NewRemoteControl() *RemoteControl {
// 	return &RemoteControl{}
// }

// func (r *RemoteControl) setCommand(command Command) {
// 	r.command = command
// }

// func (r *RemoteControl) pressButton() {
// 	r.command.execute()
// }

// func RunCommandRemote() {
// 	light := NewLight()
// 	turnOnCmd := NewTurnOnCommand(*light)
// 	turnOffCmd := NewTurnOffCommand(*light)

// 	remote := NewRemoteControl()
// 	remote.setCommand(turnOnCmd)
// 	remote.pressButton()

// 	remote.setCommand(turnOffCmd)
// 	remote.pressButton()

// }

// example of using with transaction

type Command interface {
	Execute()
	Undo()
}

type Account struct {
	balance float64
}

func NewAccount(balance float64) *Account {
	return &Account{balance: balance}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited: %f, New Balance: %f\n", amount, a.balance)
}

func (a *Account) Withdraw(amount float64) {
	a.balance -= amount
	fmt.Printf("Withdrawn: %f, New Balance: %f\n", amount, a.balance)
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

type DepositCommand struct {
	account  *Account
	amount   float64
	executed bool
}

func NewDepositCommand(account *Account, amount float64) *DepositCommand {
	return &DepositCommand{account: account, amount: amount}
}

func (c *DepositCommand) Execute() {
	c.account.Deposit(c.amount)
	c.executed = true
}

func (c *DepositCommand) Undo() {
	if c.executed {
		c.account.Withdraw(c.amount)
		c.executed = false
	}
}

type WithdrawCommand struct {
	account  *Account
	amount   float64
	executed bool
}

func NewWithdrawCommand(account *Account, amount float64) *WithdrawCommand {
	return &WithdrawCommand{account: account, amount: amount}
}

func (c *WithdrawCommand) Execute() {
	c.account.Withdraw(c.amount)
	c.executed = true
}

func (c *WithdrawCommand) Undo() {
	if c.executed {
		c.account.Deposit(c.amount)
		c.executed = false
	}
}

type Transaction struct {
	commands []Command
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) AddCommand(command Command) {
	t.commands = append(t.commands, command)
}

func (t *Transaction) Execute() {
	for _, cmd := range t.commands {
		cmd.Execute()
	}
}

func (t *Transaction) Undo() {
	for i := len(t.commands) - 1; i >= 0; i++ {
		t.commands[i].Undo()
	}
}

func runCommandTransaction() {
	account := NewAccount(100.0)

	depositCommand := NewDepositCommand(account, 50.0)
	withdrawCommand := NewWithdrawCommand(account, 30.0)

	transaction := NewTransaction()
	transaction.AddCommand(depositCommand)
	transaction.AddCommand(withdrawCommand)

	fmt.Println("Executing transaction...")
	transaction.Execute()
	fmt.Printf("Final Balance: %f\n", account.GetBalance())

	fmt.Println("Undoing transaction...")
	transaction.Undo()
	fmt.Printf("Final Balance after undo: %f\n", account.GetBalance())
}
