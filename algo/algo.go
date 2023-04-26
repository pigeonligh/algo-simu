package algo

type IAlgorithm interface {
	Init(args ...string) error
	Solve() error
	Result() string
}

type Checker func(string) string
