package todo

// В этом блоке импортируются необходимые пакеты.
import (
	"context"
	"net/http"
	"time"
)

// Тип Server представляет HTTP-сервер приложения.
type Server struct {
	// Поле httpServer содержит указатель на стандартный HTTP-сервер из пакета net/http.
	httpServer *http.Server
}

// Метод Run запускает HTTP-сервер и слушает входящие запросы на указанном порту.
func (s *Server) Run(port string, handler http.Handler) error {
	// Создаем новый HTTP-сервер с указанными параметрами.
	s.httpServer = &http.Server{
		Addr:           ":" + port, // Адрес и порт для прослушивания.
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // Максимальный размер заголовка запроса (1 MB).
		ReadTimeout:    10 * time.Second, // Таймаут на чтение запроса (10 секунд).
		WriteTimeout:   10 * time.Second, // Таймаут на запись ответа (10 секунд).
	}

	// Запускаем сервер и возвращаем ошибку, если она возникла.
	return s.httpServer.ListenAndServe()
}

// Метод Shutdown останавливает HTTP-сервер и освобождает все ресурсы.
func (s *Server) Shutdown(ctx context.Context) error {
	// Вызываем метод Shutdown стандартного HTTP-сервера и возвращаем ошибку, если она возникла.
	return s.httpServer.Shutdown(ctx)
}
