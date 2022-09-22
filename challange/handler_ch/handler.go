package handler_ch

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"challange/m/database_ch"
	"challange/m/modules_ch"
)

//Nova receita
func CreateNewRevenue(w http.ResponseWriter, r *http.Request) {
	var receita modules_ch.Revenue
	json.NewDecoder(r.Body).Decode(&receita)
	database_ch.DB.Create(&receita)
	json.NewEncoder(w).Encode(receita)

}

//Nova despesa
func CreateNewExpenses(w http.ResponseWriter, r *http.Request) {
	var despesa modules_ch.Expense
	json.NewDecoder(r.Body).Decode(&despesa)
	if despesa.Categoria == "" {
		despesa.Categoria = "outros"
	}
	database_ch.DB.Create(&despesa)
	json.NewEncoder(w).Encode(despesa)
}

//Todas as receitas
func GettAllRevenue(w http.ResponseWriter, r *http.Request) {
	var todasReceitas []modules_ch.Revenue
	database_ch.DB.Find(&todasReceitas)
	json.NewEncoder(w).Encode(todasReceitas)
}

//Todas as despesa
func GettAllExpense(w http.ResponseWriter, r *http.Request) {
	var todasDespesas []modules_ch.Expense
	database_ch.DB.First(&todasDespesas)
	json.NewEncoder(w).Encode(todasDespesas)
}

//Info por ID receita
func RevenueInfo(w http.ResponseWriter, r *http.Request) {
	var info modules_ch.Revenue
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&info, id)
	json.NewEncoder(w).Encode(info)
}

//Info por ID despesa
func ExpenseInfo(w http.ResponseWriter, r *http.Request) {
	var info modules_ch.Expense
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&info, id)
	json.NewEncoder(w).Encode(info)
}

//Atualizar receita
func UpdateRevenue(w http.ResponseWriter, r *http.Request) {
	var editar modules_ch.Revenue
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&editar, id)
	json.NewDecoder(r.Body).Decode(&editar)
	database_ch.DB.Save(&editar)
	json.NewEncoder(w).Encode(editar)
}

//Atualizar despesa
func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	var atualizar modules_ch.Expense
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&atualizar, id)
	json.NewDecoder(r.Body).Decode(&atualizar)
	database_ch.DB.Save(&atualizar)
	json.NewEncoder(w).Encode(atualizar)
}

//Deletar receita
func DeleteRevenue(w http.ResponseWriter, r *http.Request) {
	var deletar modules_ch.Revenue
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&deletar, id)
	json.NewDecoder(r.Body).Decode(&deletar)
	database_ch.DB.Delete(&deletar, id)
	json.NewEncoder(w).Encode(deletar)
}

//Deletar despesa
func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	var deletar modules_ch.Expense
	vars := mux.Vars(r)
	id := vars["id"]
	database_ch.DB.First(&deletar, id)
	json.NewDecoder(r.Body).Decode(&deletar)
	database_ch.DB.Delete(&deletar, id)
	json.NewEncoder(w).Encode(deletar)
}

//Get receita especifica
func GetRevenueDesc(w http.ResponseWriter, r *http.Request) {
	var des modules_ch.Revenue
	vars := mux.Vars(r)
	ep := vars["descricao"]
	database_ch.DB.Find(&des, ep)
	json.NewEncoder(w).Encode(ep)
}
