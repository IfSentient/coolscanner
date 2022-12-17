package scanners

import "coolscanner/pkg/models"

type SimpleMatcher struct {
}

func (s *SimpleMatcher) Process(data *models.SystemData) ([]models.Problem, error) {
	return []models.Problem{}, nil
}
