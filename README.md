# Сервис авторизации на gRPC
## Стек технологий
- **Go**
- **PostgreSQL**
- **gRPC**
- **JWT**
- **Taskfile**
## Особенности проекта
1. Работа с базой данных осуществляется с помощью чистых SQL-запросов (без ORM).
2. Гибкая архитектура, обеспечивающая лёгкость в модификации и расширении функционала.
## Структура проекта
- **cmd/**: Исходный код для запуска приложения.
- **config/**: Файлы конфигурации проекта.
- **docker/**: Файлы для настройки и запуска Docker-контейнеров.
- **internal/**: Внутренние пакеты и логика приложения.
- **migrations/**: SQL-файлы для миграций базы данных.
- **tests/**: Тесты для проверки функциональности.
- **Taskfile.yaml**: Файл конфигурации Taskfile.
## Запуск проекта
1. **Запустить базу данных в Docker**
   ```sh
   docker-compose up -d
   ```
2. **Установить Taskfile (если не установлен)**
   ```sh
   brew install go-task
   ```
3. **Запустить сервис**
   ```sh
   task start
   ```
4. **(Опционально) Выполнить миграции**
   ```sh
   task migrate
   ```
## Флаги и переменные окружения
Для управления сервисом можно использовать различные флаги и переменные окружения.
## О миграциях
По умолчанию миграции создают необходимые таблицы без использования отдельной таблицы `migrations`. При необходимости код можно легко модифицировать для добавления этой функциональности.
## Дополнительная информация
Проект использует чистые SQL-запросы для взаимодействия с базой данных, избегая использования ORM, что обеспечивает высокую производительность и контроль над выполняемыми операциями.
Архитектура сервиса спроектирована таким образом, чтобы облегчить его модификацию и расширение, что позволяет адаптировать его под различные требования и сценарии использования.
Если у вас возникнут вопросы или потребуется дополнительная информация, пожалуйста, обратитесь к документации проекта или свяжитесь с автором репозитория.
