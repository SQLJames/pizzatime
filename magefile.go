//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	helpers "pizzatime/internal/mage"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	goexe = "go"
	dirs  = []string{"bin"}

	targets = []helpers.Target{
		{GOOS: "linux", GOARCH: "amd64"},
		{GOOS: "darwin", GOARCH: "amd64"},
	}
)

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}
}

func flags() string {
	f := helpers.LDFlags{}

	f["version.ApplicationName"] = filepath.Base(helpers.ModulePath())
	f["version.BuildDate"] = time.Now().Format(time.RFC3339)
	f["version.BuildTag"] = helpers.GitTag()
	f["version.CommitHash"] = helpers.GitCommitHash()

	return f.String()
}

func ensureDirs() error {
	fmt.Println("--> Ensuring output directories")

	for _, dir := range dirs {
		if !helpers.FileExists("./" + dir) {
			fmt.Printf("    creating './%s'\n", dir)
			if err := os.MkdirAll("./"+dir, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}

// Clean up after yourself
func Clean() {
	fmt.Println("--> Cleaning output directories")

	for _, dir := range dirs {
		fmt.Printf("    removing './%s'\n", dir)
		os.RemoveAll("./" + dir)
	}
}

// Vendor dependencies with go modules
func Vendor() {
	fmt.Println("--> Updating dependencies")
	sh.Run(goexe, "mod", "tidy")
}
func commands() []string {
	c := []string{}

	if files, err := ioutil.ReadDir("./cmd"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				c = append(c, file.Name())
			}
		}
	}

	return c
}

// Build the application for local running
func Build() error {
	mg.SerialDeps(Vendor, ensureDirs)

	for _, command := range commands() {
		fmt.Printf("--> Building '%s'\n", command)

		binaryPath := filepath.Join("./bin", command)
		sourcePath := filepath.Join(helpers.ModulePath(), "/cmd", command)
		if err := sh.Run(goexe, "build", "-o", binaryPath, "-ldflags="+flags(), sourcePath); err != nil {
			return err
		}
	}

	return nil

}

// Release the application for all defined targets
func Release() error {
	mg.SerialDeps(Vendor, ensureDirs)

	cmds := commands()

	cgoEnabled := os.Getenv("CGO_ENABLED") == "1"

	var wg sync.WaitGroup
	wg.Add(len(targets) * len(cmds))
	for _, c := range cmds {
		fmt.Printf("--> Building '%s' for release\n", c)
		for _, t := range targets {
			t.SourceDir = c
			go func(t helpers.Target) {
				defer wg.Done()

				env := map[string]string{
					"GOOS":   t.GOOS,
					"GOARCH": t.GOARCH,
				}

				if cgoEnabled && runtime.GOOS != env["GOOS"] {
					fmt.Printf("      CGO is enabled, skipping compilation of %s for %s\n", t.Name(), env["GOOS"])
					return
				}
				fmt.Printf("      Building %s\n", t.Name())

				binaryPath := filepath.Join("./bin", t.Name())
				sourcePath := filepath.Join(helpers.ModulePath(), "/cmd", t.SourceDir)

				err := sh.RunWith(env, goexe, "build", "-o", binaryPath, "-ldflags="+flags(), sourcePath)
				if err != nil {
					fmt.Printf("compilation failed: %s\n", err.Error())
					return
				}

			}(t)
		}
	}
	wg.Wait()

	return nil
}

// Lint the codebase, checking for common errors
func Lint() {
	fmt.Println("--> Linting codebase")

	c := exec.Command("gometalinter", "-e", "internal", "-e", "go/pkg/mod", "./...")
	c.Env = os.Environ()
	out, err := c.CombinedOutput()
	if err == nil {
		fmt.Println("    no issues detected")
	} else {
		fmt.Print("    ")
		fmt.Println(strings.Replace(string(out), "\n", "\n    ", -1))
	}
}

// Test the codebase
func Test() error {
	mg.SerialDeps(Vendor, ensureDirs)

	fmt.Println("--> Testing codebase")
	results, err := sh.Output(goexe, "test", "-cover", "-e", "internal", "-e", "cache", "./...")
	fmt.Print("    ")
	fmt.Println(strings.Replace(results, "\n", "\n    ", -1))

	return err
}
