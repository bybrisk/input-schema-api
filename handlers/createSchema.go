
package handlers

import (
	"net/http"
	"fmt"
	"github.com/bybrisk/input-schema-api/data"
)

// swagger:route POST /schema/create schema createSchema
// Create a bot schema for a business 
//
// responses:
//	200: createSchemaResponse
//  422: errorValidation
//  501: errorResponse

func (p *Input_Schema) Create_Bot_Schema (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> input-schema-api Module")
	registeration := &data.CreateSchemaRequest{}

	err:=registeration.FromJSONToCreateSchemaRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = registeration.ValidateCreateSchemaRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> input-schema-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add data to mongo
	response := data.CreateBotSchemaMongo(registeration)

	//writing to the io.Writer
	w.Header().Set("Content-Type", "application/json")
	
	err = response.CreateSchemaResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}