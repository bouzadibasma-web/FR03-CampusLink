package sorting

import (
	"campuslink/internal/incident"
	"campuslink/internal/priority"
)

func BubbleSort(incidents []incident.Incident) {

	n := len(incidents)

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			score1 := priority.CalculateScore(incidents[j])
			score2 := priority.CalculateScore(incidents[j+1])

			if score1 < score2 {

				incidents[j], incidents[j+1] =
					incidents[j+1], incidents[j]
			}
		}
	}
}
