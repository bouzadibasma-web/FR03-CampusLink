package priority

import (
	"testing"

	"campuslink/internal/incident"
)

func TestCalculateScore(t *testing.T) {

	tests := []struct {
		nom       string
		urgence   int
		criticite int
		impact    int
		attendu   int
	}{
		{"Incident prioritaire", 5, 3, 4, 25},
		{"Incident moyen", 3, 3, 4, 19},
		{"Incident faible", 1, 2, 2, 9},
	}

	for _, test := range tests {

		t.Run(test.nom, func(t *testing.T) {

			i := incident.Incident{
				Urgency:     test.urgence,
				Criticality: test.criticite,
				Impact:      test.impact,
			}

			resultat := CalculateScore(i)

			if resultat != test.attendu {
				t.Errorf(
					"Score incorrect : obtenu %d, attendu %d",
					resultat,
					test.attendu,
				)
			}
		})
	}
}
