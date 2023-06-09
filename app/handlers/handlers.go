//	@title			Augustin Swagger
//	@version		0.0.1
//	@description	This swagger describes every endpoint of this project.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	GNU Affero General Public License
//	@license.url	https://www.gnu.org/licenses/agpl-3.0.txt

//	@host		localhost:3000
//	@BasePath	/api/

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

package handlers

import (
	"augustin/database"
	"augustin/structs"
	"encoding/json"
	"net/http"

	_ "github.com/swaggo/files"        // swagger embed files
	_ "github.com/swaggo/http-swagger" // http-swagger middleware

	log "github.com/sirupsen/logrus"
)

// ReturnHelloWorld godoc
//
//	 	@Summary 		Return HelloWorld
//		@Description	Return HelloWorld as sample API call
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Router			/hello/ [get]
//
// HelloWorld API Handler fetching data from database
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	greeting, err := database.Db.GetHelloWorld()
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(greeting))
}

// CreatePayments godoc
//
//	 	@Summary 		Get all payments
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	structs.Payment
//		@Router			/payments [get]
func GetPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := database.Db.GetPayments()
	if err != nil {
		log.Errorf("GetPayments DB Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshal_struct, err := json.Marshal(payments)
	if err != nil {
		log.Errorf("JSON conversion failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(marshal_struct))
}

// CreatePayments godoc
//
//	 	@Summary 		Create a set of payments
//		@Description    {"Payments":[{"Sender": 1, "Receiver":1, "Type":1,"Amount":1.00]}
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	structs.PaymentType
//		@Router			/payments [post]
func CreatePayments(w http.ResponseWriter, r *http.Request) {
	var paymentBatch structs.PaymentBatch
	err := json.NewDecoder(r.Body).Decode(&paymentBatch)
	if err != nil {
		log.Errorf("JSON conversion failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = database.Db.CreatePayments(paymentBatch.Payments)
	if err != nil {
		log.Errorf("CreatePayments query failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ReturnSettings godoc
//
//	 	@Summary 		Return settings
//		@Description	Return settings about the web-shop
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	structs.Settings
//		@Router			/settings/ [get]
//
// Get Settings API Handler fetching data without database
func GetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := database.Db.GetSettings()
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshal_struct, err := json.Marshal(settings)
	if err != nil {
		log.Errorf("JSON conversion failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(marshal_struct))
}

// ReturnVendorInformation godoc
//
//	 	@Summary 		Return vendor information
//		@Description	Return information for the vendor
//		@Tags			core
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	structs.Vendor
//		@Router			/vendor/ [get]
//
// Vendor API Handler fetching data without database
func Vendors(w http.ResponseWriter, r *http.Request) {
	vendors, err := database.Db.GetVendorSettings()
	if err != nil {
		log.Errorf("QueryRow failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(vendors))
}
