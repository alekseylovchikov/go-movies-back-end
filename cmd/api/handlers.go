package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	links, err := app.DB.AllLinks()
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(links)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
