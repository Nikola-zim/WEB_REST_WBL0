package WEB_REST_exm0302

import (
	"WEB_REST_exm0302/pkg/mynats"
	//"WEB_REST_exm0302/pkg/mynats"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func ContinuousPrint() {
	for {
		fmt.Println("Constant print")
		time.Sleep(3000 * time.Millisecond)
	}
}

// , subNats *mynats.SubsNats
func (s *Server) Run(port string, handler http.Handler, subNats *mynats.SubsNats) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	//Получение по mynats и отправка Json сервису для записи
	go subNats.GetAndWriteJson()

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
