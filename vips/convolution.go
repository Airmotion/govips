package vips

// #include "convolution.h"
import "C"

type SharpenMode int

const (
	SharpenModeLuminescence SharpenMode = C.VIPS_SHARPEN_MODE_LUMINESCENCE
	SharpenModeRGB          SharpenMode = C.VIPS_SHARPEN_MODE_RGB
)

// https://libvips.github.io/libvips/API/current/libvips-convolution.html#vips-gaussblur
func vipsGaussianBlur(in *C.VipsImage, sigma float64, precision Precision) (*C.VipsImage, error) {
	incOpCounter("gaussblur")
	var out *C.VipsImage

	if err := C.gaussian_blur_image(in, &out, C.double(sigma), C.int(precision)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}

// https://libvips.github.io/libvips/API/current/libvips-convolution.html#vips-sharpen
func vipsSharpen(in *C.VipsImage, sigma float64, x1 float64, m2 float64, mode SharpenMode) (*C.VipsImage, error) {
	incOpCounter("sharpen")
	var out *C.VipsImage

	if err := C.sharpen_image(in, &out, C.double(sigma), C.double(x1), C.double(m2), C.VipsSharpenMode(mode)); err != 0 {
		return nil, handleImageError(out)
	}

	return out, nil
}
