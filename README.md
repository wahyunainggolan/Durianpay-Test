# fullstack app

Explain your service in here. This is fulltsack project related Payment using golang as backend and nuxt as frontend....

list of tools version of your machine:

```bash
go version go1.25.5 darwin/arm64
node v24.13.1
```

Install all related requirements:

--- backend 
```bash
cd backend
go mod tidy
```
--- frontend 
```bash
cd frontend
npm install
```

How to run backend server on local:
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

How to run backend server on production build:
1. Copy and rename env.example from root to backend and frontend folder/directory:

--- in root 
```bash
cp env.example .env
```
--- frontend
```bash
cp env.example frontend/.env
```
--- backend
```bash
cp env.example backend/.env
```
2. Configure your environment variables in .env, especially update JWT_SECRET=supersecret
3. Build and run the backend using Docker: 
```bash
docker-compose build --no-cache
docker-compose up
```

Or build manually:
```bash
cd backend
docker build -t go-backend .
docker run -p 8080:8080 go-backend
```

How to run frontend on local:
1. Copy and rename env.example from root to frontend directory
```bash
cp env.example frontend/.env
```
2. move to frontend directory and run backend
```bash
cd frontend
npm run dev
```
Frontend will run at: http://localhost:3000

How to run frontend on production build:

A. Run using Docker (Nginx) : 
```bash
docker build -t vue-frontend .
docker run -p 3000:80 vue-frontend
```

B. Or via docker-compose:
1. Copy and rename env.example to backend and frontend folder:
--- frontend
```bash
cp env.example frontend/.env
```
--- backend
```bash
cp env.example backend/.env
```
2. Configure your environment variables in .env, especially update JWT_SECRET=supersecret
3. Build and run the frontend using Docker: 
```bash
docker-compose build --no-cache
docker-compose up
```

To checking openapi documentations, you can visit this url after backend running.
After backend is running, you can access:
```bash
http://localhost:8080/openapi.yaml
```

Login to frontend by visiting:

```bash
http://localhost:3000/login
```

evidences: Add video evidences of your service
see backend [README.md](backend/README.md)
see frontend [README.md](frontend/README.md)
