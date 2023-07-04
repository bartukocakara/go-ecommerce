ASCII Treedir

```
.
├── internal/
│   ├── repository/
│   │   └── user.repository.go
│   ├── service/
│   │   └── user.service.go
│   ├── route/
│   │   └── user.route.go
│   ├── handler/
│   │   └── user.handler.go
│   ├── middleware/
│   │   ├── jwt.middleware.go
│   │   ├── role-permission.middleware.go
│   │   └── cors.middleware.go
│   ├── seeder/
│   │   └── user.seeder.go
│   │   └── role.seeder.go
│   └── migration/
│   │   └── user.migration.go
│   │   └── role.migration.go
├── container/
│   ├── user.container.go
│   └── role.container.go
├── config/
│   ├── database.go
│   └── app.go
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
