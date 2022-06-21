package controllers

func (s *Server) InitializeRouter() {
	s.Router.HandleFunc("/", s.Home).Methods("GET")
}
