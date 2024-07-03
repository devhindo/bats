package api

import (
	"github.com/joho/godotenv"
)

func init() {
	
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}


	mailAuth = initMailService()

}

func RUN() {

	

	/*

	internalDB := db.NewMySQL()
	err := internalDB.Connect()
	if err != nil {
		// error couldn't connect to internal DB
	}
	
	externalDB := db.NewSupaBase()
	same as above

	api := APIServer{
		internalDB :
		externalDB :
	}
	*/

	api := APIServer{
		ListenAddress: "3000",
	}
	api.RUN()

	//INIT()
}




// in-memory db for now
/*

func INIT(port string, internalDB db.Database, externalDB db.Database) {
	api := APIServer{
		ListenAddress: port,
		InternalDB: internalDB,
		ExternalDB: externalDB,
	}
}

*/