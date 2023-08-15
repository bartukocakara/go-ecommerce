package container

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	DB                        *gorm.DB
	UserHandler               handler.UserHandler
	UserService               service.UserService
	UserRepository            repository.UserRepository
	RoleHandler               handler.RoleHandler
	RoleService               service.RoleService
	RoleRepository            repository.RoleRepository
	ForgotPasswordRepository  repository.ForgotPasswordTokenRepository
	AuthHandler               handler.AuthHandler
	AuthService               service.AuthService
	ProductHandler            handler.ProductHandler
	ProductService            service.ProductService
	ProductRepository         repository.ProductRepository
	CategoryHandler           handler.CategoryHandler
	CategoryService           service.CategoryService
	CategoryRepository        repository.CategoryRepository
	ProductCategoryHandler    handler.ProductCategoryHandler
	ProductCategoryService    service.ProductCategoryService
	ProductCategoryRepository repository.ProductCategoryRepository
	BasketHandler             handler.BasketHandler
	BasketService             service.BasketService
	BasketRepository          repository.BasketRepository
	BasketItemHandler         handler.BasketItemHandler
	BasketItemService         service.BasketItemService
	BasketItemRepository      repository.BasketItemRepository
	OrderHandler              handler.OrderHandler
	OrderService              service.OrderService
	OrderRepository           repository.OrderRepository
	OrderItemHandler          handler.OrderItemHandler
	OrderItemService          service.OrderItemService
	OrderItemRepository       repository.OrderItemRepository
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		DB: db,
	}
}
