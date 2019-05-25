package controller

import (
	"github.com/app-sre/qontract-operator/pkg/controller/qontractapply"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, qontractapply.Add)
}
