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
│   │   ├── forgot-password.go
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
├── stubs/
│   ├── module.container.stub
│   ├── module.dto.stub
│   ├── module.entity.stub
│   ├── module.handler.stub
│   ├── module.migration.stub
│   ├── module.repository.stub
│   ├── module.route.stub
│   ├── module.seeder.stub
│   └── module.service.stub
├── container/
│   ├── user.container.go
│   └── role.container.go
├── public/
│   └── mails
│         ├── forget-password-mail.html
│         └── welcome-mail.html
├── config/
│   ├── database.go
│   ├── mail.go
│   └── app.go
├── database/
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
create module command (entity, container, dto, route, handler, service, repository, migration, seeder)
Then you should add  module layers to container/container.go in Container struct
```
go run main.go genrate-file {MODUL_NAME}
```
