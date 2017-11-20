package workflow

import (
	"encoding/base64"
	"encoding/json"
	"time"
)

func Base64ToBytes(d Data) Monad {
	dString := d.(string)
	return func(e error) (Data, error) {
		return base64.StdEncoding.DecodeString(dString)
	}
}

func BytesToData(d Data) Monad {
	dBytes := d.([]byte)
	return func(e error) (Data, error) {
		data := &CarData{}
		err := json.Unmarshal(dBytes, &data)
		return data, err
	}
}

func TimestampData(d Data) Monad {
	data := d.(*CarData)
	return func(e error) (Data, error) {
		data.Timestamp = time.Now().Format("20060102150405")
		return data, nil
	}
}

func DataToJson(d Data) Monad {
	data := d.(*CarData)
	return func(e error) (Data, error) {
		car, err := json.Marshal(data)
		return string(car), err
	}
}

