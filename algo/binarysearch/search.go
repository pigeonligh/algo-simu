package binarysearch

import (
	"fmt"
	"strconv"

	"github.com/pigeonligh/algo-simu/algo"
)

type searcher struct {
	checker algo.Checker
	result  int

	min int
	max int
}

func (s *searcher) Init(args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("please input min and max")
	}
	s.min, _ = strconv.Atoi(args[0])
	s.max, _ = strconv.Atoi(args[1])
	return nil
}

func (s *searcher) check(val int) bool {
	return s.checker(fmt.Sprint(val)) == "yes"
}

func (s *searcher) Solve() error {
	left := s.min
	right := s.max + 1
	for left+1 < right {
		// current interval: [left, right)
		mid := (left + right) / 2
		if s.check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	s.result = left
	return nil
}

func (s *searcher) Result() string {
	return fmt.Sprint(s.result)
}

func New(checker algo.Checker) (algo.IAlgorithm, error) {
	return &searcher{checker: checker}, nil
}
