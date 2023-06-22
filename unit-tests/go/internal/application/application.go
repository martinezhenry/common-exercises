package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/martinezhenry/common-exercises/unit-tests/go/internal/storage"
	"github.com/martinezhenry/common-exercises/unit-tests/go/internal"
)

type App interface {
	Run() error
}

type app struct {
	userRepository internal.UserRepository
	server		 *http.ServeMux
}

func (a *app) Run() error {
	// Initialize the HTTP server
	a.server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Set up the routes for the user repository
	a.routes()

	// Start the server
	if err := http.ListenAndServe(":8080", a.server); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func NewApplication() App {
	// Initialize the storage
	storage := storage.NewMemoryStorage[string, internal.User]()
	// Initialize the user repository with the storage
	userRepository := internal.NewUserRepository(storage)

	server := http.NewServeMux()

	// Return a new instance of the application with the user repository
	// and the storage.
	return &app{
		userRepository: userRepository,
		server:         server,
	}
}

func (a *app) routes() {
	a.server.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			users, err := a.userRepository.GetAllUsers()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			var user internal.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := a.userRepository.SetUser(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	a.server.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/user/"):]

		switch r.Method {
		case http.MethodGet:
			user, err := a.userRepository.GetUserByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			json.NewEncoder(w).Encode(user)
		case http.MethodDelete:
			if err := a.userRepository.DeleteUser(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
