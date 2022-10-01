# Asynchronous logger

Асинхронный логгер

## Аргументы и переменные окружения

--threads_count [int] [default: 3]
Кол-во потоков для записи логов


--log_count [int] [default: 50]
Кол-во логов записанных в каждом потоке

--async_logger_enabled [bool] [default: true]
Флаг для использщования обычного логгера или асинхронного

## Примеры запуска

```bash
go run cmd/main.go --threads_count 3 --log_count 50
```

## Примечание
func (f *File) Sync() не используется специально.