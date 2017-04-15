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
func (t *Note) GetOneNote() {
	NoteID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(NoteID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	Note := &models.Note{}
	t.Ctx.DB.Where("id = ?", ID).Find(&Note)
	t.Ctx.Data["Last"] = Note
	t.Ctx.Template = "single"
	t.JSON(http.StatusOK)
}

func (t *Note) Modify() {
	Note := &models.Note{}
	NoteID := t.Ctx.Params["id"]
	if NoteID=="new" {

	} else {
		ID, err := strconv.Atoi(NoteID)
		if err != nil {
			t.Ctx.Data["Message"] = err.Error()
			t.Ctx.Template = "error"
			t.HTML(http.StatusInternalServerError)
			return
		}
		t.Ctx.DB.Where("id = ?", ID).Find(&Note)
	}
	
	req := t.Ctx.Request()
	_ = req.ParseForm()
	Note.Body=req.PostForm["note[body]"][0]
	Note.Title=req.PostForm["note[title]"][0]
	Note.NotebookID,_=strconv.Atoi(req.PostForm["note[notebook]"][0])
	t.Ctx.DB.Save(&Note)
	t.Ctx.Redirect("/#editNote/"+NoteID, http.StatusFound)
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
	t.Ctx.Redirect("/#notes", http.StatusFound)
}

func NewNote() controller.Controller {
	return &Note{
		Routes: []string{
			"get;/;Home",
			"get;/api/data.json;GetAll",
			"get;/api/get-one/{id};GetOneNote",
			"post;/api/modify-note/{id};Modify",
			"get;/api/delete-note/{id};Delete",
		},
	}
}
