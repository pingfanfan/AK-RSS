package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/opmlwatch/opmlwatch/internal/config"
	"github.com/opmlwatch/opmlwatch/internal/core"
	"github.com/opmlwatch/opmlwatch/internal/opml"
	"github.com/opmlwatch/opmlwatch/internal/rules"
	"github.com/opmlwatch/opmlwatch/internal/sinks"
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

	// MVP placeholder entries so sinks and rule flow are testable end-to-end.
	entries := make([]core.Entry, 0, len(feeds))
	now := time.Now().UTC()
	for _, f := range feeds {
		entries = append(entries, core.Entry{
			FeedTitle: f.Title,
			Title:     fmt.Sprintf("Feed discovered: %s", f.Title),
			Link:      f.XMLURL,
			Published: now,
			Tags:      []string{"feed-discovery"},
		})
	}

	selected := rules.Apply(entries, cfg.Rules)
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
			"feeds_total": fmt.Sprintf("%d", len(feeds)),
			"summarizer":  summarizer.Name(),
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
