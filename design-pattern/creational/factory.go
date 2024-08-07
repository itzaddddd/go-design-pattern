package creational

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
	Wallet    = 3
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	case Wallet:
		return new(WalletPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}
type WalletPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using debit card\n", amount)

}

func (c *WalletPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using walled\n", amount)

}

func (c *WalletPM) Add(amount float32) string {
	return fmt.Sprintf("%0.2f add  wallet\n", amount)

}

func runFactory() {
	factory, err := GetPaymentMethod(Wallet)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	wallet, ok := factory.(*WalletPM)

	if !ok {
		fmt.Printf("err: assert payment method failed\n")
	}

	paymentResult := wallet.Pay(2999.51)
	fmt.Printf("result of payment: %s", paymentResult)

	addResult := wallet.Add(3000.50)
	fmt.Printf("result of add: %s", addResult)

}
