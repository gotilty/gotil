package config

// 			PRIVATE METHODS

var defaultSeperator string = ","

// 			PRIVATE METHODS

// 			PUBLIC METHODS

//GetDefaultSeperator variable for arrays toString()
func GetDefaultSeperator() string {
	return defaultSeperator
}

// 			PUBLIC METHODS

// Primitive types
var Types struct {
	Int64_slice  []int64
	Int32_slice  []int32
	Int16_slice  []int16
	Int8_slice   []int8
	Int_slice    []int
	Uint64_slice []uint64
	Uint32_slice []uint32
	Uint16_slice []uint16
	Uint8_slice  []uint8
	Uint_slice   []uint
	String_slice []string
}
