package vips

// #cgo pkg-config: vips
// #include "convolution.h"
import "C"

// https://libvips.github.io/libvips/API/current/libvips-convolution.html#vips-gaussblur
func vipsGaussianBlur(in *C.VipsImage, sigma float64, precision Precision) (*C.VipsImage, error) {
	incOpCounter("gaussblur")
	var out *C.VipsImage

	if err := C.gaussian_blur_image(in, &out, C.double(sigma), C.int(precision)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}
