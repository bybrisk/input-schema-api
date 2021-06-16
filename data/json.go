package data

import (
	"encoding/json"
	"io"
)	

func (d *CreateSchemaResponse) CreateSchemaResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *CreateSchemaRequest) FromJSONToCreateSchemaRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}