package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"linkedin-automation-go/stealth"
	"linkedin-automation-go/config"
)

func main() {
	fmt.Println("LinkedIn Automation Proof-of-Concept")
	fmt.Println("Educational use only")
	cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
	    log.Fatal(err)
    }

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
	stealth.ApplyFingerprintMask(page)
	stealth.MoveMouseHumanLike(
	page,
	stealth.Point{X: 100, Y: 120},
	stealth.Point{X: 600, Y: 420},
)

stealth.Think(cfg.Stealth.MinThinkMs, cfg.Stealth.MaxThinkMs)
stealth.ScrollHumanLike(page)

	log.Println("Browser launched successfully")

	// Simulate human think time
	time.Sleep(2 * time.Second)

	log.Println("Demo completed")
}
