package k8szoo

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//var store = sessions.NewCookieStore(os.Getenv("SESSION_KEY"))

var store = sessions.NewCookieStore([]byte("GO_SESS"))

func HealthHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type", "text/plain")
	fmt.Fprint(response, "I'm okay jack!")
}

func NotFoundHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("X-Template-File", "html"+request.URL.Path)
	tmpl := template.Must(template.ParseFiles("html/404.html"))
	tmpl.Execute(response, nil)
}

func CSSHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-type", "text/css")
	tmpl := template.Must(template.ParseFiles("html/cover.css"))
	tmpl.Execute(response, nil)
}

func getAnimalFromSession(response http.ResponseWriter, request *http.Request) (int, error) {
	session, err := store.Get(request, "session-name")
	if err != nil {
		return -1, err
	}
	var animalID int

	if session.Values["chosenAnimal"] == nil {
		animalID, _ = ReserveRandomAnimal()
		//name := Animals[animalID].AnimalName
		fmt.Printf("[Existing session] Reserving %s (%d) for %s\n", "something", animalID, session.ID)
		session.Values["chosenAnimal"] = animalID
	} else {
		sessionAnimalID := session.Values["chosenAnimal"]
		var ok bool
		if animalID, ok = sessionAnimalID.(int); !ok {
			return -1, errors.New("Conversion error")
		}
		if IsAnimalReserved(Animals[animalID].AnimalName) == false {
			fmt.Printf("[New session] Reserving %s (%d) for %s\n", Animals[animalID].AnimalName, animalID, session.ID)
			err = ReserveAnimalByName(Animals[animalID].AnimalName)
			if err != nil {
				animalID, _ = ReserveRandomAnimal()
				session.Values["chosenAnimal"] = animalID
			}
		} else {
			fmt.Printf("Session %s has %s (%d) reserved.\n", session.ID, Animals[animalID].AnimalName, animalID)
		}
		return animalID, nil
	}
	err = session.Save(request, response)
	if err != nil {
		return -1, err
	}
	return animalID, nil
}

func RandomAnimalHandler(response http.ResponseWriter, request *http.Request) {
	type TemplateData struct {
		Animal     AnimalData
		Templates  []string
		TotalCount int
		AvailCount int
	}
	animalID, err := getAnimalFromSession(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles(filepath.FromSlash("html/index.html")))
	animal := Animals[animalID]
	templates, err := filepath.Glob(filepath.FromSlash("./html/example/*.yaml"))
	if err != nil {
		templates = []string{}
	}
	for i, _ := range templates {
		thistemplate := strings.Split(templates[i], string(os.PathSeparator))
		templates[i] = thistemplate[len(thistemplate)-1]
	}
	data := TemplateData{
		Animal:     animal,
		Templates:  templates,
		TotalCount: len(Animals),
		AvailCount: len(AvailableAnimals),
	}
	tmpl.Execute(response, data)
}

func TemplateHandler(response http.ResponseWriter, request *http.Request) {
	animalID, err := getAnimalFromSession(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Add("X-Template-File", "html"+request.URL.Path)

	// tmpl, err := template.New("html"+request.URL.Path).Funcs(template.FuncMap{
	// 	"ToUpper": strings.ToUpper,
	// 	"ToLower": strings.ToLower,
	// }).ParseFiles("html"+request.URL.Path)
	_ = strings.ToLower("Hello")
	if strings.Index(request.URL.Path, "/") < 0 {
		http.Error(response, "No slashes wat - "+request.URL.Path, http.StatusInternalServerError)
		return
	}

	basenameSlice := strings.Split(request.URL.Path, "/")
	basename := basenameSlice[len(basenameSlice)-1]
	//fmt.Fprintf(response, "%q", basenameSlice)
	tmpl, err := template.New(basename).Funcs(template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}).ParseFiles("html" + request.URL.Path)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
		// NotFoundHandler(response, request)
		// return
	}
	err = tmpl.Execute(response, Animals[animalID])
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReleaseHandler(response http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "session-name")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	animalID, err := getAnimalFromSession(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	animal := Animals[animalID]
	err = ReleaseAnimalByName(animal.AnimalName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["chosenAnimal"] = nil
	err = session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(response, request, "/", http.StatusSeeOther)
	return
}

func HandleHTTP() {
	r := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	r.HandleFunc("/", RandomAnimalHandler)
	r.HandleFunc("/healthz", HealthHandler)
	r.HandleFunc("/release", ReleaseHandler)
	r.HandleFunc("/example/{.*}", TemplateHandler)
	r.HandleFunc("/cover.css", CSSHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         "0.0.0.0:5353",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on 0.0.0.0:5353")
	copy(AvailableAnimals, Animals)
	srv.ListenAndServe()
}
