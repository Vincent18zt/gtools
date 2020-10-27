package gtools

import (
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
				v1, _ := strconv.ParseFloat(v.(string), 64)
				v2, _ := strconv.ParseFloat(value.(string), 64)
				if v1 == v2 {
					return true
				}
			} else if strings.ToLower(v.(string)) == strings.ToLower(value.(string)) {
				return true
			}
		case int8:
			switch v.(type) {
			case bool:

			case int8:
				v1 := float64(v.(int8))
				v2 := float64(value.(int8))
				if v1 == v2 {
					return true
				}
			}

		case bool:
			// todo if v is digit
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
			}

		case map[string]interface{}:
			if v == value {
				return true
			}
		case []string:
			if v == value {
				return true
			}
		}
	}
	return false
}
