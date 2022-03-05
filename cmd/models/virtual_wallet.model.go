package models

import (
	"models"

	"github.com/google/uuid"
)

type VirtualWallet struct {
	ID               string
	Balance          int64
	LastTransactions []models.Transaction
}

func (v *VirtualWallet) SetID() error {
	id, err := uuid.UUID()
	if err != nil {
		return err
	}
	v.ID = id
	return nil
}

func (v *VirtualWallet) SetBalance() {}

func (v *VirtualWallet) GetBalance() int64 {
	return v.Balance
}

func (v *VirtualWallet) GetLastTransactions() []models.Transaction {
	return v.LastTransactions
}
