package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.SetOutput(os.Stderr)

	goToolPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatalln("finding go executable")
	}

	log.Printf("using go at %s", goToolPath)

	tmpDir, err := ioutil.TempDir("", "generated-go-workspace")
	if err != nil {
		log.Fatalf("making workspace dir: %s", err)
	}

	os.Setenv("GOPATH", tmpDir)

	installPackage("github.com/golang/lint/golint")
	installPackage("github.com/onsi/ginkgo/ginkgo")

	variablesString := fmt.Sprintf("GOPATH=%s PATH=%s/bin:$PATH GO15VENDOREXPERIMENT=1", tmpDir, tmpDir)
	expandedVariablesString := os.ExpandEnv(variablesString)

	fmt.Println(expandedVariablesString)
	os.Exit(0)
}

func installPackage(packageName string) {
	installCmd := exec.Command(
		goToolPath,
		"get", packageName,
	)

	outBytes, err := installCmd.CombinedOutput()
	if err != nil {
		log.Fatalf("installing %s: %s", packageName, string(outBytes))
	}
}
