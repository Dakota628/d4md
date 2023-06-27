package d4data

import (
	"errors"
	"github.com/alphadose/haxmap"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

const (
	dataFilePattern = "*SecretCellar*.qst.json" // TODO: back to *.json
)

type FilesById = *haxmap.Map[int64, *File]

func parseWorker(
	c chan string,
	wg *sync.WaitGroup,
	files FilesById,
	progress *atomic.Uint64,
	total uint64,
) {
	defer wg.Done()

	for {
		filePath, ok := <-c
		if !ok {
			return
		}

		// Open the file
		dataFile, err := os.Open(filePath)
		if err != nil {
			slog.Error("Failed to open data file %q", err, slog.String("filePath", filePath)) // TODO: probably want to write out to channels for multierr
			_ = dataFile.Close()
			continue
		}

		// Get json bytes
		jsonBytes, err := io.ReadAll(dataFile)
		_ = dataFile.Close()
		if err != nil {
			slog.Error("Failed to parse json for  %q", err, slog.String("filePath", filePath)) // TODO: probably want to write out to channels for multierr
			continue
		}

		// Parse json
		json := gjson.ParseBytes(jsonBytes)

		// Parse file
		file, err := NewFile(json)
		if err != nil {
			slog.Error("Failed to parse data file", err, slog.String("filePath", filePath)) // TODO: probably want to write out to channels for multierr
			continue
		}

		slog.Info(
			"Parsed data file",
			slog.String("filePath", filePath),
			slog.Uint64("progress", progress.Add(1)),
			slog.Uint64("total", total),
		)

		files.Set(file.SnoId, file)
	}
}

func Parse(dataPath string, workers uint) (FilesById, error) {
	// Assure workers >0
	if workers < 1 {
		return nil, errors.New("must have at least one worker")
	}

	// Get all data file names
	matchesArr, err := doublestar.FilepathGlob(filepath.Join(dataPath, "**", dataFilePattern))
	if err != nil {
		return nil, err
	}

	// Add matches to a buffered channel
	total := uint64(len(matchesArr))
	matches := make(chan string, total)
	for _, match := range matchesArr {
		matches <- match
	}
	close(matches)

	// Spawn workers to parse the files
	wg := &sync.WaitGroup{}
	files := haxmap.New[int64, *File]()
	progress := &atomic.Uint64{}

	for i := uint(0); i < workers; i++ {
		wg.Add(1)
		go parseWorker(matches, wg, files, progress, total)
	}

	wg.Wait()

	return files, nil
}
