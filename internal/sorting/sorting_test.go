package sorting

import (
	"testing"
	"time"

	"campuslink/internal/incident"
	"campuslink/internal/priority"
)

func verifierTri(t *testing.T, incidents []incident.Incident) {

	t.Helper()

	for i := 0; i < len(incidents)-1; i++ {

		score1 := priority.CalculateScore(incidents[i])
		score2 := priority.CalculateScore(incidents[i+1])

		if score1 < score2 {
			t.Errorf(
				"Liste mal triee : score %d avant score %d",
				score1,
				score2,
			)
		}
	}
}

func creerIncidents() []incident.Incident {

	return []incident.Incident{
		{
			ID:          1,
			Title:       "Incident A",
			Urgency:     2,
			Criticality: 2,
			Impact:      2,
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Incident B",
			Urgency:     5,
			Criticality: 5,
			Impact:      5,
			CreatedAt:   time.Now(),
		},
		{
			ID:          3,
			Title:       "Incident C",
			Urgency:     1,
			Criticality: 1,
			Impact:      1,
			CreatedAt:   time.Now(),
		},
	}
}

func TestBubbleSort(t *testing.T) {

	incidents := creerIncidents()

	BubbleSort(incidents)

	verifierTri(t, incidents)
}

func TestMergeSort(t *testing.T) {

	incidents := creerIncidents()

	MergeSort(incidents)

	verifierTri(t, incidents)
}

func TestListeVide(t *testing.T) {

	var incidents []incident.Incident

	BubbleSort(incidents)
	MergeSort(incidents)

	if len(incidents) != 0 {
		t.Error("La liste devrait rester vide")
	}
}

func TestUnSeulIncident(t *testing.T) {

	incidents := []incident.Incident{
		{
			ID:          1,
			Title:       "Incident unique",
			Urgency:     3,
			Criticality: 3,
			Impact:      3,
			CreatedAt:   time.Now(),
		},
	}

	BubbleSort(incidents)
	MergeSort(incidents)

	if len(incidents) != 1 {
		t.Error("La liste devrait contenir un incident")
	}
}

func TestScoresEgaux(t *testing.T) {

	incidents := []incident.Incident{
		{
			ID:          1,
			Title:       "Incident A",
			Urgency:     3,
			Criticality: 3,
			Impact:      3,
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Incident B",
			Urgency:     3,
			Criticality: 3,
			Impact:      3,
			CreatedAt:   time.Now(),
		},
	}

	BubbleSort(incidents)
	verifierTri(t, incidents)

	MergeSort(incidents)
	verifierTri(t, incidents)
}

func TestBubbleEtMergeMemeResultat(t *testing.T) {

	incidents1 := creerIncidents()
	incidents2 := creerIncidents()

	BubbleSort(incidents1)
	MergeSort(incidents2)

	if len(incidents1) != len(incidents2) {
		t.Fatal("Les deux listes ont des tailles differentes")
	}

	for i := range incidents1 {

		score1 := priority.CalculateScore(incidents1[i])
		score2 := priority.CalculateScore(incidents2[i])

		if score1 != score2 {
			t.Errorf("Difference a la position %d", i)
		}
	}
}
