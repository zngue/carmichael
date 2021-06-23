package main

import (
	"testing"

	"github.com/zngue/go_helper/pkg"

	"github.com/zngue/go_helper/pkg/code"
)

func TestCode(t *testing.T) {
	pkg.NewConfig()
	code.CodeModelList()
}
