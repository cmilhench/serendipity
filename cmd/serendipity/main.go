package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/cmilhench/serendipity/pkg"
)

var revision = "latest"

func main() {
	global := flag.NewFlagSet("serendipity", flag.ExitOnError)
	global.Usage = func() {
		fmt.Fprintf(global.Output(), "usage: %s <command> [<args>]\n", global.Name())
		global.PrintDefaults()
	}
	global.Parse(os.Args[1:])
	if len(global.Args()) == 0 {
		global.Usage()
		os.Exit(1)
	}

	log.SetFlags(log.LstdFlags)

	switch global.Args()[0] {
	case "person":
		personCommand()
	default:
		fmt.Printf("%q is not valid command.\n", global.Args()[0])
		os.Exit(2)
	}
}

func personCommand() {
	command := flag.NewFlagSet("person", flag.ExitOnError)
	command.Usage = func() {
		fmt.Fprintf(command.Output(), "usage: %s %s <command> [<args>]\n", os.Args[0], command.Name())
		fmt.Fprintf(command.Output(), "Creates random user profiles\n")
		command.PrintDefaults()
	}
	count := command.Uint("n", 1, "Number of people")
	seed := command.String("s", "", "Random seed")
	command.Parse(os.Args[2:])
	if len(command.Args()) != 0 {
		command.Usage()
		os.Exit(1)
	}

	h := md5.New()
	h.Write([]byte(*seed))
	num := binary.BigEndian.Uint64(h.Sum(nil))

	r := pkg.New()
	r.Seed(int64(num))

	var p *pkg.Person

	fmt.Print("[\n")
	for i := 0; i < int(*count); i++ {
		p = r.Person()
		buf, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}
		if i != 0 {
			fmt.Print(",\n")
		}
		fmt.Printf("%s", buf)
	}
	fmt.Print("\n]")
}
