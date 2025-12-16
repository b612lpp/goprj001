package midlware

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware логирует информацию о каждом запросе и ответе
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем входящий запрос
		slog.Info(
			"Входящий запрос",
			"method", r.Method,
			"path", r.RequestURI,
			"remote_addr", r.RemoteAddr,
		)

		// Оборачиваем ResponseWriter для перехвата статус кода
		wrappedWriter := &statusCodeWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Вызываем следующий обработчик
		next.ServeHTTP(wrappedWriter, r)

		// Логируем ответ
		duration := time.Since(start)
		slog.Info(
			"Ответ отправлен",
			"method", r.Method,
			"path", r.RequestURI,
			"status", wrappedWriter.statusCode,
			"duration_ms", duration.Milliseconds(),
		)
	})
}

// statusCodeWriter перехватывает статус код из ResponseWriter
type statusCodeWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusCodeWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
