package main

// Cette structure correspond à ce que l'API envoie pour un artiste.
// On utilise des "tags" `json:"..."` pour dire à Go quel champ JSON va dans quelle variable.
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    string   `json:"relations"` // URL vers les concerts
}
