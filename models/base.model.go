package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (b *Base) SetID() error {
	id, err := uuid.UUID()
	if err != nil {
		return err
	}
	b.ID = id
	return nil
}

func (b *Base) IsEmptyId() bool {
	return len(b.ID) == 0
}

func (b *Base) SetCreatedAt(t time.Time) {
	b.CreatedAt = t.String()
}

func (b *Base) SetUpdatedAt(t time.Time) {
	b.UpdatedAt = t.String()
}

func (b *Base) SetDeletedAt(t time.Time) {
	b.DeletedAt = t.String()
}
