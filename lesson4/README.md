# Lesson 4: Queries

## Concepts

- **Query** — reads the current state of a running workflow without interrupting it.
- **Query handler** — a function registered inside the workflow via `workflow.SetQueryHandler`.
- **Signals vs Queries** — signals send data in and change state; queries read state and change nothing.

## What this lesson does

An order workflow that tracks status through a lifecycle driven by signals, queryable at any point.

```
OrderWorkflow("ORD-001") starts with status = "pending"

query  → "pending"
signal → "processing"
query  → "processing"
signal → "shipped"
query  → "shipped"
signal → "delivered"   → workflow completes
```

## Key rules

- Query handlers are closures — they capture workflow-local variables directly.
- Queries are synchronous — the caller waits for the result.
- Queries cannot modify workflow state — they are read-only.
- The workflow loops until it receives `"delivered"` via signal.

## Structure

```
lesson4/
├── order/
│   └── workflow.go    — tracks status, registers query handler, loops on signal channel
├── worker/
│   └── main.go
├── starter/
│   └── main.go
├── querier/
│   └── main.go        — queries current status
└── signaler/
    └── main.go        — advances status (usage: go run ... <status>)
```

## Run

**Terminal 1:**
```bash
go run lesson4/worker/main.go
```

**Terminal 2:**
```bash
go run lesson4/starter/main.go
```

**Query and advance status:**
```bash
go run lesson4/querier/main.go                        # → pending
go run lesson4/signaler/main.go processing
go run lesson4/querier/main.go                        # → processing
go run lesson4/signaler/main.go shipped
go run lesson4/querier/main.go                        # → shipped
go run lesson4/signaler/main.go delivered             # workflow completes
```
