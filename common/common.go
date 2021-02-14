package common

//ValueSlice slice of values
type ValueSlice []interface{}

//Has checks whether value list has param "a" value.
func (l ValueSlice) Has(a interface{}) bool {
	for _, b := range l {
		if b == a {
			return true
		}
	}
	return false
}
