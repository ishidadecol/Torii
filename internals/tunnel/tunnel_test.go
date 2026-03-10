package tunnel

import (
	"testing"
)

// Test GenerateID produces a non-empty string
func TestGenerateID(t *testing.T) {
	id := GenerateID()
	if id == "" {
		t.Fatal("GenerateID() returned empty string")
	}
}

// Test registering a tunnel
func TestRegisterAndGetTunnel(t *testing.T) {
	manager := NewManager()

	// Create a fake tunnel
	tunnel := &Tunnel{
		ID:   "abc123",
		Conn: nil, // we don't need a real connection for this test
	}

	manager.Register(tunnel)

	// Check that the tunnel is retrievable
	got := manager.Get("abc123")
	if got == nil {
		t.Fatal("Expected to get tunnel, got nil")
	}

	if got.ID != "abc123" {
		t.Fatalf("Expected tunnel ID abc123, got %s", got.ID)
	}
}

// Test removing a tunnel
func TestRemoveTunnel(t *testing.T) {
	manager := NewManager()

	tunnel := &Tunnel{
		ID:   "abc123",
		Conn: nil,
	}

	manager.Register(tunnel)
	manager.Remove("abc123")

	got := manager.Get("abc123")
	if got != nil {
		t.Fatal("Expected tunnel to be removed, but it still exists")
	}
}

// Test concurrent access
func TestConcurrentRegister(t *testing.T) {
	manager := NewManager()
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i int) {
			id := GenerateID()
			manager.Register(&Tunnel{ID: id})
			done <- true
		}(i)
	}

	// wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}

	if len(manager.tunnels) != 10 {
		t.Fatalf("Expected 10 tunnels, got %d", len(manager.tunnels))
	}
}
