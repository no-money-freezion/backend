package main

import (
	"backend/internal/app"
	"backend/internal/repo"
	"backend/internal/repo/inmemory"
	"backend/internal/repo/postgres"
	services "backend/internal/service"
	"backend/pkg"
	"fmt"

	"log"
	"os"
)

const dbTypeInmemory = "inmemory"

func main() {
	// TODO
	// 1) ctx for graceful shutdown etc
	// 2) metrics
	// 3) logging
	// 4) config from env/file
	// 5) signal processing
	// 6) real tests
	// 7) golangci-lint
	// 8) errs with stack traces
	// 9) TTL for urls maybe
	// 10) spam/ddos protection

	var err error
	var repoInstance repo.Repository

	dbType := os.Getenv("DB_TYPE")

	if dbType == dbTypeInmemory {
		repoInstance = inmemory.NewRepository()
	} else {

		host := pkg.GetEnvWithDefault("DB_HOST", "db")
		user := pkg.GetEnvWithDefault("DB_USER", "postgres")
		password := pkg.GetEnvWithDefault("DB_PASSWORD", "postgres")
		dbname := pkg.GetEnvWithDefault("DB_NAME", "postgres")
		sslmode := pkg.GetEnvWithDefault("DB_SSLMODE", "disable")

		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
			host, user, password, dbname, sslmode)

		repoInstance, err = postgres.NewRepository(connStr)
		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}
	}

	slotService := services.NewService(repoInstance)
	newApp := app.NewApp(slotService, repoInstance)

	err = newApp.Start()
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
}
