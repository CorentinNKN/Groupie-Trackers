package main

import "strings"

func FilterBySearch(query string, allArtists []Artist) []Artist {
	var results []Artist
	q := strings.ToLower(query) // On met tout en minuscule pour ignorer la casse

	for _, artist := range allArtists {
		// On vérifie si le nom de l'artiste contient la recherche
		if strings.Contains(strings.ToLower(artist.Name), q) {
			results = append(results, artist)
		}
	}
	return results
}
