package main

import (
	//"errors"
	"net/http"

	"monkey.anbirilo.net/internal/models"
	// "github.com/julienschmidt/httprouter"
	// "snippetbox.anbirilo.net/internal/models"
	// "snippetbox.anbirilo.net/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 1. Сначала получаем готовую строку, передавая модель и ID пользователя (например, 1)
	expensesText := models.GetExpensesString(app.expenses, 1)

	// 2. Превращаем строку в срез байт и отдаем в ответ
	w.Write([]byte(expensesText))
}
