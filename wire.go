package main

/*
func InitializeItemApiController(db *sql.DB, opts ...openapi.ItemApiOption) openapi.Router {
	wire.Build(
		dbs.New,
		wire.Bind(new(dbs.DBTX), new(*sql.DB)),
		openapi.NewItemApiService,
		openapi.NewItemApiController)
	return &openapi.ItemApiController{}
}

func InitializeItemsApiController(db *sql.DB,  opts ...openapi.ItemsApiOption) openapi.Router {
	wire.Build(
		dbs.New,
		wire.Bind(new(dbs.DBTX), new(*sql.DB)),
		openapi.NewItemsApiService,
		openapi.NewItemsApiController)
	return &openapi.ItemsApiController{}
}
*/