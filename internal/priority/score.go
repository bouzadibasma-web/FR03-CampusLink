package priority

import "campuslink/internal/incident"

func CalculateScore(i incident.Incident) int {

	score := (i.Urgency * 3) +
		(i.Criticality * 2) +
		i.Impact

	return score
}
