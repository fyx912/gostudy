package controllers

import (
	"github.com/astaxie/beego"
)

type NavigationController struct {
	beego.Controller
}

func (this *NavigationController) Dashboard() {
	this.TplName = "index.html"
}

func (this *NavigationController) Index() {
	this.TplName = "index.html"
}

func (this *NavigationController) MeCenter() {
	this.TplName = "meCenter.html"
}

func (this *NavigationController) System() {
	this.TplName = "system.html"
}
func (this *NavigationController) Forms() {
	this.TplName = "forms.html"
}
func (this *NavigationController) Tables() {
	this.TplName = "tables.html"
}
func (this *NavigationController) Charts() {
	this.TplName = "charts.html"
}
func (this *NavigationController) Typography() {
	this.TplName = "typography.html"
}
func (this *NavigationController) Elements() {
	this.TplName = "elements.html"
}
