# Analyse des performances et complexité – FR03

## 1. Objectif

L'objectif est de comparer deux algorithmes de tri permettant de classer les incidents informatiques selon leur score de priorité décroissant.

Les algorithmes étudiés sont Bubble Sort et Merge Sort.

## 2. Résultats expérimentaux

Un test a été réalisé avec 10 000 incidents.

| Algorithme | Temps d'exécution |
|---|---:|
| Bubble Sort | 230,3787 ms |
| Merge Sort | 4,698 ms |

Merge Sort est environ 49 fois plus rapide que Bubble Sort lors de cette expérience.

## 3. Analyse Big-O

### Bubble Sort

- Meilleur cas : O(n), avec une optimisation permettant de détecter l'absence d'échanges.
- Cas moyen : O(n²).
- Pire cas : O(n²).

Bubble Sort compare les éléments deux à deux et effectue des échanges jusqu'à obtenir un classement correct.

### Merge Sort

- Meilleur cas : O(n log n).
- Cas moyen : O(n log n).
- Pire cas : O(n log n).

Merge Sort divise la liste en plusieurs parties, trie chaque partie puis fusionne les résultats.

## 4. Optimisation et comparaison

Bubble Sort est simple à comprendre et à programmer, mais devient lent lorsque le nombre d'incidents augmente.

Merge Sort est plus performant pour traiter un grand volume de données.

Il nécessite cependant une mémoire supplémentaire pour effectuer les fusions.

## 5. Conclusion

Les résultats expérimentaux montrent que Merge Sort est environ 49 fois plus rapide que Bubble Sort lors du test réalisé avec 10 000 incidents.

Nous retenons donc Merge Sort comme solution optimisée pour le classement des incidents dans CampusLink.