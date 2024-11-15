package paymentsdk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LoggingRoundTripper struct {
	log     Logger
	proxied http.RoundTripper
}

func NewLoggingRoundTripper(log Logger, proxied http.RoundTripper) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		log:     log,
		proxied: proxied,
	}
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Логирование запроса
	if lrt.log.Enabled() {

		lrt.log.Info(fmt.Sprintf("Request URL: %s", req.URL))
		lrt.log.Info(fmt.Sprintf("Request Headers: %v", req.Header))

		//пропускаем лог для загружаемых файлов и если тела нет впринципе
		if req.Body != nil && !strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") {
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}

			lrt.log.Info(fmt.Sprintf("Request Body: %s", string(bodyBytes)))

			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	// Выполнение запроса
	resp, err := lrt.proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Логирование ответа
	if lrt.log.Enabled() {

		lrt.log.Info(fmt.Sprintf("Response Status: %s", resp.Status))
		lrt.log.Info(fmt.Sprintf("Response Headers: %v", resp.Header))

		if resp.Body != nil {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			lrt.log.Info(fmt.Sprintf("Response Body: %s", string(bodyBytes)))

			// Восстановление тела ответа после чтения
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	return resp, nil
}
