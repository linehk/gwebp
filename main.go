package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	root = flag.String("r", ".", "root directory")
	ext  = flag.String("e", "jpg", "filename extension")
	args = flag.String("a", "-q 50", "cwebp arguments")
	keep = flag.Bool("k", true, "keep")
)

func main() {
	flag.Parse()
	err := filepath.Walk(*root, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk err: %v", err)
		return
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if !info.IsDir() && strings.Contains(info.Name(), "."+*ext) {
		newPath := strings.TrimSuffix(path, "."+*ext) + ".webp"
		allArgs := append([]string{path}, strings.Split(*args, " ")...)
		allArgs = append(allArgs, "-o", newPath)
		cmd := exec.Command("cwebp", allArgs...)
		err := cmd.Run()
		if err != nil {
			return err
		}
		if !*keep {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
