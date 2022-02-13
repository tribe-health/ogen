// Binary ogen generates go source code from OAS.
package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen"
)

type formattedSource struct {
	Format bool
	Root   string
}

func (t formattedSource) WriteFile(name string, content []byte) error {
	out := content
	if t.Format {
		buf, err := format.Source(content)
		if err != nil {
			return err
		}
		out = buf
	}
	return os.WriteFile(filepath.Join(t.Root, name), out, 0600)
}

func main() {
	var (
		specPath       = flag.String("schema", "", "Path to openapi spec file")
		targetDir      = flag.String("target", "api", "Path to target dir")
		packageName    = flag.String("package", "api", "Target package name")
		performFormat  = flag.Bool("format", true, "perform code formatting")
		filterPath     = flag.String("filter-path", "", "Filter operations by path regex")
		filterMethods  = flag.String("filter-methods", "", "Filter operations by HTTP methods (comma-separated)")
		clean          = flag.Bool("clean", false, "Clean generated files before generation")
		verbose        = flag.Bool("v", false, "verbose")
		generateTests  = flag.Bool("generate-tests", false, "Generate tests based on schema examples")
		skipTestsRegex = flag.String("skip-tests", "", "Skip tests matched by regex")
		inferTypes     = flag.Bool("infer-types", false, "Infer schema types, if type is not defined explicitly")

		debugIgnoreNotImplemented = flag.String("debug.ignoreNotImplemented", "", "Ignore methods having functionality which is not implemented (all, oneOf, anyOf, allOf, nullable types, complex parameter types)")
		debugNoerr                = flag.Bool("debug.noerr", false, "Ignore all errors")
	)

	flag.Parse()
	if *specPath == "" {
		panic("no spec provided")
	}
	data, err := os.ReadFile(*specPath)
	if err != nil {
		panic(err)
	}

	spec, err := ogen.Parse(data)
	if err != nil {
		panic(err)
	}
	files, err := os.ReadDir(*targetDir)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(*targetDir, 0750); err != nil {
			panic(err)
		}
	}
	if *clean {
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			name := f.Name()
			if !strings.HasSuffix(name, "_gen.go") {
				continue
			}
			if !(strings.HasPrefix(name, "openapi") || strings.HasPrefix(name, "oas")) {
				continue
			}
			if err := os.Remove(filepath.Join(*targetDir, name)); err != nil {
				panic(err)
			}
		}
	}

	fs := formattedSource{
		Root:   *targetDir,
		Format: *performFormat,
	}

	var filters gen.Filters
	{
		if filterPath != nil {
			filters.PathRegex, err = regexp.Compile(*filterPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid filter.path flag value: %s", err)
				return
			}
		}

		if filterMethods != nil {
			for _, m := range strings.Split(*filterMethods, ",") {
				m = strings.TrimSpace(m)
				if m == "" {
					continue
				}
				filters.Methods = append(filters.Methods, m)
			}
		}
	}

	opts := gen.Options{
		VerboseRoute:         *verbose,
		GenerateExampleTests: *generateTests,
		SkipTestRegex:        nil,
		InferSchemaType:      *inferTypes,
		Filters:              filters,
		IgnoreNotImplemented: strings.Split(*debugIgnoreNotImplemented, ","),
	}
	if expr := *skipTestsRegex; expr != "" {
		r, err := regexp.Compile(expr)
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
		opts.SkipTestRegex = r
	}

	if *debugNoerr {
		opts.IgnoreNotImplemented = []string{"all"}
	}

	g, err := gen.NewGenerator(spec, opts)
	if err != nil {
		panic(fmt.Sprintf("%s: %+v", *specPath, err))
	}

	if err := g.WriteSource(fs, *packageName); err != nil {
		panic(fmt.Sprintf("%s: %+v", *specPath, err))
	}
}
