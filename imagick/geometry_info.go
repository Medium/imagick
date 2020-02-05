package imagick

/*
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type GeometryInfo struct {
	gi C.GeometryInfo
}
