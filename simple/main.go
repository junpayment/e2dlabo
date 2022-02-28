package main

import (
	"e2dlabo/simple/lib1"
	"e2dlabo/simple/lib2"
	"e2dlabo/simple/lib3"
	"log"
)

func main() {
	lib := lib1.Lib{
		Lib: &lib2.Lib{
			Lib: &lib3.Lib{},
		},
	}
	err := lib.Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("yeah")
}
