package main

import (
	"errors"
	"fmt"
	"os"
	"wgetNew/service"

	"github.com/spf13/pflag"
)

var (
	PathFlag      = pflag.StringP("path", "p", ".", "path to save files")
	DeepFlag      = pflag.IntP("deep", "d", 0, "recursion deep")
	MirroringFlag = pflag.BoolP("mirroring", "m", false, "mirror downloaded site")
)

func main() {
	pflag.Parse()

	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "to little args: main.go url path-to-save(default \".\") recursion-deep(default 0)")
		os.Exit(1)
	}

	u := args[1]
	path := *PathFlag
	deep := *DeepFlag

	// parse path and create folder if not exists
	if path != "." {
		err := os.Mkdir(path, 0755)
		if err != nil && !errors.Is(err, os.ErrExist) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if *MirroringFlag {
		err := os.Mkdir(path+"/static", 0755)
		if err != nil && !errors.Is(err, os.ErrExist) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// new entitie parser
	srv := service.NewService(path, *MirroringFlag)

	// start parsing
	err := srv.Start(u, deep)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *MirroringFlag {
		srv.Mirror()
	}
}
