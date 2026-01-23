package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	AllArtists  []Artist
	DisplayGrid *fyne.Container
	// On stocke les références des widgets pour lire leurs valeurs
	memberChecks   = make(map[int]*widget.Check)
	creationSlider *widget.Slider
	albumEntry     *widget.Entry
	citySelect     *widget.Select
)

func SetupUI(w fyne.Window) {
	// Initialisation de la grille d'affichage
	DisplayGrid = container.NewGridWithColumns(3) // 3 colonnes d'artistes

	// Création du panneau de filtres
	filterPanel := createFilterPanel()

	// Layout principal : Filtres à gauche (Scrollable), Artistes à droite
	mainContent := container.NewHSplit(
		container.NewVScroll(filterPanel),
		container.NewVScroll(DisplayGrid),
	)
	mainContent.Offset = 0.2 // 20% pour les filtres, 80% pour les artistes

	w.SetContent(mainContent)
	RefreshDisplay() // Premier affichage
}

func createFilterPanel() *fyne.Container {
	// 1. Membres
	membersBox := container.NewVBox(widget.NewLabel("Nombre de membres :"))
	for i := 1; i <= 7; i++ {
		val := i
		cb := widget.NewCheck(fmt.Sprintf("%d membres", i), func(b bool) { RefreshDisplay() })
		memberChecks[val] = cb
		membersBox.Add(cb)
	}

	// 2. Date de création
	creationSlider = widget.NewSlider(1950, 2024)
	creationLabel := widget.NewLabel("Année min : 1950")
	creationSlider.OnChanged = func(v float64) {
		creationLabel.SetText(fmt.Sprintf("Année min : %.0f", v))
		RefreshDisplay()
	}

	// 3. Premier Album
	albumEntry = widget.NewEntry()
	albumEntry.SetPlaceHolder("Année album...")
	albumEntry.OnChanged = func(s string) { RefreshDisplay() }

	// 4. Localisation
	citySelect = widget.NewSelect(GetAllCities(), func(s string) { RefreshDisplay() })
	citySelect.PlaceHolder = "Choisir une ville"

	return container.NewVBox(
		widget.NewLabel("--- FILTRES ---"),
		membersBox,
		widget.NewSeparator(),
		creationLabel, creationSlider,
		widget.NewSeparator(),
		widget.NewLabel("Premier Album :"),
		albumEntry,
		widget.NewSeparator(),
		widget.NewLabel("Ville :"),
		citySelect,
	)
}

func RefreshDisplay() {
	DisplayGrid.Objects = nil // On vide tout

	for _, artist := range AllArtists {
		// Logique de filtrage
		if !matchMembers(artist) {
			continue
		}
		if float64(artist.CreationDate) < creationSlider.Value {
			continue
		}
		if albumEntry.Text != "" && !strings.Contains(artist.FirstAlbum, albumEntry.Text) {
			continue
		}
		if !matchCity(artist) {
			continue
		}

		// Si OK, on ajoute la carte de l'artiste
		DisplayGrid.Add(widget.NewLabel(artist.Name)) // Remplace par ta fonction CreateArtistCard
	}
	DisplayGrid.Refresh()
}

func matchMembers(a Artist) bool {
	checkedCount := 0
	for _, cb := range memberChecks {
		if cb.Checked {
			checkedCount++
		}
	}
	if checkedCount == 0 {
		return true
	} // Si rien n'est coché, on affiche tout
	return memberChecks[len(a.Members)].Checked
}

func matchCity(a Artist) bool {
	if citySelect.Selected == "" || citySelect.Selected == "Toutes" {
		return true
	}
	// Ici tu dois comparer citySelect.Selected avec les données de l'API Locations
	// que tu auras pré-chargées.
	return true
}
