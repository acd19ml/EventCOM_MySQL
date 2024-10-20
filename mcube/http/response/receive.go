package response

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"github.com/acd19ml/EventCOM_MySQL/mcube/exception"
)

// GetDataFromBody 获取body中的数据
func GetDataFromBody(body io.ReadCloser, v interface{}) error {
	defer body.Close()

	bytesB, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	data := NewData(v)

	if err := json.Unmarshal(bytesB, data); err != nil {
		return err
	}

	if data.Code == nil {
		return errors.New("reponse code is nil")
	}

	if *data.Code != 0 {
		return exception.NewAPIException(data.Namespace, *data.Code, data.Reason, data.Message)
	}

	return nil
}
