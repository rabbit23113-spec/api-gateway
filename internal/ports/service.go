package ports

type Service interface {
	ValidateJWT() error
}
