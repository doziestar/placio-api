package controller

import "placio-app/service"

type MenuController struct {
	menuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{menuService: menuService}
}
