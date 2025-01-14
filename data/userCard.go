package data

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// UserCard represents a single user card
type UserCard struct {
	ID    int
	Name  string
	Email string
	Phone string
	Image string
}

// Random string generator (for name, email, etc.)
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// GenerateUserData generates random user data for a user card
func generateUserData(id int) UserCard {
	// Random name
	name := randomString(5) + " " + randomString(7)

	// Random email
	email := randomString(5) + "@example.com"

	// Random phone number
	phone := "+1-" + strconv.Itoa(rand.Intn(1000)) + "-" + strconv.Itoa(rand.Intn(1000)) + "-" + strconv.Itoa(rand.Intn(10000))

	// Random image URL
	image := "https://randomuser.me/api/portraits/men/" + strconv.Itoa(rand.Intn(10)+1) + ".jpg"

	return UserCard{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
		Image: image,
	}
}

// PageOne generates the main page with user cards
func PageOne() *Page {
	return &Page{
		TemplateName: "pageone.html", // Template file name
		RenderFunc:   renderPageOne,
	}
}

// Page struct for rendering
type Page struct {
	TemplateName string
	RenderFunc   func(w http.ResponseWriter, r *http.Request) error
}

// Render function to render the page and generate user cards
func renderPageOne(w http.ResponseWriter, r *http.Request) error {
	// Parse and render the HTML template for the main page
	tmpl, err := template.ParseFiles("view/pages/pageone.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	// Generate 1000 user cards
	cards := []UserCard{}
	for i := 1; i <= 1000; i++ {
		cards = append(cards, generateUserData(i))
	}

	// Execute the template and pass the user cards data to it
	return tmpl.Execute(w, cards)
}

// LoadUserCardsPartial handles HTMX requests for dynamically loading more user cards
func LoadUserCardsPartial(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a batch of 20 user cards
	cards := []UserCard{}
	for i := 1; i <= 1000; i++ {
		cards = append(cards, generateUserData(i))
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
}
