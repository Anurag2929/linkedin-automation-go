package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// Think simulates human thinking delay
func Think(minMs, maxMs int) {
	rand.Seed(time.Now().UnixNano())
	delay := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// ScrollHumanLike simulates natural scrolling
func ScrollHumanLike(page *rod.Page) {
	rand.Seed(time.Now().UnixNano())

	scrolls := rand.Intn(5) + 3
	for i := 0; i < scrolls; i++ {
		_, err := page.Eval(`() => {
			window.scrollBy({
				top: Math.random() * 400 + 100,
				behavior: "smooth"
			});
		}`)
		if err != nil {
			return
		}
		Think(400, 900)
	}
}

