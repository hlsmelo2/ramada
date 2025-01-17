package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func GetRootPath() string {
	app, _error := os.Executable()

	if _error != nil {
		return ""
	}

	return filepath.Dir(app)
}

func RequestToJSON(writer http.ResponseWriter, request http.Request) map[string]interface{} {
	var (
		bodyRequest []byte
		jsonMap     map[string]interface{}
	)

	bodyRequest, _error := io.ReadAll(request.Body)

	if _error != nil {
		writer.Write([]byte(_error.Error()))

		return nil
	}

	jsonString := string(bodyRequest)
	json.Unmarshal([]byte(jsonString), &jsonMap)

	return jsonMap
}

func StrToUint(aNumber string) *uint64 {
	aUint, _error := strconv.ParseUint(aNumber, 10, 64)

	if _error != nil {
		fmt.Println(_error)

		return nil
	}

	return &aUint
}

func StrToFloat(aNumber string) *float64 {
	aFloat, _error := strconv.ParseFloat(aNumber, 10)

	if _error != nil {
		fmt.Println(_error)

		return nil
	}

	return &aFloat
}

func StrToDateTime(aDate string) *time.Time {
	if aDate == "" {
		fmt.Println("Empty date")
		var aTime time.Time = time.Now()

		return &aTime
	}

	dateTime, _error := time.Parse("Y-m-d h:i:s", aDate)

	if _error != nil {
		fmt.Println(_error)

		return nil
	}

	return &dateTime
}

func DBClose(db *gorm.DB) {
	if db, _error := db.DB(); _error != nil {
		db.Close()
	}
}
