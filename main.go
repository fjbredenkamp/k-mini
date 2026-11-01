package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
)

func versionString(version, commit string) string {
	if commit == "" {
		return fmt.Sprintf("k-mini %s", version)
	}
	return fmt.Sprintf("k-mini %s (%s)", version, commit)
}

func main() {
	fmt.Println(versionString(version, commit))
}
