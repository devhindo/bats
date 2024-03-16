package api

func INIT() {
	api := APIServer{
		listenAddress: ":8080",
	}
	api.RUN()
}