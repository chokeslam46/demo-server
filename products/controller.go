package products

import (
	"github.com/astaxie/beego"
)

// Controller of product modules
type Controller struct {
	beego.Controller
}

// List products
func (c *Controller) List() {
	products := GetProducts()
	c.Data["json"] = products
	c.ServeJSON()
}

// Get products
func (c *Controller) Get() {
	id, _ := c.GetInt64(":id")
	product := GetProduct(id)
	c.Data["json"] = product
	c.ServeJSON()
}
