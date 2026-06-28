# SpotSync API

A smart parking and EV charging reservation system for airports and malls.

**Live URL:** `https://spot-sync-q2uf.onrender.com`

---

## Features

- User registration and login with JWT authentication
- Role-based access control (driver / admin)
- Parking zone management (create, list, get by ID)
- Real-time available spot calculation per zone
- Spot reservation with concurrency-safe booking (transaction + row-level locking)
- View and cancel personal reservations
- Admin view of all reservations in the system

---

## Tech Stack

| Technology | Purpose |
|---|---|
| Go 1.26 | Primary language |
| Echo v4 | HTTP web framework |
| GORM | ORM for database access |
| PostgreSQL (NeonDB) | Relational database |
| JWT (golang-jwt/jwt v5) | Authentication tokens |
| bcrypt | Password hashing |
| go-playground/validator | Request validation |
| godotenv | Environment variable loading |

---

## Architecture

This project follows **Clean Architecture** with strict layer separation:

