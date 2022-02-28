package lib1

import "log"

//go:generate mockgen -source=lib1.go -destination=mocks/lib1.go -package=mocks

type Lib2 interface {
	Do() error
}

type Lib struct {
	Lib Lib2
}

func (l *Lib) Do() error {
	_ = l.Lib.Do()
	log.Println(`this is lib1`)
	return nil
}
