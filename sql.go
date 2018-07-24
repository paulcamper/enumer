package main

// Arguments to format are:
//	[1]: type name
const sqlValueMethod = `func (i %[1]s) Value() (driver.Value, error) {
	return i.String(), nil
}
`

// Arguments to format are:
//  [1]: type name
//  [2]: null value
const sqlValueMethodWithNull = `func (i %[1]s) Value() (driver.Value, error) {
	if i == %[2]s {
		return nil, nil
	}
	return i.String(), nil
}
`

const sqlScanMethod = `func (i *%[1]s) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := %[1]sString(str)
	if err != nil {
		return err
	}
	
	*i = val
	return nil
}
`

const sqlListMethod = `func %[1]sSqlEnumString() string {
	list := make([]string, len(_%[1]sNameToValue_map))
	idx := 0
	for k := range _%[1]sNameToValue_map {
		list[idx] = k
		idx++
	}
	return strings.Join(list, ",")
}
`

func (g *Generator) addSQLMethods(typeName string, nullValue string) {
	g.Printf("\n")
	if nullValue != "" {
		g.Printf(sqlValueMethodWithNull, typeName, nullValue)
	} else {
		g.Printf(sqlValueMethod, typeName)
	}
	g.Printf("\n\n")
	g.Printf(sqlScanMethod, typeName)
	g.Printf("\n\n")
	g.Printf(sqlListMethod, typeName)
}
