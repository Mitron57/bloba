package main

import (
    "github.com/go-chi/chi/v5"
    "context"
    "github.com/google/uuid"
    "seems/application/services"
    "seems/domain/models"
)

func main() {
    router := chi.NewRouter()
    users := services.Users{Storage: make(map[uuid.UUID]*models.User)}
    ctx := context.WithValue(context.Background(), "users", &users)
    router.
}
