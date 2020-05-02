package controller

import (
	"github.com/BuddhiWathsala/helloworld-k8s-operator/pkg/controller/helloworld"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, helloworld.Add)
}
