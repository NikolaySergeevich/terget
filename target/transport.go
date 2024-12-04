package target

type Transport interface {
	ListenAndServe(host, port string)
	RegisterRoutes()
}