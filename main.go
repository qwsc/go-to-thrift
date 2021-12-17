package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"github.com/qwsc/go-to-thrift/generator"
)

// Arguments contains command line arguments
type Arguments struct {
	InputFile string

	OutputFile string
	AstFile    string
}

var (
	Args Arguments
)

func handleArgs() error {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.StringVar(&Args.OutputFile, "o", "./output.thrift", "")
	f.StringVar(&Args.AstFile, "a", "", "")
	if err := f.Parse(os.Args[1:]); err != nil {
		return err
	}
	rest := f.Args()
	if len(rest) != 1 {
		return errors.New("argument nums error")
	}
	Args.InputFile = rest[0]
	return nil
}

func main() {
	if err := handleArgs(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, Args.InputFile, nil, parser.ParseComments)
	if err != nil || astFile == nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if Args.AstFile != "" {
		builder := strings.Builder{}
		_ = ast.Fprint(&builder, fileSet, astFile, ast.NotNilFilter)
		if err = ioutil.WriteFile(Args.AstFile, []byte(builder.String()), 0644); err != nil { // ignore_security_alert
			fmt.Println(err)
			os.Exit(2)
		}
	}

	result := generator.NewGenerator(astFile).Generate()
	if err = ioutil.WriteFile(Args.OutputFile, []byte(result), 0644); err != nil { // ignore_security_alert
		fmt.Println(err)
		os.Exit(2)
	}
}
