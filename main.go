package main

import "flag"
import "fmt"
import "golang.org/x/tools/cover"
import "os"
import "text/template"

func main() {
	flagCoverProfile := flag.String("coverprofile", "", "Coverage profile")
	flagFormat := flag.String("format", "{{ .FileName }};{{ .Coverage }}", "Output format")

	flag.Parse()
	if *flagCoverProfile == "" || *flagFormat == "" {
		flag.Usage()
		os.Exit(-1)
	}

	tpl, err := template.New("out").Parse(*flagFormat)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid format:", err)
		os.Exit(-1)
	}

	profiles, err := cover.ParseProfiles(*flagCoverProfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid coverage profiles:", err)
		os.Exit(-1)
	}

	for _, prof := range profiles {
		var count float32
		for _, b := range prof.Blocks {
			if b.Count > 0 {
				count++
			}
		}

		coverage := (count / float32(len(prof.Blocks))) * float32(100)
		fc := fileCoverage{
			FileName: prof.FileName,
			Coverage: coverage,
		}
		if err = tpl.Execute(os.Stdout, fc); err != nil {
			fmt.Fprintln(os.Stderr, "Unable to write formatted output:", err)
		}

		fmt.Fprintln(os.Stdout)
	}
}

type fileCoverage struct {
	FileName string
	Coverage float32
}
