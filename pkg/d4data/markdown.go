package d4data

import (
	"d4md/pkg/d4data/exp"
	"d4md/pkg/markdown"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/windler/dotgraph/renderer"
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
		if err := fieldsToMarkdown(x.Fields, f, associations, g, "Fields"); err != nil {
			return err
		}

		// Quest phases
		if x.Type.Name == fileTypeQuestDef {
			if err := questToMarkdown(x.Fields, f, associations, g); err != nil {
				return err
			}
		}

		return nil
	case Object:
		return fieldsToMarkdown(x.Fields, f, associations, g, x.Type.String()) // TODO: add special cases for phases
	case Ref:
		switch x.Type.Name {
		case "DT_SNO":
			if linkedFile, ok := associations.Get(x.Raw); ok {
				baseDir, _ := filepath.Split(f.Name)
				relPath, err := filepath.Rel(baseDir, linkedFile.Name+".md")
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

const (
	fileTypeQuestDef = "QuestDefinition"

	qdStartPhase      = "unk_942bcdb" // Maybe
	qdPhases          = "arQuestPhases"
	phaseCallbackSets = "arCallbackSets"
	qdLink            = "ptLink"
	linkNextPhase     = "unk_d17aff0"
	qdPhaseId         = "dwUID"
)

func questToMarkdown(fields map[string]any, f *File, associations FilesById, g *markdown.Generator) error {
	if err := g.H2("Quest Details (WIP)"); err != nil {
		return err
	}

	if err := g.H3("Phase Order"); err != nil {
		return err
	}

	qd, err := exp.ParseQuestDef(f.Raw__.Raw)
	if err != nil {
		return err
	}

	r := &renderer.PNGRenderer{
		OutputFile: "graph.png",
	}
	r.Render(qd.Graph().String())

	spew.Dump(qd)
	return nil

	//// Map all quest phases
	//questPhaseById := make(map[int64]Object, 0)
	//phaseGraph := graph.New("Phase Order")
	//
	//phases, ok := fields[qdPhases].([]any)
	//if !ok {
	//	return fmt.Errorf("failed to parse quest details (a): %#v", phases)
	//}
	//
	//for _, phase := range phases {
	//	phaseObj, ok := phase.(Object)
	//	if !ok {
	//		return errors.New("failed to parse quest details (b)")
	//	}
	//
	//	currentPhase, ok := phaseObj.Fields[qdPhaseId].(float64)
	//	if !ok {
	//		return errors.New("failed to parse quest details (c)")
	//	}
	//
	//	for _, callbackSet := range phaseObj.Fields[phaseCallbackSets].([]any) {
	//		callbackSetObj := callbackSet.(Object)
	//
	//		links, ok := callbackSetObj.Fields[qdLink].([]any)
	//		if !ok {
	//			return errors.New("failed to parse quest details (d)")
	//		}
	//
	//		for _, link := range links {
	//			nextPhase, ok := link.(Object).Fields[linkNextPhase].(float64)
	//			if !ok {
	//				return errors.New("failed to parse quest details (e)")
	//			}
	//
	//			questPhaseById[int64(currentPhase)] = phaseObj
	//			phaseGraph.AddDirectedEdge(
	//				strconv.FormatFloat(currentPhase, 'f', -1, 64),
	//				strconv.FormatFloat(nextPhase, 'f', -1, 64),
	//				"", // TODO: add attributes
	//			)
	//		}
	//	}
	//}
	//
	//r := &renderer.PNGRenderer{
	//	OutputFile: "graph.png",
	//}
	//r.Render(phaseGraph.String())
	//
	////// Follow quest phase graph
	////curr := fields[qdStartPhase].(int64)
	////for {
	////	// Print current
	////	if err := ToMarkdown(questPhaseById[curr], f, associations, g); err != nil {
	////		return err
	////	}
	////
	////	// Get next
	////	curr, ok = edges[curr]
	////	if !ok {
	////		break
	////	}
	////}
	//
	//return nil
}
