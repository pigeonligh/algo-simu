package algo

type IAlgorithm interface {
	Init() error
	Solve() error
	Result() string
}

type Checker func(string) string
