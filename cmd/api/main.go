package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	var app application

	app.Domain = "example.com"

	/**
	 * Эта функция определяет строковый флаг с указанным именем (в данном случае "dsn"), значением по умолчанию и описанием. Когда программа будет запущена, если пользователь предоставит аргумент -dsn=свое_значение, то значение этого аргумента будет сохранено в переменной app.DSN. Если пользователь не предоставит аргумент -dsn, то в app.DSN будет сохранено значение по умолчанию ("host=localhost port=5432 ... connect_timeout=5")
	 */
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=links sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	/**
	 * В этом Go коде, ключевое слово defer используется для того, чтобы гарантировать,
	 * что метод Close будет вызван на объекте app.DB непосредственно перед тем, как функция
	 * main завершит свою работу.
	 */

	log.Println("Starting server on port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
