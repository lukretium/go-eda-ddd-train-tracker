package domain

type WorkoutRepository interface {
	Save(workout *Workout) error
	FindByID(id string) (*Workout, error)
	// ListByUserID(userID string) ([]*Workout, error) // for future use
}
