# 📦 Resilient Microservice with Flagsmith & Go

## 🎯 Objective & Use Cases

Develop a REST API in Go acting as a **Dynamic Routing Engine (Canary Engine)**, consuming Feature Flags and Remote Config in a non-blocking, resilient, and secure manner for high-throughput environments.

The microservice must be able to make real-time decisions about traffic destinations (e.g., `v1-legacy` vs `v2-new`) based on user context and platform-defined rules, without the need for additional deployments.

## 🛠️ Technical Stack

- **Language:** Go (1.22+)
- **Router:** `gin-gonic/gin`
- **SDK:** Flagsmith Go SDK (with local evaluation cache and background polling)
- **Container:** Docker (Multi-stage build with `distroless` or `alpine` image)

## 📌 Main Deliverables & Requirements

- [X] **SDK Integration:** Configure the Flagsmith client with in-memory caching and synchronization intervals.
- [X] **Safe Fallback Pattern:** Ensure that network failures or API timeouts do not cause HTTP 500 errors (utilizing predefined default values).
- [X] **Observability & Health Checks:** `/healthz` (liveness) and `/readyz` (readiness) endpoints accurately reflecting application integrity and cache hydration state.
- [ ] **Dynamic Canary Engine:** Create the `router` package and the `/api/v1/route` endpoint to route traffic deterministically (via hashing over user context) based on Remote Configs.
- [ ] **Automated Testing:** Unit tests utilizing clean *interfaces* (e.g., `flags.Reader`) and *mocks* to simulate routing scenarios and failure behaviors.
