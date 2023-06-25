package d4data

import (
	"d4md/pkg/markdown"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func generateMarkdownWorker(
	c chan *File,
	files FilesById,
	wg *sync.WaitGroup,
	progress *atomic.Uint64,
	total uint64,
) {
	defer wg.Done()

	for {
		file, ok := <-c
		if !ok {
			return
		}

		mdFilePath := filepath.Join(".", file.Name+".md")
		mdDir, _ := filepath.Split(mdFilePath)

		if err := os.MkdirAll(mdDir, 0700); err != nil {
			slog.Error("Error writing markdown file", err, slog.String("mdFilePath", mdFilePath))
		}

		mdFile, err := os.Create(mdFilePath)
		if err != nil {
			slog.Error("Error writing markdown file", err, slog.String("mdFilePath", mdFilePath))
		}

		err = ToMarkdownWriter(file, files, mdFile)
		if err != nil {
			slog.Error("Error writing markdown file", err, slog.String("mdFilePath", mdFilePath))
		}
		_ = mdFile.Close()

		slog.Info(
			"Generated markdown file",
			slog.String("mdFilePath", mdFilePath),
			slog.Uint64("progress", progress.Add(1)),
			slog.Uint64("total", total),
		)
	}
}

func GenerateMarkdown(files FilesById, workers uint) error {
	// Assure workers >0
	if workers < 1 {
		return errors.New("must have at least one worker")
	}

	// Add files to channel
	c := make(chan *File, 1000)

	go func() {
		files.ForEach(func(_ int64, file *File) bool {
			c <- file
			return true
		})
		close(c)
	}()

	// Start generator workers
	wg := &sync.WaitGroup{}
	progress := &atomic.Uint64{}
	total := uint64(files.Len())

	for i := uint(0); i < workers; i++ {
		wg.Add(1)
		go generateMarkdownWorker(c, files, wg, progress, total)
	}

	wg.Wait()
	return nil
}

func ToMarkdownWriter(f *File, associations FilesById, w io.Writer) error {
	g := markdown.NewGenerator(w)

	return ToMarkdown(f, f, associations, g)
}

func ToMarkdown(v any, f *File, associations FilesById, g *markdown.Generator) error {
	switch x := v.(type) {
	case *File:
		// Header
		if err := g.H1(x.Name); err != nil {
			return err
		}

		// Metadata
		if err := g.TableWideHeader("Metadata"); err != nil {
			return err
		}

		if err := g.TableRow(
			func() error {
				return g.Bold("Name")
			},
			func() error {
				return g.Write(x.Name)
			},
		); err != nil {
			return err
		}

		if err := g.TableRow(
			func() error {
				return g.Bold("Type")
			},
			func() error {
				return g.Write(x.Type.String())
			},
		); err != nil {
			return err
		}

		if err := g.TableRow(
			func() error {
				return g.Bold("SNO ID")
			},
			func() error {
				return g.Write(strconv.FormatInt(x.SnoId, 10))
			},
		); err != nil {
			return err
		}

		if err := g.TableFooter(); err != nil {
			return err
		}

		// Fields
		return fieldsToMarkdown(x.Fields, f, associations, g, "Fields")
	case Object:
		return fieldsToMarkdown(x.Fields, f, associations, g, x.Type.String()) // TODO: add special cases for phases
	case Ref:
		switch x.Type.Name {
		case "DT_SNO":
			if linkedFile, ok := associations.Get(x.Raw); ok {
				baseDir, _ := filepath.Split(f.Name)
				relPath, err := filepath.Rel(baseDir, linkedFile.Name)
				if err != nil {
					return err
				}

				return g.Link(x.String(), relPath)
			}
		case "DT_SNO_NAME":
			// TODO
		case "DT_GBID":
			// TODO
		}
		return g.Link(x.String(), "#UKNOWN")
	case Vector2:
		return g.Write(x.String())
	case Vector3:
		return g.Write(x.String())
	case Type:
		return g.Write(x.String())
	case string:
		if strings.ContainsRune(x, '\n') {
			return g.CodeBlock(x, "lua") // It's probably a script
		}
		return g.InlineCode(x)
	case int64:
		return g.InlineCode(strconv.FormatInt(x, 10))
	case float64:
		return g.InlineCode(strconv.FormatFloat(x, 'f', -1, 64))
	case bool:
		return g.InlineCode(strconv.FormatBool(x))
	case map[string]any:
		return fieldsToMarkdown(x, f, associations, g, "Fields")
	case []any:
		for _, value := range x {
			if value == nil {
				continue
			}
			if err := ToMarkdown(value, f, associations, g); err != nil {
				return err
			}
			if err := g.Write("\n"); err != nil {
				return err
			}
		}
	case nil:
		break // Don't display nil
	default:
		return fmt.Errorf("cannot convert %T to markdown", x)
	}

	return nil
}

func fieldsToMarkdown(
	x map[string]any,
	f *File,
	associations FilesById,
	g *markdown.Generator,
	tableHeader string,
) error {
	if err := g.TableWideHeader(tableHeader); err != nil {
		return err
	}

	for key, value := range x {
		if value == nil {
			continue
		}

		if err := g.TableRow(
			func() error {
				return g.Bold(key)
			},
			func() error {
				return ToMarkdown(value, f, associations, g)
			},
		); err != nil {
			return err
		}
	}

	if err := g.TableFooter(); err != nil {
		return err
	}

	return nil
}
