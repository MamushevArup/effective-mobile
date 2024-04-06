package cars

type UseCase struct {
	carManager carManager
}

func New(c carManager) *UseCase {
	return &UseCase{
		carManager: c,
	}
}
