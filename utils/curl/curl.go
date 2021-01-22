package curl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/martinyonathann/bookstore_users-api/domain/items"
	"github.com/martinyonathann/bookstore_users-api/logger"
	"github.com/martinyonathann/bookstore_users-api/utils/errors"
	"go.uber.org/zap"
)

func RequestToGateway(methodReq, url string, reqBody io.ReadCloser) (*items.Item, *errors.RestErr) {

	var errorsCreate errors.RestErr
	var itemsDomain items.Item

	client := &http.Client{}
	req, err := http.NewRequest(methodReq, url, reqBody)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	bodyString := string(body)

	json.Unmarshal([]byte(bodyString), &errorsCreate)

	if errorsCreate.Status == 0 {
		json.Unmarshal([]byte(bodyString), &itemsDomain)
		// logger.ResponseLog("Response from "+url, zap.Any("data_response", itemsDomain))
		logger.ResponseLog("Response", zap.Any("data_response", itemsDomain))
		return &itemsDomain, nil
	}

	// logger.ResponseLog("Response from "+url, zap.Any("data_response", errorsCreate))
	logger.ResponseLog("Response", zap.Any("data_response", errorsCreate))
	return nil, &errorsCreate

}
