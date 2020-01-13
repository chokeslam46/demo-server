package users

import (
	"github.com/astaxie/beego"
)

// Controller of user modules
type Controller struct {
	beego.Controller
}

// List users
func (c *Controller) List() {
	users := GetUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

// Get user
func (c *Controller) Get() {
	id, _ := c.GetInt64(":id")
	user := GetUser(id)
	c.Data["json"] = user
	c.ServeJSON()
}
