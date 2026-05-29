# Lesson 3: Signals

## Concepts

- **Signal** — sends data into a running workflow from the outside. Changes workflow state.
- **Signal channel** — a named channel the workflow listens on via `workflow.GetSignalChannel`.
- **Receive** — blocks the workflow at that point until a signal arrives.

## What this lesson does

An approval workflow that starts and waits indefinitely for an approval or rejection signal.

```
starter → ApprovalWorkflow("REQ-001")
              └── blocks at Receive() ... waiting ...

signaler → sends { Approved: true }
              └── workflow unblocks → returns "Request REQ-001 approved"
```

## Key rules

- A workflow blocked on `Receive` consumes no CPU — it just waits in Temporal's state.
- The workflow can wait for days or weeks — there is no timeout unless you add one.
- Signals are delivered reliably — if the worker is down when a signal arrives, it will be delivered when the worker comes back.
- Signals change state. Queries (lesson 4) read state.

## Structure

```
lesson3/
├── approval/
│   └── workflow.go    — waits for signal, returns approved/rejected
├── worker/
│   └── main.go
├── starter/
│   └── main.go        — starts the workflow
└── signaler/
    └── main.go        — sends the approval signal
```

## Run

**Terminal 1:**
```bash
go run lesson3/worker/main.go
```

**Terminal 2** (workflow starts and blocks):
```bash
go run lesson3/starter/main.go
```

Check Web UI — workflow is in **Running** state.

**Terminal 3** (send the signal):
```bash
go run lesson3/signaler/main.go
```

Workflow completes immediately after the signal arrives.
