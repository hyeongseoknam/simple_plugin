package simple_plugin

import (
	"fmt"
	"time"

	"github.com/whatap/go-api/common/lang/value"
)

type addFields func(
	measurement string,
	tags map[string]string,
	fields map[string]interface{},
	t time.Time,
)

func PopulateAllTags(whatapTags *value.MapValue, tags map[string]string) {
	for k, v := range tags {
		whatapTags.PutString(k, v)
	}
}

func PopulateAllFields(whatapFields *value.MapValue, fields map[string]interface{}) {
	for k, v := range fields {
		whatapFields.Put(k, toValue(v))
	}
}

func toValue(src interface{}) (ret value.Value) {

	switch src.(type) {
	case int:
		ret = value.NewDecimalValue(int64(src.(int)))
	case int16:
		ret = value.NewDecimalValue(int64(src.(int16)))
	case int32:
		ret = value.NewDecimalValue(int64(src.(int32)))
	case int64:
		ret = value.NewDecimalValue(int64(src.(int64)))
	case float32:
		ret = value.NewFloatValue(float32(src.(float32)))
	case float64:
		ret = value.NewFloatValue(float32(src.(float64)))
	case string:
		ret = value.NewTextValue(src.(string))
	default:
		ret = value.NewTextValue(fmt.Sprint(src))
	}

	return
}
