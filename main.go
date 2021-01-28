package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	needed := []string{"fzf", "git", "go", "kubernetes-cli", "kustomize", "node", "r", "zsh-autosuggestions", "zsh", "awscli", "terraform", "java", "helm"}
	var answer string

	for {
		diff := diff(getLeaves(), needed)

		if len(diff) == 0 {
			fmt.Println("No packages to uninstall")
			os.Exit(0)
		}

		fmt.Println("Plan to uninstall: \n", diff, "\nType 'yes' if you agree")

		fmt.Scanln(&answer)
		if answer != "yes" {
			os.Exit(1)
		}
		for _, pkg := range diff {
			uninstall(pkg)
		}
	}
}
func getLeaves() []string {
	out, err := exec.Command("brew", "leaves").CombinedOutput()

	if err != nil {
		fmt.Println("cmd.Run() failed with", err)
		os.Exit(1)
	}
	return strings.Split(string(out), "\n")
}

func uninstall(pkg string) {
	fmt.Println("Trying to uninstall", pkg)
	out, err := exec.Command("brew", "uninstall", "--force", pkg).CombinedOutput()
	if err != nil {
		fmt.Println("failed to uninstall ", pkg, " with error ", err)
		os.Exit(2)
	}
	fmt.Println(string(out))
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
