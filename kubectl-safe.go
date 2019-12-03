package main

import (
	"fmt"
	"os"
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
		//log.Println(err)
	}
	fmt.Printf("Found kubectl here: %s\n", kubectlBinary)
	return kubectlBinary
}

func main() {
	// argsWithoutProg := os.Args[1:]

	// fmt.Println(argsWithoutProg)
	findKubectlBinary()

}

/*

kubectlbin="/usr/bin/kubectl"

if [ -z $1 ] || [ "x$1" == "x" ]; then
  exit 0
fi

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
