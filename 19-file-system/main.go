package main

import (
	"fmt"
	"io/fs"
	"os"
)

func getEmptyFiles(args []string) {
	files, err := fs.ReadDir(os.DirFS(args[0]), ".")
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, file := range files {
		info, error := file.Info()

		if error != nil {
			fmt.Println(error)
			return
		}

		if info.Size() == 0 {
			//1 byte for newline char
			total++
			total += len(file.Name())
		}
	}
	fmt.Printf("Total space needed: %d bytes\n", total)

	//optimisation of allocation of the slice
	names := make([]byte, 0, total)
	for _, file := range files {
		fmt.Println(file.Name())
		info, error := file.Info()

		if error != nil {
			fmt.Println(error)
			return
		}

		if info.Size() == 0 {
			name := file.Name()
			fmt.Println("Empty file:" + name)
			names = append(names, file.Name()...)
			names = append(names, '\n')
		}
	}

	// last arg = file permission
	err = os.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}

func printDirStructure(args []string, curDir string) {
	files, err := fs.ReadDir(os.DirFS(args[0]), ".")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			printDirStructure([]string{args[0] + "/" + file.Name()}, curDir+file.Name()+"/")
		} else {
			fmt.Println(curDir + file.Name())
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}
	getEmptyFiles(args)
	printDirStructure(args, "")
}
