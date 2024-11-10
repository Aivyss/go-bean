package test

import (
	"go-bean"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	bean.Clean()
}
