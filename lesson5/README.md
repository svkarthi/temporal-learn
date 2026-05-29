# Lesson 5: Timers

## Concepts

- **`workflow.NewTimer`** — a durable timer. Survives worker crashes and restarts.
- **`workflow.NewSelector`** — waits for whichever of multiple futures/channels fires first.
- **`workflow.WithCancel`** — creates a cancellable context to cancel a timer before it fires.

## What this lesson does

A reminder workflow that waits 10 seconds before firing. Can be cancelled before the timer fires.

**Scenario A — timer fires:**
```
ReminderWorkflow("Buy groceries", 10)
    └── waits 10 seconds → "Reminder fired: Buy groceries"
```

**Scenario B — cancelled before timer fires:**
```
ReminderWorkflow("Buy groceries", 10)
    └── cancel signal arrives → timer cancelled → "Reminder cancelled: Buy groceries"
```

## Key rules

- `workflow.NewTimer` is not `time.Sleep` — it is tracked by Temporal and resumes correctly after a worker restart.
- `workflow.NewSelector` picks whichever fires first — timer or signal. The other is ignored.
- Cancelling a timer via `cancelTimer()` does not cancel the workflow — just the timer future.

## Structure

```
lesson5/
├── reminder/
│   └── workflow.go    — sets timer, uses selector to wait for timer or cancel signal
├── worker/
│   └── main.go
├── starter/
│   └── main.go        — starts workflow with task="Buy groceries", delay=10s
└── canceller/
    └── main.go        — sends cancel signal before timer fires
```

## Run

**Scenario A — let it fire:**
```bash
# Terminal 1
go run lesson5/worker/main.go

# Terminal 2 — wait 10 seconds
go run lesson5/starter/main.go
```

**Scenario B — cancel it:**
```bash
# Terminal 1
go run lesson5/worker/main.go

# Terminal 2
go run lesson5/starter/main.go

# Terminal 3 — within 10 seconds
go run lesson5/canceller/main.go
```
