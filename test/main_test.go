package test

import (
	"github.com/aivyss/go-bean"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	bean.Clean()
}
