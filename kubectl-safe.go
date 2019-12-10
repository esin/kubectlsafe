package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func getEnvironmentPath() string {
	return os.Getenv("PATH")
}

func boldString(inString string) string {
	return "\033[1m" + inString + "\033[0m"
}

func findKubectlBinary() string {
	pathDirs := strings.Split(getEnvironmentPath(), ":")
	kubectlBinary := ""
	for _, envPath := range pathDirs {

		filepath.Walk(envPath, func(fullFileName string, info os.FileInfo, err error) error {
			if info != nil {
				if !info.IsDir() {
					if path.Base(fullFileName) == "kubectl" || path.Base(fullFileName) == "kubectl.exe" {
						kubectlBinary = fullFileName
						return nil
					}
				}
			}
			return nil
		})
	}
	return kubectlBinary
}

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	var writeCommands = []string{"apply", "create", "delete", "edit", "patch", "replace", "scale"}

	argsList := os.Args[1:]

	for _, kubeCommand := range writeCommands {

		// for example kubectl apply ......
		if kubeCommand == argsList[0] {
			kubectlBin := findKubectlBinary()

			if kubectlBin == "" {
				os.Exit(1)
			}

			// Get current context. TODO: try to use config parser from kubectl
			cmd := exec.Command(kubectlBin, "config1", "current-context")
			var outb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				os.Exit(1)
			}

			currentContext := strings.TrimSuffix(outb.String(), "\n")

			var userChoise string
			fmt.Printf("Current context is %s. Show must go on? [y/N] ", currentContext) // for Windows
			n, err := fmt.Scanf("%s", &userChoise)
			if err != nil || n != 1 {
				os.Exit(1)
			}

			if (strings.ToLower(userChoise) == "y") || (strings.ToLower(userChoise) == "yes") {
				cmd := exec.Command(kubectlBin, argsList...)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				cmd.Run()
			}
			os.Exit(0)
		}
	}
}
