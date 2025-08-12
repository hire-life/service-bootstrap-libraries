package arch

type Application interface {
	Module
	Register()
	Run()
}
