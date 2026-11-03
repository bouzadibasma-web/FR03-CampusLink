package main

import (
	"fmt"
	"time"

	"campuslink/internal/incident"
	"campuslink/internal/priority"
	"campuslink/internal/sorting"
)

func main() {

	fmt.Println("===================================")
	fmt.Println("          CAMPUSLINK")
	fmt.Println("     Gestion des incidents")
	fmt.Println("===================================")

	incidents := []incident.Incident{
		{ID: 1, Title: "Connexion impossible", Urgency: 5, Criticality: 3, Impact: 4, CreatedAt: time.Now()},
		{ID: 2, Title: "Bug affichage notes", Urgency: 2, Criticality: 2, Impact: 3, CreatedAt: time.Now()},
		{ID: 3, Title: "Serveur indisponible", Urgency: 5, Criticality: 5, Impact: 5, CreatedAt: time.Now()},
		{ID: 4, Title: "Mot de passe oublie", Urgency: 1, Criticality: 2, Impact: 2, CreatedAt: time.Now()},
		{ID: 5, Title: "Application lente", Urgency: 3, Criticality: 3, Impact: 4, CreatedAt: time.Now()},
	}

	fmt.Println("\n===== LISTE INITIALE =====")

	for _, inc := range incidents {
		fmt.Println("ID :", inc.ID)
		fmt.Println("Incident :", inc.Title)
		fmt.Println("Score :", priority.CalculateScore(inc))
		fmt.Println("-------------------------")
	}

	incidentsBubble := make([]incident.Incident, len(incidents))
	copy(incidentsBubble, incidents)

	sorting.BubbleSort(incidentsBubble)

	fmt.Println("\n===== CLASSEMENT BUBBLE SORT =====")

	for position, inc := range incidentsBubble {
		fmt.Println("Position :", position+1)
		fmt.Println("Incident :", inc.Title)
		fmt.Println("Score :", priority.CalculateScore(inc))
		fmt.Println("-------------------------")
	}

	incidentsMerge := make([]incident.Incident, len(incidents))
	copy(incidentsMerge, incidents)

	sorting.MergeSort(incidentsMerge)

	fmt.Println("\n===== CLASSEMENT MERGE SORT =====")

	for position, inc := range incidentsMerge {
		fmt.Println("Position :", position+1)
		fmt.Println("Incident :", inc.Title)
		fmt.Println("Score :", priority.CalculateScore(inc))
		fmt.Println("-------------------------")
	}

	copieBubble := make([]incident.Incident, len(incidents))
	copy(copieBubble, incidents)

	debutBubble := time.Now()
	sorting.BubbleSort(copieBubble)
	dureeBubble := time.Since(debutBubble)

	copieMerge := make([]incident.Incident, len(incidents))
	copy(copieMerge, incidents)

	debutMerge := time.Now()
	sorting.MergeSort(copieMerge)
	dureeMerge := time.Since(debutMerge)

	fmt.Println("\n===== COMPARAISON SUR 5 INCIDENTS =====")
	fmt.Println("Bubble Sort :", dureeBubble)
	fmt.Println("Merge Sort  :", dureeMerge)

	grandsIncidents := make([]incident.Incident, 10000)

	for i := 0; i < 10000; i++ {
		grandsIncidents[i] = incident.Incident{
			ID:          i + 1,
			Title:       fmt.Sprintf("Incident %d", i+1),
			Urgency:     (i % 5) + 1,
			Criticality: ((i + 2) % 5) + 1,
			Impact:      ((i + 3) % 5) + 1,
			CreatedAt:   time.Now(),
		}
	}

	copieGrandeBubble := make([]incident.Incident, len(grandsIncidents))
	copy(copieGrandeBubble, grandsIncidents)

	debut := time.Now()
	sorting.BubbleSort(copieGrandeBubble)
	tempsBubble := time.Since(debut)

	copieGrandeMerge := make([]incident.Incident, len(grandsIncidents))
	copy(copieGrandeMerge, grandsIncidents)

	debut = time.Now()
	sorting.MergeSort(copieGrandeMerge)
	tempsMerge := time.Since(debut)

	fmt.Println("\n===== TEST AVEC 10000 INCIDENTS =====")
	fmt.Println("Bubble Sort :", tempsBubble)
	fmt.Println("Merge Sort  :", tempsMerge)

	fmt.Println("\nFin du programme.")
}
