package main

import (
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

	stealth.RandomDelay(
		cfg.Timing.MinDelayMs,
		cfg.Timing.MaxDelayMs,
	)

	time.Sleep(3 * time.Second)
}
