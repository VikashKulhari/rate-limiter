# 🚦 Go Rate Limiter Middleware (Sliding Window + Redis)

> A fast, flexible and JWT-aware rate limiter for Go APIs using Redis sorted sets and the sliding window algorithm.

---

## 📸 Overview

This package helps you rate limit requests in Go REST APIs by user (JWT) and IP address. It supports:

- 🚀 Sliding window algorithm for smoother request limits
- 🔒 JWT claims support: `userID`, `role`, `email`, `exp`
- 🌐 Per-user and per-IP tracking
- 🧠 Redis-based key expiration to prevent memory leaks
- 🧪 Full unit tests with `redismock`

---

## ⚡ Quick Start

### 🔧 Install

```bash
go get github.com/VikashKulhari/ratelimiter
