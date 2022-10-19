package singleton

type Singleton struct {
	count int
}

var instance *Singleton

func init() {
	instance = &Singleton{}
}

func GetInstance() *Singleton {
	return nil
}

func (s *Singleton) AddOne() int {
	return 0
}
