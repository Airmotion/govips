#include "convolution.h"

int gaussian_blur_image(VipsImage *in, VipsImage **out, double sigma, int precision) {
	return vips_gaussblur(in, out, sigma, "precision", precision, NULL);
}
