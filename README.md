# go-flagsmith

A REST microservice in Go designed to demonstrate resilience patterns, observability, and dynamic feature flag consumption using the Flagsmith SDK.

This project is part of building a foundation for an Internal Developer Platform (IDP).

---

## 🎯 Project Objectives

- **Feature Flags & Remote Config:** Highly efficient consumption focused on low latency and network failure resilience.
- **Lifecycle Probes:** `/healthz` (liveness) and `/readyz` (readiness) endpoints reflecting the actual state of application subsystems.
- **Resilience:** Safe fallback strategies to prevent downtime during external service instability.
- **Go Best Practices:** Safe concurrency (`sync.RWMutex`), dependency injection, and automated testing.

---

## 🛠️ Technical Stack

- **Language:** Go (1.22+)
- **HTTP Router:** Gin Web Framework
- **SDK:** Flagsmith Go Client (Local Evaluation / Polling)
- **Containerization:** Docker

---

## 🚀 How to Run

### Prerequisites

- Go 1.22+
- Make (optional, for task automation)

### Environment Variables

Create a `.env` file or export the Flagsmith API key:

```bash
export FLAGSMITH_API_KEY="your_key_here"
```

### Running Locally

```bash
# Formats, vets, checks envs, and runs the application!
make dev
```

---

## 📌 Project Status

> ⚠️ **Under active development.** Package structures, routes, and integrations are being continuously refined.