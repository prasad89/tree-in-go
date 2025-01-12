package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}
	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil
	}
	fis, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}
	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}
	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Printf("%s", indent + "└──")
			add = "   "
		} else {
			fmt.Printf("%s", indent + "├──")
		}
		if err := tree(filepath.Join(root, name), indent+add); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, arg := range args {
		if err := tree(arg, ""); err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}
