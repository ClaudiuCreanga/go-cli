package main
import (
        flag "github.com/jessevdk/go-flags"
        "fmt"
        //"os"
        //"strings"
		"os"
)


type Defaults struct {
        Verbose []bool `short:"v" long:"verbose" description:"Show verbose    debug information"`
        Terse   bool   `short:"t" long:"terse" description:"Shows terse output"`
}

type options struct {
        Config string `short:"c" long:"config" description:"config file"`
		Help     bool     `short:"h" long:"help" description:"show help message"`
}

func main() {
        defaultOptions := Defaults{}
        opts := options{}

        parser := flag.NewParser(&defaultOptions, flag.Default)
        parser.AddCommand("trash", "manage trash", "", &opts)

        parser.Parse()
        fmt.Println(opts.Config)

		_, err := parser.Parse()
		if err != nil {
			fmt.Printf("fail to parse args: %v", err)
			os.Exit(1)
		}

		if opts.Help {
			parser.WriteHelp(os.Stdout)
			os.Exit(0)
		}

        }
