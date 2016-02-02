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

	installGinkgo := exec.Command(
		goToolPath,
		"get", "github.com/onsi/ginkgo/ginkgo",
	)

	outBytes, err := installGinkgo.CombinedOutput()
	if err != nil {
		log.Fatalf("installing ginkgo: %s", string(outBytes))
	}

	variablesString := fmt.Sprintf("GOPATH=%s PATH=%s/bin:$PATH GO15VENDOREXPERIMENT=1", tmpDir, tmpDir)
	expandedVariablesString := os.ExpandEnv(variablesString)

	fmt.Println(expandedVariablesString)
	os.Exit(0)
}
