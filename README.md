# Dmart
E-Commerce platform

# 🛒 Go E-Commerce Microservices (Gin-Based)

A **production-grade microservices architecture** for an e-commerce platform built using **Go (Golang)** and **Gin**, following clean architecture and event-driven design.

---

## 🚀 Architecture Overview

- Distributed Microservices Architecture  
- API Gateway as a single entry point  
- Event-driven communication using Kafka  
- Database per service  
- Stateless services with JWT authentication  

### 🔷 High-Level Flow

Client (Web / Mobile / Postman)  
↓ HTTPS (JSON)  
API Gateway (Gin)  
↓  
Microservices (Gin-based services)  
↓  
PostgreSQL (per service) + Redis (cache)  
↓  
Kafka (event bus)  
↓  
Background Workers  

---

## 🧱 Core Design Principles

- Clean Architecture (Handler → Service → Repository)  
- Loose coupling via Kafka events  
- Database per service  
- Eventual consistency using Saga pattern  
- Stateless services with JWT  
- Resilience:
  - Retries  
  - Circuit breakers  
  - Idempotency  
- Observability:
  - Structured logging (Zap)  
  - Correlation IDs  
  - Distributed tracing  

---

## 🏗️ Services

- API Gateway → Routing, Auth, Middleware  
- User Service → Authentication & Profile  
- Product Service → Product Catalog  
- Cart Service → Shopping Cart  
- Order Service → Order Management  
- Notification Service → Event-based notifications  

**Future:**
- Payment Service  
- Inventory Service  

---

## ⚙️ Tech Stack

- Language: Go 1.23+  
- Framework: Gin  
- Database: PostgreSQL  
- Cache: Redis  
- Messaging: Kafka  
- Auth: JWT  
- Logging: Zap  
- Config: Viper + .env  
- Containerization: Docker  
- Orchestration: Kubernetes  
- Migrations: golang-migrate  
- Validation: validator.v10  
- Observability: OpenTelemetry  

---

## 📁 Project Structure

### Monorepo Layout
go-ecommerce-microservices/
├── api-gateway/
├── user-service/
├── product-service/
├── cart-service/
├── order-service/
├── notification-service/
├── docker-compose.yml
├── k8s/
├── shared/
└── README.md


---

### Per-Service Structure
user-service/
├── cmd/
│ └── server/main.go
├── internal/
│ ├── config/
│ ├── middleware/
│ ├── handler/
│ ├── service/
│ ├── repository/
│ ├── model/
│ ├── cache/
│ ├── queue/
│ ├── worker/
│ └── utils/
├── pkg/logger/
├── migrations/
├── docker/
├── test/
├── go.mod
├── .env.example
└── README.md


---

## 🌐 API Gateway

### Responsibilities

- Request routing  
- JWT authentication  
- Rate limiting (Redis)  
- Logging & monitoring  
- CORS handling  
- Panic recovery  

### Routes
/api/v1/auth/*
/api/v1/products/*
/api/v1/cart/*
/api/v1/orders/*


---

## 🔐 Authentication Flow

1. User logs in via `/auth/login`  
2. JWT token is issued  
3. Client sends token in headers  
4. API Gateway:
   - Validates JWT  
   - Injects user claims into context  
5. Services use context data  

---

## 🔄 Communication Patterns

### Synchronous
- HTTP/REST (Gin)  
- Optional gRPC for internal calls  

### Asynchronous
- Kafka Events:
  - UserCreated  
  - ProductUpdated  
  - OrderCreated  

---

## 🧠 Caching Strategy

- Cache-Aside Pattern  
- Redis used for:
  - Product caching  
  - Cart caching  
  - Rate limiting  

---

## 🔁 Event Flow Example (Order)

1. Order Service creates order  
2. Publishes `OrderCreated` event  
3. Consumers:
   - Notification Service → Sends notification  
   - Inventory Service → Updates stock  
   - Payment Service → Processes payment  

---

## 🛠️ Implementation Roadmap

### Phase 0: Foundation
- Go modules setup  
- Viper config  
- Zap logging  
- PostgreSQL setup  
- Basic Gin server  
- Docker setup  

### Phase 1: API Gateway
- Middleware chain:
  - Request ID  
  - Logging  
  - Recovery  
  - CORS  
  - Rate limiting  
  - JWT auth  

### Phase 2: User Service
- Register/Login  
- JWT generation  
- Protected routes  

### Phase 3: Product Service
- CRUD APIs  
- Filtering & pagination  
- Kafka events  

### Phase 4: Cart Service
- Cart operations  
- Transactions  

### Phase 5: Order Service
- Order creation  
- Saga pattern  
- Idempotency  

### Phase 6: Redis
- Cache implementation  

### Phase 7–8: Kafka & Workers
- Consumers  
- Retry logic  

### Phase 9: Production
- Error handling  
- OpenAPI docs  
- Circuit breakers  
- Observability  

### Phase 10: Testing & Deployment
- Unit & integration tests  
- Docker optimization  
- Kubernetes setup  

---

## 🧪 Testing Strategy

- Unit tests (service layer)  
- Integration tests (DB + API)  
- API testing (Postman / Curl)  

---

## 🐳 Running the Project

```bash
docker-compose up --build

Services
API Gateway → localhost:8080
User Service → localhost:8081
Product Service → localhost:8082
Cart Service → localhost:8083
Order Service → localhost:8084


🩺 Health Check
GET /health

Response:

{
  "status": "ok"
}
