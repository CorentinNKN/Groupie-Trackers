package main

import ( //outils hbt de go 1et2
	"fmt" //1format: permet de mélanger nombre+mot
	"image/color"
	"strings" //2 que pour les phrases/mots

	"fyne.io/fyne/v2"           //unité de mesure fyne largeur/ hauteur
	"fyne.io/fyne/v2/app"       // permet de lancer l'applicqation avec "app.New()"
	"fyne.io/fyne/v2/canvas"    //sert à afficher l'image de l'artiste:"canvas.NewImageFromURI"
	"fyne.io/fyne/v2/container" //
	"fyne.io/fyne/v2/dialog"    // Permet d'afficher un msg d'erreur si l'api ne répond pas
	"fyne.io/fyne/v2/storage"   // permet de gérer les liens internet
	"fyne.io/fyne/v2/widget"    // widget: barre de recherche, bouton, cardre (card)
)

// VARIABLES GLOBALES
// Je les mets ici pour qu'elles soient accessibles partout dans le code
// listeArtistes :  copie complète données reçues de l'API
// grilleArtistes : zone visuelle où on affiche les cartes
var listeArtistes []Artist
var grilleArtistes *fyne.Container
var maFenetre fyne.Window

func main() {
	//CRÉATION DE LA FENÊTRE, initialise l'application -> titre à ma fenêtre.

	monApp := app.New()
	maFenetre = monApp.NewWindow("Groupie Tracker - Mon Projet B1")
	maFenetre.Resize(fyne.NewSize(1000, 800)) // Je définis une taille de départ confortable

	// RÉCUPÉRATION DES DONNÉES (BACKEND)
	// Ici, j'appelle la fonction GetArtistes() (dans api-manager.go)
	// Cela me permet de récupérer la vraie liste depuis Internet
	var err error
	listeArtistes, err = GetArtistes()

	// Si Internet ne marche pas, j'affiche une erreur à l'utilisateur
	if err != nil {
		dialog.ShowError(err, maFenetre)
	}

	//LA BARRE DE RECHERCHE
	entreeRecherche := widget.NewEntry()
	entreeRecherche.SetPlaceHolder("Tapez le nom d'un groupe (ex: Queen)...")

	// C'est ici que je gère l'interaction :
	// À chaque fois que l'utilisateur tape une lettre (OnChanged),
	// je lance ma fonction de filtre.
	entreeRecherche.OnChanged = func(texte string) {
		filtrerArtistes(texte)
	}

	// LA GRILLE D'AFFICHAGE
	// J'utilise un GridWrap : c'est un conteneur intelligent qui place
	// les cartes à la ligne automatiquement quand il n'y a plus de place.
	grilleArtistes = container.NewGridWrap(fyne.NewSize(200, 300))

	// Au démarrage, j'appelle ma fonction avec un texte vide ("")
	// pour dire : "Affiche tout le monde sans filtre".
	filtrerArtistes("")

	// J'ajoute un ascenseur (Scroll) car il y a 52 artistes, ça dépasse l'écran.
	zoneDefilement := container.NewVScroll(grilleArtistes)

	// MISE EN PAGE (LAYOUT)
	// J'organise l'onglet Accueil avec un "BorderLayout" :
	// - En Haut : La barre de recherche (avec un peu de marge/padding)
	// - Au Centre : La grille défilante
	contenuAccueil := container.NewBorder(
		container.NewPadded(entreeRecherche),
		nil, nil, nil,
		zoneDefilement,
	)

	// LES ONGLETS (NAVIGATION)
	// créationonglets pour la future carte+les favoris
	labelCarte := widget.NewLabel("La carte du monde sera affichée ici.")
	labelCarte.Alignment = fyne.TextAlignCenter

	labelFavoris := widget.NewLabel("La liste des favoris sera affichée ici.")
	labelFavoris.Alignment = fyne.TextAlignCenter

	lesOnglets := container.NewAppTabs(
		container.NewTabItem("Catalogue", contenuAccueil),
		container.NewTabItem("Carte", labelCarte),
		container.NewTabItem("Favoris", labelFavoris),
	)

	// LANCEMENT
	// J'intègre les onglets dans la fenêtre et je lance la boucle infinie sinon prog s'arrete
	maFenetre.SetContent(lesOnglets)
	maFenetre.ShowAndRun()
}

// LE FILTRE
// Cette fonction est appelée quand on tape dans la barre de recherche
func filtrerArtistes(recherche string) {
	// 1. Je commence par vider la grille pour ne pas empiler les résultats
	grilleArtistes.Objects = nil

	// 2. Je parcours ma liste complète d'artistes (boucle For Range)
	for _, artiste := range listeArtistes {

		// tout en minuscule pour que la recherche ne soit pas sensible à la casse
		// (Comme ça "queen" trouve bien "Queen").
		nomMinuscule := strings.ToLower(artiste.Name)
		rechercheMinuscule := strings.ToLower(recherche)

		// 3. Condition : Si le nom contient le texte cherché OU si la recherche est vide...
		if strings.Contains(nomMinuscule, rechercheMinuscule) || recherche == "" {
			// ... Alors je crée la carte visuelle et je l'ajoute à la grille
			carte := creerUneCarte(artiste)
			grilleArtistes.Add(carte)
		}
	}

	// 4. Important : Je demande à Fyne de rafraîchir l'affichage pour voir les changements
	grilleArtistes.Refresh()
}

// -> FONCTION VISUELLE : CRÉATION D'UNE CARTE COLORÉE
func creerUneCarte(a Artist) fyne.CanvasObject {
	// 1. CRÉATION DE LA BOÎTE
	contenu := container.NewVBox()

	// 2. GESTION DE L'IMAGE (inchangée)
	lien, err := storage.ParseURI(a.Image)
	if err == nil {
		image := canvas.NewImageFromURI(lien)
		image.FillMode = canvas.ImageFillContain
		image.SetMinSize(fyne.NewSize(150, 150))
		contenu.Add(image)
	}

	// --- COULEUR 1 : VIOLET pour le NOM ---
	// R=Rouge, G=Vert, B=Bleu, A=Opacité (255 = visible à fond)
	couleurViolette := color.NRGBA{R: 138, G: 43, B: 226, A: 255}

	// On utilise canvas.NewText au lieu de widget.NewLabel
	texteNom := canvas.NewText(a.Name, couleurViolette)
	texteNom.Alignment = fyne.TextAlignCenter
	texteNom.TextSize = 20                          // On grossit un peu le texte
	texteNom.TextStyle = fyne.TextStyle{Bold: true} // En gras

	contenu.Add(texteNom) // On ajoute le texte violet

	// COULEUR 2 : VERT POMME pour la DATE
	couleurPomme := color.NRGBA{R: 50, G: 205, B: 50, A: 255}

	strDate := fmt.Sprintf("Création : %d", a.CreationDate)
	texteDate := canvas.NewText(strDate, couleurPomme)
	texteDate.Alignment = fyne.TextAlignCenter
	texteDate.TextSize = 14

	contenu.Add(texteDate) // On ajoute le texte vert

	// 5. GESTION DU BOUTON
	bouton := widget.NewButton("Voir Détails", func() {
		message := "Nom : " + a.Name + "\n" + "Date : " + fmt.Sprint(a.CreationDate)
		dialog.ShowInformation("Détails", message, maFenetre)
	})
	contenu.Add(bouton)

	// 6. FINITION
	card := widget.NewCard("", "", contenu)
	return card
}
