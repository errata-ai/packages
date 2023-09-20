package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmdtest"
)

func Test(t *testing.T) {
	styles, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	for _, style := range styles {
		if !style.IsDir() {
			continue
		}

		cases := filepath.Join(style.Name(), "testdata")
		if _, err := os.Stat(cases); os.IsNotExist(err) {
			fmt.Printf("skipping '%s': no testdata\n", style.Name())
			continue
		}

		doTest(cases, t)
	}
}

func doTest(name string, t *testing.T) {
	ts, err := cmdtest.Read(name)
	if err != nil {
		t.Fatal(err)
	}

	ts.Setup = func(_ string) error {
		_, testFileName, _, ok := runtime.Caller(0)
		if !ok {
			return fmt.Errorf("failed get real working directory from caller")
		}

		projectRootDir := filepath.Join(filepath.Dir(testFileName), name)
		if err := os.Setenv("ROOTDIR", projectRootDir); err != nil {
			return fmt.Errorf("failed change 'ROOTDIR' to caller working directory: %v", err)
		}

		return nil
	}

	path, err := exec.LookPath("vale")
	if err != nil {
		t.Fatal(err)
	}

	ts.Commands["vale"] = cmdtest.Program(path)
	ts.Commands["cdf"] = cmdtest.InProcessProgram("cdf", cdf)

	ts.Run(t, false)
}
