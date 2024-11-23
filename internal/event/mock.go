package event

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/khulnasoft/artifacthub/internal/hub"
	"github.com/stretchr/testify/mock"
)

// ManagerMock is a mock implementation of the EventManager interface.
type ManagerMock struct {
	mock.Mock
}

// GetPending implements the EventManager interface.
func (m *ManagerMock) GetPending(ctx context.Context, tx pgx.Tx) (*hub.Event, error) {
	args := m.Called(ctx, tx)
	data, _ := args.Get(0).(*hub.Event)
	return data, args.Error(1)
}
