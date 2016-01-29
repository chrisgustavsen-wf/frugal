package frugal

import (
	"fmt"
	"time"
)

// FTransportMonitor watches and heals an FTransport.
type FTransportMonitor interface {
	// OnClosedCleanly is called when the transport is closed cleanly by a call to Close()
	OnClosedCleanly()

	// OnClosedUncleanly is called when the transport is closed for a reason *other* than a call to Close().
	// Returns whether to try reopening the transport and, if so, how long to wait before making the attempt.
	OnClosedUncleanly() (reopen bool, wait time.Duration)

	// OnReopenFailed is called when an attempt to reopen the transport fails.
	// Given the number of previous attempts to re-open the transport and the length of the previous wait,
	// Returns whether to attempt to re-open the transport, and how long to wait before making the attempt.
	OnReopenFailed(prevAttempts uint, prevWait time.Duration) (reopen bool, wait time.Duration)

	// ReopenSucceeded is called after the transport has been successfully re-opened.
	OnReopenSucceeded()
}

type fTransportMonitor struct {
	maxReopenAttempts uint
	initialWait       time.Duration
	maxWait           time.Duration
}

// NewFTransportMonitor returns an impplementation that attempts to re-open uncleanly-closed
// transports with exponential backoff behavior.
func NewFTransportMonitor(maxReopenAttempts uint, initialWait, maxWait time.Duration) FTransportMonitor {
	return &fTransportMonitor{
		maxReopenAttempts: maxReopenAttempts,
		initialWait:       initialWait,
		maxWait:           maxWait,
	}
}

func (m *fTransportMonitor) OnClosedUncleanly() (bool, time.Duration) {
	return m.maxReopenAttempts > 0, m.initialWait
}

func (m *fTransportMonitor) OnReopenFailed(prevAttempts uint, prevWait time.Duration) (bool, time.Duration) {
	if prevAttempts >= m.maxReopenAttempts {
		return false, 0
	}

	nextWait := prevWait * 2
	if nextWait > m.maxWait {
		nextWait = m.maxWait
	}
	return true, nextWait
}

func (m *fTransportMonitor) OnClosedCleanly() {}

func (m *fTransportMonitor) OnReopenSucceeded() {}

type monitorRunner struct {
	monitor       FTransportMonitor
	transport     FTransport
	closedChannel <-chan bool
}

// Starts a monitoring
func (r *monitorRunner) run() {
	fmt.Println("FTransport Monitor: Beginning to monitor transport...")
	for {
		wasClean := <-r.closedChannel

		if wasClean {
			r.handleCleanClose()
			return
		} else if shouldContinue := r.handleUncleanClose(); !shouldContinue {
			return
		}
	}
}

// Handle a clean close of the transport.
func (r *monitorRunner) handleCleanClose() {
	fmt.Println("FTransport Monitor: FTransport was closed cleanly. Terminating...")
	r.monitor.OnClosedCleanly()
}

// Handle an unclean close of the transport.
func (r *monitorRunner) handleUncleanClose() bool {
	fmt.Println("FTransport Monitor: FTransport was closed uncleanly!")

	reopen, initialWait := r.monitor.OnClosedUncleanly()
	if !reopen {
		fmt.Println("FTransport Monitor: Instructed not to reopen. Terminating...")
		return false
	}

	return r.attemptReopen(initialWait)
}

// Attempt to reopen the uncleanly closed transport.
func (r *monitorRunner) attemptReopen(initialWait time.Duration) bool {
	wait := initialWait
	reopen := true
	prevAttempts := uint(0)

	for reopen {
		fmt.Printf("FTransport Monitor: Attempting to reopen after %v\n", wait)
		time.Sleep(wait)

		if err := r.transport.Open(); err != nil {
			fmt.Printf("FTransport Monitor: Failed to re-open transport due to: %v\n", err)
			prevAttempts++

			reopen, wait = r.monitor.OnReopenFailed(prevAttempts, wait)
			continue
		}
		fmt.Printf("FTransport Monitor: Successfully re-opened!")
		r.monitor.OnReopenSucceeded()
		return true
	}

	fmt.Println("FTransport Monitor: ReopenFailed callback instructed not to reopen. Terminating...")
	return false
}
