package oapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type APIRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Data    interface{}
	Result  interface{}
}

func NewRequest(method, url string) *APIRequest {
	return &APIRequest{
		Method: method,
		URL:    url,
	}
}

func (apiReq *APIRequest) Do() (*APIResponse, error) {
	var request *http.Request
	apiResp := new(APIResponse)

	if apiReq.Data != nil {
		DataJSON, err := json.Marshal(apiReq.Data)
		if err != nil {
			return apiResp, fmt.Errorf("failed marshal Data of %s %s request. Error: %v", apiReq.Method, apiReq.URL, err)
		}
		request, _ = http.NewRequest(apiReq.Method, apiReq.URL, bytes.NewBuffer(DataJSON))
	} else {
		request, _ = http.NewRequest(apiReq.Method, apiReq.URL, nil)
	}
	request.Close = true
	request.Header.Set("Content-Type", "application/json")
	for key, value := range apiReq.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResp, fmt.Errorf("%s %s request failed with error: %v", apiReq.Method, apiReq.URL, err)
	}

	apiResp.Response = response

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		code, _ := strconv.Atoi(response.Header.Get("Error-Code"))
		apiResp.Code = code
		apiResp.ErrMessage = string(body)
		return apiResp, fmt.Errorf("%s %s request failed. Err: %s", apiReq.Method, apiReq.URL, string(body))
	}

	if apiReq.Result != nil {
		apiResp.Data = &apiReq.Result
		if err := json.NewDecoder(response.Body).Decode(&apiResp.Data); err != nil {
			return apiResp, err
		}
	}

	return apiResp, nil
}

func (apiReq *APIRequest) SendPDF(pdfPath string) (*APIResponse, error) {
	file, err := os.Open(pdfPath)
	var request *http.Request
	apiResp := new(APIResponse)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	apiReq.Data = bytes.NewBuffer(fileBytes)
	// apiReq.Headers["Content-Type"] = "application/pdf"
	request, err = http.NewRequest(apiReq.Method, apiReq.URL, bytes.NewBuffer(fileBytes))

	if err != nil {
		return nil, err
	}
	// if apiReq.Data != nil {
	// 	request, _ = http.NewRequest(apiReq.Method, apiReq.URL, apiReq.Data.(*bytes.Buffer))
	// } else {
	// 	request, _ = http.NewRequest(apiReq.Method, apiReq.URL, nil)
	// }

	request.Close = true
	// request.Header.Set("Content-Type", "application/pdf")
	for key, value := range apiReq.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return apiResp, fmt.Errorf("%s %s request failed with error: %v", apiReq.Method, apiReq.URL, err)
	}

	apiResp.Response = response

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		code, _ := strconv.Atoi(response.Header.Get("Error-Code"))
		apiResp.Code = code
		apiResp.ErrMessage = string(body)
		return apiResp, fmt.Errorf("%s %s request failed. Err: %s %d", apiReq.Method, apiReq.URL, string(body), response.StatusCode)
	}

	if apiReq.Result != nil {
		apiResp.Data = &apiReq.Result
		if err := json.NewDecoder(response.Body).Decode(&apiResp.Data); err != nil {
			return apiResp, err
		}
	}

	return apiResp, nil
}
