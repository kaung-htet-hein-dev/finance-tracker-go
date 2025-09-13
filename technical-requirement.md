## Technical Requirements - Backend (Go) for AI-Powered Personal Finance App

### 1. API Architecture

- Use Go with a minimal web framework (Echo).
- Expose RESTful endpoints for all core operations.

### 2. Authentication

- Implement JWT-based authentication for secure user sessions.
- Endpoints for sign up, login, and logout.
- Passwords must be hashed and salted before storage.

### 3. Database

- Use SQLite (for simplicity).
- Tables:
  - `users`: id, email, password_hash, created_at
  - `transactions`: id, user_id, amount, category, type (income/expense), date, note, created_at
  - `categories`: id, name, user_id (optional, if supporting custom categories)

### 4. Endpoints

- `POST /api/register`: Create new user.
- `POST /api/login`: Authenticate user.
- `GET /api/dashboard`: Get user balance and recent transactions.
- `POST /api/transactions`: Add new transaction.
- `GET /api/transactions`: List transactions (support query filters: date range, category).
- `DELETE /api/transactions/:id`: Delete transaction.
- `GET /api/insights`: Return AI-powered monthly summary and suggestions.
- (Optional) `GET/POST/DELETE /api/categories`: Manage categories.

### 5. AI/Insights Logic

- Implement basic analysis in Go:
  - Aggregate spending by category/month.
  - Generate simple tips (e.g., "You spent most on X last month. Try to save more.").
- (Optional) Use a lightweight ML library for pattern detection or automated categorization.

### 6. Validation & Error Handling

- Validate all incoming data (amount, date, etc.).
- Return appropriate HTTP status codes and error messages.

### 7. Security

- Secure all endpoints; require authentication except registration/login.
- Sanitize inputs to prevent SQL injection and other vulnerabilities.

### 8. Documentation

- Provide API documentation (OpenAPI spec or markdown).
- Include setup instructions for local development.

### 9. Testing (Optional)

- Write unit tests for core business logic and API endpoints.

---

**Note:**  
Keep the codebase minimal and easy to understand for portfolio purposes. Focus on clarity, security, and maintainability.
