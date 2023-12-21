package proof

import (
	"reflect"
	"testing"
)

type Proof struct {
	r *testing.T
	s any
	o any
}

func New(t *testing.T) *Proof {
	return &Proof{r: t}
}

func (p *Proof) Prove(anything any) *Proof {
	p.s = anything
	return p
}

func (p *Proof) Equal(anything any) {
	p.o = anything
	if p.s == nil {
		p.r.Fail()
	}

	eq := areEqual(p)
	if eq == false {
		p.r.Fail()
	}
}

func areEqual(p *Proof) bool {
	if isNil(p.s) && isNil(p.o) {
		return true
	}

	if isNil(p.s) || isNil(p.o) {
		return false
	}

	reflectS := reflect.ValueOf(p.s)
	reflectO := reflect.ValueOf(p.o)
	return reflectS.Equal(reflectO)
}

func isNil(obj any) bool {
	return obj == nil
}
