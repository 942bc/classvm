package classfile

// tag 常量
const (
	CONSTANT_Class              = 7
	CONSTANT_Integer            = 3
	CONSTANT_Long               = 5
	CONSTANT_Float              = 4
	CONSTANT_Double             = 6
	CONSTANT_String             = 8
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
)

type ConstantIntegerInfo struct {
	tag   uint8
	bytes int32
}

type ConstantLongInfo struct {
	tag   uint8
	bytes int64
}

type ConstantFloatInfo struct {
	tag   uint8
	bytes float32
}

type ConstantDoubleInfo struct {
	tag   uint8
	bytes float64
}
