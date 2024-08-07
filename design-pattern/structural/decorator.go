package structural

import "fmt"

type FetchService interface {
	fetchData(url string) string
}

type ConcreteFetchService struct{}

func NewConcreteFetchService() ConcreteFetchService {
	return ConcreteFetchService{}
}

func (c *ConcreteFetchService) fetchData(url string) string {
	return fmt.Sprintf("Data from %s", url)
}

type FetchServiceDecorator struct {
	wrappe FetchService
}

func NewFetchServiceDecorator(fetchService FetchService) FetchServiceDecorator {
	return FetchServiceDecorator{
		wrappe: fetchService,
	}
}

func (f *FetchServiceDecorator) fetchData(url string) string {
	return f.wrappe.fetchData(url)
}

type LoggingFetchServiceDecorator struct {
	fetServiceDecorator FetchServiceDecorator
}

func NewLoggingFetchServiceDecorator(fetchService FetchService) LoggingFetchServiceDecorator {
	return LoggingFetchServiceDecorator{
		fetServiceDecorator: NewFetchServiceDecorator(fetchService),
	}
}

func (l *LoggingFetchServiceDecorator) fetchData(url string) string {
	fmt.Printf("Fetching data from %s\n", url)
	data := l.fetServiceDecorator.fetchData(url)
	fmt.Printf("Fetched data: %s", data)
	return data
}

func clientCode(fetchService FetchService) {
	data := fetchService.fetchData("https://example.com/api/v1/data")
	fmt.Println(data)
}

func RunDecorator() {
	simpleFetchService := NewConcreteFetchService()
	loggedFetchService := NewLoggingFetchServiceDecorator(&simpleFetchService)

	clientCode(&loggedFetchService)
}
