package ports

type API interface {
	Run(Service) error
}
