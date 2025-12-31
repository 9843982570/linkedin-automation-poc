package main

import (
	"fmt"
	"log"
	"time"

	"linkedin-automation-poc/internal/browser"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/stealth"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	b, err := browser.Start(cfg.Browser.Headless)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	page := b.MustPage("https://example.com")
	page.MustWaitLoad()

	fmt.Println("Page loaded. Starting human-like scrolling...")

	// Scroll DOWN
	for i := 0; i < 5; i++ {
		page.MustEval(`() => window.scrollBy(0, 300)`)
		stealth.RandomDelay(
			cfg.Timing.MinDelayMs,
			cfg.Timing.MaxDelayMs,
		)
	}

	// Scroll UP
	for i := 0; i < 5; i++ {
		page.MustEval(`() => window.scrollBy(0, -300)`)
		stealth.RandomDelay(
			cfg.Timing.MinDelayMs,
			cfg.Timing.MaxDelayMs,
		)
	}

	fmt.Println("Human-like scrolling completed.")

	time.Sleep(3 * time.Second)
}

