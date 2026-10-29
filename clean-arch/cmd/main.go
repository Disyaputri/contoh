package main

import (
	"log"

	"clean-arch/internal/config"
	"clean-arch/internal/handlers"
	dbrepo "clean-arch/internal/repository/db"
	"clean-arch/internal/usecase"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := dbrepo.NewUserRepository(db)

	uc := usecase.NewUserUsecase(repo)

	handler := handlers.NewUserHandler(uc)

	handler.Run()
}