package utilities

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/ramadoiranedar/user_micro/internal/constants"
)

func GetClientIPv4(request *http.Request) (ip string, err error) {
	ip = request.Header.Get("X-Real-IP") // some proxies
	if ip == "" {
		ip = request.Header.Get("X-Forwarded-For") // some proxies
	}

	if ip == "" {
		ip = request.RemoteAddr
	}

	ip, _, err = net.SplitHostPort(ip)
	if err != nil {
		return "", err
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil || parsedIP.To4() == nil {
		return "", fmt.Errorf("not a valid IPv4 address: %s", ip)
	}

	return
}

func readRequestBodyForLogger(r *http.Request, c constants.Constants) (result string) {
	result = "parse body only content-type: [application/json, multipart/form-data]"

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			result = fmt.Sprintf("error read body: %s", err.Error())
			return
		}
		result = string(body)
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		formData, err := parseFormDataBody(r, c)
		if err != nil {
			result = fmt.Sprintf("error parsing form-data: %s", err.Error())
			return
		}

		var bodyFormData strings.Builder
		for key, value := range formData.Text {
			bodyFormData.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}

		for key, filename := range formData.Files {
			bodyFormData.WriteString(fmt.Sprintf("%s: %s\n", key, filename))
		}

		result = bodyFormData.String()
	}

	return
}

type FormData struct {
	Text  map[string]string
	Files map[string]string
}

func parseFormDataBody(r *http.Request, c constants.Constants) (*FormData, error) {
	formData := &FormData{
		Text:  make(map[string]string),
		Files: make(map[string]string),
	}

	err := r.ParseMultipartForm(c.GetAppMaxUploadMB() << 20) // size MB limit
	if err != nil {
		return nil, err
	}

	for name, values := range r.MultipartForm.Value {
		if len(values) > 0 {
			formData.Text[name] = values[0]
		}
	}

	for name, headers := range r.MultipartForm.File {
		for _, header := range headers {
			filename := header.Filename
			formData.Files[name] = filename
		}
	}

	return formData, nil
}
