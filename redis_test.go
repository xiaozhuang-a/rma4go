package main

import (
	"github.com/winjeg/rma4go/cmder"

	"testing"
)

func TestPrintStat(t *testing.T) {
	printKeyStat(cmder.GetHost(), cmder.GetAuth(), cmder.GetPort())
}
