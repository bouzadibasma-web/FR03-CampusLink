package incident

import (
	"fmt"
	"strings"
	"time"
)

// Incident represente un probleme signale sur CampusLink.
type Incident struct {
	ID          int
	Title       string
	Urgency     int
	Criticality int
	Impact      int
	CreatedAt   time.Time
}

// Validate verifie les informations de l'incident.
func (i Incident) Validate() error {

	if i.ID <= 0 {
		return fmt.Errorf("identifiant invalide")
	}

	if strings.TrimSpace(i.Title) == "" {
		return fmt.Errorf("le titre est obligatoire")
	}

	if i.Urgency < 1 || i.Urgency > 5 {
		return fmt.Errorf("urgence invalide")
	}

	if i.Criticality < 1 || i.Criticality > 5 {
		return fmt.Errorf("criticite invalide")
	}

	if i.Impact < 1 || i.Impact > 5 {
		return fmt.Errorf("impact invalide")
	}

	if i.CreatedAt.IsZero() {
		return fmt.Errorf("date de creation invalide")
	}

	return nil
}
