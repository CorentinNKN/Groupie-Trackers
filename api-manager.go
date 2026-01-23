package main

import (
	"encoding/json"
	"net/http"
)

// GetArtistes va chercher la liste sur le web
func GetArtistes() ([]Artist, error) {
	// On appelle l'API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // On ferme la connexion à la fin de la fonction

	// On transforme le JSON en liste d'objets Artiste
	var artistes []Artist
	err = json.NewDecoder(resp.Body).Decode(&artistes)
	return artistes, err
}

// GetRelations fonctionne de la même manière mais pour le lien concerts/dates
func GetRelations() (RelationData, error) {
	var data RelationData
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	return data, err
}

func GetAllCities() []string {
	// 1. Appelle l'endpoint /locations
	// 2. Parcoure tous les artistes
	// 3. Stocke les villes dans une Map pour éviter les doublons
	cityMap := make(map[string]bool)
	cityMap["Toutes"] = true

	// Exemple de boucle sur tes données récupérées
	for _, data := range LoadedLocations.Index {
		for _, city := range data.Locations {
			cityMap[city] = true
		}
	}

	var cities []string
	for city := range cityMap {
		cities = append(cities, city)
	}
	return cities
}
