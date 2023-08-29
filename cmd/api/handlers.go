package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/**
 * w представляет собой интерфейс, который обеспечивает методы для
 * записи ответа на HTTP-запрос.
 * Вы можете использовать этот интерфейс для задания статус-кода ответа,
 * установки заголовков ответа и записи тела ответа.
 * ---------
 * r представляет собой указатель на структуру http.Request, которая содержит всю информацию о входящем HTTP-запросе.
 * Эта структура включает в себя такие данные, как HTTP-метод (GET, POST и т.д.), заголовки запроса, тело запроса,
 * параметры запроса и так далее.
 * В этом конкретном примере r не используется, но если бы вы хотели обработать какие-либо детали запроса
 * (например, извлечь параметры из URL или прочитать тело запроса), вы бы использовали этот объект.
 */

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Links app",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllLinks(w http.ResponseWriter, r *http.Request) {
	var links []models.Link

	// maybe I need this in the future
	//added_at, _ := time.Parse("2023-01-01", "2023-01-01")

	doctor := models.Link{
		ID:          1,
		Title:       "Doctor",
		Description: "Hello world this is a description",
		Phone:       "89991231212",
		Email:       "hello@test.com",
		Category:    "health",
		Links:       "https://google.com",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	links = append(links, doctor)

	cat_doctor := models.Link{
		ID:          2,
		Title:       "Animal doctor",
		Description: "Hello world this is a description",
		Phone:       "89991231212",
		Email:       "hello123@test.com",
		Category:    "animal",
		Links:       "https://google.com",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	links = append(links, cat_doctor)

	out, err := json.Marshal(links)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
