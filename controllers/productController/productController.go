package productcontroller

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
	"github.com/paratonsp/learning-go/helper"
	"github.com/paratonsp/learning-go/models"
)

func Get(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	models.DB.Find(&products)
	helper.ResponseJSON(w, http.StatusOK, products)

}

func GetById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := mux.Vars(r)["id"]

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Data tidak ditemukan"}
			helper.ResponseJSON(w, http.StatusNotFound, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	helper.ResponseJSON(w, http.StatusOK, product)

}

func Post(w http.ResponseWriter, r *http.Request) {

	var productInput models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&productInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&productInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := mux.Vars(r)["id"]

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		response := map[string]string{"message": "tidak dapat mengupdate product"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "Data berhasil diperbarui"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	id := mux.Vars(r)["id"]

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		response := map[string]string{"message": "Tidak dapat menghapus product"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "Data berhasil dihapus"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
