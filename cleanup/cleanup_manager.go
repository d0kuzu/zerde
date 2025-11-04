package cleanup

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type CleanupManager struct {
	mu       sync.Mutex
	cleanups []func()
}

func (m *CleanupManager) Add(fn func()) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cleanups = append(m.cleanups, fn)
}

func (m *CleanupManager) Cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()
	log.Println("🧹 Cleaning up resources...")
	for i := len(m.cleanups) - 1; i >= 0; i-- {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Cleanup panic: %v\n", r)
				}
			}()
			m.cleanups[i]()
		}()
	}
}

func (m *CleanupManager) Start() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	m.Cleanup()
	fmt.Println("✅ Graceful shutdown complete")
	os.Exit(0)
}
