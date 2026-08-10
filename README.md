# dispatch-api

The dispatch service for **Waggle Freight**. It assigns delivery jobs to available drivers.

This repository is used as the sample component in the CloudBees Unify tutorial [Create a custom action](https://docs.cloudbees.com/docs/cloudbees-unify/latest/continuous-integration/tutorials/create-a-custom-action), where a CI workflow builds and tests the service and then runs a custom action against it.

## What it does

The service exposes a single scheduling rule: given a number of available drivers, it returns how many deliveries can be scheduled for a shift. Each driver handles a fixed route size (`deliveriesPerDriver`, currently 4), and driver counts of zero or fewer schedule no deliveries.

## Requirements

- Go 1.20 or later

## Run it

```bash
go run .
```

Expected output:

```
Waggle dispatch-api starting
Drivers available: 3
Deliveries scheduled: 12
```

## Test it

```bash
go test ./...
```

The tests cover the standard case, a single driver, zero drivers, and negative input.

## Project layout

| File           | Purpose                                            |
| -------------- | -------------------------------------------------- |
| `main.go`      | The dispatch service and `AssignDeliveries` logic. |
| `main_test.go` | Table-driven tests for `AssignDeliveries`.         |
| `go.mod`       | Go module definition.                              |
