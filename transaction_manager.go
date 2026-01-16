package txmanager

import (
	"context"
	"database/sql"
)

type Querier interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type txKey struct{}

// TransactionManager defines the interface for managing database transactions.
type TransactionManager interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// transactionManager implements the TransactionManager interface.
type transactionManager struct {
	db *sql.DB
}

// New creates a new TransactionManager instance.
func New(db *sql.DB) TransactionManager {
	return &transactionManager{db: db}
}

// WithinTransaction executes a function within a database transaction.
// If the function returns an error, the transaction is rolled back.
// Otherwise, the transaction is committed.
func (tm *transactionManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	err = fn(context.WithValue(ctx, txKey{}, tx))
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetQuerier(ctx context.Context, db *sql.DB) Querier {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return db
}
