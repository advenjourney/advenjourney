package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/advenjourney/api/graph"
	api "github.com/advenjourney/api/graph/generated"
	"github.com/advenjourney/api/internal/auth"
	database "github.com/advenjourney/api/internal/pkg/db/mysql"
	"github.com/advenjourney/api/pkg/config"
	"github.com/advenjourney/api/pkg/version"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Printf("api %s %s", version.Info(), version.BuildContext())

	cfg := config.Load()
	_ = godotenv.Load()
	if addr, ok := os.LookupEnv("API_SERVER_PORT"); ok {
		cfg.Server.Addr = addr
	}
	dsn, ok := os.LookupEnv("API_DB_DSN")
	if !ok || dsn == "" {
		log.Fatal("DSN (API_DB_DSN) not provided")
	}
	cfg.Database.DSN = dsn

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Middleware())

	database.InitDB(*cfg)
	database.Migrate()
	server := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", CorsMiddleware(server))

	filesDir := http.Dir(filepath.Join("./static"))
	FileServer(router, "/", filesDir)

	log.Printf("Server running at %s", cfg.Server.Addr)
	log.Fatal(http.ListenAndServe(cfg.Server.Addr, router))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "localhost:8080" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		} else if r.Host == "staging.advenjourney.com" {
			w.Header().Set("Access-Control-Allow-Origin", "https://staging.advenjourney.com")
		}

		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
