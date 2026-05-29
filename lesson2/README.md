# Lesson 2: Retries and Failures

## Concepts

- **RetryPolicy** — controls how many times and how often Temporal retries a failed activity.
- **Exponential backoff** — Temporal automatically increases the wait between retries (1s, 2s, 4s...).
- **Attempt number** — each retry gets an incremented attempt number accessible via `activity.GetInfo(ctx).Attempt`.
- **Event history** — Temporal records only the final attempt outcome, not each individual failure.

## What this lesson does

Simulates a flaky activity that fails on attempts 1 and 2, succeeds on attempt 3.

```
GreetingWorkflow("Karthik")
    └── ComposeGreeting — attempt 1: FAIL
                        — attempt 2: FAIL
                        — attempt 3: SUCCESS → "Hello, Karthik! Welcome to Temporal."
```

## Key rules

- By default Temporal retries indefinitely with exponential backoff — no retry code needed.
- `MaximumAttempts` caps the total number of tries. If all attempts fail, the workflow receives the error.
- The workflow itself never sees the failures — it only sees success or final failure.
- Activity retries are **not** recorded as separate events in the workflow history. Only the final `ActivityTaskStarted` (with `attempt` and `lastFailure`) and `ActivityTaskCompleted` are recorded.

## Structure

```
lesson2/
├── greeting/
│   ├── activity.go    — fails if attempt < 3
│   ├── workflow.go    — sets RetryPolicy with MaximumAttempts: 5
│   └── constants.go   — shared task queue name
├── worker/
│   └── main.go
└── starter/
    └── main.go
```

## Run

**Terminal 1:**
```bash
go run lesson2/worker/main.go
```

**Terminal 2:**
```bash
go run lesson2/starter/main.go
```

Watch the worker logs — you will see attempt 1 and 2 fail, attempt 3 succeed.
