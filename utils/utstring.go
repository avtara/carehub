package utils

import (
	"encoding/json"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetEnv(nameConfig, def string) string {
	if viper.IsSet(nameConfig) {
		return viper.GetString(nameConfig)
	}

	return def
}

func IsNil(value interface{}) (res bool) {
	return (value == nil || (reflect.TypeOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()))
}

func ToString(value interface{}) (res string) {
	if !IsNil(value) {
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			res = val.String()

		case reflect.Ptr:
			res = ToString(reflect.Indirect(val))

		default:
			switch valx := value.(type) {
			case []byte:
				res = string(valx)

			case time.Time:
				res = valx.Format(time.RFC3339Nano)

			default:
				byt, err := json.Marshal(value)
				if err == nil {
					res = string(byt)
				}
			}
		}
	}
	return
}

func ToBool(value interface{}, def bool) bool {
	vx := reflect.ValueOf(value)
	switch vx.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		switch vx.Int() {
		case 1:
			return true

		case 0:
			return false
		}

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		switch vx.Uint() {
		case 1:
			return true

		case 0:
			return false
		}

	case reflect.Bool:
		return vx.Bool()

	default:
		switch strings.ToLower(ToString(value)) {
		case "true", "1":
			return true

		case "false", "0":
			return false
		}
	}

	return def
}

func ToFloat(value interface{}, def float64) float64 {
	r, err := strconv.ParseFloat(ToString(value), 64)
	if err != nil {
		r = def
	}
	return r
}

func ToInt(value interface{}, def int64) int64 {
	r, err := strconv.ParseInt(ToString(value), 10, 64)
	if err != nil {
		r = def
	}
	return r
}

func GenerateRandomString(n int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CapitalizeString(str string) (strCap string) {
	caser := cases.Title(language.English)
	strCap = caser.String(str)
	return strCap
}
