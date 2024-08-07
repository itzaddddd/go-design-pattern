package behavior

import "fmt"

type NotificationStrategy interface {
	send(user string, message string)
}

type EmailNoti struct{}

func (n *EmailNoti) send(user string, message string) {
	fmt.Printf("send email to %s , message: %s\n", user, message)
}

type PushNoti struct{}

func (n *PushNoti) send(user string, message string) {
	fmt.Printf("send push notification to %s , message: %s\n", user, message)
}

type SmsNoti struct{}

func (n *SmsNoti) send(user string, message string) {
	fmt.Printf("send sms to %s , message: %s\n", user, message)
}

type NotiContext struct {
	strategy NotificationStrategy
}

func NewNotiContext(strategy NotificationStrategy) *NotiContext {
	return &NotiContext{
		strategy: strategy,
	}
}

func (n *NotiContext) setStrategy(strategy NotificationStrategy) {
	n.strategy = strategy
}

func (n *NotiContext) notify(user string, message string) {
	n.strategy.send(user, message)
}

func RunStrategy() {
	user := "mbk49"
	message := "Hello, welcome!"

	emailStrategy := EmailNoti{}
	ctx := NewNotiContext(&emailStrategy)
	ctx.notify(user, message)

	smeStrategy := SmsNoti{}
	ctx.setStrategy(&smeStrategy)
	ctx.notify(user, message)
}
