Chat gpt  prompt
```
can you provide me full module layer for basket module with interface
```

ASCII Treedir

```
.
├── internal/
│   ├── dto/
│   │   └── auth.dto.go
│   ├── enums/
│   │   └── forgot-password.enum.go
│   ├── repository/
│   │   └── user.repository.go
│   ├── service/
│   │   └── user.service.go
│   ├── route/
│   │   └── user.route.go
│   ├── handler/
│   │   └── user.handler.go
│   ├── lib/
│   │   └── mail.go
│   ├── mail/
│   │   └── welcome.mail.go
│   ├── middleware/
│   │   ├── jwt.middleware.go
│   │   ├── role.middleware.go
│   │   ├── permission.middleware.go
│   │   └── cors.middleware.go
│   ├── seeder/
│   │   └── user.seeder.go
│   │   └── role.seeder.go
│   ├── migration/
│   │   └── user.migration.go
│   │   └── role.migration.go
├── container/
│   ├── user.container.go
│   └── role.container.go
├── config/
│   ├── database.go
│   └── app.go
├── databse/
│   ├── migration/
│   │    └── user.migration.go
│   ├── seeder/
│   │    └── user.seeder.go
├── main.go
└── .env
```

migration command 
```
go run main.go migrate
```

seed command 
```
go run main.go seed
```
