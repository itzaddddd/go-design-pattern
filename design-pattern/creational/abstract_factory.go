package creational

import "fmt"

// abstract product
type UserManagement interface {
	createUser()
	deleteUser()
}

type ContentManagement interface {
	createPost()
	deletePost()
}

// concrete product for english
type EnglishUserManagement struct{}

func (e *EnglishUserManagement) createUser() {
	fmt.Println("Creating user in english")
}

func (e *EnglishUserManagement) deleteUser() {
	fmt.Println("Deleting user in english")
}

type EnglishContentManagement struct{}

func (e *EnglishContentManagement) createPost() {
	fmt.Println("Creating post in english")
}

func (e *EnglishContentManagement) deletePost() {
	fmt.Println("Deleting post in english")
}

// concrete product for thai
type ThaiUserManagement struct{}

func (e *ThaiUserManagement) createUser() {
	fmt.Println("Creating user in thai")
}

func (e *ThaiUserManagement) deleteUser() {
	fmt.Println("Deleting user in thai")
}

type ThaiContentManagement struct{}

func (e *ThaiContentManagement) createPost() {
	fmt.Println("Creating post in thai")
}

func (e *ThaiContentManagement) deletePost() {
	fmt.Println("Deleting post in thai")
}

// abstract factory interface
type LocalFactory interface {
	createUserManagement() UserManagement
	createContentManagement() ContentManagement
}

// concrete factory
type EnglishLocaleFactory struct{}

func (e *EnglishLocaleFactory) createUserManagement() UserManagement {
	return new(EnglishUserManagement)
}

func (e *EnglishLocaleFactory) createContentManagement() ContentManagement {
	return new(EnglishContentManagement)
}

type ThaiLocaleFactory struct{}

func (e *ThaiLocaleFactory) createUserManagement() UserManagement {
	return new(ThaiUserManagement)
}

func (e *ThaiLocaleFactory) createContentManagement() ContentManagement {
	return new(ThaiContentManagement)
}

// locale type
const (
	EN = "en"
	TH = "th"
)

// usage based on locale
func getFactoryForLocale(locale string) LocalFactory {
	switch locale {
	case EN:
		return new(EnglishLocaleFactory)
	case TH:
		return new(ThaiLocaleFactory)
	default:
		return nil
	}
}

func runAbstract() {
	factory := getFactoryForLocale("th")

	thFactory := factory.(*ThaiLocaleFactory)
	thUserManagement := thFactory.createUserManagement()
	thUserManagement.createUser()

	contentManagement := factory.createContentManagement()
	contentManagement.createPost()
}
