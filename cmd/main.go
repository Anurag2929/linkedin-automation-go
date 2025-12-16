package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation-go/config"
	"linkedin-automation-go/state"
	"linkedin-automation-go/stealth"
)

func main() {
	fmt.Println("LinkedIn Automation Proof-of-Concept")
	fmt.Println("Educational use only")

	// Load config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	// Initialize logger
	logger := config.NewLogger()

	// Initialize application state
	appState := state.NewState()
	appState.AddAction("Application started")

	// Launch browser
	url := launcher.New().
		Leakless(false).
		Headless(cfg.Browser.Headless).
		MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()
	defer browser.MustClose()

	// Open a safe demo page
	page := browser.MustPage("https://example.com")
	page.MustWaitLoad()

	// Apply stealth techniques
	stealth.ApplyFingerprintMask(page)

	stealth.MoveMouseHumanLike(
		page,
		stealth.Point{X: 100, Y: 120},
		stealth.Point{X: 600, Y: 420},
	)

	stealth.Think(cfg.Stealth.MinThinkMs, cfg.Stealth.MaxThinkMs)
	stealth.ScrollHumanLike(page)

	logger.Info("Browser launched successfully")
	appState.AddAction("Browser launched")

	// Simulate human think time
	time.Sleep(2 * time.Second)

	logger.Info("Demo completed")
	appState.AddAction("Demo completed")

	// Save state to disk
	err = appState.Save("state.json")
	if err != nil {
		logger.Warn("Failed to save state file")
	}
}
