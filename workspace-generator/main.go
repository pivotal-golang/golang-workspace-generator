package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Usage: `eval $(workspace-generator)`
func main() {
	// set log output to STDERR, because STDOUT is what the caller will eval.
	log.SetOutput(os.Stderr)

	tmpDir, err := ioutil.TempDir("", "generated-go-workspace")
	if err != nil {
		log.Fatalf("making workspace dir: %s", err.Error())
	}

	os.Setenv("GOPATH", tmpDir)

	err = installPackages([]string{
		"github.com/golang/lint/golint",
		"github.com/onsi/ginkgo/ginkgo",
	})
	if err != nil {
		log.Fatalf("failed to install packages: %s", err.Error())
	}

	fmt.Printf("export GOPATH=%s\n", tmpDir)
	fmt.Printf("export PATH=%s/bin:$PATH\n", tmpDir)
	fmt.Printf("export GO15VENDOREXPERIMENT=1\n")

	os.Exit(0)
}

func installPackages(packageNames []string) error {
	goBinary, err := exec.LookPath("go")
	if err != nil {
		return fmt.Errorf("Failed to locate go executable in PATH: %s", err.Error())
	}

	for _, packageName := range packageNames {
		log.Printf("Installing package: %s\n", packageName)
		installCmd := exec.Command(
			goBinary,
			"get", packageName,
		)

		outBytes, err := installCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("Failed to install %s:\nError: %s\nOutput: %s\n",
				packageName,
				err.Error(),
				string(outBytes),
			)
		}
	}

	return nil
}
