package filters

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/filters"
	"reflect"
	"strings"
)

func BuildQueryFromFilter(filter filters.Filter) string {
	var query []string
	// reflect allows accessing type metadata (ex: struct tags)
	fields := reflect.TypeOf(filter)
	for _, name := range filter.GetFieldNames() {
		field, ok := fields.FieldByName(name)
		if !ok {
			continue
		}
		fieldValue := reflect.ValueOf(filter).FieldByName(name)
		if !fieldValue.IsZero() {
			var qs string
			switch {
			case field.Tag.Get("array") == "true":
				qs = fmt.Sprintf("%s IN (%v)", field.Tag.Get("column"), fieldValue.Interface())
			case field.Tag.Get("like") == "true":
				qs = fmt.Sprintf("%s LIKE '%s%v%s'", field.Tag.Get("column"), "%", fieldValue.Interface(), "%")
			default:
				qs = fmt.Sprintf("%s = '%v'", field.Tag.Get("column"), fieldValue.Interface())
			}
			query = append(query, qs)
		}
	}
	return strings.Join(query, " AND ")
}
