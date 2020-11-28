package main

import (
	"flag"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ernado/tl"

	"github.com/ernado/td/gen"
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
	return ioutil.WriteFile(filepath.Join(t.Root, name), out, 0600)
}

func main() {
	schemaPath := flag.String("schema", "", "Path to .tl file")
	targetDir := flag.String("target", "td", "Path to target dir")
	packageName := flag.String("package", "td", "Target package name")
	performFormat := flag.Bool("format", true, "perform code formatting")
	clean := flag.Bool("clean", false, "Clean generated files before generation")
	flag.Parse()
	if *schemaPath == "" {
		panic("no schema provided")
	}
	f, err := os.Open(*schemaPath)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	schema, err := tl.Parse(f)
	if err != nil {
		panic(err)
	}
	files, err := ioutil.ReadDir(*targetDir)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if os.IsNotExist(err) {
		if err := os.Mkdir(*targetDir, 0600); err != nil {
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
			if !strings.HasPrefix(name, "tl_") && name != "rpc_gen.go" {
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
	g, err := gen.NewGenerator(schema)
	if err != nil {
		panic(err)
	}
	if err := g.WriteSource(fs, *packageName, gen.Template()); err != nil {
		panic(err)
	}
}