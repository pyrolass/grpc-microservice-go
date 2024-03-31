package handlers

type GRPCDriverHandlerInterface interface {
	InsertDriverLog() error
}

type GRPCDriverHandler struct {
}

func NewDriverHandler() GRPCDriverHandlerInterface {
	return &GRPCDriverHandler{}
}

func (d *GRPCDriverHandler) InsertDriverLog() error {
	return nil
}
