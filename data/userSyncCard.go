package data

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// SyncPoolUserCard represents a single user card
type SyncPoolUserCard struct {
	ID    int
	Name  string
	Email string
	Phone string
	Image string
}

// Pool for reusing SyncPoolUserCard instances
var syncPoolUserCardPool = sync.Pool{
	New: func() interface{} {
		return &SyncPoolUserCard{}
	},
}

// SyncPoolRandomString generates a random string of a given length
func SyncPoolRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// SyncPoolGenerateUserData generates random user data for a SyncPoolUserCard
func SyncPoolGenerateUserData(id int) *SyncPoolUserCard {
	// Get a user card instance from the pool
	card := syncPoolUserCardPool.Get().(*SyncPoolUserCard)

	// Populate user data
	card.ID = id
	card.Name = SyncPoolRandomString(5) + " " + SyncPoolRandomString(7)
	card.Email = SyncPoolRandomString(5) + "@example.com"
	card.Phone = "+1-" + strconv.Itoa(rand.Intn(1000)) + "-" + strconv.Itoa(rand.Intn(1000)) + "-" + strconv.Itoa(rand.Intn(10000))
	card.Image = "https://randomuser.me/api/portraits/women/" + strconv.Itoa(rand.Intn(10)+1) + ".jpg"

	return card
}

// SyncPoolReleaseUserCard releases a SyncPoolUserCard back to the pool
func SyncPoolReleaseUserCard(card *SyncPoolUserCard) {
	// Clear sensitive fields before releasing the object
	card.Name = ""
	card.Email = ""
	card.Phone = ""
	card.Image = ""

	// Put the card back into the pool
	syncPoolUserCardPool.Put(card)
}

// SyncPoolPageOne generates the main page with user cards
func SyncPoolPageOne() *SyncPoolPage {
	return &SyncPoolPage{
		TemplateName: "pageone.html", // Template file name
		RenderFunc:   SyncPoolRenderPageOne,
	}
}

// SyncPoolPage struct for rendering
type SyncPoolPage struct {
	TemplateName string
	RenderFunc   func(w http.ResponseWriter, r *http.Request) error
}

// SyncPoolRenderPageOne renders the page and generates SyncPoolUserCards
func SyncPoolRenderPageOne(w http.ResponseWriter, r *http.Request) error {
	// Parse and render the HTML template for the main page
	tmpl, err := template.ParseFiles("view/pages/pageone.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Generate 1000 SyncPoolUserCards
	cards := make([]*SyncPoolUserCard, 1000)
	for i := 0; i < 1000; i++ {
		cards[i] = SyncPoolGenerateUserData(i + 1)
	}

	// Execute the template and pass the user cards data to it
	err = tmpl.Execute(w, cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Release user cards back to the pool after rendering
	for _, card := range cards {
		SyncPoolReleaseUserCard(card)
	}

	return nil
}

// SyncPoolLoadUserCardsPartial handles HTMX requests for dynamically loading more user cards
func SyncPoolLoadUserCardsPartial(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate 1000 SyncPoolUserCards (instead of just 20)
	cards := make([]*SyncPoolUserCard, 1000)
	for i := 0; i < 1000; i++ {
		cards[i] = SyncPoolGenerateUserData(i + 1)
	}

	// Parse the HTML snippet to render user cards
	tmpl, err := template.New("cards").Parse(`
		{{range .}}
		<div class="w-[180px] h-[300px] border-[#fc03c6] border-[2px] shadow-md shadow-white/20 drop-shadow-md flex flex-col items-center justify-center text-center text-white font-mono" id="card-{{.ID}}">
			<img src="{{.Image}}" alt="User Image">
			<h3>{{.Name}}</h3>
			<p>Email: {{.Email}}</p>
			<p>Phone: {{.Phone}}</p>
		</div>
		{{end}}
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the user cards for HTMX
	err = tmpl.Execute(w, cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Release user cards back to the pool after rendering
	for _, card := range cards {
		SyncPoolReleaseUserCard(card)
	}
}
