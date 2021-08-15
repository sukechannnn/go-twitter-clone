package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/sukechannnn/go-twitter-clone/graph"
	"github.com/sukechannnn/go-twitter-clone/graph/generated"
	"github.com/sukechannnn/go-twitter-clone/graph/model"
	db "github.com/sukechannnn/go-twitter-clone/infrastructure"
	"gorm.io/gorm"
)

const defaultPort = "8080"

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignedIn struct {
	Result string
}

func authenticate(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		var signin SignIn
		json.Unmarshal(body, &signin)
		user, err := model.FindBy(db, "email", signin.Email)
		if err != nil || user == nil {
			http.Error(w, "Invalid login error", http.StatusForbidden)
		}
		cookie := http.Cookie{
			Name:     "auth",
			Value:    user.ID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		signedIn := SignedIn{"ok"}
		res, _ := json.Marshal(signedIn)
		w.Write(res)
	})
}

func main() {
	db := db.ConnectDb()
	router := chi.NewRouter()
	router.Use(graph.Middleware(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/sign_in", authenticate(db))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
