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

```bash
cd backend
go run main.go
```
Backend will run at : http://localhost:8080

How to run backend server on production build:
1. Rename the environment file:
```bash
mv env.example .env
```
2. Configure your environment variables in .env, especially update JWT_SECRET=supersecret
3. Build and run the backend using Docker: 
```bash
docker-compose up --build -d
```
4. Verify containers are running:
```bash
docker ps
```

Or build manually:
```bash
cd backend
docker build -t go-backend .
docker run -p 8080:8080 go-backend
```

How to run frontend on local:

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
1. Rename the environment file:
```bash
mv env.example .env
```
2. Configure your environment variables in .env, especially update JWT_SECRET=supersecret
3. Build and run the backend using Docker: 
```bash
docker-compose up --build -d
```
4. Verify containers are running:
```bash
docker ps
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
