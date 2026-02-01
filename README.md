🇺🇸 English Guide
📝 About The Project

Watcher is a robust monolithic monitoring system. It combines the Management API, background workers, and Auth system into a single high-performance binary, making it easy to scale and maintain.

Key Features:

Monolithic Architecture: One service, one binary, simplified orchestration.

Clean Architecture: Strict separation of concerns for maintainability.

Kubernetes Ready: Ready-to-use manifests for cluster deployment.

Testing: Comprehensive unit tests for Middleware and Auth logic using GoMock.

Security: JWT-based route protection and salted password hashing.

🚀 How to Download and Run

Clone the Repository:
Bash

    git clone https://github.com/your-username/proekt.git
    cd proekt

Run via Docker Compose:
Bash

    docker-compose up --build

Deploy to Kubernetes:
Bash

    kubectl apply -f ./k8s/

Run Tests:
Bash

    go test -v ./pkg/handler/

🛠 Tech Stack

Language: Go (Golang)

Architecture: Monolith

API: Gin Gonic

Database: PostgreSQL + Sqlx

Orchestration: Kubernetes & Docker

Mocking: GoMock & Testify

🇷🇺 Инструкция на русском языке
📝 О проекте

Watcher — это классический монолитный сервис мониторинга. Он объединяет в себе API для управления целями, воркер для проверки статусов и систему авторизации в одном бинарном файле. Это обеспечивает высокую производительность и простоту развертывания.

Основные особенности:

Monolithic Architecture: Всё 
приложение упаковано в один
эффективный сервис.

Clean Architecture: Четкое
разделение на слои (Handler
Service, Repository).

Kubernetes Ready: Набор
манифестов для деплоя в кластер.

Testing: Полное покрытие
критической логики модульными
тестами с использованием моков.

Security: JWT-авторизация
безопасное хранение паролей.

🚀 Как скачать и запустить

Скачивание проекта:
Bash

    git clone https://github.com/your-username/proekt.git
    cd proekt

Запуск через Docker Compose:
Bash

    docker-compose up --build

Развертывание в Kubernetes:
Bash

    kubectl apply -f ./k8s/

Запуск тестов:
Bash

    go test -v ./pkg/handler/

