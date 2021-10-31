package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("usage: %s PACKAGE_NAMES...", os.Args[0]))
	}

	rootPackages, err := getPackagesInfo(os.Args[1:])
	if err != nil {
		panic(err)
	}
	allPackages, err := getPackagesInfo([]string{"..."})
	if err != nil {
		panic(err)
	}

	allPackagesMap := map[string]PackageInfo{}
	for _, pkg := range allPackages {
		allPackagesMap[pkg.ImportPath] = pkg
	}
	for _, pkg := range rootPackages {
		printDeps(pkg, allPackagesMap, 0)
	}
}

func printDeps(pkg PackageInfo, allPackagesMap map[string]PackageInfo, depth int) {
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), pkg.ImportPath)
	for _, name := range pkg.Deps {
		pkg, ok := allPackagesMap[name]
		if !ok {
			continue
		}

		printDeps(pkg, allPackagesMap, depth+1)
	}
}

type PackageInfo struct {
	ImportPath string
	Deps       []string
}

func getPackagesInfo(names []string) ([]PackageInfo, error) {
	cmd := exec.Command("go", append([]string{"list", "-json", "-e"}, names...)...)
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(&stdout)
	packagesInfo := []PackageInfo{}
	for {
		info := PackageInfo{}
		err = dec.Decode(&info)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		packagesInfo = append(packagesInfo, info)
	}

	return packagesInfo, nil
}
