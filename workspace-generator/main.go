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

	log.Printf("using go at %s", goBinary())

	tmpDir, err := ioutil.TempDir("", "generated-go-workspace")
	if err != nil {
		log.Fatalf("making workspace dir: %s", err)
	}

	os.Setenv("GOPATH", tmpDir)

	installPackage("github.com/golang/lint/golint")
	installPackage("github.com/onsi/ginkgo/ginkgo")

	fmt.Printf("export GOPATH=%s\n", tmpDir)
	fmt.Printf("export PATH=%s/bin:$PATH\n", tmpDir)
	fmt.Printf("export GO15VENDOREXPERIMENT=1\n")

	os.Exit(0)
}

func installPackage(packageName string) {
	log.Printf("Installing package: %s\n", packageName)
	installCmd := exec.Command(
		goBinary(),
		"get", packageName,
	)

	outBytes, err := installCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("installing %s: %s", packageName, string(outBytes))
	}
}

func goBinary() string {
	goToolPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatalln("finding go executable")
	}
	return goToolPath
}
