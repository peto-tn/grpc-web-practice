package interfaces

import (
	"net"

	"blog/app/application"
	"blog/app/config"
	api "blog/app/proto/api"

	"google.golang.org/grpc"
)

type Server struct {
	conf      *config.Config
	systemApp application.SystemInteractor
	server    *grpc.Server
}

func NewServer(
	conf *config.Config,
	systemApp application.SystemInteractor,
	article api.ArticleServer,
) *Server {
	server := grpc.NewServer()

	api.RegisterArticleServer(server, article)

	return &Server{
		conf:      conf,
		systemApp: systemApp,
		server:    server,
	}
}

func (s *Server) Run() error {
	listenPort, err := net.Listen("tcp", ":"+s.conf.Server.ListenPort)
	if err != nil {
		return err
	}

	return s.server.Serve(listenPort)
}
