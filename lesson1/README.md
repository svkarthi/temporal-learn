# Lesson 1: Workflows and Activities

## Concepts

- **Workflow** — orchestrates the work. Decides what to run, in what order.
- **Activity** — does the actual work. Any code that touches the outside world (I/O, APIs, DBs) belongs here.
- **Worker** — a process that registers workflows and activities, then polls a task queue for work.
- **Starter** — triggers a workflow execution by sending it to the Temporal server.
- **Task Queue** — the named channel connecting starters and workers. Both must reference the same name.

## What this lesson does

Runs three greeting activities in **parallel** and returns all three results.

```
GreetingWorkflow("Karthik")
    ├── ComposeGreeting1 → "Hello, Karthik!"
    ├── ComposeGreeting2 → "Good to see you, Karthik!"
    └── ComposeGreeting3 → "Welcome to Temporal, Karthik!"
```

Activities are fired simultaneously and the workflow waits for all three to complete.

## Key rules

- Workflows must be **deterministic** — no direct I/O, no `time.Now()`, no random numbers.
- Activities can do anything — they run outside the workflow and are retried independently.
- Each activity must be registered with the worker or it cannot be executed.

## Structure

```
lesson1/
├── greeting/
│   ├── activity.go    — three greeting activities
│   ├── workflow.go    — parallel execution of all three
│   └── constants.go   — shared task queue name
├── worker/
│   └── main.go        — registers workflow + activities, polls task queue
└── starter/
    └── main.go        — triggers the workflow and prints results
```

## Run

**Terminal 1:**
```bash
go run lesson1/worker/main.go
```

**Terminal 2:**
```bash
go run lesson1/starter/main.go
```

**Expected output:**
```
Greeting 1: Hello, Karthik!
Greeting 2: Good to see you, Karthik!
Greeting 3: Welcome to Temporal, Karthik!
```
