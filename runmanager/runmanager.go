package runmanager

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Runner incorporates a Run and a Shutdown method to control a service.
type Runner struct {
	Run      func() error
	Shutdown func() error
}

// RunManager can Run services while managing startup and Shutdown.
type RunManager struct {
	Context              context.Context
	WaitingGroupShutdown *sync.WaitGroup
	ErrChan              chan<- error
	Cancel               func()
}

// Run runs the Runner.
func (m *RunManager) Run(r *Runner) {
	const oneRun = 1

	m.WaitingGroupShutdown.Add(oneRun)

	go func() {
		if err := r.Run(); err != nil {
			m.ErrChan <- err
		}
	}()
	go func() {
		<-m.Context.Done()

		if err := r.Shutdown(); err != nil {
			m.ErrChan <- err
		}

		m.WaitingGroupShutdown.Done()
	}()
}

// RunInterruptHandler runs an interrupt handler ontop of the RunManager.
func (m *RunManager) RunInterruptHandler() {
	RunInterruptHandler(m)
}

// RunInterruptHandler runs an interrupt handler ontop of the RunManager.
func RunInterruptHandler(m *RunManager) {
	// Signal Runner.
	c := make(chan os.Signal, 1)

	m.Run(&Runner{
		Run: func() error {
			signal.Notify(c, syscall.SIGTERM, os.Interrupt)
			defer signal.Stop(c)
			for sig := range c {
				switch sig {
				case syscall.SIGTERM, os.Interrupt:
					fmt.Println("Canceling due to interrupt")
					m.Cancel()
				default:
					log.Println("Unknown signal:", sig)
				}
			}
			return nil
		},
		Shutdown: func() error {
			signal.Stop(c)
			return nil
		},
	})
}

// Service is a runnable service.
type Service func(m *RunManager)

// RunServices starts up all given services.
func RunServices(services []Service) {
	ctx, cancel := context.WithCancel(context.Background())

	const defaultErrorChannelSize = 100
	errChan := make(chan error, defaultErrorChannelSize)

	// Shutdown waiting group.
	var wgShutdown sync.WaitGroup

	manager := &RunManager{
		Context:              ctx,
		WaitingGroupShutdown: &wgShutdown,
		ErrChan:              errChan,
		Cancel:               cancel,
	}

	for _, service := range services {
		service(manager)
	}

	select {
	case err := <-errChan:
		fmt.Println("Shutdown-cause=error:", err)
		cancel()
	case <-ctx.Done():
		// Shutting down without error as the reason.
		// Can happen due to interrupt.
		break
	}

	// wait for Shutdown to finish.
	wgShutdown.Wait()

printExistingErrors:
	for {
		select {
		case err := <-errChan:
			fmt.Println("error:", err)
		default:
			break printExistingErrors
		}
	}
}
