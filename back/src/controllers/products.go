package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"ramada/api/src/db"
	"ramada/api/src/models"
	"ramada/api/src/utils"
	"strings"

	"github.com/gorilla/mux"
)

type UploadedFileData struct {
	file   multipart.File
	header *multipart.FileHeader
}

func UpinsertProduct(writer http.ResponseWriter, request *http.Request) {
	var (
		product     map[string]interface{}
		model       models.Product
		db                 = db.GetDB()
		requestJSON        = utils.RequestToJSON(writer, *request)
		id          uint64 = 0
	)

	product = map[string]interface{}{
		"Name":        requestJSON["Name"],
		"Description": requestJSON["Description"],
		"Price":       requestJSON["Price"],
		"Category":    requestJSON["Category"],
	}

	if idString := mux.Vars(request)["id"]; idString != "" {
		id = *utils.StrToUint(idString)
	}

	if id != 0 {
		db.Find(&model, models.Product{ID: id})
	}

	if model.ID != 0 {
		db.Model(model).Updates(product)
		utils.DBClose(db)
		json.NewEncoder(writer).Encode(model)

		return
	}

	db.Create(product).Last(&model)
	utils.DBClose(db)
	json.NewEncoder(writer).Encode(model)
}

func GetProducts(writer http.ResponseWriter, request *http.Request) {
	var (
		db       = db.GetDB()
		products *[]models.Product
	)

	queryParams := request.URL.Query()

	if queryParams.Get("nome") != "" {
		db = db.Where("name = ?", queryParams.Get("nome"))
	}

	if queryParams.Get("categoria") != "" {
		db = db.Where("category = ?", queryParams.Get("categoria"))
	}

	if queryParams.Get("preco_min") != "" {
		db = db.Where("price >= ?", utils.StrToFloat(queryParams.Get("preco_min")))
	}

	if queryParams.Get("preco_max") != "" {
		db = db.Where("price <= ?", utils.StrToFloat(queryParams.Get("preco_max")))
	}

	db.Find(&products)
	utils.DBClose(db)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(products)
}

func GetProduct(writer http.ResponseWriter, request *http.Request) {
	var (
		db    = db.GetDB()
		model *models.Product
		id    = *utils.StrToUint(mux.Vars(request)["id"])
	)

	db.First(&model, models.Product{ID: id})
	utils.DBClose(db)

	if model.ID == 0 {
		http.Error(writer, "product not exists", http.StatusOK)

		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(model)
}

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	var (
		db    = db.GetDB()
		model *models.Product
		id    = *utils.StrToUint(mux.Vars(request)["id"])
	)

	db.First(&model, models.Product{ID: id})

	if model.ID == 0 {
		http.Error(writer, "product not exists", http.StatusOK)
		utils.DBClose(db)

		return
	}

	db.Delete(model, id)
	utils.DBClose(db)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(model)
}

func getUploadedFileData(fileKey string, request *http.Request) UploadedFileData {
	empty := UploadedFileData{
		file:   nil,
		header: nil,
	}

	_error := request.ParseMultipartForm(32 << 10)

	if _error != nil {
		return empty
	}

	file, header, _error := request.FormFile(fileKey)

	if _error != nil {
		fmt.Println(_error)
		return empty
	}

	defer file.Close()

	return UploadedFileData{
		file:   file,
		header: header,
	}
}

func saveUploadedFile(filename string, uploadedFile multipart.File) *os.File {
	path := utils.GetRootPath() + "/src/files/"
	file, _error := os.Create(path + filename)

	if _error != nil {
		fmt.Println(_error)
		return nil
	}

	defer file.Close()

	io.Copy(file, uploadedFile)

	return file
}

func getFileLines(request *http.Request) []string {
	// upload := request.MultipartForm.File
	fileData := getUploadedFileData("file", request)

	if fileData.file == nil {
		fmt.Println("error reading uploaded file")
		return nil
	}

	// saveUploadedFile("import.csv", fileData.file);

	fileContent, _error := io.ReadAll(fileData.file)

	if _error != nil {
		fmt.Println(_error)
		return nil
	}

	return strings.Split(string(fileContent), "\n")
}

func ImportProducts(writer http.ResponseWriter, request *http.Request) {
	lines := getFileLines(request)

	if lines == nil {
		http.Error(writer, "error reading uploaded file", http.StatusOK)

		return
	}

	for index, line := range lines {
		if index == 0 {
			continue
		}

		data := strings.Split(line, ",")

		var product = models.Product{
			ID:          *utils.StrToUint(data[0]),
			Name:        data[1],
			Description: data[2],
			// Price:       *utils.StrToFloat(data[3]),
			Price:     data[3],
			Category:  data[4],
			CreatedAt: *utils.StrToDateTime(data[5]),
			UpdatedAt: *utils.StrToDateTime(data[6]),
		}

		marshal, _error := json.Marshal(product)

		if _error != nil {
			fmt.Println(_error)

			return
		}

		reader := bytes.NewReader(marshal)
		body := io.NopCloser(reader)

		if _error != nil {
			fmt.Println(_error)

			return
		}

		UpinsertProduct(writer, &http.Request{
			Body: body,
		})
	}

	json.NewEncoder(writer).Encode(lines)
}
