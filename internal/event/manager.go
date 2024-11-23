package event

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v4"
	"github.com/khulnasoft/artifacthub/internal/hub"
)

const (
	// Database queries
	getPendingEventDBQ = `select get_pending_event()`
)

// Manager provides an API to manage events.
type Manager struct{}

// NewManager creates a new Manager instance.
func NewManager() *Manager {
	return &Manager{}
}

// GetPending returns a pending event to be processed if available.
func (m *Manager) GetPending(ctx context.Context, tx pgx.Tx) (*hub.Event, error) {
	var dataJSON []byte
	if err := tx.QueryRow(ctx, getPendingEventDBQ).Scan(&dataJSON); err != nil {
		return nil, err
	}
	var e *hub.Event
	if err := json.Unmarshal(dataJSON, &e); err != nil {
		return nil, err
	}
	return e, nil
}
