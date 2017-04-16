package controllers

import (
	"net/http"
	"strconv"
	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"../models"
)

var decoder = schema.NewDecoder()

type Notebook struct {
	controller.BaseController
	Routes []string
}

func (t *Notebook) Modify() {
	Notebook := &models.Notebook{}
	NotebookID := t.Ctx.Params["id"]
	if NotebookID=="new" {

	} else {
		ID, err := strconv.Atoi(NotebookID)
		if err != nil {
			t.Ctx.Data["Message"] = err.Error()
			t.Ctx.Template = "error"
			t.HTML(http.StatusInternalServerError)
			return
		}
		t.Ctx.DB.Where("id = ?", ID).Find(&Notebook)
	}
	
	req := t.Ctx.Request()
	_ = req.ParseForm()
	Notebook.Title=req.PostForm["notebook[title]"][0]
	t.Ctx.DB.Save(&Notebook)
	t.Ctx.Redirect("/#editNotebook/"+NotebookID, http.StatusFound)
}

func (t *Notebook) Delete() {
	NotebookID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(NotebookID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	// we must delete all child notes first!
	t.Ctx.DB.Delete(&models.Note{}, "notebook_id = ?", ID)
	t.Ctx.DB.Delete(&models.Notebook{ID: ID})
	t.Ctx.Redirect("/#notebooks", http.StatusFound)
}

func NewNotebook() controller.Controller {
	return &Notebook{
		Routes: []string{
			"post;/api/modify-notebook/{id};Modify",
			"get;/api/delete-notebook/{id};Delete",
		},
	}
}
