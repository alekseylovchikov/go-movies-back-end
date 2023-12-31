package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/**
* Это определение метода routes для типа application.
* (app *application) - это получатель метода. Это означает,
* что мы можем вызывать этот метод на экземплярах (или, точнее,
* на указателях экземпляров) типа application.
 */

func (app *application) routes() http.Handler {
	/**
	* Создается новый роутер с помощью пакета chi.
	* chi - это популярная сторонняя библиотека для роутинга в Go.
	 */
	mux := chi.NewRouter()

	/**
	* Это добавляет промежуточное ПО (или middleware) к роутеру.
	* В этом конкретном случае промежуточное ПО называется Recoverer
	* и, вероятно, предназначено для восстановления после паник в Go,
	* чтобы предотвратить падение всего сервера из-за неожиданных ошибок.
	* Паника в Go - это событие, которое может прервать выполнение программы,
	* и Recoverer помогает грациозно обрабатывать такие паники, преобразуя их,
	* например, в HTTP-ответы с кодом ошибки 500.
	 */
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/links", app.AllLinks)

	return mux
}
