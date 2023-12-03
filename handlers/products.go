package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nvrl0vd/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}
// 	rw.WriteHeader(http.StatusMethodNotAllowed)

// 	if r.Method == http.MethodPost {
// 		p.AddProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		if len(g[0]) != 2 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		p.updateProducts(id, rw, r)
// 	}

// }

func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle GET products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	// prod := &data.Product{}
	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	// }
	prod := r.Context().Value(KeyProd{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Problem with ID", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT product", id)

	prod := r.Context().Value(KeyProd{}).(data.Product)
	// err = prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	// }
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProd struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "[ERROR] problem reading product", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProd{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
