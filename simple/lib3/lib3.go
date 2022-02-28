package lib3

import "log"

type Lib struct{}

func (l *Lib) Do() error {
	log.Println(`this is lib3`)
	return nil
}
