package controllers

import (
	"net/http"
	"strconv"
	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"../models"
)

var decoder = schema.NewDecoder()

type Note struct {
	controller.BaseController
	Routes []string
}

func (t *Note) Home() {
	t.Ctx.Template = "index"
	t.HTML(http.StatusOK)
}
func (t *Note) GetAll() {
	notes := []*models.Note{}
	t.Ctx.DB.Order("created_at desc").Find(&notes)
	t.Ctx.Data["List"] = notes[:len(notes)-1]
	t.Ctx.Data["Last"] = notes[len(notes)-1]
	t.Ctx.Template = "notes"
	t.JSON(http.StatusOK)
}

func (t *Note) Create() {
	Note := &models.Note{}
	req := t.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(Note, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(Note)
	t.Ctx.Redirect("/", http.StatusFound)
}

func (t *Note) Delete() {
	NoteID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(NoteID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	t.Ctx.DB.Delete(&models.Note{ID: ID})
	t.Ctx.Redirect("/", http.StatusFound)
}

func NewNote() controller.Controller {
	return &Note{
		Routes: []string{
			"get;/;Home",
			"get;/notes.json;GetAll",
			//"post;/create;Create",
			//"get;/delete/{id};Delete",
		},
	}
}
