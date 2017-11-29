package po

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// HTTPServer starts our webserver on a given port
func (p *Po) HTTPServer(port, auth string) {
	p.token = auth
	r := mux.NewRouter()

	// setup our routes
	r.HandleFunc("/{entity}", p.HTTPAuth(p.HTTPGetEntity)).Methods("GET")
	r.HandleFunc("/{entity}", p.HTTPAuth(p.HTTPCreateEntity)).Methods("POST")
	r.HandleFunc("/{entity}/{attribute}", p.HTTPAuth(p.HTTPSetEntityAttribute)).Methods("POST")
	r.HandleFunc("/{entity}/{attribute}", p.HTTPAuth(p.HTTPGetEntityAttribute)).Methods("GET")
	r.HandleFunc("/{entity}/{attribute}/{type}/{value}", p.HTTPAuth(p.HTTPCreateEntityAttribute)).Methods("POST")
	r.HandleFunc("/{entity}/{attribute}/increment", p.HTTPAuth(p.HTTPIncrementAttribute)).Methods("POST")
	r.HandleFunc("/{entity}/{attribute}/reset", p.HTTPAuth(p.HTTPResetEntityAttribute)).Methods("POST")
	r.HandleFunc("/{entity}/{attribute}/{value}", p.HTTPAuth(p.HTTPSetEntityAttribute)).Methods("POST")

	// set some defaults
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	// start serving
	log.Fatal(srv.ListenAndServe())
}

// HTTPGetEntity handler for web requests with new messages
func (p *Po) HTTPGetEntity(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	entity, err := p.Entity(vars["entity"])
	if err == nil {
		response.WriteHeader(http.StatusOK)
		v, _ := json.Marshal(entity.Export())
		fmt.Fprintf(response, string(v))
	} else {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, err.Error())
	}
}

// HTTPIncrementAttribute handler for web requests to increment a counter
func (p *Po) HTTPIncrementAttribute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	e, err := p.Entity(vars["entity"])
	if err == nil {
		// entity is set, is attribute?
		if a, ok := e.Attribute(vars["attribute"]); ok == nil {
			if i, converted := a.Get().(int); converted {
				i++
				a.Set(strconv.Itoa(i))
				v, _ := json.Marshal(i)
				fmt.Fprintf(response, string(v))
			} else {
				response.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(response, err.Error())
			}
		} else {
			response.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(response, err.Error())
		}
	} else {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, err.Error())
	}
}

// HTTPSetEntityAttribute handler for web requests with new messages
func (p *Po) HTTPSetEntityAttribute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	e := p.E(vars["entity"], Attributes{vars["attribute"]: "string"})

	if a, exists := e.Attribute(vars["attribute"]); exists == nil {
		if v, exists := vars["value"]; exists {
			a.Set(v)
			v, _ := json.Marshal(e.Export())
			fmt.Fprintf(response, string(v))
		} else {
			body, err := ioutil.ReadAll(request.Body)
			defer request.Body.Close()
			if err == nil {
				a.Set(string(body))
				response.WriteHeader(http.StatusOK)
				v, _ := json.Marshal(e.Export())
				fmt.Fprintf(response, string(v))
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(response, err.Error())
			}
		}
	} else {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, exists.Error())
	}
}

// HTTPGetEntityAttribute handler for web requests with new messages
func (p *Po) HTTPGetEntityAttribute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	e, err := p.Entity(vars["entity"])
	if err == nil {
		// entity is set, is attribute?
		if a, ok := e.Attribute(vars["attribute"]); ok == nil {
			response.WriteHeader(http.StatusOK)
			v, _ := json.Marshal(a.Get())
			fmt.Fprintf(response, string(v))
		} else {
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(response, err.Error())
		}
	} else {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, err.Error())
	}
}

// HTTPResetEntityAttribute handler for web requests with resetting Entityes
func (p *Po) HTTPResetEntityAttribute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	e, err := p.Entity(vars["entity"])
	if err == nil {
		// entity is set, is attribute?
		if a, ok := e.Attribute(vars["attribute"]); ok == nil {
			a.Reset()
			response.WriteHeader(http.StatusOK)
			v, _ := json.Marshal(e.Export())
			fmt.Fprintf(response, string(v))
		} else {
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(response, err.Error())
		}
	} else {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, err.Error())
	}
}

// HTTPCreateEntityAttribute handler for web requests to create new entity attributes
func (p *Po) HTTPCreateEntityAttribute(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var e *Entity
	var exists error
	if e, exists = p.Entity(vars["entity"]); exists == nil {
		e.AddAttribute(vars["attribute"], vars["type"])
	} else {
		e = p.E(vars["entity"], Attributes{vars["attribute"]: vars["type"]})
	}

	a := e.A(vars["attribute"])
	a.Set(vars["value"])
	a.Set(vars["value"])
	response.WriteHeader(http.StatusOK)
	v, _ := json.Marshal(e.Export())
	fmt.Fprintf(response, string(v))
}

// HTTPCreateEntity handler for web requests with new messages
func (p *Po) HTTPCreateEntity(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	e := p.AddEntity(vars["entity"], Attributes{})
	response.WriteHeader(http.StatusOK)
	v, _ := json.Marshal(e.Export())
	fmt.Fprintf(response, string(v))
}

// HTTPAuth wraps each request to check for a valid token
func (p *Po) HTTPAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		token, _, _ := r.BasicAuth()
		if p.token != "" && p.token != token {
			http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}
}
