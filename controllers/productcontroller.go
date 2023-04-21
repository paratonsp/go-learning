package controllers

import (
	"fmt"
	"learning-go/helper"
	"learning-go/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	models.DB.Find(&products)
	helper.ResponseJSON(w, http.StatusOK, products)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := mux.Vars(r)["id"]

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Data Not Found"}
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var productInput models.Product

	r.ParseMultipartForm(32 << 20)
	_, _, err := r.FormFile("assets")
	if err == nil {

		imageToS3, err := helper.FileUploadS3(models.S3S, r, "products")
		if err != nil {
			fmt.Print(err)
		}

		productInput.Gambar = imageToS3
	}

	s, err := strconv.ParseInt(r.Form.Get("stok"), 10, 64)
	if err != nil {
		fmt.Print(err)

	}

	productInput.NamaProduk = r.Form.Get("nama_produk")
	productInput.Deskripsi = r.Form.Get("deskripsi")
	productInput.Stok = s

	if err := models.DB.Create(&productInput).Error; err != nil {
		response := map[string]string{"message": "Cannot Create Data"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Success Create Data"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	id := mux.Vars(r)["id"]

	r.ParseMultipartForm(32 << 20)
	_, _, err := r.FormFile("assets")
	if err == nil {

		imageToS3, err := helper.FileUploadS3(models.S3S, r, "products")
		if err != nil {
			fmt.Print(err)
		}

		product.Gambar = imageToS3
	}

	s, err := strconv.ParseInt(r.Form.Get("stok"), 10, 64)
	if err != nil {
		fmt.Print(err)
	}

	product.NamaProduk = r.Form.Get("nama_produk")
	product.Deskripsi = r.Form.Get("deskripsi")
	product.Stok = s

	if err := models.DB.Model(&product).Where("id = ?", id).Updates(&product).Error; err != nil {
		fmt.Print(err)
		response := map[string]string{"message": "Cannot Update Data"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]string{"message": "Success Update Data"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	id := mux.Vars(r)["id"]

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		response := map[string]string{"message": "Cannot Delete Data"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]string{"message": "Success Delete Data"}
	helper.ResponseJSON(w, http.StatusOK, response)

}
