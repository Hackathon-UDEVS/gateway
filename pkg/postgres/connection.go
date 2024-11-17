package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Hackaton-UDEVS/gateway/internal/config"
	"github.com/Hackaton-UDEVS/gateway/internal/logs"
)

func ConnectionDB(cfg config.Config) (*sql.DB, error) {

	log, _ := logs.NewLogger()

	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHOST, cfg.DBPORT, cfg.DBUSER, cfg.DBPASSWORD, cfg.DBNAME)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Error("Error while connection postgres")
	}
	err = db.Ping()
	if err != nil {
		log.Error("Error while pinging postgres")
	}
	log.Info("Successfully connected to postgres")
	return db, nil
}
