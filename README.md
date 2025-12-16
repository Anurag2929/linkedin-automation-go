# LinkedIn Automation – Go (Technical Assignment)

## Disclaimer

This project was created strictly for educational and internship evaluation purposes.

Automating LinkedIn violates LinkedIn’s Terms of Service. This repository is not intended for real usage, production deployment, or use with real LinkedIn accounts.

---

## Project Description

This repository contains a Go-based LinkedIn automation proof-of-concept developed as part of an internship technical assignment.

The objective of this project was to quickly understand browser automation using Go and implement the required concepts such as stealth behavior, configuration handling, logging, and state persistence.

The focus of this implementation is on system structure and behavior simulation rather than actual LinkedIn automation.

---

## What This Project Implements

### Browser Automation
- Uses the Rod library to control a Chromium browser
- Demonstrates basic page navigation and interaction

### Stealth and Human-like Behavior
- Human-like mouse movement using Bézier curves
- Randomized delays and scrolling behavior
- Browser fingerprint masking
  - Custom User-Agent
  - Viewport randomization
  - Suppression of navigator.webdriver

### Configuration Management
- YAML-based configuration file (config.yaml)
- Sensible default values
- Configurable headless mode and timing behavior

### Logging
- Structured logging with levels (INFO, WARN, ERROR)
- Timestamped logs for clarity during execution

### State Persistence
- Application state stored in a JSON file (state.json)
- Tracks lifecycle events such as:
  - Application start
  - Browser launch
  - Demo completion
- Allows inspection of previous runs

---

## Project Structure

linkedin-automation-go/
├── cmd/
│ └── main.go
├── stealth/
│ ├── mouse.go
│ ├── timing.go
│ └── fingerprint.go
├── config/
│ ├── config.go
│ └── logger.go
├── state/
│ └── store.go
├── config.yaml
├── state.json
├── go.mod
└── README.md

---

## How to Run

### Requirements
- Go 1.20 or later
- Git

### Steps
```bash
git clone https://github.com/Anurag2929/linkedin-automation-go.git
cd linkedin-automation-go
go run ./cmd
```
When executed, the program launches a browser, applies stealth techniques, simulates human-like interaction, logs actions to the console, and saves execution state to state.json.

Configuration
The application behavior can be modified using the config.yaml file:

browser:
  headless: false

stealth:
  min_think_ms: 800
  max_think_ms: 1500
  
Demo Video
A short demo video is provided as part of the submission. It shows program execution, browser automation behavior, logging output, and the generated state file.

Notes
This project is intentionally limited to safe demo behavior. No real LinkedIn accounts are used. The goal is to demonstrate technical understanding and problem-solving ability rather than real-world automation.

Author
K Anurag Gouda
B.Tech Computer Science
Internship Technical Assignment

---

## Final Step
Commit it:

```bash
git add README.md
git commit -m "Update README with project description and usage instructions"
git push
