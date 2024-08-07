package creational

type Singleton interface {
	AddOne()
	GetCount() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) GetCount() int {
	return s.count
}

func (s *singleton) AddOne() {
	s.count++
}
