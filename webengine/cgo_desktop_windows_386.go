package webengine

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads
#cgo CXXFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_WEBENGINE_LIB -DQT_WIDGETS_LIB -DQT_WEBENGINEWIDGETS_LIB -DQT_WEBCHANNEL_LIB -DQT_NETWORK_LIB -DQT_WEBENGINECORE_LIB -DQT_QUICK_LIB -DQT_GUI_LIB -DQT_QML_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/mingw53_32/include -I/usr/local/Qt5.7.0/5.7/mingw53_32/mkspecs/win32-g++
#cgo CXXFLAGS: -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtWebEngine -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtWidgets -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtWebEngineWidgets -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtWebChannel -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtNetwork -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtWebEngineCore -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtQuick -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtGui -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtQml -I/usr/local/Qt5.7.0/5.7/mingw53_32/include/QtCore

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition
#cgo LDFLAGS: -L/usr/local/Qt5.7.0/5.7/mingw53_32/lib -lQt5WebEngine -lQt5Widgets -lQt5WebEngineWidgets -lQt5WebChannel -lQt5Network -lQt5WebEngineCore -lQt5Quick -lQt5Gui -lQt5Qml -lQt5Core -lmingw32 -lqtmain -lshell32
*/
import "C"
