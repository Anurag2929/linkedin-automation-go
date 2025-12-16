# linkedin-automation-go

# LinkedIn Automation Proof of Concept (Go)

⚠️ Disclaimer  
This project is created strictly for **educational and evaluation purposes only**.  
Automating LinkedIn violates LinkedIn’s Terms of Service. This tool is **not meant to be used on real LinkedIn accounts** and should not be deployed in production.

---

## Project Description

This repository contains a **Go-based proof-of-concept** for LinkedIn automation built as part of an internship technical assignment.

The goal of this project is **not real-world automation**, but to demonstrate:
- Browser automation skills
- Understanding of stealth and anti-detection concepts
- Clean Go project structure
- Ability to quickly learn and implement unfamiliar technologies

The implementation focuses on **architecture and behavior simulation**, rather than actual LinkedIn interaction.

---

## Key Features Implemented

### Browser Automation
- Uses the **Rod** library to control a Chromium browser
- Demonstrates page navigation and interaction

### Stealth & Human-like Behavior
- Human-like mouse movement using Bézier curves
- Randomized delays and scrolling behavior
- Browser fingerprint masking:
  - Custom User-Agent
  - Viewport randomization
  - `navigator.webdriver` suppression

### Configuration Management
- YAML-based configuration file (`config.yaml`)
- Sensible default values
- Configurable headless mode and timing behavior

### Logging
- Structured logging with levels (INFO / WARN / ERROR)
- Timestamped logs for clarity

### State Persistence
- Application state stored in a JSON file (`state.json`)
- Tracks lifecycle events such as:
  - Application start
  - Browser launch
  - Demo completion
- Enables resuming or inspecting previous runs

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

yaml
Copy code

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
When executed, the program:

Launches a browser

Applies stealth techniques

Simulates human-like interaction

Logs actions

Saves execution state to state.json
