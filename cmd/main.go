package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"linkedin-automation-go/stealth"
)

func main() {
	fmt.Println("LinkedIn Automation Proof-of-Concept")
	fmt.Println("Educational use only")

	// Launch browser
	url := launcher.New().
	Leakless(false).
	Headless(false).
	MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()
	defer browser.MustClose()

	// Open a safe demo page
	page := browser.MustPage("https://example.com")
	page.MustWaitLoad()
	stealth.MoveMouseHumanLike(
	page,
	stealth.Point{X: 100, Y: 120},
	stealth.Point{X: 600, Y: 420},
)

	log.Println("Browser launched successfully")

	// Simulate human think time
	time.Sleep(2 * time.Second)

	log.Println("Demo completed")
}
