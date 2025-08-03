package f

import "sync"

// WithLock run the provided function under the lock.
func WithLock(mu sync.Locker, action func()) {
        mu.Lock()
        defer mu.Unlock()

        action()
}
