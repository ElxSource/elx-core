package modules

import (
	"reflect"

	"github.com/jose78/gansible/conf"
)

type Debug struct {
	msg string
}

func init() {
	instance := conf.GetInstance()
	instance.Set("debug", reflect.TypeOf(new(Debug)))
}
