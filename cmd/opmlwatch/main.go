package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/config"
	"github.com/opmlwatch/opmlwatch/internal/core"
	"github.com/opmlwatch/opmlwatch/internal/fetcher"
	"github.com/opmlwatch/opmlwatch/internal/opml"
	"github.com/opmlwatch/opmlwatch/internal/rules"
	"github.com/opmlwatch/opmlwatch/internal/sinks"
	"github.com/opmlwatch/opmlwatch/internal/state"
	"github.com/opmlwatch/opmlwatch/internal/summarizers"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		runCmd(os.Args[2:])
	case "validate-config":
		validateConfigCmd(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}

func runCmd(args []string) {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	configPath := fs.String("config", "config.yaml", "Path to config YAML")
	dryRun := fs.Bool("dry-run", false, "Do not perform network sink delivery")
	_ = fs.Parse(args)

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := config.Validate(cfg); err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	feeds, err := opml.LoadFeeds(cfg.Feeds.OPMLPath)
	if err != nil {
		log.Fatalf("load OPML: %v", err)
	}

	now := time.Now().UTC()
	f := fetcher.New(cfg.Feeds.RequestTimeoutSeconds)
	entries, err := f.Fetch(context.Background(), feeds, fetcher.Options{
		MaxItemsPerFeed: cfg.Feeds.MaxItemsPerFeed,
		TimeoutSeconds:  cfg.Feeds.RequestTimeoutSeconds,
	})
	if err != nil {
		log.Fatalf("fetch entries: %v", err)
	}

	seenStore, err := state.New(cfg.Store.SeenFile)
	if err != nil {
		log.Fatalf("init seen store: %v", err)
	}
	newEntries := seenStore.FilterNew(entries)

	selected := rules.Apply(newEntries, cfg.Rules)
	selected = limitRecent(selected, cfg.Feeds.MaxEntriesPerRun)
	summarizer, err := summarizers.BuildFromConfig(cfg.Summarizer)
	if err != nil {
		log.Fatalf("build summarizer: %v", err)
	}

	for i := range selected {
		analysis, err := summarizer.Summarize(context.Background(), selected[i])
		if err != nil {
			log.Printf("summarizer %s failed for %s: %v", summarizer.Name(), selected[i].Link, err)
			continue
		}
		selected[i].Analysis = &analysis
	}

	batch := core.Batch{
		RunAt:   now,
		Entries: selected,
		Meta: map[string]string{
			"feeds_total":     fmt.Sprintf("%d", len(feeds)),
			"entries_total":   fmt.Sprintf("%d", len(entries)),
			"entries_new":     fmt.Sprintf("%d", len(newEntries)),
			"entries_matched": fmt.Sprintf("%d", len(selected)),
			"summarizer":      summarizer.Name(),
		},
	}

	deliverySinks, err := sinks.BuildFromConfig(cfg.Sinks, *dryRun)
	if err != nil {
		log.Fatalf("build sinks: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, sink := range deliverySinks {
		if err := sink.Send(ctx, batch); err != nil {
			log.Fatalf("sink %s failed: %v", sink.Name(), err)
		}
		log.Printf("sink %s delivered %d entries", sink.Name(), len(batch.Entries))
	}

	seenStore.MarkSeen(entries)
	if err := seenStore.Save(); err != nil {
		log.Fatalf("save seen state: %v", err)
	}
}

func limitRecent(entries []core.Entry, max int) []core.Entry {
	if max <= 0 || len(entries) <= max {
		return entries
	}
	cp := append([]core.Entry(nil), entries...)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Published.After(cp[j].Published)
	})
	return cp[:max]
}

func validateConfigCmd(args []string) {
	fs := flag.NewFlagSet("validate-config", flag.ExitOnError)
	configPath := fs.String("config", "config.yaml", "Path to config YAML")
	_ = fs.Parse(args)

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	if err := config.Validate(cfg); err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	fmt.Println("config valid")
}

func usage() {
	fmt.Fprintf(os.Stderr, "opmlwatch commands:\n")
	fmt.Fprintf(os.Stderr, "  run --config examples/config.basic.yaml [--dry-run]\n")
	fmt.Fprintf(os.Stderr, "  validate-config --config examples/config.basic.yaml\n")
}
