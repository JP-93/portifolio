package route_ch

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"challange/m/handler_ch"
	"challange/m/middleware"
)

func HandlerRoute() {

	r := mux.NewRouter()
	r.Use(middleware.ContentType)
	r.HandleFunc("/receita", handler_ch.CreateNewRevenue).Methods("POST")
	r.HandleFunc("/despesa", handler_ch.CreateNewExpenses).Methods("POST")
	r.HandleFunc("/total-receitas", handler_ch.GettAllRevenue).Methods("GET")
	r.HandleFunc("/total-despesas", handler_ch.GettAllExpense).Methods("GET")
	r.HandleFunc("/receita/{id}", handler_ch.RevenueInfo).Methods("GET")
	r.HandleFunc("/despesa/{id}", handler_ch.ExpenseInfo).Methods("GET")
	r.HandleFunc("/receita/atualizar/{id}", handler_ch.UpdateRevenue).Methods("PUT")
	r.HandleFunc("/despesa/atualizar/{id}", handler_ch.UpdateExpense).Methods("PUT")
	r.HandleFunc("/receita/{id}", handler_ch.DeleteRevenue).Methods("DELETE")
	r.HandleFunc("/despesa/{id}", handler_ch.DeleteExpense).Methods("DELETE")
	r.HandleFunc("/receita/descricao", handler_ch.GetRevenueDesc).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))

}
