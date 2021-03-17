#include "convolution.h"

int gaussian_blur_image(VipsImage *in, VipsImage **out, double sigma, int precision) {
	return vips_gaussblur(in, out, sigma, "precision", precision, NULL);
}

int sharpen_image(VipsImage *in, VipsImage **out, double sigma, double x1, double m2) {
	return vips_sharpen(in, out, "sigma", sigma, "x1", x1, "m2", m2, NULL);
}
