package main

import (
	"GoREST/internals/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
