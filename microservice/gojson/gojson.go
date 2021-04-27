package gojson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Js struct {
	data interface{}
}

//Initialize the json configruation
func Json(data string) *Js {
	j := new(Js)
	var f interface{}
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		return j
	}
	j.data = f
	return j
}

//According to the key of the returned data information,return js.data
func (j *Js) Get(key string) *Js {
	m := j.Getdata()
	if v, ok := m[key]; ok {
		j.data = v
		return j
	}
	j.data = nil
	return j
}

//判断是否有郊
func (j *Js) IsValid() bool {
	if nil == j.data {
		return false
	} else {
		return true
	}
}

//return json data
func (j *Js) Getdata() map[string]interface{} {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

func (j *Js) Getindex(i int) *Js {

	num := i - 1
	if m, ok := (j.data).([]interface{}); ok {
		//防止越界
		if num < len((j.data).([]interface{})) {
			v := m[num]
			j.data = v
		} else {
			j.data = nil
		}
		return j
	}

	if m, ok := (j.data).(map[string]interface{}); ok {
		var n = 0
		var data = make(map[string]interface{})
		for i, v := range m {
			if n == num {
				switch vv := v.(type) {
				case float64:
					data[i] = strconv.FormatFloat(vv, 'f', -1, 64)
					j.data = data
					return j
				case string:
					data[i] = vv
					j.data = data
					return j
				case []interface{}:
					j.data = vv
					return j
				}

			}
			n++
		}

	}
	j.data = nil
	return j
}

// When the data {"result":["username","password"]} can use arrayindex(1) get the username
func (j *Js) Arrayindex(i int) string {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		data := errors.New("index out of range list").Error()
		return data
	}
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num]
		switch vv := v.(type) {
		case float64:
			return strconv.FormatFloat(vv, 'f', -1, 64)
		case string:
			return vv
		default:
			return ""
		}

	}

	if _, ok := (j.data).(map[string]interface{}); ok {
		return "error"
	}
	return "error"

}

//The data must be []interface{} ,According to your custom number to return your key and array data
func (j *Js) Getkey(key string, i int) *Js {
	num := i - 1
	if i > len((j.data).([]interface{})) {
		j.data = errors.New("index out of range list").Error()
		return j
	}
	if m, ok := (j.data).([]interface{}); ok {
		v := m[num].(map[string]interface{})
		if h, ok := v[key]; ok {
			j.data = h
			return j
		}

	}
	j.data = nil
	return j
}

//According to the custom of the PATH to find the PATH
func (j *Js) Getpath(args ...string) *Js {
	d := j
	for i := range args {
		m := d.Getdata()

		if val, ok := m[args[i]]; ok {
			d.data = val
		} else {
			d.data = nil
			return d
		}
	}
	return d
}

func (j *Js) Tostring() string {
	if m, ok := j.data.(string); ok {
		return m
	}
	if m, ok := j.data.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	if m, ok := j.data.(bool); ok {
		return strconv.FormatBool(m)
	}
	return ""
}

func (j *Js) ToArray() (k, d []string) {
	var key, data []string
	if m, ok := (j.data).([]interface{}); ok {
		for _, value := range m {
			for index, v := range value.(map[string]interface{}) {
				switch vv := v.(type) {
				case float64:
					data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
					key = append(key, index)
				case string:
					data = append(data, vv)
					key = append(key, index)

				}
			}
		}

		return key, data
	}

	if m, ok := (j.data).(map[string]interface{}); ok {
		for index, v := range m {
			switch vv := v.(type) {
			case float64:
				data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
				key = append(key, index)
			case string:
				data = append(data, vv)
				key = append(key, index)
			}
		}
		return key, data
	}

	return nil, nil
}

func (j *Js) StringtoArray() []string {
	var data []string
	for _, v := range j.data.([]interface{}) {
		switch vv := v.(type) {
		case string:
			data = append(data, vv)
		case float64:
			data = append(data, strconv.FormatFloat(vv, 'f', -1, 64))
		}
	}
	return data
}

func (j *Js) Type() {
	fmt.Println(reflect.TypeOf(j.data))
}