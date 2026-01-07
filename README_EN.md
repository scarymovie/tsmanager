# Transaction Manager for PostgreSQL in Go

This library provides a simple and effective transaction manager for working with PostgreSQL in Go. It simplifies transaction management, ensuring that database operations are performed within a transactional context, promoting data consistency and integrity.

## Features

- Simple API for starting, committing, and rolling back transactions.
- Easy integration with existing Go projects.
- Context-based transaction management.
- Supports transaction isolation levels.

## Installation

```bash
go get github.com/scarymovie/tsmanager
```

## Usage

Here is a simple example of how to use the transaction manager:

```go
package main

import (
    "context"
    "database/sql"
    "log"
    
    "github.com/scarymovie/tsmanager"
    
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "your-connection-string")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    tm := tsmanager.New(db)

    err = tm.WithinTransaction(context.Background(), func(ctx context.Context) error {
        // Use getQuerier to get a Querier that will use the transaction
        querier := tsmanager.GetQuerier(ctx, db)
        
        // Perform database operations using the querier
        _, err := querier.ExecContext(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
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