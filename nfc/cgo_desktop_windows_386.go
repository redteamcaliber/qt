package nfc

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads
#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_NFC_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/mingw53_32/include -I/usr/local/Qt5.7.0/5.7/mingw53_32/mkspecs/win32-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtNfc -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtCore

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition
#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/mingw53_32/lib -lQt5Nfc -lQt5Core -lmingw32 -lqtmain -lshell32
*/
import "C"
