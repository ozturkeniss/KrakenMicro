# GoMicro

GoMicro is a modern, microservices-based backend system designed for e-commerce platforms. The architecture leverages Go for high performance and scalability, and integrates PostgreSQL, Redis, and RabbitMQ to ensure robust data management, caching, and asynchronous communication. API Gateway functionality is provided by Krakend, enabling unified and secure access to all services.

## Architecture Overview

```mermaid
flowchart TD
    %% Nodes
    AGW[Krakend API Gateway]
    US[User Service]
    PS[Product Service]
    BS[Basket Service]
    PAY[Payment Service]
    DB_USER[(User PostgreSQL)]
    DB_PRODUCT[(Product PostgreSQL)]
    DB_PAYMENT[(Payment PostgreSQL)]
    REDIS[(Redis Cache)]
    MQ[(RabbitMQ)]

    %% API Gateway routes
    AGW -- "gRPC/HTTP" --> US
    AGW -- "gRPC/HTTP" --> PS
    AGW -- "gRPC/HTTP" --> BS
    AGW -- "gRPC/HTTP" --> PAY

    %% Service to DB
    US -- "SQL" --> DB_USER
    PS -- "SQL" --> DB_PRODUCT
    PAY -- "SQL" --> DB_PAYMENT
    BS -- "Cache" --> REDIS

    %% Message Queue
    PS -- "Publishes Events" --> MQ
    PAY -- "Publishes Events" --> MQ
    MQ -- "Stock Updates" --> PS

    %% Inter-service communication
    BS -- "Product Info" --> PS
    PAY -- "Basket Info" --> BS

    %% Styling
    style AGW fill:#b2ebf2,stroke:#0097a7,stroke-width:2px,color:#111
    style US fill:#e0f7fa,stroke:#0097a7,stroke-width:2px,color:#111
    style PS fill:#b2ebf2,stroke:#0097a7,stroke-width:2px,color:#111
    style BS fill:#e0f7fa,stroke:#0097a7,stroke-width:2px,color:#111
    style PAY fill:#b2ebf2,stroke:#0097a7,stroke-width:2px,color:#111
    style DB_USER fill:#80deea,stroke:#0097a7,stroke-width:2px,color:#111
    style DB_PRODUCT fill:#4dd0e1,stroke:#0097a7,stroke-width:2px,color:#111
    style DB_PAYMENT fill:#26c6da,stroke:#0097a7,stroke-width:2px,color:#111
    style REDIS fill:#00bcd4,stroke:#0097a7,stroke-width:2px,color:#111
    style MQ fill:#4dd0e1,stroke:#0097a7,stroke-width:2px,color:#111
    %% Set all text to black
    linkStyle default stroke:#0097a7,stroke-width:1.5px,color:#111
```

## Services and Responsibilities

- **User Service**: Manages user registration, authentication, and profile operations.
- **Product Service**: Handles product CRUD operations and inventory management.
- **Basket Service**: Manages user shopping baskets with high-performance access via Redis.
- **Payment Service**: Processes payments and updates inventory asynchronously using RabbitMQ.
- **API Gateway (Krakend)**: Provides a single entry point for all client requests, routing them to the appropriate microservice.

## Technology Stack

- Go (1.24+)
- PostgreSQL
- Redis
- RabbitMQ
- gRPC
- Krakend (API Gateway)
- Docker & Docker Compose

## Project Directory Structure

```mermaid
flowchart TD
    A[cmd] --> B[api-gateway]
    A --> C[user-service]
    A --> D[product-service]
    A --> E[basket-service]
    A --> F[payment-service]
    G[internal] --> H[user]
    G --> I[product]
    G --> J[basket]
    G --> K[payment]
    L[api] --> M[proto]
    N[deployments] --> O[docker]
    P[tests]
```

## Getting Started

1. **Prerequisites:**
   - Docker & Docker Compose
   - Go 1.24+

2. **Clone the repository:**
   ```bash
   git clone <repo-url>
   cd gomicro
   ```

3. **Start all services:**
   ```bash
   docker-compose up --build
   ```

4. **Access the API Gateway:**
   - http://localhost:8085

## API Gateway Endpoints

| Endpoint         | Method | Service           |
|------------------|--------|-------------------|
| /users           | GET    | User Service      |
| /users           | POST   | User Service      |
| /products        | GET    | Product Service   |
| /products        | POST   | Product Service   |
| /basket          | GET    | Basket Service    |
| /basket          | POST   | Basket Service    |
| /payments        | POST   | Payment Service   |

## Testing

Unit tests for all services are located in the `tests/` directory. To run all tests:

```bash
go test ./tests/...
```

## Protocol Definitions

All inter-service communication is implemented using gRPC. Protocol buffer definitions are located in the `api/proto/` directory.

## Contributing & License

Contributions are welcome via pull requests or issues. Licensed under the MIT License. 