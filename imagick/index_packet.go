// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#cgo !no_pkgconfig pkg-config: MagickWand MagickCore
#include <magick/MagickCore.h>
*/
import "C"

type IndexPacket Quantum
