package cars

type Service struct {
	carRepo inserter
}

func New(i inserter) *Service {
	return &Service{
		carRepo: i,
	}
}
