package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Epic55/booking_with_timezones/internal/handlers"
	"github.com/Epic55/booking_with_timezones/internal/repository"
)

func (a *App) initServer() error {
	a.repository = repository.NewRepository(a.logger, a.dbConn)

	err := a.repository.Migrate()
	if err != nil {
		return fmt.Errorf("failed to migrate db: %w", err)
	}

	a.handlers = handlers.NewHandlers(a.logger, a.repository)

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/api/v1/bookings/:workshop_id", a.handlers.CreateBooking)
	router.GET("/api/v1/bookings/:workshop_id", a.handlers.ListBookings)

	a.router = router

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", "8080"),
		Handler: router,
	}

	a.http = server
	a.closers = append(a.closers, func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return a.http.Shutdown(ctx)
	})

	return nil
}
