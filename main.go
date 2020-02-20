package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	needed := []string{"fzf", "git", "go", "kubernetes-cli", "kustomize", "node", "r", "zsh-autosuggestions", "zsh"}
	cmd := exec.Command("brew", "leaves")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal("cmd.Run() failed with %s\n", err)
	}
	leaves := strings.Split(string(out), "\n")
	diff := diff(leaves, needed)

	for _, pkg := range diff {
		uninstall(pkg)
	}
}

func uninstall(pkg string) {
	fmt.Printf("Trying to uninstall %s...", pkg)
	out, err := exec.Command("brew", "uninstall", pkg).CombinedOutput()
	if err != nil {
		log.Fatal("failed to uninstall %s with error %s\n", pkg, err)
	}
	fmt.Printf("%s", out)
}

func contains(str string, arr []string) bool {
	for _, x := range arr {
		if str == x {
			return true
		}
	}
	return false
}

func diff(arr1 []string, arr2 []string) []string {
	res := []string{}
	for _, x := range arr1 {
		if !contains(x, arr2) && len(x) > 1 {
			res = append(res, x)
		}
	}
	return res
}
