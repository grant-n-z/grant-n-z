package endpoint

type TestIF interface {
	Init()
	Run()
	Close()
}
