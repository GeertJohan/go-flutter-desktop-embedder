package flutter

/*
// Linux Build Tags
// ----------------
#cgo linux CFLAGS: -I${SRCDIR}/library -I/usr/include -D_GLFW_X11
#cgo linux LDFLAGS: -lflutter_engine -Wl,-rpath,$ORIGIN -lglfw

// Windows Build Tags
// ----------------
#cgo windows CFLAGS: -I${SRCDIR}/library
#cgo windows LDFLAGS: -lflutter_engine

// Darwin Build Tags
// ----------------
#cgo darwin CFLAGS: -I${SRCDIR}/library
#cgo darwin LDFLAGS: -framework FlutterEmbedder

*/
import "C"

import (
	// prevents dep from stripping out the c source files in flutter library.
	_ "github.com/Drakirus/go-flutter-desktop-embedder/flutter/library"
)
