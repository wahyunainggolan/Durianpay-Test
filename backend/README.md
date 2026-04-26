A. How to run backend server on local:
1. Copy and rename env.example from root to backend directory
```bash
cp env.example backend/.env
```
2. move to backend directory and run backend
```bash
cd backend
go run main.go
```
Backend will run at : http://localhost:8080

B. API Type:
- POST /dashboard/v1/auth/login {email,password}
- GET /dashboard/v1/payments?sort=sort,status=status,id=id

C. How to run backend unit testing : 
Run all tests: 
```bash
go test ./...
```
Run tests with coverage report:
```bash
go test ./... -cover
```