package utils

import (
	"fmt"
	"strings"
)

type Strings struct {
	build *strings.Builder
}

func StringBuilder() *Strings {
	return &Strings{build: new(strings.Builder)}
}

func (s *Strings) String() string {
	return s.build.String()
}

func (s *Strings) Len() int {
	return s.build.Len()
}

func (s *Strings) Empty() bool {
	return s.build.Len() == 0
}

func (s *Strings) Cap() int {
	return s.build.Cap()
}

func (s *Strings) Reset() {
	s.build.Reset()
}

func (s *Strings) Grow(n int) {
	s.build.Grow(n)
}

func (s *Strings) Write(p []byte) (int, error) {
	return s.build.Write(p)
}

func (s *Strings) WriteByte(c byte) error {
	return s.build.WriteByte(c)
}

func (s *Strings) WriteRune(r rune) (int, error) {
	return s.build.WriteRune(r)
}

func (s *Strings) WriteString(str string) (int, error) {
	return s.build.WriteString(str)
}

type Condition func() bool

func (s *Strings) WriteOnCondition(cond Condition, v string) *Strings {
	if cond() {
		_, _ = s.build.WriteString(v)
	}
	return s
}

type Int struct {
	v int
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.v)
}

func (i Int) Load() int {
	return i.v
}

func NewInt(v int) Int {
	return Int{v: v}
}
