package sorting

import (
	"campuslink/internal/incident"
	"campuslink/internal/priority"
)

func MergeSort(incidents []incident.Incident) {

	if len(incidents) <= 1 {
		return
	}

	milieu := len(incidents) / 2

	gauche := make([]incident.Incident, milieu)
	droite := make([]incident.Incident, len(incidents)-milieu)

	copy(gauche, incidents[:milieu])
	copy(droite, incidents[milieu:])

	MergeSort(gauche)
	MergeSort(droite)

	merge(incidents, gauche, droite)
}

func merge(
	resultat []incident.Incident,
	gauche []incident.Incident,
	droite []incident.Incident,
) {

	i := 0
	j := 0
	k := 0

	for i < len(gauche) && j < len(droite) {

		scoreGauche := priority.CalculateScore(gauche[i])
		scoreDroite := priority.CalculateScore(droite[j])

		if scoreGauche >= scoreDroite {
			resultat[k] = gauche[i]
			i++
		} else {
			resultat[k] = droite[j]
			j++
		}

		k++
	}

	for i < len(gauche) {
		resultat[k] = gauche[i]
		i++
		k++
	}

	for j < len(droite) {
		resultat[k] = droite[j]
		j++
		k++
	}
}
