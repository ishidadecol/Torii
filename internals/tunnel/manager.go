package tunnel

import "sync"

type Manager struct {
	tunnels map[string]*Tunnel
	mu      sync.RWMutex
}

// Creates a new manager and returns a pointer to it
func NewManager() *Manager {
	return &Manager{
		tunnels: make(map[string]*Tunnel),
	}
}

/*Manager function: receives a pointer to a tunnel
* locks itself and registers the tunnel to the manager */
func (m *Manager) Register(t *Tunnel) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tunnels[t.ID] = t
}

/*Manager functionl: get a specific tunnel pointer
* by its id*/
func (m *Manager) Get(id string) *Tunnel {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.tunnels[id]
}

/*Manager function: Remove a specific tunnel from the
* manager by its id*/
func (m *Manager) Remove(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.tunnels, id)
}
