package gtools

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

type Gmap struct {
	M map[string]interface{}
}

func (g *Gmap) ContainsKey(key string) bool {
	if _, ok := g.M[key]; ok {
		return true
	}
	return false
}

// have to be the same type
func (g *Gmap) ContainsValue(value interface{}) bool {
	for _, v := range g.M {
		if v == value {
			return true
		}
	}
	return false
}

// regardless of type and capital or small letter
func (g *Gmap) ContainsValueWithoutType(value interface{}) bool {
	pattern := "\\d+"
	for _, v := range g.M {
		switch value.(type) {
		case string:
			if isDigit, _ := regexp.MatchString(pattern, value.(string)); isDigit {
				var v1 float64
				v2, _ := strconv.ParseFloat(value.(string), 64)
				switch v.(type) {
				case bool:
					d := int(v2) / 1 % 10
					b1 := digitToBool(d)
					if b1 == v.(bool) {
						return true
					}
				case string:
					if v.(string) == value.(string) {
						return true
					}
				case int8:
					v1 = float64(v.(int8))
					if v1 == v2 {
						return true
					}
				case int16:
					v1 = float64(v.(int16))
					if v1 == v2 {
						return true
					}
				case int32:
					v1 = float64(v.(int32))
					if v1 == v2 {
						return true
					}
				case int:
					v1 = float64(v.(int))
					if v1 == v2 {
						return true
					}
				case int64:
					v1 = float64(v.(int64))
					if v1 == v2 {
						return true
					}
				case float32:
					v1 = float64(v.(float32))
					if v1 == v2 {
						return true
					}
				case float64:
					v1 = v.(float64)
					if v1 == v2 {
						return true
					}
				}
			} else if strings.ToLower(v.(string)) == strings.ToLower(value.(string)) {
				return true
			}
		case int8:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := value.(int8) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(int8))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(int8))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(int8))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(int8))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(int8))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(int8))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(int8))
			}
			if v1 == v2 {
				return true
			}
		case int16:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := value.(int16) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(int16))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(int16))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(int16))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(int16))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(int16))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(int16))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(int16))
			}
			if v1 == v2 {
				return true
			}
		case int:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := value.(int) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(int))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(int))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(int))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(int))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(int))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(int))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(int))
			}
			if v1 == v2 {
				return true
			}
		case int32:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := value.(int32) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(int32))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(int32))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(int32))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(int32))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(int32))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(int32))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(int32))
			}
			if v1 == v2 {
				return true
			}
		case int64:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := value.(int64) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(int64))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(int64))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(int64))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(int64))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(int64))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(int64))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(int64))
			}
			if v1 == v2 {
				return true
			}
		case float32:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := int(value.(float32)) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = float64(value.(float32))
			case int16:
				v1 = float64(v.(int16))
				v2 = float64(value.(float32))
			case int:
				v1 = float64(v.(int))
				v2 = float64(value.(float32))
			case int32:
				v1 = float64(v.(int32))
				v2 = float64(value.(float32))
			case int64:
				v1 = float64(v.(int64))
				v2 = float64(value.(float32))
			case float32:
				v1 = float64(v.(float32))
				v2 = float64(value.(float32))
			case float64:
				v1 = v.(float64)
				v2 = float64(value.(float32))
			}
			if v1 == v2 {
				return true
			}
		case float64:
			var v1, v2 float64
			switch v.(type) {
			case bool:
				d := int(value.(float64)) / 1 % 10
				b1 := digitToBool(d)
				if b1 == v.(bool) {
					return true
				}
			case int8:
				v1 = float64(v.(int8))
				v2 = value.(float64)
			case int16:
				v1 = float64(v.(int16))
				v2 = value.(float64)
			case int:
				v1 = float64(v.(int))
				v2 = value.(float64)
			case int32:
				v1 = float64(v.(int32))
				v2 = value.(float64)
			case int64:
				v1 = float64(v.(int64))
				v2 = value.(float64)
			case float32:
				v1 = float64(v.(float32))
				v2 = value.(float64)
			case float64:
				v1 = v.(float64)
				v2 = value.(float64)
			}
			if v1 == v2 {
				return true
			}
		case bool:
			switch v.(type) {
			case bool:
				if v.(bool) == value.(bool) {
					return true
				}
			case string:
				b1, err := strconv.ParseBool(v.(string))
				if err != nil {
					return false
				}
				if b1 == value.(bool) {
					return true
				}
			case int8:
				d := v.(int8) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case int16:
				d := v.(int16) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case int32:
				d := v.(int32) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case int:
				d := v.(int) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case int64:
				d := v.(int64) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case float32:
				d := int(v.(float32)) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			case float64:
				d := int(v.(float64)) / 1 % 10
				b1 := digitToBool(d)
				if b1 == value.(bool) {
					return true
				}
			}
		case map[string]interface{}:
			switch v.(type) {
			case map[string]interface{}:
				var equal = true
				for k, v1 := range value.(map[string]interface{}) {
					if _, ok := v.(map[string]interface{})[k]; !ok {
						equal = false
					}
					if v1 != v.(map[string]interface{})[k] {
						equal = false
					}
				}
				if equal {
					return true
				}
			}
		case []string:
			switch v.(type) {
			case []string:
				var equal = true
				for i := range value.([]string) {
					var exist = false
					for j := range v.([]string) {
						if value.([]string)[i] == v.([]string)[j] {
							exist = true
						}
					}
					if !exist {
						equal = false
					}
				}
				if equal {
					return true
				}
			}
		}
	}
	return false
}

func (g *Gmap) ToString() string {
	bytes, _ := json.Marshal(g.M)
	return string(bytes)
}

func digitToBool(d interface{}) bool {
	var b1 bool
	switch d.(type) {
	case int8:
		if d.(int8) == 0 {
			b1 = false
		} else {
			b1 = true
		}
	case int16:
		if d.(int16) == 0 {
			b1 = false
		} else {
			b1 = true
		}
	case int32:
		if d.(int32) == 0 {
			b1 = false
		} else {
			b1 = true
		}
	case int:
		if d.(int) == 0 {
			b1 = false
		} else {
			b1 = true
		}
	case int64:
		if d.(int64) == 0 {
			b1 = false
		} else {
			b1 = true
		}
	}
	return b1
}
