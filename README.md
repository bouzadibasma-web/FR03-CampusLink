# FR03 – CampusLink

## Gestion et priorisation des incidents

### 1. Présentation du projet

Ce projet consiste à développer un programme en langage Go permettant de gérer et de prioriser des incidents informatiques.

L'objectif est de classer les incidents selon leur niveau de priorité afin de faciliter leur traitement.

### 2. Technologies utilisées

- Langage : Go
- Éditeur : Visual Studio Code
- Terminal : PowerShell

### 3. Structure du projet

```text
FR03-CampusLink/
├── data/
├── docs/
│   └── analyse.md
├── internal/
│   ├── incident/
│   │   ├── incident.go
│   │   └── incident_test.go
│   ├── priority/
│   │   ├── score.go
│   │   └── score_test.go
│   └── sorting/
│       ├── bubble.go
│       ├── merge.go
│       └── sorting_test.go
├── go.mod
├── main.go
└── README.md
```

### 4. Calcul du score de priorité

Le score est calculé avec la formule :

Score = (Urgence × 3) + (Criticité × 2) + Impact

### 5. Algorithmes utilisés

**Bubble Sort :**

Algorithme de tri dont la complexité moyenne et maximale est O(n²).

**Merge Sort :**

Algorithme de tri basé sur la division et la fusion, avec une complexité O(n log n).

### 6. Tests et performances

Le programme a été testé avec 10 000 incidents.

- Bubble Sort : 230,3787 ms
- Merge Sort : 4,698 ms

Merge Sort est environ 49 fois plus rapide lors de cette expérience.

### 7. Exécution du programme

Pour lancer le programme :

```powershell
go run .
```

Pour lancer les tests :

```powershell
go test ./...
```

### 8. Conclusion

Ce projet permet de comprendre les notions d'algorithmie, de complexité, de tri et d'optimisation en utilisant le langage Go.