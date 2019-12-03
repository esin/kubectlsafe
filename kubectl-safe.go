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
	//fmt.Printf("Found kubectl here: %s\n", kubectlBinary)
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

			cmd := exec.Command(kubectlBin, "config", "current-context")
			var outb, errb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = &errb
			err := cmd.Run()
			if err != nil {
				fmt.Print(errb.String())
				os.Exit(2)
			}

			fmt.Println("out:", outb.String(), "err:", errb.String())
		}
	}

}

/*

// kubectlbin="/usr/bin/kubectl"

// if [ -z $1 ] || [ "x$1" == "x" ]; then
//   exit 0
// fi

if [ $1 == "apply" ] || [ $1 == "create" ] || [ $1 == "delete" ] || [ $1 == "edit" ] || [ $1 == "patch" ] || [ $1 == "replace" ]; then
  cc=$($kubectlbin config current-context);
  read -p $'Current context is \e[1m'$cc$'\e[0m. Show must go on? [y/N] ' -n1 -r
  echo
  if [[ $REPLY =~ ^[Yy]$ ]]; then
    $kubectlbin $@
    exit 0
  fi
else
  $kubectlbin $@
fi */
