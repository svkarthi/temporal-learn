# Lesson 6: Nexus

## Concepts

- **Nexus** — a Temporal feature for cross-service communication. One workflow calls an operation exposed by a completely separate service.
- **Nexus Service** — a named group of operations, like a micro-API.
- **Nexus Operation** — a single callable unit. Can be sync (returns immediately) or async (backed by a workflow).
- **Nexus Endpoint** — a routing config created in Temporal that maps an endpoint name to a target task queue.
- **Handler worker** — registers the Nexus service and executes its operations.
- **Caller workflow** — calls the Nexus operation via an endpoint name and service name.

## What this lesson does

A caller workflow calls a `compose` operation on a `greeting-service` Nexus service via a Nexus endpoint.

```
starter
  → CallerWorkflow (nexus-caller-task-queue)
      → "greeting-endpoint" → Temporal routes to nexus-handler-task-queue
          → ComposeOperation runs → "Hello from Nexus, Karthik!"
      ← result
  ← "Hello from Nexus, Karthik!"
```

## Key rules

- The caller and handler are fully decoupled — they share only the operation definition, not implementation.
- The endpoint must be created in Temporal before running (see setup below).
- Two separate workers are needed: one for the handler, one for the caller workflow.
- `RegisterNexusService` is used instead of `RegisterActivity` on the handler worker.

## Structure

```
lesson6/
├── greetingsvc/
│   └── service.go       — defines ComposeOperation and NewService()
├── caller/
│   └── workflow.go      — calls ComposeOperation via Nexus client
├── callerworker/
│   └── main.go          — registers CallerWorkflow on nexus-caller-task-queue
├── handler/
│   └── main.go          — registers greeting-service on nexus-handler-task-queue
└── starter/
    └── main.go          — starts CallerWorkflow
```

## Setup — create the Nexus endpoint (once)

```bash
temporal operator nexus endpoint create \
  --name greeting-endpoint \
  --target-namespace default \
  --target-task-queue nexus-handler-task-queue
```

## Run

**Terminal 1 — handler worker:**
```bash
go run lesson6/handler/main.go
```

**Terminal 2 — caller worker:**
```bash
go run lesson6/callerworker/main.go
```

**Terminal 3 — start:**
```bash
go run lesson6/starter/main.go
```

**Expected output:**
```
Result: Hello from Nexus, Karthik!
```
