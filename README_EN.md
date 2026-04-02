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

### Basic Example

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

### Connection Pool Configuration

The `pgxpool` library provides flexible settings for managing the connection pool:

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/scarymovie/txmanager"
)

func main() {
    ctx := context.Background()
    
    // Create pool configuration
    poolConfig, err := pgxpool.ParseConfig("postgres://user:password@localhost:5432/dbname")
    if err != nil {
        log.Fatal(err)
    }

    // Configure pool settings
    poolConfig.MaxConns = 25              // Maximum number of connections
    poolConfig.MinConns = 5               // Minimum number of connections
    poolConfig.MaxConnLifetime = time.Hour // Maximum lifetime of a connection
    poolConfig.MaxConnIdleTime = 30 * time.Minute // Maximum idle time of a connection

    // Optional: health check configuration
    poolConfig.HealthCheckPeriod = time.Minute

    // Create pool with configuration
    pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()

    tm := txmanager.New(pool)

    err = tm.WithinTransaction(ctx, func(ctx context.Context) error {
        querier := txmanager.GetQuerier(ctx, pool)
        _, err := querier.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
        return err
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Transaction completed successfully")
}
```

#### Main Configuration Parameters:

| Parameter | Description |
|-----------|-------------|
| `MaxConns` | Maximum number of connections in the pool (default: 4) |
| `MinConns` | Minimum number of connections in the pool (default: 0) |
| `MaxConnLifetime` | Maximum time a connection can live before being closed |
| `MaxConnIdleTime` | Maximum time a connection can be idle before being closed |
| `HealthCheckPeriod` | Period for checking the health of connections in the pool |

## Contribution

Contributions are welcome! Please create an issue or submit a pull request.