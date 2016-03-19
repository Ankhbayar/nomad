package structs

import (
	"fmt"

	cgroupConfig "github.com/opencontainers/runc/libcontainer/configs"
)

const (
	// The default user that the executor uses to run tasks
	DefaultUnpriviledgedUser = "nobody"
)

// WaitResult stores the result of a Wait operation.
type WaitResult struct {
	ExitCode int
	Signal   int
	Err      error
}

func NewWaitResult(code, signal int, err error) *WaitResult {
	return &WaitResult{
		ExitCode: code,
		Signal:   signal,
		Err:      err,
	}
}

func (r *WaitResult) Successful() bool {
	return r.ExitCode == 0 && r.Signal == 0 && r.Err == nil
}

func (r *WaitResult) String() string {
	return fmt.Sprintf("Wait returned exit code %v, signal %v, and error %v",
		r.ExitCode, r.Signal, r.Err)
}

// IsolationConfig has information about the isolation mechanism the executor
// uses to put resource constraints and isolation on the user process
type IsolationConfig struct {
	Cgroup *cgroupConfig.Cgroup
}

// RecoverableError wraps an error and marks whether it is recoverable and could
// be retried or it is fatal.
type RecoverableError struct {
	Err         error
	Recoverable bool
}

// NewRecoverableError is used to wrap an error and mark it as recoverable or
// not.
func NewRecoverableError(e error, recoverable bool) *RecoverableError {
	return &RecoverableError{
		Err:         e,
		Recoverable: recoverable,
	}
}

func (r *RecoverableError) Error() string {
	return r.Err.Error()
}

type CPUStats struct {
	CPU    string
	User   float64
	System float64
	Idle   float64
}

type TaskResourceStats struct {
	CPUStats *CPUStats
}
