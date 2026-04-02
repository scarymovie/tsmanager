# Transaction Manager for PostgreSQL in Go

This library provides a simple and effective transaction manager for working with PostgreSQL in Go. It simplifies transaction management, ensuring that database operations are performed within a transactional context, promoting data consistency and integrity.

## Features

- Simple API for starting, committing, and rolling back transactions.
- Easy integration with existing Go projects.
- Context-based transaction management.
- Supports transaction isolation levels.
- Uses `pgxpool` for efficient connection pooling.

## Installation

```bash
go get github.com/scarymovie/txmanager
```

## Usage

Here is a simple example of how to use the transaction manager:

```go
package main

import (
    "context"
    "log"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/scarymovie/txmanager"
)

func main() {
    ctx := context.Background()
    
    pool, err := pgxpool.New(ctx, "postgres://user:password@localhost:5432/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()

    tm := txmanager.New(pool)

    err = tm.WithinTransaction(ctx, func(ctx context.Context) error {
        // Use GetQuerier to get a Querier that will use the transaction
        querier := txmanager.GetQuerier(ctx, pool)

        // Perform database operations using the querier
        _, err := querier.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
        if err != nil {
            return err
        }

        // More operations can be performed here
        // If any operation returns an error, the transaction will be rolled back
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Transaction completed successfully")
}
```

## Contribution

Contributions are welcome! Please create an issue or submit a pull request.