package e2d

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go -package=mocks

type Lib3 interface {
	Do() error
}
