package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/takashabe/go-ddd-sample/application"
	router "github.com/takashabe/go-router"
)

// printDebugf behaves like log.Printf only in the debug env
func printDebugf(format string, args ...interface{}) {
	if env := os.Getenv("GO_SERVER_DEBUG"); len(env) != 0 {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// ErrorResponse is Error response template
type ErrorResponse struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("reason: %s, error: %s", e.Message, e.Error.Error())
}

// Respond is response write to ResponseWriter
func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, err, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *ErrorResponse, ErrorResponse:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to parse json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}

// Error is wrapped Respond when error response
func Error(w http.ResponseWriter, code int, err error, msg string) {
	e := &ErrorResponse{
		Message: msg,
		Error:   err,
	}
	printDebugf("%s", e.String())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

// JSON is wrapped Respond when success response
func JSON(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, src)
}

// Routes returns the initialized router
func Routes() *router.Router {
	r := router.NewRouter()
	r.Get("/user/:id", getUser)
	r.Get("/users", getUsers)
	r.Post("/user", createUser)

	return r
}

// Run start server
func Run(port int) error {
	log.Printf("Server running at http://localhost:%d/", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), Routes())
}

func getUser(w http.ResponseWriter, r *http.Request, id int) {
	user, err := application.GetUser(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get user")
		return
	}
	JSON(w, http.StatusOK, user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := application.GetUsers()
	if err != nil {
		Error(w, http.StatusNotFound, err, "failed to get user list")
		return
	}
	JSON(w, http.StatusOK, users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Name string `json:"name"`
	}
	var p payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Error(w, http.StatusBadRequest, err, "failed to parse request")
		return
	}

	err = application.AddUser(p.Name)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}
	JSON(w, http.StatusCreated, nil)
}
