package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

type SelectMode int

const (
	RandomSelect SelectMode = iota
	RoundRobinSelect
)

type Discovery interface {
	Refresh() error
	Update(servers []string) error
	Get(mode SelectMode) (string, error)
	GetAll() ([]string, error)
}

// MultiServersDiscovery is a discovery for multi servers without a registry center
// user provides the server addresses explicitly instead
type MultiServerDiscovery struct {
	r       *rand.Rand
	mu      sync.RWMutex
	servers []string
	index   int
}

// NewMultiServerDiscovery creates a MultiServersDiscovery instance
func NewMultiServerDiscovery(servers []string) *MultiServerDiscovery {
	d := MultiServerDiscovery{
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
		servers: servers,
	}
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return &d
}

var _ Discovery = (*MultiServerDiscovery)(nil)

// Refresh doesn't make sense for MultiServersDiscovery, so ignore it
func (m *MultiServerDiscovery) Refresh() error {
	return nil
}

// Update the servers of discovery dynamically if needed
func (m *MultiServerDiscovery) Update(servers []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.servers = servers
	return nil
}

// Get a server according to mode
func (m *MultiServerDiscovery) Get(mode SelectMode) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	n := len(m.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case RandomSelect:
		return m.servers[m.r.Intn(n)], nil
	case RoundRobinSelect:
		s := m.servers[m.index%n]
		m.index = (m.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

// GetAll returns all servers in discovery
func (m *MultiServerDiscovery) GetAll() ([]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	// return a copy of d.servers
	servers := make([]string, len(m.servers), len(m.servers))
	copy(servers, m.servers)
	return servers, nil
}
