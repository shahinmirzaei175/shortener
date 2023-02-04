package ports

type DB interface {
	Connect()
	Migration()
}
