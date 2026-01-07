# Менеджер транзакций для PostgreSQL на Go

Эта библиотека предоставляет простой и эффективный менеджер транзакций для работы с PostgreSQL на языке Go. Она упрощает управление транзакциями, обеспечивая выполнение операций с базой данных в транзакционном контексте, способствуя согласованности и целостности данных.

## Особенности

- Простой API для начала, фиксации и отката транзакций.
- Легкая интеграция с существующими проектами на Go.
- Управление транзакциями на основе контекста.
- Поддержка уровней изоляции транзакций.

## Установка

```bash
go get github.com/scarymovie/tsmanager
```

## Использование

Вот простой пример того, как использовать менеджер транзакций:

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
        // Используйте getQuerier для получения Querier, который будет использовать транзакцию
        querier := tsmanager.GetQuerier(ctx, db)
        
        // Выполняйте операции с базой данных с помощью querier
        _, err := querier.ExecContext(ctx, "INSERT INTO users (name, email) VALUES ($1, $2)", "John Doe", "john@example.com")
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