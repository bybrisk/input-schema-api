package data

func CreateBotSchemaMongo(d *CreateSchemaRequest) *CreateSchemaResponse{
	var response CreateSchemaResponse

	_ = CreateSchemaDB(d)
	
	response = CreateSchemaResponse{
		BusinessID:d.BusinessID,
		Message:"Successfully created business schema!",
		Status:200,
	}

	return &response
}		