package helper

import (
	"fmt"
	"strings"
)

// InQueryPlaceholders mengganti tanda tanya dalam query dengan placeholders sesuai dengan panjang slice.
func InQueryPlaceholders(query string, length int) string {
	var qword string
	for i := 0; i < length; i++ {
		qword += "?,"
	}
	if strings.HasSuffix(qword,","){
		qword = qword[:len(qword)-1]
	}
	query = strings.Replace(query,"(?)",fmt.Sprintf("(%v)",qword),-1)
	return query
}

// SliceToInterface mengonversi slice menjadi slice dari interface{}.
func SliceToInterface(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
