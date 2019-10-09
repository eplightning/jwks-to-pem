package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s [options] file_or_url\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	output := flag.String("output", "", "file to write to instead of STDOUT")
	format := flag.String("format", "pkix", "output PEM type: pkix, pkcs1")
	flag.Parse()

	if len(flag.Arg(0)) == 0 {
		fmt.Fprintln(os.Stderr, "input file path or URL is required")
		os.Exit(64)
	}

	if *format != "pkix" && *format != "pkcs1" {
		fmt.Fprintln(os.Stderr, "unknown format")
		os.Exit(64)
	}

	writer := os.Stdout
	if len(*output) > 0 {
		var err error
		writer, err = os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open specified output file for writing: %v\n", err)
		}
		defer writer.Close()
	}

	err := writePEM(flag.Arg(0), writer, *format)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func writePEM(input string, output io.Writer, format string) error {
	json, err := fetchJWKSData(input)
	if err != nil {
		return err
	}

	jwk, err := convertToJWK(json)
	if err != nil {
		return err
	}

	pem, err := convertToPEM(jwk, format)
	if err != nil {
		return err
	}

	output.Write(pem)
	return nil
}
