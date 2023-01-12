package main

import (
	"github.com/google/go-cmdtest"
	"github.com/stretchr/testify/require"
	"os/exec"
	"regexp"
	"testing"
)

func TestCli(t *testing.T) {
	ts, err := cmdtest.Read("scripts")
	require.NoError(t, err)
	binary := "/Users/aricart/go/bin/nsc"
	ts.Commands["nsc"] = func(args []string, inputFile string) ([]byte, error) {
		cmd := exec.Command(binary, args...)
		out, err := cmd.CombinedOutput()
		require.NoError(t, err)
		out = filterNkeys(out)
		return out, err
	}

	ts.KeepRootDirs = false
	ts.Run(t, false)
}

func filterNkeys(data []byte) []byte {
	var re = regexp.MustCompile(`"S?[OAU][A-Z0-9]{55,56}"`)
	data = re.ReplaceAll(data, []byte(`"[HIDDEN KEY]"`))

	return data
}
