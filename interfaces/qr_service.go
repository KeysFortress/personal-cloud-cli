package interfaces

type QRService interface {
	Get() (string, error)
}
