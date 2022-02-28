package lib2

import "log"

//go:generate mockgen -source=lib2.go -destination=mocks/lib2.go -package=mocks

type Lib3 interface {
	Do() error
}

type Lib struct {
	Lib Lib3
}

func (l *Lib) Do() error {
	_ = l.Lib.Do()
	log.Println(`this is lib2`)
	return nil
}
