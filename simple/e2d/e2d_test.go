package e2d

import (
	"e2dlabo/simple/e2d/mocks"
	"e2dlabo/simple/lib1"
	"e2dlabo/simple/lib2"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)
	lib3Mock := mocks.NewMockLib3(ctrl)
	lib3Mock.EXPECT().Do().DoAndReturn(func() error {
		log.Println("this is mock")
		return nil
	})
	lib := lib1.Lib{
		Lib: &lib2.Lib{
			Lib: lib3Mock,
		},
	}
	err := lib.Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("yeah")
}
