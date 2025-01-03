package lib

var dotconfname string

type Variables struct {
	Dotconfname string
}

func Initialize(vars Variables) {
	dotconfname = vars.Dotconfname
}
