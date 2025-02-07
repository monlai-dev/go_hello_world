# Cinema Ticket Booking System 🎬

A full-featured cinema ticket booking system built with Golang, featuring real-time seat availability updates and payment integration.

## Features ✨

### Core Functionality
- **User Authentication** 🔐
  - JWT-based registration/login
  - Bcrypt password hashing
- **Movie Management** 🎥
  - CRUD operations for movies/theaters
  - Showtime scheduling
- **Seat Booking** 💺
  - Real-time seat availability (Socket.IO)
  - 10-minute booking hold system
  - Concurrent booking prevention
- **Payment Integration** 💳
  - PayOS/VNPay integration
  - Payment status tracking
- **Ticket Management** 🎫
  - Booking cancellation
  - Booking history

### Advanced Features
- Redis caching for high-frequency data
- Transactional database operations
- Role-based access control (Admin/User)
- Dockerized development environment

## Tech Stack 🛠️

| Component              | Technology                          |
|------------------------|-------------------------------------|
| **Backend**            | Go 1.20+                           |
| **Database**           | PostgreSQL                          |
| **Authentication**      | JWT                                 |
| **Caching**            | Redis                               |
| **Real-time**          | Socket.IO                           |
| **Containerization**   | Docker & Docker Compose             |
| **Payment Gateway**     | PayOS/VNPay (REST API integration)  |
| **Logging**            | Zap/Slog                            |

## System Architecture 🏗️

### Service Breakdown
```mermaid
graph TD
    A[API Gateway] --> B[User Service]
    A --> C[Movie Service]
    A --> D[Booking Service]
    A --> E[Payment Service]
    D --> F[Redis]
    E --> G[Payment Gateways]
