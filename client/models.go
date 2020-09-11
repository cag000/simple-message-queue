package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Dummy struct {
	Datas []string		
}

func (d *Dummy) ReadFile(name string) error {
	var M []interface{}
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &M)
	if err != nil {
		return err
	}
	for _, v := range M {
		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(v)
		if err != nil {
			return err
		}
		d.Datas = append(d.Datas, buf.String())
	}
	return nil
}
