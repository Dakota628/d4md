package main

import (
	"d4md/pkg/d4data"
	"golang.org/x/exp/slog"
	"path/filepath"
)

const (
	parseWorkers    = 1000
	generateWorkers = 100
)

var (
	d4DataPath = filepath.Join(".", "d4data", "json")
)

func main() {
	// Parse the data into memory; yes this uses a lot of memory, but it's a lot faster than the alternatives
	files, err := d4data.Parse(d4DataPath, parseWorkers)
	if err != nil {
		slog.Error("Error parsing data path", err)
	}

	// Start markdown generation
	if err := d4data.GenerateMarkdown(files, generateWorkers); err != nil {
		slog.Error("Error generating markdown", err)
	}
}
