package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"

	"github.com/olafal0/configinator/configinator"
)

func main() {
	specFile := flag.String("specfile", "", "The path of the config spec file (in toml format)")
	outDir := flag.String("outdir", "", "The output directory for generated code")
	outFile := flag.String("outfile", "config.go", "The output go file for generated code")
	flag.Parse()

	configCtx, err := configinator.ConfigCtxFromFile(*specFile)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = configinator.ExecuteTemplate(buf, configCtx)
	if err != nil {
		log.Fatal(err)
	}
	if *outDir == "" {
		saveFormatted(*outFile, buf)
	} else {
		saveFormatted(filepath.Join(*outDir, *outFile), buf)
	}
}

func saveFormatted(filename string, content *bytes.Buffer) {
	formatted, err := format.Source(content.Bytes())
	if err != nil {
		fmt.Println(content.String())
		log.Fatal(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		log.Fatal(err)
	}
}
