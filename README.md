# Менеджер транзакций для PostgreSQL на Go

Эта библиотека предоставляет простой и эффективный менеджер транзакций для работы с PostgreSQL на языке Go. Она упрощает управление транзакциями, обеспечивая выполнение операций с базой данных в транзакционном контексте, способствуя согласованности и целостности данных.

## Особенности

- Простой API для начала, фиксации и отката транзакций.
- Легкая интеграция с существующими проектами на Go.
- Управление транзакциями на основе контекста.
- Поддержка уровней изоляции транзакций.
- Использование `pgxpool` для эффективного управления соединениями.

## Установка

```bash
go get github.com/scarymovie/txmanager
```

## Использование

Вот простой пример того, как использовать менеджер транзакций:

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
        // Используйте GetQuerier для получения Querier, который будет использовать транзакцию
        querier := txmanager.GetQuerier(ctx, pool)

        // Выполняйте операции с базой данных с помощью querier
        _, err := querier.Exec(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
        if err != nil {
            return err
        }

        // Здесь можно выполнять дополнительные операции
        // Если какая-либо операция возвращает ошибку, транзакция будет откачена
        return nil
    })

    if err != nil {
        log.Fatal(err)
    }

    log.Println("Транзакция успешно завершена")
}
```

## Участие в разработке

Приветствуется участие! Пожалуйста, создайте issue или отправьте pull request.