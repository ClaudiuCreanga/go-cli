package main
import (
        flag "github.com/jessevdk/go-flags"
        //"fmt"
        //"os"
        //"strings"
)


type Defaults struct {
        Verbose []bool `short:"v" long:"verbose" description:"Show verbose    debug information"`
        Terse   bool   `short:"t" long:"terse" description:"Shows terse output"`
}

type ListCommand struct {
        Config string `short:"c" long:"config" description:"config file" optional:"yes"`
}

func main() {
        defaultOptions := Defaults{}
        listCmd := ListCommand{}

        parser := flag.NewParser(&defaultOptions, flag.Default)
        parser.AddCommand("list", "lists something", "", &listCmd)

        parser.Parse()
}
