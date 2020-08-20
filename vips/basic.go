package vips

// #cgo pkg-config: vips
// #include <vips/vips.h>
import "C"

// Precision represents VipsPrecision type https://libvips.github.io/libvips/API/current/libvips-basic.html#VipsPrecision
type Precision int

// Precision enum
const (
	PrecisionInteger     Precision = C.VIPS_PRECISION_INTEGER
	PrecisionFloat       Precision = C.VIPS_PRECISION_FLOAT
	PrecisionApproximate Precision = C.VIPS_PRECISION_APPROXIMATE
	PrecisionLast        Precision = C.VIPS_PRECISION_LAST
)
