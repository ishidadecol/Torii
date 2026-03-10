package main

import (
	"net"
	"testing"
)

func Test_handleConnection(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple echo", "hello", "Echo - hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serverConn, clientConn := net.Pipe()
			defer func() {
				err := clientConn.Close()
				if err != nil {
					t.Logf("Failed to close connection: %v", err)
				}
			}()

			go handleConnection(serverConn)

			_, err := clientConn.Write([]byte(tt.input))
			if err != nil {
				t.Logf("Failed to write to connection: %v", err)
			}
			buffer := make([]byte, 1024)
			n, err := clientConn.Read(buffer)
			if err != nil {
				t.Fatalf("Failed reading response %v", err)
			}

			if string(buffer[:n]) != tt.expected {
				t.Errorf("expected %q got %q", tt.expected, string(buffer[:n]))
			}
		})
	}
}
