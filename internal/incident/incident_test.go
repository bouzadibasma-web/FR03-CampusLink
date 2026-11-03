package incident

import (
	"testing"
	"time"
)

func TestIncidentValide(t *testing.T) {
	i := Incident{
		ID:          1,
		Title:       "Connexion impossible",
		Urgency:     5,
		Criticality: 3,
		Impact:      4,
		CreatedAt:   time.Now(),
	}

	err := i.Validate()

	if err != nil {
		t.Errorf("Incident valide refuse : %v", err)
	}
}

func TestTitreVide(t *testing.T) {
	i := Incident{
		ID:          1,
		Title:       "",
		Urgency:     3,
		Criticality: 3,
		Impact:      3,
		CreatedAt:   time.Now(),
	}

	err := i.Validate()

	if err == nil {
		t.Error("Le titre vide devrait etre refuse")
	}
}

func TestUrgenceInvalide(t *testing.T) {
	i := Incident{
		ID:          1,
		Title:       "Incident test",
		Urgency:     10,
		Criticality: 3,
		Impact:      3,
		CreatedAt:   time.Now(),
	}

	err := i.Validate()

	if err == nil {
		t.Error("Une urgence invalide devrait etre refusee")
	}
}
