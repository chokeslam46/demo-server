package routers

import (
	"github.com/astaxie/beego"
	"github.com/chokeslam46/demo-server/products"
	"github.com/chokeslam46/demo-server/users"
)

func init() {
	userController := &users.Controller{}
	beego.Router("/users", userController, "get:List")
	beego.Router("/users/:id", userController, "get:Get")

	productController := &products.Controller{}
	beego.Router("/products", productController, "get:List")
	beego.Router("/products/:id", productController, "get:Get")
}
