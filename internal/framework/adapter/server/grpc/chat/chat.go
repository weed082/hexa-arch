package chat

type Server struct {
}

func (s *Server) ChatService(csi ChatService_ChatServiceServer) error {
	return nil
}
