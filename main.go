package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Masterminds/semver"
)

func main() {
	var (
		infile  = flag.String("i", "", "Input file (default: stdin)")
		outfile = flag.String("o", "", "Output file (default: stdout)")
		kind    = flag.String("kind", "patch", "Kind of update: major, minor, or patch (default: patch)")

		oldVersion string
	)
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *infile != "" {
		buf, err := ioutil.ReadFile(*infile)
		if err != nil {
			log.Fatal(err)
		}
		oldVersion = string(buf)
	} else {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		oldVersion = string(buf)
	}

	oldVersion = strings.TrimSpace(oldVersion)
	newVersion, err := bump(oldVersion, *kind)
	if err != nil {
		log.Fatal(err)
	}

	// Write to output
	if *outfile != "" {
		err = ioutil.WriteFile(*outfile, []byte(newVersion), 0644)
	} else {
		_, err = fmt.Fprintf(os.Stdout, "%s", newVersion)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of bump:\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "\tbump [flags]\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func bump(oldVersion string, kind string) (string, error) {
	var newv semver.Version

	hasVPrefix := strings.HasPrefix(oldVersion, "v")

	oldv, err := semver.NewVersion(oldVersion)
	if err != nil {
		return "", err
	}

	switch kind {
	default:
		return "", errors.New("Invalid kind")
	case "major":
		newv = oldv.IncMajor()
	case "minor":
		newv = oldv.IncMinor()
	case "patch":
		newv = oldv.IncPatch()
	}

	var sb strings.Builder
	if hasVPrefix {
		sb.WriteByte('v')
	}
	sb.WriteString(newv.String())
	return sb.String(), nil
}
