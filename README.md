# Go Race Condition Example

This repository demonstrates a race condition in Go.  The program uses goroutines to concurrently increment a shared variable, `count`. Without proper synchronization, this leads to incorrect results because the goroutines can race to update `count`.

The `bug.go` file contains the buggy code. The `bugSolution.go` file shows how to fix the race condition using a mutex.