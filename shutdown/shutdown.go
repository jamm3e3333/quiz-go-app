package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	SIGINT = iota
	SIGTERM
)

var (
	initiated       = false
	shutdownChan    chan os.Signal
	shutdownSignals = []os.Signal{
		SIGINT:  syscall.SIGINT,
		SIGTERM: syscall.SIGTERM,
	}
)

func SetupShutdownContext() context.Context {
	if initiated {
		panic("cannot create signal context twice")
	}

	initiated = true
	shutdownChan = make(chan os.Signal, 1)

	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(shutdownChan, shutdownSignals...)
	go func() {
		<-shutdownChan
		cancel()
	}()

	return ctx
}

func SignalShutdown() bool {
	if shutdownChan != nil {
		select {
		case shutdownChan <- shutdownSignals[SIGINT]:
			return true
		case shutdownChan <- shutdownSignals[SIGTERM]:
			return true
		default:
		}
	}

	return false
}
