package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/plordb/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers set app confg for helpers
func NewHelpers(a *config.AppConfig, infoW io.Writer, errW io.Writer) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)

	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {

	//trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	trace := fmt.Sprintf(fmt.Sprintf("helpers.go-28 Error Interno del Server: %s\n", err.Error()))
	app.ErrorLog.Println(trace)
	//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// PrintStruct turns a struct into json and prints it.
func PrintStruct(item interface{}) {
	data, _ := json.MarshalIndent(item, "", "    ")
	fmt.Println(string(data))
}
