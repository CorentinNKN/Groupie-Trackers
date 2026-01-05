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
