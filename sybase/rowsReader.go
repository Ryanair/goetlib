package sybase

type Rows interface {
	Columns() ([]string, error)
	Next() bool
	Scan(dest ...interface{}) error
}

func ToMap(rows Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	final_result := []map[string]interface{}{}
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		tmp_struct := map[string]interface{}{}

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmp_struct[col] = v
		}
		final_result = append(final_result, tmp_struct)
	}

	return final_result, nil
}
