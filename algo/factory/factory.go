package factory

import (
	"fmt"

	"github.com/pigeonligh/algo-simu/algo"
	"github.com/pigeonligh/algo-simu/algo/binarysearch"
)

type newFn func(algo.Checker) (algo.IAlgorithm, error)

var factory map[string]newFn

func init() {
	factory = make(map[string]newFn)
	Register("binarysearch", binarysearch.New)
}

func Register(name string, fn newFn) {
	factory[name] = fn
}

func New(kind string, checker algo.Checker) (algo.IAlgorithm, error) {
	newFunc, ok := factory[kind]
	if ok {
		return newFunc(checker)
	}
	return nil, fmt.Errorf("unknown algorithm")
}
