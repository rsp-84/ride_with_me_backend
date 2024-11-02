package resolvers

import (
	"fmt"

	"github.com/rsp84/ride-with-me-backend/graph/models"
)

func UserQuery(r *queryResolver) ([]*models.User, error) {
	var users []*models.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", result.Error)
	}
	return users, nil
}
