package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         string
	Date       string
	Type       string
	Payer      string
	PayerID    string
	Receiver   string
	ReceiverID string
	Sum        int64
}

func NewTransaction(
	typeOfTransaction,
	payer,
	receiver,
	payerID,
	receiverID string,
	sum int64,
	date time.Time) (*Transaction, error) {
	id, err := uuid.UUID()
	if err != nil {
		return nil, err
	}
	time := date.String()
	transaction := &Transaction{
		ID:         id,
		Date:       time,
		Type:       typeOfTransaction,
		Payer:      payer,
		PayerID:    payerID,
		Receiver:   receiver,
		ReceiverID: receiverID,
		Sum:        sum,
	}
	return transaction, nil
}

func (t *Transaction) GetTransactionID() string {
	return t.ID
}

func (t *Transaction) GetSum() int64 {
	return t.Sum
}

func (t *Transaction) GetDate() string {
	return t.Date
}
