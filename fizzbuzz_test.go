package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (s *FizzBuzzSuite) TestOneUnchanged() {
	r := Run([]int{1})

	s.Equal([]string{"1"}, r)
}
