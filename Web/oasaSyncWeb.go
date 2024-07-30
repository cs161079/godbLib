package web

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
)

const (
	oasaApplicationHost = "http://telematics.oasa.gr"
	testApplicationHost = "http://localhost:8080"
	geoapifyApplication = "https://api.geoapify.com"
)

type OpswHttpRequest struct {
	Method   string
	Headers  map[string]string
	Body     io.Reader
	Endpoint string
}

func getProperty(v interface{}, property string) any {
	if v != nil {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			return nil
		} else {
			result := v.(map[string]any)[property]
			return result
		}
	}
	return nil
}

func checkFields(request *OpswHttpRequest) error {
	if request.Endpoint == "" {
		return fmt.Errorf("REQUEST ENDPOINT IS NOT SET")
	}
	if request.Method == "" {
		return fmt.Errorf("REQUEST HTTP METHOD IS NOT SET")
	}
	return nil
}

func httpRequest(request *OpswHttpRequest) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	if request == nil {
		return nil, fmt.Errorf("REQUEST OBJECT-STRUCT IS NIL OR IS NOT SET CORRECTLY")
	}

	var err = checkFields(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(request.Method, request.Endpoint, request.Body)
	if err != nil {
		return nil, err
	}

	if request.Headers != nil && len(request.Headers) > 0 {
		for key, value := range request.Headers {
			req.Header.Set(key, value)
		}
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	logger.INFO(fmt.Sprintf("%s %s %d", response.Request.Method, response.Request.URL.String(), response.StatusCode))

	return response, nil
}

func getRequest(url string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	response, err := client.Do(req)
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	if err != nil {
		return nil, err
	}
	//fmt.Printf("%s %d %s %s %s \n", time.Now().Format("2006-01-02 15:04:05"), response.StatusCode, strings.Split(url, "?")[1], req.Method, req.Host)
	logger.INFO(response.Status + " " + url)
	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", response.StatusCode)
	return response, nil
}

func TestRequestFunction() (*string, error) {
	// keys := make([]int, len(extraParams))

	var req OpswHttpRequest = OpswHttpRequest{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/%s/%s", testApplicationHost, "api", "cronjob"),
	}

	//Error Code for error occured in Request Creation
	response, err := httpRequest(&req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf(models.INTERNALL_SERVER_ERROR)
	}

	reader, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	} else {
		responseStr := string(reader)
		return &responseStr, nil
	}
}

func TestMakeRequest(action string) (*string, error) {
	var req OpswHttpRequest = OpswHttpRequest{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/api/v1/sync/%s", testApplicationHost, action),
	}
	// req.Headers = map[string]string{
	// 	"Accept-Encoding": "gzip, deflate"}

	//Error Code for error occured in Request Creation
	response, err := httpRequest(&req)

	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf(models.INTERNALL_SERVER_ERROR)
	}
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respStr := string(resp)

	return &respStr, nil
}

func MakeRequest(action string) (*string, error) {
	var extraparamUrl string
	// keys := make([]int, len(extraParams))

	var req OpswHttpRequest = OpswHttpRequest{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/api/?act=%s%s", oasaApplicationHost, action, extraparamUrl),
	}
	req.Headers = map[string]string{
		"Accept-Encoding": "gzip, deflate"}
	//Error Code for error occured in Request Creation
	response, err := httpRequest(&req)

	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf(models.INTERNALL_SERVER_ERROR)
	}

	reader, err := gzip.NewReader(response.Body)

	if err != nil {
		return nil, err
	} else {
		defer reader.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		responseStr := buf.String()

		return &responseStr, nil
	}
}

func OasaRequestApi(action string, extraParams map[string]interface{}) *OasaResponse {
	var oasaResult OasaResponse = OasaResponse{}
	var extraparamUrl string = ""
	// keys := make([]int, len(extraParams))
	for k := range extraParams {
		extraparamUrl = extraparamUrl + "&" + k + "=" + strconv.FormatInt(extraParams[k].(int64), 10)
	}
	var req OpswHttpRequest = OpswHttpRequest{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("%s/api/?act=%s%s", oasaApplicationHost, action, extraparamUrl),
	}
	//Error Code for error occured in Request Creation
	//var tries = 1
	var resp []byte = nil
	//var retry bool = false
	var err error = nil
	//for tries <= 3 {
	//	resp, err, retry = internalHttpRequest(req)
	//	if !retry {
	//		break
	//	}
	//	if tries > 1 {
	//		logger.WARN(fmt.Sprintf("Try make request for %d time beacause of error %+v", tries, err))
	//	}
	//	tries = tries + 1
	//}
	resp, err, _ = internalHttpRequest(req)
	if err != nil {
		oasaResult.Error = err
		return &oasaResult
	}

	var tmpResult interface{}
	err = json.Unmarshal(resp, &tmpResult)
	if err != nil {
		oasaResult.Error = fmt.Errorf("AN ERROR OCCURED WHEN CONVERT JSON STRING TO INTERFACE. %s \n %+v", err.Error(), resp)
		return &oasaResult
	}
	hasError := getProperty(tmpResult, "error")
	if hasError != nil {
		oasaResult.Error = fmt.Errorf("SERVER RESPONSES ERROR. %s", hasError)
		return &oasaResult
	}

	oasaResult.Data = tmpResult
	return &oasaResult
}

func internalHttpRequest(req OpswHttpRequest) ([]byte, error, bool) {
	response, err := httpRequest(&req)
	if err != nil {
		logger.ERROR(fmt.Sprintf("An error occured in httpRequest(). %s", err.Error()))
		return nil, err, false
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		var returnedError = fmt.Errorf("AN ERROR OCCURED ANALYZE RESPONSE BODY. %s", err.Error())
		return nil, returnedError, false
	}
	if response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons {
		//fmt.Println("Client Error Response from Server")
		var returnedError = fmt.Errorf("%s %s", response.Status, responseBody)
		return nil, returnedError, true
	}
	if response.StatusCode >= http.StatusInternalServerError && response.StatusCode <= http.StatusNetworkAuthenticationRequired {
		var returnedError = fmt.Errorf("%s %s", response.Status, responseBody)
		//logger.ERROR(string(responseBody))
		return nil, returnedError, false
	}
	return responseBody, nil, false
}

func GeneralRequest(req OpswHttpRequest, responseModel any) *OasaResponse {
	var oasaResult OasaResponse = OasaResponse{}
	response, err := httpRequest(&req)
	if err != nil {
		oasaResult.Error = err
		return &oasaResult
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		oasaResult.Error = fmt.Errorf("AN ERROR OCCURED ANALYZE RESPONSE BODY. %s", err.Error())
		return &oasaResult
	}
	if response.StatusCode >= http.StatusBadRequest && response.StatusCode <= http.StatusUnavailableForLegalReasons {
		//fmt.Println("Client Error Response from Server")
		oasaResult.Error = fmt.Errorf("%s %s", response.Status, responseBody)
		return &oasaResult
	}
	if response.StatusCode >= http.StatusInternalServerError && response.StatusCode <= http.StatusNetworkAuthenticationRequired {
		oasaResult.Error = fmt.Errorf("%s %s", response.Status, responseBody)
		//logger.ERROR(string(responseBody))
		return &oasaResult
	}

	err = json.Unmarshal(responseBody, responseModel)
	if err != nil {
		oasaResult.Error = fmt.Errorf("AN ERROR OCCURED WHEN CONVERT JSON STRING TO INTERFACE. %s", err.Error())
		return &oasaResult
	}

	return &oasaResult
}
