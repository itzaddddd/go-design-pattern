package structural

import "fmt"

// import "fmt"

// abstract for message display
type IMessageDisplay interface {
	display()
}

// type MessageDisplay struct {
// 	messageSource IMessageSource
// }

func NewAlertMessageDisplay(messageSource IMessageSource) *alertMessageDisplay {
	return &alertMessageDisplay{
		messageSource: messageSource,
	}
}

// implementor for message source
type IMessageSource interface {
	getMessage() string
}

// refined display alert
type alertMessageDisplay struct {
	messageSource IMessageSource
}

func (a *alertMessageDisplay) display() {
	message := a.messageSource.getMessage()
	fmt.Println(message)
}

// refined display modal
type modalMessageDisplay struct {
	messageSource IMessageSource
}

func NewModalMessageDisplay(messageSource IMessageSource) *modalMessageDisplay {
	return &modalMessageDisplay{
		messageSource: messageSource,
	}
}

func (m *modalMessageDisplay) display() {
	message := m.messageSource.getMessage()
	fmt.Printf("modal: %s\n", message)
}

// concrete implementor for message source

type StaticMessageSource struct {
	message string
}

func NewStaticMessageSource(message string) StaticMessageSource {
	return StaticMessageSource{
		message: message,
	}
}

func (s *StaticMessageSource) getMessage() string {
	return s.message
}

type ApiMessageSource struct{}

func NewApiMessageSource() ApiMessageSource {
	return ApiMessageSource{}
}

func (a *ApiMessageSource) getMessage() string {
	return "message from api"
}

func runBridge() {
	staticSource := NewStaticMessageSource("static message")
	apiSource := NewApiMessageSource()

	alertStaticDisplay := NewAlertMessageDisplay(&staticSource)
	alertApiDisplay := NewAlertMessageDisplay(&apiSource)
	alertStaticDisplay.display()
	alertApiDisplay.display()

	modalStaticDisplay := NewModalMessageDisplay(&staticSource)
	modalApiDisplay := NewModalMessageDisplay(&apiSource)
	modalStaticDisplay.display()
	modalApiDisplay.display()

}
