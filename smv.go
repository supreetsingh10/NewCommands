package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// smv <files not to be copied> <final location>
// Example
// ~> smv ~/Download/lotr.txt ~/Music/lz.go ~/Backup

var directories []string
var doNotMoveFiles map[string]bool

func cliargs() (map[string]bool, string) {
	arglength := len(os.Args[0:])
	doNotMoveFiles = make(map[string]bool)

	// This function creates the hashmap for the command line arguments.
	// The hash map created will prevent selected files to be moved.
	for i := range os.Args[1:] {
		doNotMoveFiles[filepath.Dir(os.Args[i+1])+"/"+filepath.Base(os.Args[i+1])] = true
		if i != arglength-2 {
			directories = append(directories, filepath.Dir(os.Args[i+1]))
		}
	}
	return doNotMoveFiles, os.Args[len(os.Args)-1]
}

func moveIndividual(filename string, finalDir string) error {
	err := os.Rename(filename, finalDir+"/"+filename)
	return err
}

func validFinal(finalDir string) string {
	if string(finalDir[0]) == "~" || string(finalDir[0]) == "/" || string(finalDir[0]) == ".." {
		_, err := os.Stat(finalDir)
		if err == nil {
			return finalDir
		}
		log.Fatal(err)
	}
	wd, _ := os.Getwd()
	finalpath := wd + "/" + finalDir
	_, err := os.Stat(finalpath)
	if err != nil {
		log.Fatal(err)
	}
	return finalpath
}

func moveAllExcept(doNotMoveFiles map[string]bool, finalDir string) {
	for _, ffs := range directories {
		allFiles, err := ioutil.ReadDir(ffs)
		finalpath := validFinal(finalDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range allFiles {
			if doNotMoveFiles[ffs+"/"+f.Name()] {
				continue
			} else {
				err := moveIndividual(ffs+"/"+f.Name(), finalpath)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	doNotMoveFiles, finalDir := cliargs()
	moveAllExcept(doNotMoveFiles, finalDir)
	log.Print("Completed")
}
