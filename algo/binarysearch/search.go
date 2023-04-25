package binarysearch

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pigeonligh/algo-simu/algo"
)

type searcher struct {
	checker algo.Checker
	result  int

	min int
	max int
}

func (s *searcher) Init() error {
	if len(os.Args) < 4 {
		return fmt.Errorf("please input min and max")
	}
	s.min, _ = strconv.Atoi(os.Args[2])
	s.max, _ = strconv.Atoi(os.Args[3])
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
