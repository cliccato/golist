package main

import (
    "os"
    "fmt"
    "log"
    "io/ioutil"
)

const(
    colorReset = "\033[0m"
    colorBlue = "\033[34m"
    max_arg = 1
    max_line = 3
    max_len = 17
)

func main() {

    var dir = "." 
    var line = 0
    var j int
    var narg = len(os.Args[1:])
    log.SetFlags(0)

    if narg > max_arg {
        log.Fatal("Too many arguments")
    }
    if narg == 1 {
        dir = os.Args[1]
    }

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fname := file.Name()
        if line == max_line {
            line = 0
	    fmt.Printf("\n")
        }
        if string(fname[0:1]) != "." {
            cdir, err := os.Stat(fname)
            if err != nil {
		log.Fatal(err)
	    }

            if cdir.IsDir() {
                fmt.Printf(string(colorBlue))
                fmt.Printf("%s\t", fname)

                fmt.Printf(string(colorReset))
            }else {
                fmt.Printf("%s\t", fname)
            }
        line++
        }
    }
    fmt.Printf("\n")
}
