package main

import (
	"github.com/google/go-cmdtest"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCli(t *testing.T) {
	ts, err := cmdtest.Read("scripts")
	require.NoError(t, err)
	ts.Commands["nsc"] = cmdtest.Program("/Users/aricart/go/bin/nsc")
	ts.KeepRootDirs = true
	ts.Run(t, true)
}
