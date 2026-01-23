package main

var LoadedLocations AllLocations
var LoadedRelations AllRelations

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"` // URL vers les lieux
}

type LocationData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type AllLocations struct {
	Index []LocationData `json:"index"`
}

// Structure pour une relation unique (liée à un artiste)
type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Structure pour l'index global des relations (si tu en as besoin)
type AllRelations struct {
	Index []RelationData `json:"index"`
}
