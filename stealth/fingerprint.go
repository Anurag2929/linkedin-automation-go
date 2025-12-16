package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// ApplyFingerprintMask applies basic browser fingerprint masking
func ApplyFingerprintMask(page *rod.Page) {
	rand.Seed(time.Now().UnixNano())

	// Random realistic viewport
	width := rand.Intn(400) + 1200
	height := rand.Intn(300) + 700

	page.MustSetViewport(width, height, 1, false)

	// Disable webdriver flag
	_, _ = page.Eval(`() => {
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	}`)

	// Set custom User-Agent
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/120.0.0.0 Safari/537.36"

	page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: ua,
	})
}
