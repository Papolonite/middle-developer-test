package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	// setting up pgx to database
	pool, err := pgxpool.New(ctx, "postgresql://dev:dev@localhost:5102/postgres?sslmode=disable")

	if err != nil {
		log.Error().Err(err).Msg("Unable to connect do database")
	}

	// setting up fiber
	app := fiber.New()

}
