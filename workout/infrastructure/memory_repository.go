package infrastructure

import (
	"sync"

	"workout/domain"
)

type MemoryWorkoutRepository struct {
	mu       sync.RWMutex
	workouts map[string]*domain.Workout
}

var _ domain.WorkoutRepository = (*MemoryWorkoutRepository)(nil) // compile-time check

func NewMemoryWorkoutRepository() *MemoryWorkoutRepository {
	return &MemoryWorkoutRepository{
		workouts: make(map[string]*domain.Workout),
	}
}

func (r *MemoryWorkoutRepository) Save(workout *domain.Workout) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.workouts[workout.ID] = workout
	return nil
}

func (r *MemoryWorkoutRepository) FindByID(id string) (*domain.Workout, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	w, ok := r.workouts[id]
	if !ok {
		return nil, nil
	}
	return w, nil
}

func (r *MemoryWorkoutRepository) ListByUserID(userID string) ([]*domain.Workout, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.Workout
	for _, w := range r.workouts {
		if w.UserID == userID {
			result = append(result, w)
		}
	}
	return result, nil
}
