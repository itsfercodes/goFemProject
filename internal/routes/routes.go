package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/itsfercodes/goFemProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	// Workouts routes
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.HandleGetWorkoutByID)

	return r
}
