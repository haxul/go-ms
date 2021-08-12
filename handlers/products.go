package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/haxul/go-ms/data"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.SendAsJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle Put Products")
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Id parsing error", http.StatusBadRequest)
		return
	}

	incomingProduct := req.Context().Value("ProductCtx").(data.Product)
	err = data.UpdateProduct(id, &incomingProduct)

	if err != nil {
		http.Error(rw, "product is not found", http.StatusNotFound)
		return
	}

	list := data.GetProducts()
	err = list.SendAsJSON(rw)
	if err != nil {
		http.Error(rw, "Json parsing error", http.StatusInternalServerError)
		return
	}
}

func (p *Products) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle Post request")

	incProduct := request.Context().Value("ProductCtx").(data.Product)
	data.AddProduct(&incProduct)
	err := incProduct.SendAsJson(writer)
	if err != nil {
		http.Error(writer, "Send json error", http.StatusInternalServerError)
		return
	}
}

func (p *Products) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		incProduct := &data.Product{}
		err := incProduct.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Json parsing error", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "ProductCtx", incProduct)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func (p *Products) MiddlewareAddJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
