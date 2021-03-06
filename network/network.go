// +build !minimal

package network

//#include <stdint.h>
//#include <stdlib.h>
//#include "network.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func (p *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return p
}

func (p *QAbstractNetworkCache) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractNetworkCache) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	return n
}
func NewQAbstractNetworkCache(parent core.QObject_ITF) *QAbstractNetworkCache {
	defer qt.Recovering("QAbstractNetworkCache::QAbstractNetworkCache")

	var tmpValue = NewQAbstractNetworkCacheFromPointer(C.QAbstractNetworkCache_NewQAbstractNetworkCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQAbstractNetworkCache_CacheSize
func callbackQAbstractNetworkCache_CacheSize(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QAbstractNetworkCache::cacheSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::cacheSize"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(0)
}

func (ptr *QAbstractNetworkCache) ConnectCacheSize(f func() int64) {
	defer qt.Recovering("connect QAbstractNetworkCache::cacheSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::cacheSize", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCacheSize() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::cacheSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::cacheSize")
	}
}

func (ptr *QAbstractNetworkCache) CacheSize() int64 {
	defer qt.Recovering("QAbstractNetworkCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractNetworkCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractNetworkCache_Clear
func callbackQAbstractNetworkCache_Clear(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::clear")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractNetworkCache) ConnectClear(f func()) {
	defer qt.Recovering("connect QAbstractNetworkCache::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::clear", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectClear() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::clear")
	}
}

func (ptr *QAbstractNetworkCache) Clear() {
	defer qt.Recovering("QAbstractNetworkCache::clear")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

//export callbackQAbstractNetworkCache_Data
func callbackQAbstractNetworkCache_Data(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractNetworkCache::data")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::data"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*core.QUrl) *core.QIODevice)(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(nil)
}

func (ptr *QAbstractNetworkCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	defer qt.Recovering("connect QAbstractNetworkCache::data")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::data", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectData(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::data")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::data")
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Insert
func callbackQAbstractNetworkCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::insert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::insert"); signal != nil {
		signal.(func(*core.QIODevice))(core.NewQIODeviceFromPointer(device))
	}

}

func (ptr *QAbstractNetworkCache) ConnectInsert(f func(device *core.QIODevice)) {
	defer qt.Recovering("connect QAbstractNetworkCache::insert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::insert", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectInsert(device core.QIODevice_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::insert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::insert")
	}
}

func (ptr *QAbstractNetworkCache) Insert(device core.QIODevice_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::insert")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Insert(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

//export callbackQAbstractNetworkCache_MetaData
func callbackQAbstractNetworkCache_MetaData(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractNetworkCache::metaData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData(signal.(func(*core.QUrl) *QNetworkCacheMetaData)(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(nil)
}

func (ptr *QAbstractNetworkCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	defer qt.Recovering("connect QAbstractNetworkCache::metaData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaData", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectMetaData(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::metaData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaData")
	}
}

func (ptr *QAbstractNetworkCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	defer qt.Recovering("QAbstractNetworkCache::metaData")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QAbstractNetworkCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Prepare
func callbackQAbstractNetworkCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractNetworkCache::prepare")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::prepare"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*QNetworkCacheMetaData) *core.QIODevice)(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(nil)
}

func (ptr *QAbstractNetworkCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	defer qt.Recovering("connect QAbstractNetworkCache::prepare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::prepare", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectPrepare(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::prepare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::prepare")
	}
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::prepare")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQAbstractNetworkCache_Remove
func callbackQAbstractNetworkCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractNetworkCache::remove")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl) bool)(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QAbstractNetworkCache) ConnectRemove(f func(url *core.QUrl) bool) {
	defer qt.Recovering("connect QAbstractNetworkCache::remove")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::remove", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectRemove(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::remove")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::remove")
	}
}

func (ptr *QAbstractNetworkCache) Remove(url core.QUrl_ITF) bool {
	defer qt.Recovering("QAbstractNetworkCache::remove")

	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_UpdateMetaData
func callbackQAbstractNetworkCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::updateMetaData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::updateMetaData"); signal != nil {
		signal.(func(*QNetworkCacheMetaData))(NewQNetworkCacheMetaDataFromPointer(metaData))
	}

}

func (ptr *QAbstractNetworkCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	defer qt.Recovering("connect QAbstractNetworkCache::updateMetaData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::updateMetaData", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectUpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("disconnect QAbstractNetworkCache::updateMetaData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::updateMetaData")
	}
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::updateMetaData")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

//export callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache
func callbackQAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::~QAbstractNetworkCache")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::~QAbstractNetworkCache"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DestroyQAbstractNetworkCacheDefault()
	}
}

func (ptr *QAbstractNetworkCache) ConnectDestroyQAbstractNetworkCache(f func()) {
	defer qt.Recovering("connect QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::~QAbstractNetworkCache", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDestroyQAbstractNetworkCache() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::~QAbstractNetworkCache")
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	defer qt.Recovering("QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCacheDefault() {
	defer qt.Recovering("QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCacheDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractNetworkCache_TimerEvent
func callbackQAbstractNetworkCache_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::timerEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::timerEvent")
	}
}

func (ptr *QAbstractNetworkCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractNetworkCache_ChildEvent
func callbackQAbstractNetworkCache_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::childEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::childEvent")
	}
}

func (ptr *QAbstractNetworkCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractNetworkCache_ConnectNotify
func callbackQAbstractNetworkCache_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractNetworkCache::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::connectNotify", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::connectNotify")
	}
}

func (ptr *QAbstractNetworkCache) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_CustomEvent
func callbackQAbstractNetworkCache_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::customEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::customEvent")
	}
}

func (ptr *QAbstractNetworkCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractNetworkCache_DeleteLater
func callbackQAbstractNetworkCache_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractNetworkCache) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QAbstractNetworkCache::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::deleteLater", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::deleteLater")
	}
}

func (ptr *QAbstractNetworkCache) DeleteLater() {
	defer qt.Recovering("QAbstractNetworkCache::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) DeleteLaterDefault() {
	defer qt.Recovering("QAbstractNetworkCache::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractNetworkCache_DisconnectNotify
func callbackQAbstractNetworkCache_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractNetworkCache) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractNetworkCache::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::disconnectNotify", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::disconnectNotify")
	}
}

func (ptr *QAbstractNetworkCache) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractNetworkCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractNetworkCache_Event
func callbackQAbstractNetworkCache_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractNetworkCache::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractNetworkCache) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractNetworkCache::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::event", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::event")
	}
}

func (ptr *QAbstractNetworkCache) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractNetworkCache::event")

	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractNetworkCache) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractNetworkCache::event")

	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_EventFilter
func callbackQAbstractNetworkCache_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractNetworkCache::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractNetworkCacheFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractNetworkCache) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractNetworkCache::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::eventFilter", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::eventFilter")
	}
}

func (ptr *QAbstractNetworkCache) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractNetworkCache::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractNetworkCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractNetworkCache::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractNetworkCache_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractNetworkCache_MetaObject
func callbackQAbstractNetworkCache_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractNetworkCache::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractNetworkCache::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractNetworkCacheFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractNetworkCache) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QAbstractNetworkCache::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaObject", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractNetworkCache::metaObject")
	}
}

func (ptr *QAbstractNetworkCache) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QAbstractNetworkCache::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractNetworkCache_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QAbstractNetworkCache::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractNetworkCache_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   = QAbstractSocket__BindFlag(0x4)
)

//QAbstractSocket::NetworkLayerProtocol
type QAbstractSocket__NetworkLayerProtocol int64

const (
	QAbstractSocket__IPv4Protocol                = QAbstractSocket__NetworkLayerProtocol(0)
	QAbstractSocket__IPv6Protocol                = QAbstractSocket__NetworkLayerProtocol(1)
	QAbstractSocket__AnyIPProtocol               = QAbstractSocket__NetworkLayerProtocol(2)
	QAbstractSocket__UnknownNetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(-1)
)

//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors = QAbstractSocket__PauseMode(0x1)
)

//QAbstractSocket::SocketError
type QAbstractSocket__SocketError int64

const (
	QAbstractSocket__ConnectionRefusedError           = QAbstractSocket__SocketError(0)
	QAbstractSocket__RemoteHostClosedError            = QAbstractSocket__SocketError(1)
	QAbstractSocket__HostNotFoundError                = QAbstractSocket__SocketError(2)
	QAbstractSocket__SocketAccessError                = QAbstractSocket__SocketError(3)
	QAbstractSocket__SocketResourceError              = QAbstractSocket__SocketError(4)
	QAbstractSocket__SocketTimeoutError               = QAbstractSocket__SocketError(5)
	QAbstractSocket__DatagramTooLargeError            = QAbstractSocket__SocketError(6)
	QAbstractSocket__NetworkError                     = QAbstractSocket__SocketError(7)
	QAbstractSocket__AddressInUseError                = QAbstractSocket__SocketError(8)
	QAbstractSocket__SocketAddressNotAvailableError   = QAbstractSocket__SocketError(9)
	QAbstractSocket__UnsupportedSocketOperationError  = QAbstractSocket__SocketError(10)
	QAbstractSocket__UnfinishedSocketOperationError   = QAbstractSocket__SocketError(11)
	QAbstractSocket__ProxyAuthenticationRequiredError = QAbstractSocket__SocketError(12)
	QAbstractSocket__SslHandshakeFailedError          = QAbstractSocket__SocketError(13)
	QAbstractSocket__ProxyConnectionRefusedError      = QAbstractSocket__SocketError(14)
	QAbstractSocket__ProxyConnectionClosedError       = QAbstractSocket__SocketError(15)
	QAbstractSocket__ProxyConnectionTimeoutError      = QAbstractSocket__SocketError(16)
	QAbstractSocket__ProxyNotFoundError               = QAbstractSocket__SocketError(17)
	QAbstractSocket__ProxyProtocolError               = QAbstractSocket__SocketError(18)
	QAbstractSocket__OperationError                   = QAbstractSocket__SocketError(19)
	QAbstractSocket__SslInternalError                 = QAbstractSocket__SocketError(20)
	QAbstractSocket__SslInvalidUserDataError          = QAbstractSocket__SocketError(21)
	QAbstractSocket__TemporaryError                   = QAbstractSocket__SocketError(22)
	QAbstractSocket__UnknownSocketError               = QAbstractSocket__SocketError(-1)
)

//QAbstractSocket::SocketOption
type QAbstractSocket__SocketOption int64

const (
	QAbstractSocket__LowDelayOption                = QAbstractSocket__SocketOption(0)
	QAbstractSocket__KeepAliveOption               = QAbstractSocket__SocketOption(1)
	QAbstractSocket__MulticastTtlOption            = QAbstractSocket__SocketOption(2)
	QAbstractSocket__MulticastLoopbackOption       = QAbstractSocket__SocketOption(3)
	QAbstractSocket__TypeOfServiceOption           = QAbstractSocket__SocketOption(4)
	QAbstractSocket__SendBufferSizeSocketOption    = QAbstractSocket__SocketOption(5)
	QAbstractSocket__ReceiveBufferSizeSocketOption = QAbstractSocket__SocketOption(6)
)

//QAbstractSocket::SocketState
type QAbstractSocket__SocketState int64

const (
	QAbstractSocket__UnconnectedState = QAbstractSocket__SocketState(0)
	QAbstractSocket__HostLookupState  = QAbstractSocket__SocketState(1)
	QAbstractSocket__ConnectingState  = QAbstractSocket__SocketState(2)
	QAbstractSocket__ConnectedState   = QAbstractSocket__SocketState(3)
	QAbstractSocket__BoundState       = QAbstractSocket__SocketState(4)
	QAbstractSocket__ListeningState   = QAbstractSocket__SocketState(5)
	QAbstractSocket__ClosingState     = QAbstractSocket__SocketState(6)
)

//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         = QAbstractSocket__SocketType(1)
	QAbstractSocket__UnknownSocketType = QAbstractSocket__SocketType(-1)
)

type QAbstractSocket struct {
	core.QIODevice
}

type QAbstractSocket_ITF interface {
	core.QIODevice_ITF
	QAbstractSocket_PTR() *QAbstractSocket
}

func (p *QAbstractSocket) QAbstractSocket_PTR() *QAbstractSocket {
	return p
}

func (p *QAbstractSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QAbstractSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQAbstractSocket(ptr QAbstractSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSocketFromPointer(ptr unsafe.Pointer) *QAbstractSocket {
	var n = new(QAbstractSocket)
	n.SetPointer(ptr)
	return n
}
func NewQAbstractSocket(socketType QAbstractSocket__SocketType, parent core.QObject_ITF) *QAbstractSocket {
	defer qt.Recovering("QAbstractSocket::QAbstractSocket")

	var tmpValue = NewQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QAbstractSocket) Abort() {
	defer qt.Recovering("QAbstractSocket::abort")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) AtEnd() bool {
	defer qt.Recovering("QAbstractSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode QAbstractSocket__BindFlag) bool {
	defer qt.Recovering("QAbstractSocket::bind")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Bind(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Bind2(port uint16, mode QAbstractSocket__BindFlag) bool {
	defer qt.Recovering("QAbstractSocket::bind")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Bind2(ptr.Pointer(), C.ushort(port), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) BytesAvailable() int64 {
	defer qt.Recovering("QAbstractSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) BytesToWrite() int64 {
	defer qt.Recovering("QAbstractSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) CanReadLine() bool {
	defer qt.Recovering("QAbstractSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Close() {
	defer qt.Recovering("QAbstractSocket::close")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Close(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_ConnectToHost2
func callbackQAbstractSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	defer qt.Recovering("connect QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost2", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost2() {
	defer qt.Recovering("disconnect QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost2")
	}
}

func (ptr *QAbstractSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QAbstractSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQAbstractSocket_ConnectToHost
func callbackQAbstractSocket_ConnectToHost(ptr unsafe.Pointer, hostName *C.char, port C.ushort, openMode C.longlong, protocol C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectToHostDefault(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	defer qt.Recovering("connect QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectToHost() {
	defer qt.Recovering("disconnect QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectToHost")
	}
}

func (ptr *QAbstractSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QAbstractSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QAbstractSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QAbstractSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QAbstractSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQAbstractSocket_Connected
func callbackQAbstractSocket_Connected(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::connected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connected")
	}
}

func (ptr *QAbstractSocket) Connected() {
	defer qt.Recovering("QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Connected(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_DisconnectFromHost
func callbackQAbstractSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::disconnectFromHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectFromHost", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectFromHost")
	}
}

func (ptr *QAbstractSocket) DisconnectFromHost() {
	defer qt.Recovering("QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_Disconnected
func callbackQAbstractSocket_Disconnected(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::disconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnected")
	}
}

func (ptr *QAbstractSocket) Disconnected() {
	defer qt.Recovering("QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQAbstractSocket_Error2
func callbackQAbstractSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::error")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::error2"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QAbstractSocket) ConnectError2(f func(socketError QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::error2", f)
	}
}

func (ptr *QAbstractSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::error2")
	}
}

func (ptr *QAbstractSocket) Error2(socketError QAbstractSocket__SocketError) {
	defer qt.Recovering("QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {
	defer qt.Recovering("QAbstractSocket::error")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QAbstractSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Flush() bool {
	defer qt.Recovering("QAbstractSocket::flush")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractSocket_HostFound
func callbackQAbstractSocket_HostFound(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::hostFound")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::hostFound"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {
	defer qt.Recovering("connect QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectHostFound(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::hostFound", f)
	}
}

func (ptr *QAbstractSocket) DisconnectHostFound() {
	defer qt.Recovering("disconnect QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectHostFound(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::hostFound")
	}
}

func (ptr *QAbstractSocket) HostFound() {
	defer qt.Recovering("QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_HostFound(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) IsSequential() bool {
	defer qt.Recovering("QAbstractSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsValid() bool {
	defer qt.Recovering("QAbstractSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) LocalAddress() *QHostAddress {
	defer qt.Recovering("QAbstractSocket::localAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QAbstractSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) LocalPort() uint16 {
	defer qt.Recovering("QAbstractSocket::localPort")

	if ptr.Pointer() != nil {
		return uint16(C.QAbstractSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PauseMode() QAbstractSocket__PauseMode {
	defer qt.Recovering("QAbstractSocket::pauseMode")

	if ptr.Pointer() != nil {
		return QAbstractSocket__PauseMode(C.QAbstractSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PeerAddress() *QHostAddress {
	defer qt.Recovering("QAbstractSocket::peerAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QAbstractSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) PeerName() string {
	defer qt.Recovering("QAbstractSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSocket) PeerPort() uint16 {
	defer qt.Recovering("QAbstractSocket::peerPort")

	if ptr.Pointer() != nil {
		return uint16(C.QAbstractSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Proxy() *QNetworkProxy {
	defer qt.Recovering("QAbstractSocket::proxy")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QAbstractSocket_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQAbstractSocket_ProxyAuthenticationRequired
func callbackQAbstractSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::proxyAuthenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkProxy, *QAuthenticator))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QAbstractSocket) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	defer qt.Recovering("connect QAbstractSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::proxyAuthenticationRequired", f)
	}
}

func (ptr *QAbstractSocket) DisconnectProxyAuthenticationRequired() {
	defer qt.Recovering("disconnect QAbstractSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::proxyAuthenticationRequired")
	}
}

func (ptr *QAbstractSocket) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {
	defer qt.Recovering("QAbstractSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ProxyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkProxy(proxy), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QAbstractSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QAbstractSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	defer qt.Recovering("QAbstractSocket::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxlen)))
	}
	return 0
}

//export callbackQAbstractSocket_Resume
func callbackQAbstractSocket_Resume(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::resume")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QAbstractSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QAbstractSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::resume", f)
	}
}

func (ptr *QAbstractSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QAbstractSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::resume")
	}
}

func (ptr *QAbstractSocket) Resume() {
	defer qt.Recovering("QAbstractSocket::resume")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ResumeDefault() {
	defer qt.Recovering("QAbstractSocket::resume")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) SetLocalAddress(address QHostAddress_ITF) {
	defer qt.Recovering("QAbstractSocket::setLocalAddress")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetLocalAddress(ptr.Pointer(), PointerFromQHostAddress(address))
	}
}

func (ptr *QAbstractSocket) SetLocalPort(port uint16) {
	defer qt.Recovering("QAbstractSocket::setLocalPort")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetLocalPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QAbstractSocket) SetPauseMode(pauseMode QAbstractSocket__PauseMode) {
	defer qt.Recovering("QAbstractSocket::setPauseMode")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPauseMode(ptr.Pointer(), C.longlong(pauseMode))
	}
}

func (ptr *QAbstractSocket) SetPeerAddress(address QHostAddress_ITF) {
	defer qt.Recovering("QAbstractSocket::setPeerAddress")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPeerAddress(ptr.Pointer(), PointerFromQHostAddress(address))
	}
}

func (ptr *QAbstractSocket) SetPeerName(name string) {
	defer qt.Recovering("QAbstractSocket::setPeerName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QAbstractSocket_SetPeerName(ptr.Pointer(), nameC)
	}
}

func (ptr *QAbstractSocket) SetPeerPort(port uint16) {
	defer qt.Recovering("QAbstractSocket::setPeerPort")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPeerPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QAbstractSocket::setProxy")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

//export callbackQAbstractSocket_SetReadBufferSize
func callbackQAbstractSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::setReadBufferSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QAbstractSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setReadBufferSize", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setReadBufferSize")
	}
}

func (ptr *QAbstractSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) SetSocketError(socketError QAbstractSocket__SocketError) {
	defer qt.Recovering("QAbstractSocket::setSocketError")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketError(ptr.Pointer(), C.longlong(socketError))
	}
}

//export callbackQAbstractSocket_SetSocketOption
func callbackQAbstractSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::setSocketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QAbstractSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setSocketOption", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::setSocketOption")
	}
}

func (ptr *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SetSocketState(state QAbstractSocket__SocketState) {
	defer qt.Recovering("QAbstractSocket::setSocketState")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketState(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQAbstractSocket_SocketOption
func callbackQAbstractSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractSocket::socketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQAbstractSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QAbstractSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	defer qt.Recovering("connect QAbstractSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::socketOption", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSocketOption() {
	defer qt.Recovering("disconnect QAbstractSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::socketOption")
	}
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QAbstractSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QAbstractSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QAbstractSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAbstractSocket) SocketType() QAbstractSocket__SocketType {
	defer qt.Recovering("QAbstractSocket::socketType")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketType(C.QAbstractSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) State() QAbstractSocket__SocketState {
	defer qt.Recovering("QAbstractSocket::state")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketState(C.QAbstractSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_StateChanged
func callbackQAbstractSocket_StateChanged(ptr unsafe.Pointer, socketState C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::stateChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::stateChanged"); signal != nil {
		signal.(func(QAbstractSocket__SocketState))(QAbstractSocket__SocketState(socketState))
	}

}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {
	defer qt.Recovering("connect QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::stateChanged", f)
	}
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::stateChanged")
	}
}

func (ptr *QAbstractSocket) StateChanged(socketState QAbstractSocket__SocketState) {
	defer qt.Recovering("QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

func (ptr *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForConnected
func callbackQAbstractSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QAbstractSocket::waitForConnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QAbstractSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForConnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForConnected() {
	defer qt.Recovering("disconnect QAbstractSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForConnected")
	}
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForConnectedDefault(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQAbstractSocket_WaitForDisconnected
func callbackQAbstractSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QAbstractSocket::waitForDisconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QAbstractSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QAbstractSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForDisconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectWaitForDisconnected() {
	defer qt.Recovering("disconnect QAbstractSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::waitForDisconnected")
	}
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForDisconnectedDefault(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WriteData(data string, size int64) int64 {
	defer qt.Recovering("QAbstractSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QAbstractSocket_WriteData(ptr.Pointer(), dataC, C.longlong(size)))
	}
	return 0
}

//export callbackQAbstractSocket_DestroyQAbstractSocket
func callbackQAbstractSocket_DestroyQAbstractSocket(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::~QAbstractSocket")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::~QAbstractSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DestroyQAbstractSocketDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDestroyQAbstractSocket(f func()) {
	defer qt.Recovering("connect QAbstractSocket::~QAbstractSocket")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::~QAbstractSocket", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDestroyQAbstractSocket() {
	defer qt.Recovering("disconnect QAbstractSocket::~QAbstractSocket")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::~QAbstractSocket")
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {
	defer qt.Recovering("QAbstractSocket::~QAbstractSocket")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) DestroyQAbstractSocketDefault() {
	defer qt.Recovering("QAbstractSocket::~QAbstractSocket")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractSocket_Open
func callbackQAbstractSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QAbstractSocket::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QAbstractSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QAbstractSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::open", f)
	}
}

func (ptr *QAbstractSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QAbstractSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::open")
	}
}

func (ptr *QAbstractSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QAbstractSocket::open")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QAbstractSocket::open")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQAbstractSocket_Pos
func callbackQAbstractSocket_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QAbstractSocket::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).PosDefault())
}

func (ptr *QAbstractSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QAbstractSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::pos", f)
	}
}

func (ptr *QAbstractSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QAbstractSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::pos")
	}
}

func (ptr *QAbstractSocket) Pos() int64 {
	defer qt.Recovering("QAbstractSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PosDefault() int64 {
	defer qt.Recovering("QAbstractSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_Reset
func callbackQAbstractSocket_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractSocket::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QAbstractSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QAbstractSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::reset", f)
	}
}

func (ptr *QAbstractSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QAbstractSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::reset")
	}
}

func (ptr *QAbstractSocket) Reset() bool {
	defer qt.Recovering("QAbstractSocket::reset")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) ResetDefault() bool {
	defer qt.Recovering("QAbstractSocket::reset")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQAbstractSocket_Seek
func callbackQAbstractSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QAbstractSocket::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QAbstractSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QAbstractSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::seek", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QAbstractSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::seek")
	}
}

func (ptr *QAbstractSocket) Seek(pos int64) bool {
	defer qt.Recovering("QAbstractSocket::seek")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QAbstractSocket::seek")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQAbstractSocket_Size
func callbackQAbstractSocket_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QAbstractSocket::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQAbstractSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QAbstractSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QAbstractSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::size", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QAbstractSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::size")
	}
}

func (ptr *QAbstractSocket) Size() int64 {
	defer qt.Recovering("QAbstractSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) SizeDefault() int64 {
	defer qt.Recovering("QAbstractSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQAbstractSocket_TimerEvent
func callbackQAbstractSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::timerEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::timerEvent")
	}
}

func (ptr *QAbstractSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQAbstractSocket_ChildEvent
func callbackQAbstractSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::childEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::childEvent")
	}
}

func (ptr *QAbstractSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQAbstractSocket_ConnectNotify
func callbackQAbstractSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectNotify", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QAbstractSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::connectNotify")
	}
}

func (ptr *QAbstractSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_CustomEvent
func callbackQAbstractSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::customEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::customEvent")
	}
}

func (ptr *QAbstractSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQAbstractSocket_DeleteLater
func callbackQAbstractSocket_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QAbstractSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QAbstractSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::deleteLater", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QAbstractSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::deleteLater")
	}
}

func (ptr *QAbstractSocket) DeleteLater() {
	defer qt.Recovering("QAbstractSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) DeleteLaterDefault() {
	defer qt.Recovering("QAbstractSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQAbstractSocket_DisconnectNotify
func callbackQAbstractSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QAbstractSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectNotify", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QAbstractSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::disconnectNotify")
	}
}

func (ptr *QAbstractSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QAbstractSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QAbstractSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQAbstractSocket_Event
func callbackQAbstractSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractSocket::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QAbstractSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::event", f)
	}
}

func (ptr *QAbstractSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::event")
	}
}

func (ptr *QAbstractSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSocket::event")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSocket::event")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQAbstractSocket_EventFilter
func callbackQAbstractSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QAbstractSocket::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQAbstractSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QAbstractSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QAbstractSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::eventFilter", f)
	}
}

func (ptr *QAbstractSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QAbstractSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::eventFilter")
	}
}

func (ptr *QAbstractSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQAbstractSocket_MetaObject
func callbackQAbstractSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QAbstractSocket::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QAbstractSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQAbstractSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QAbstractSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QAbstractSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::metaObject", f)
	}
}

func (ptr *QAbstractSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QAbstractSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QAbstractSocket::metaObject")
	}
}

func (ptr *QAbstractSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QAbstractSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QAbstractSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QAbstractSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (p *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return p
}

func (p *QAuthenticator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QAuthenticator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}
func NewQAuthenticator() *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	var tmpValue = NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
	runtime.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	var tmpValue = NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QAuthenticator).DestroyQAuthenticator)
	return tmpValue
}

func (ptr *QAuthenticator) IsNull() bool {
	defer qt.Recovering("QAuthenticator::isNull")

	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	defer qt.Recovering("QAuthenticator::option")

	if ptr.Pointer() != nil {
		var optC = C.CString(opt)
		defer C.free(unsafe.Pointer(optC))
		var tmpValue = core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), optC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QAuthenticator) Password() string {
	defer qt.Recovering("QAuthenticator::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	defer qt.Recovering("QAuthenticator::realm")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Realm(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {
	defer qt.Recovering("QAuthenticator::setOption")

	if ptr.Pointer() != nil {
		var optC = C.CString(opt)
		defer C.free(unsafe.Pointer(optC))
		C.QAuthenticator_SetOption(ptr.Pointer(), optC, core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	defer qt.Recovering("QAuthenticator::setPassword")

	if ptr.Pointer() != nil {
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		C.QAuthenticator_SetPassword(ptr.Pointer(), passwordC)
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	defer qt.Recovering("QAuthenticator::setUser")

	if ptr.Pointer() != nil {
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		C.QAuthenticator_SetUser(ptr.Pointer(), userC)
	}
}

func (ptr *QAuthenticator) User() string {
	defer qt.Recovering("QAuthenticator::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	defer qt.Recovering("QAuthenticator::~QAuthenticator")

	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsDomainNameRecord struct {
	ptr unsafe.Pointer
}

type QDnsDomainNameRecord_ITF interface {
	QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord
}

func (p *QDnsDomainNameRecord) QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord {
	return p
}

func (p *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDnsDomainNameRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	defer qt.Recovering("QDnsDomainNameRecord::QDnsDomainNameRecord")

	var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
	return tmpValue
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {
	defer qt.Recovering("QDnsDomainNameRecord::QDnsDomainNameRecord")

	var tmpValue = NewQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(PointerFromQDnsDomainNameRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsDomainNameRecord).DestroyQDnsDomainNameRecord)
	return tmpValue
}

func (ptr *QDnsDomainNameRecord) Name() string {
	defer qt.Recovering("QDnsDomainNameRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecord_ITF) {
	defer qt.Recovering("QDnsDomainNameRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_Swap(ptr.Pointer(), PointerFromQDnsDomainNameRecord(other))
	}
}

func (ptr *QDnsDomainNameRecord) TimeToLive() uint {
	defer qt.Recovering("QDnsDomainNameRecord::timeToLive")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsDomainNameRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsDomainNameRecord) Value() string {
	defer qt.Recovering("QDnsDomainNameRecord::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {
	defer qt.Recovering("QDnsDomainNameRecord::~QDnsDomainNameRecord")

	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsHostAddressRecord struct {
	ptr unsafe.Pointer
}

type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (p *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord {
	return p
}

func (p *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDnsHostAddressRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	defer qt.Recovering("QDnsHostAddressRecord::QDnsHostAddressRecord")

	var tmpValue = NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
	return tmpValue
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {
	defer qt.Recovering("QDnsHostAddressRecord::QDnsHostAddressRecord")

	var tmpValue = NewQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(PointerFromQDnsHostAddressRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsHostAddressRecord).DestroyQDnsHostAddressRecord)
	return tmpValue
}

func (ptr *QDnsHostAddressRecord) Name() string {
	defer qt.Recovering("QDnsHostAddressRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsHostAddressRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {
	defer qt.Recovering("QDnsHostAddressRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_Swap(ptr.Pointer(), PointerFromQDnsHostAddressRecord(other))
	}
}

func (ptr *QDnsHostAddressRecord) TimeToLive() uint {
	defer qt.Recovering("QDnsHostAddressRecord::timeToLive")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsHostAddressRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsHostAddressRecord) Value() *QHostAddress {
	defer qt.Recovering("QDnsHostAddressRecord::value")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QDnsHostAddressRecord_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	defer qt.Recovering("QDnsHostAddressRecord::~QDnsHostAddressRecord")

	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           = QDnsLookup__Error(7)
)

//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     = QDnsLookup__Type(1)
	QDnsLookup__AAAA  = QDnsLookup__Type(28)
	QDnsLookup__ANY   = QDnsLookup__Type(255)
	QDnsLookup__CNAME = QDnsLookup__Type(5)
	QDnsLookup__MX    = QDnsLookup__Type(15)
	QDnsLookup__NS    = QDnsLookup__Type(2)
	QDnsLookup__PTR   = QDnsLookup__Type(12)
	QDnsLookup__SRV   = QDnsLookup__Type(33)
	QDnsLookup__TXT   = QDnsLookup__Type(16)
)

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func (p *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return p
}

func (p *QDnsLookup) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QDnsLookup) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	return n
}
func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.longlong(ty), nameC, PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	defer qt.Recovering("QDnsLookup::error")

	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	defer qt.Recovering("QDnsLookup::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	defer qt.Recovering("QDnsLookup::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Nameserver() *QHostAddress {
	defer qt.Recovering("QDnsLookup::nameserver")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QDnsLookup_Nameserver(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QDnsLookup) SetName(name string) {
	defer qt.Recovering("QDnsLookup::setName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDnsLookup_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	defer qt.Recovering("QDnsLookup::setNameserver")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) SetType(vqd QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::setType")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(ptr.Pointer(), C.longlong(vqd))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	defer qt.Recovering("QDnsLookup::type")

	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.longlong(ty), nameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQDnsLookup_Abort
func callbackQDnsLookup_Abort(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::abort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::abort"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) ConnectAbort(f func()) {
	defer qt.Recovering("connect QDnsLookup::abort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::abort", f)
	}
}

func (ptr *QDnsLookup) DisconnectAbort() {
	defer qt.Recovering("disconnect QDnsLookup::abort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::abort")
	}
}

func (ptr *QDnsLookup) Abort() {
	defer qt.Recovering("QDnsLookup::abort")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(ptr.Pointer())
	}
}

//export callbackQDnsLookup_Finished
func callbackQDnsLookup_Finished(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::finished")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	defer qt.Recovering("connect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	defer qt.Recovering("disconnect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::finished")
	}
}

func (ptr *QDnsLookup) Finished() {
	defer qt.Recovering("QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Finished(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) IsFinished() bool {
	defer qt.Recovering("QDnsLookup::isFinished")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQDnsLookup_Lookup
func callbackQDnsLookup_Lookup(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::lookup")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::lookup"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) ConnectLookup(f func()) {
	defer qt.Recovering("connect QDnsLookup::lookup")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::lookup", f)
	}
}

func (ptr *QDnsLookup) DisconnectLookup() {
	defer qt.Recovering("disconnect QDnsLookup::lookup")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::lookup")
	}
}

func (ptr *QDnsLookup) Lookup() {
	defer qt.Recovering("QDnsLookup::lookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(ptr.Pointer())
	}
}

//export callbackQDnsLookup_NameChanged
func callbackQDnsLookup_NameChanged(ptr unsafe.Pointer, name *C.char) {
	defer qt.Recovering("callback QDnsLookup::nameChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::nameChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	defer qt.Recovering("connect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	defer qt.Recovering("disconnect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameChanged")
	}
}

func (ptr *QDnsLookup) NameChanged(name string) {
	defer qt.Recovering("QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QDnsLookup_NameChanged(ptr.Pointer(), nameC)
	}
}

//export callbackQDnsLookup_NameserverChanged
func callbackQDnsLookup_NameserverChanged(ptr unsafe.Pointer, nameserver unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::nameserverChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::nameserverChanged"); signal != nil {
		signal.(func(*QHostAddress))(NewQHostAddressFromPointer(nameserver))
	}

}

func (ptr *QDnsLookup) ConnectNameserverChanged(f func(nameserver *QHostAddress)) {
	defer qt.Recovering("connect QDnsLookup::nameserverChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameserverChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameserverChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameserverChanged() {
	defer qt.Recovering("disconnect QDnsLookup::nameserverChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameserverChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::nameserverChanged")
	}
}

func (ptr *QDnsLookup) NameserverChanged(nameserver QHostAddress_ITF) {
	defer qt.Recovering("QDnsLookup::nameserverChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_NameserverChanged(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

//export callbackQDnsLookup_TypeChanged
func callbackQDnsLookup_TypeChanged(ptr unsafe.Pointer, ty C.longlong) {
	defer qt.Recovering("callback QDnsLookup::typeChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::typeChanged"); signal != nil {
		signal.(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
	}

}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	defer qt.Recovering("connect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::typeChanged")
	}
}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TypeChanged(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	defer qt.Recovering("QDnsLookup::~QDnsLookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDnsLookup_TimerEvent
func callbackQDnsLookup_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::timerEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::timerEvent")
	}
}

func (ptr *QDnsLookup) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQDnsLookup_ChildEvent
func callbackQDnsLookup_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::childEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::childEvent")
	}
}

func (ptr *QDnsLookup) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQDnsLookup_ConnectNotify
func callbackQDnsLookup_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDnsLookup::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::connectNotify", f)
	}
}

func (ptr *QDnsLookup) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QDnsLookup::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::connectNotify")
	}
}

func (ptr *QDnsLookup) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDnsLookup::connectNotify")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDnsLookup) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDnsLookup::connectNotify")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_CustomEvent
func callbackQDnsLookup_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::customEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::customEvent")
	}
}

func (ptr *QDnsLookup) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQDnsLookup_DeleteLater
func callbackQDnsLookup_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQDnsLookupFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QDnsLookup) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QDnsLookup::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::deleteLater", f)
	}
}

func (ptr *QDnsLookup) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QDnsLookup::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::deleteLater")
	}
}

func (ptr *QDnsLookup) DeleteLater() {
	defer qt.Recovering("QDnsLookup::deleteLater")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) DeleteLaterDefault() {
	defer qt.Recovering("QDnsLookup::deleteLater")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQDnsLookup_DisconnectNotify
func callbackQDnsLookup_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQDnsLookupFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QDnsLookup) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QDnsLookup::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::disconnectNotify", f)
	}
}

func (ptr *QDnsLookup) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QDnsLookup::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::disconnectNotify")
	}
}

func (ptr *QDnsLookup) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDnsLookup::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QDnsLookup) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QDnsLookup::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQDnsLookup_Event
func callbackQDnsLookup_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDnsLookup::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QDnsLookup) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QDnsLookup::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::event", f)
	}
}

func (ptr *QDnsLookup) DisconnectEvent() {
	defer qt.Recovering("disconnect QDnsLookup::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::event")
	}
}

func (ptr *QDnsLookup) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDnsLookup::event")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDnsLookup) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDnsLookup::event")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQDnsLookup_EventFilter
func callbackQDnsLookup_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QDnsLookup::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQDnsLookupFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QDnsLookup) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QDnsLookup::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::eventFilter", f)
	}
}

func (ptr *QDnsLookup) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QDnsLookup::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::eventFilter")
	}
}

func (ptr *QDnsLookup) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDnsLookup::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDnsLookup) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QDnsLookup::eventFilter")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQDnsLookup_MetaObject
func callbackQDnsLookup_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QDnsLookup::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QDnsLookup::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQDnsLookupFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QDnsLookup) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QDnsLookup::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::metaObject", f)
	}
}

func (ptr *QDnsLookup) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QDnsLookup::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QDnsLookup::metaObject")
	}
}

func (ptr *QDnsLookup) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QDnsLookup::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDnsLookup_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDnsLookup) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QDnsLookup::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QDnsLookup_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QDnsMailExchangeRecord struct {
	ptr unsafe.Pointer
}

type QDnsMailExchangeRecord_ITF interface {
	QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord
}

func (p *QDnsMailExchangeRecord) QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord {
	return p
}

func (p *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDnsMailExchangeRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	var tmpValue = NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
	return tmpValue
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	var tmpValue = NewQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(PointerFromQDnsMailExchangeRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsMailExchangeRecord).DestroyQDnsMailExchangeRecord)
	return tmpValue
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {
	defer qt.Recovering("QDnsMailExchangeRecord::exchange")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Exchange(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Name() string {
	defer qt.Recovering("QDnsMailExchangeRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Preference() uint16 {
	defer qt.Recovering("QDnsMailExchangeRecord::preference")

	if ptr.Pointer() != nil {
		return uint16(C.QDnsMailExchangeRecord_Preference(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecord_ITF) {
	defer qt.Recovering("QDnsMailExchangeRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_Swap(ptr.Pointer(), PointerFromQDnsMailExchangeRecord(other))
	}
}

func (ptr *QDnsMailExchangeRecord) TimeToLive() uint {
	defer qt.Recovering("QDnsMailExchangeRecord::timeToLive")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsMailExchangeRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {
	defer qt.Recovering("QDnsMailExchangeRecord::~QDnsMailExchangeRecord")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsServiceRecord struct {
	ptr unsafe.Pointer
}

type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (p *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord {
	return p
}

func (p *QDnsServiceRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDnsServiceRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsServiceRecord() *QDnsServiceRecord {
	defer qt.Recovering("QDnsServiceRecord::QDnsServiceRecord")

	var tmpValue = NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
	return tmpValue
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	defer qt.Recovering("QDnsServiceRecord::QDnsServiceRecord")

	var tmpValue = NewQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(PointerFromQDnsServiceRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsServiceRecord).DestroyQDnsServiceRecord)
	return tmpValue
}

func (ptr *QDnsServiceRecord) Name() string {
	defer qt.Recovering("QDnsServiceRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) Port() uint16 {
	defer qt.Recovering("QDnsServiceRecord::port")

	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Port(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Priority() uint16 {
	defer qt.Recovering("QDnsServiceRecord::priority")

	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {
	defer qt.Recovering("QDnsServiceRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_Swap(ptr.Pointer(), PointerFromQDnsServiceRecord(other))
	}
}

func (ptr *QDnsServiceRecord) Target() string {
	defer qt.Recovering("QDnsServiceRecord::target")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Target(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) TimeToLive() uint {
	defer qt.Recovering("QDnsServiceRecord::timeToLive")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsServiceRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsServiceRecord) Weight() uint16 {
	defer qt.Recovering("QDnsServiceRecord::weight")

	if ptr.Pointer() != nil {
		return uint16(C.QDnsServiceRecord_Weight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {
	defer qt.Recovering("QDnsServiceRecord::~QDnsServiceRecord")

	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QDnsTextRecord struct {
	ptr unsafe.Pointer
}

type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (p *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord {
	return p
}

func (p *QDnsTextRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QDnsTextRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return n
}
func NewQDnsTextRecord() *QDnsTextRecord {
	defer qt.Recovering("QDnsTextRecord::QDnsTextRecord")

	var tmpValue = NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord())
	runtime.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
	return tmpValue
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {
	defer qt.Recovering("QDnsTextRecord::QDnsTextRecord")

	var tmpValue = NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord2(PointerFromQDnsTextRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QDnsTextRecord).DestroyQDnsTextRecord)
	return tmpValue
}

func (ptr *QDnsTextRecord) Name() string {
	defer qt.Recovering("QDnsTextRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsTextRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {
	defer qt.Recovering("QDnsTextRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_Swap(ptr.Pointer(), PointerFromQDnsTextRecord(other))
	}
}

func (ptr *QDnsTextRecord) TimeToLive() uint {
	defer qt.Recovering("QDnsTextRecord::timeToLive")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QDnsTextRecord_TimeToLive(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	defer qt.Recovering("QDnsTextRecord::~QDnsTextRecord")

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_DestroyQDnsTextRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       = QHostAddress__SpecialAddress(6)
)

type QHostAddress struct {
	ptr unsafe.Pointer
}

type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (p *QHostAddress) QHostAddress_PTR() *QHostAddress {
	return p
}

func (p *QHostAddress) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHostAddress) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHostAddress(ptr QHostAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostAddress_PTR().Pointer()
	}
	return nil
}

func NewQHostAddressFromPointer(ptr unsafe.Pointer) *QHostAddress {
	var n = new(QHostAddress)
	n.SetPointer(ptr)
	return n
}
func NewQHostAddress() *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress())
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress9(C.longlong(address)))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress8(PointerFromQHostAddress(address)))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress7(address string) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress7(addressC))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func NewQHostAddress2(ip4Addr uint) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	var tmpValue = NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress2(C.uint(uint32(ip4Addr))))
	runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
	return tmpValue
}

func (ptr *QHostAddress) Clear() {
	defer qt.Recovering("QHostAddress::clear")

	if ptr.Pointer() != nil {
		C.QHostAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	defer qt.Recovering("QHostAddress::isInSubnet")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsInSubnet(ptr.Pointer(), PointerFromQHostAddress(subnet), C.int(int32(netmask))) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLoopback() bool {
	defer qt.Recovering("QHostAddress::isLoopback")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsLoopback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsMulticast() bool {
	defer qt.Recovering("QHostAddress::isMulticast")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsMulticast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsNull() bool {
	defer qt.Recovering("QHostAddress::isNull")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) Protocol() QAbstractSocket__NetworkLayerProtocol {
	defer qt.Recovering("QHostAddress::protocol")

	if ptr.Pointer() != nil {
		return QAbstractSocket__NetworkLayerProtocol(C.QHostAddress_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostAddress) ScopeId() string {
	defer qt.Recovering("QHostAddress::scopeId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ScopeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) SetAddress6(address string) bool {
	defer qt.Recovering("QHostAddress::setAddress")

	if ptr.Pointer() != nil {
		var addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
		return C.QHostAddress_SetAddress6(ptr.Pointer(), addressC) != 0
	}
	return false
}

func (ptr *QHostAddress) SetAddress(ip4Addr uint) {
	defer qt.Recovering("QHostAddress::setAddress")

	if ptr.Pointer() != nil {
		C.QHostAddress_SetAddress(ptr.Pointer(), C.uint(uint32(ip4Addr)))
	}
}

func (ptr *QHostAddress) SetScopeId(id string) {
	defer qt.Recovering("QHostAddress::setScopeId")

	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QHostAddress_SetScopeId(ptr.Pointer(), idC)
	}
}

func (ptr *QHostAddress) Swap(other QHostAddress_ITF) {
	defer qt.Recovering("QHostAddress::swap")

	if ptr.Pointer() != nil {
		C.QHostAddress_Swap(ptr.Pointer(), PointerFromQHostAddress(other))
	}
}

func (ptr *QHostAddress) ToIPv4Address() uint {
	defer qt.Recovering("QHostAddress::toIPv4Address")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QHostAddress_ToIPv4Address(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostAddress) ToIPv4Address2(ok bool) uint {
	defer qt.Recovering("QHostAddress::toIPv4Address")

	if ptr.Pointer() != nil {
		return uint(uint32(C.QHostAddress_ToIPv4Address2(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))))
	}
	return 0
}

func (ptr *QHostAddress) ToString() string {
	defer qt.Recovering("QHostAddress::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) DestroyQHostAddress() {
	defer qt.Recovering("QHostAddress::~QHostAddress")

	if ptr.Pointer() != nil {
		C.QHostAddress_DestroyQHostAddress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError = QHostInfo__HostInfoError(2)
)

type QHostInfo struct {
	ptr unsafe.Pointer
}

type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (p *QHostInfo) QHostInfo_PTR() *QHostInfo {
	return p
}

func (p *QHostInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHostInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHostInfo(ptr QHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = new(QHostInfo)
	n.SetPointer(ptr)
	return n
}
func QHostInfo_LocalHostName() string {
	defer qt.Recovering("QHostInfo::localHostName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func (ptr *QHostInfo) LocalHostName() string {
	defer qt.Recovering("QHostInfo::localHostName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {
	defer qt.Recovering("QHostInfo::QHostInfo")

	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func NewQHostInfo(id int) *QHostInfo {
	defer qt.Recovering("QHostInfo::QHostInfo")

	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(int32(id))))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func QHostInfo_AbortHostLookup(id int) {
	defer qt.Recovering("QHostInfo::abortHostLookup")

	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(int32(id)))
}

func (ptr *QHostInfo) AbortHostLookup(id int) {
	defer qt.Recovering("QHostInfo::abortHostLookup")

	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(int32(id)))
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	defer qt.Recovering("QHostInfo::error")

	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	defer qt.Recovering("QHostInfo::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_ErrorString(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_FromName(name string) *QHostInfo {
	defer qt.Recovering("QHostInfo::fromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) FromName(name string) *QHostInfo {
	defer qt.Recovering("QHostInfo::fromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQHostInfoFromPointer(C.QHostInfo_QHostInfo_FromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QHostInfo).DestroyQHostInfo)
	return tmpValue
}

func (ptr *QHostInfo) HostName() string {
	defer qt.Recovering("QHostInfo::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	defer qt.Recovering("QHostInfo::lookupHost")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var memberC = C.CString(member)
	defer C.free(unsafe.Pointer(memberC))
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(nameC, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupHost(name string, receiver core.QObject_ITF, member string) int {
	defer qt.Recovering("QHostInfo::lookupHost")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var memberC = C.CString(member)
	defer C.free(unsafe.Pointer(memberC))
	return int(int32(C.QHostInfo_QHostInfo_LookupHost(nameC, core.PointerFromQObject(receiver), memberC)))
}

func (ptr *QHostInfo) LookupId() int {
	defer qt.Recovering("QHostInfo::lookupId")

	if ptr.Pointer() != nil {
		return int(int32(C.QHostInfo_LookupId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	defer qt.Recovering("QHostInfo::setError")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	defer qt.Recovering("QHostInfo::setErrorString")

	if ptr.Pointer() != nil {
		var strC = C.CString(str)
		defer C.free(unsafe.Pointer(strC))
		C.QHostInfo_SetErrorString(ptr.Pointer(), strC)
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	defer qt.Recovering("QHostInfo::setHostName")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QHostInfo_SetHostName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	defer qt.Recovering("QHostInfo::setLookupId")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(int32(id)))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	defer qt.Recovering("QHostInfo::~QHostInfo")

	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QHostInfo_LocalDomainName() string {
	defer qt.Recovering("QHostInfo::localDomainName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}

func (ptr *QHostInfo) LocalDomainName() string {
	defer qt.Recovering("QHostInfo::localDomainName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func (p *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return p
}

func (p *QHttpMultiPart) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QHttpMultiPart) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	return n
}
func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	var tmpValue = NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.longlong(contentType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	var tmpValue = NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	defer qt.Recovering("QHttpMultiPart::append")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() string {
	defer qt.Recovering("QHttpMultiPart::boundary")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QHttpMultiPart_Boundary(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHttpMultiPart) SetBoundary(boundary string) {
	defer qt.Recovering("QHttpMultiPart::setBoundary")

	if ptr.Pointer() != nil {
		var boundaryC = C.CString(hex.EncodeToString([]byte(boundary)))
		defer C.free(unsafe.Pointer(boundaryC))
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), boundaryC)
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	defer qt.Recovering("QHttpMultiPart::setContentType")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.longlong(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	defer qt.Recovering("QHttpMultiPart::~QHttpMultiPart")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHttpMultiPart_TimerEvent
func callbackQHttpMultiPart_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::timerEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::timerEvent")
	}
}

func (ptr *QHttpMultiPart) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQHttpMultiPart_ChildEvent
func callbackQHttpMultiPart_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::childEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::childEvent")
	}
}

func (ptr *QHttpMultiPart) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQHttpMultiPart_ConnectNotify
func callbackQHttpMultiPart_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHttpMultiPart::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::connectNotify", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QHttpMultiPart::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::connectNotify")
	}
}

func (ptr *QHttpMultiPart) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHttpMultiPart::connectNotify")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHttpMultiPart::connectNotify")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_CustomEvent
func callbackQHttpMultiPart_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::customEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::customEvent")
	}
}

func (ptr *QHttpMultiPart) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQHttpMultiPart_DeleteLater
func callbackQHttpMultiPart_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQHttpMultiPartFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QHttpMultiPart) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QHttpMultiPart::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::deleteLater", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QHttpMultiPart::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::deleteLater")
	}
}

func (ptr *QHttpMultiPart) DeleteLater() {
	defer qt.Recovering("QHttpMultiPart::deleteLater")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) DeleteLaterDefault() {
	defer qt.Recovering("QHttpMultiPart::deleteLater")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQHttpMultiPart_DisconnectNotify
func callbackQHttpMultiPart_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQHttpMultiPartFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QHttpMultiPart) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QHttpMultiPart::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::disconnectNotify", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QHttpMultiPart::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::disconnectNotify")
	}
}

func (ptr *QHttpMultiPart) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHttpMultiPart::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QHttpMultiPart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QHttpMultiPart::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQHttpMultiPart_Event
func callbackQHttpMultiPart_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHttpMultiPart::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QHttpMultiPart) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QHttpMultiPart::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::event", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::event")
	}
}

func (ptr *QHttpMultiPart) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHttpMultiPart::event")

	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QHttpMultiPart) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QHttpMultiPart::event")

	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQHttpMultiPart_EventFilter
func callbackQHttpMultiPart_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QHttpMultiPart::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQHttpMultiPartFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QHttpMultiPart) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QHttpMultiPart::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::eventFilter", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QHttpMultiPart::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::eventFilter")
	}
}

func (ptr *QHttpMultiPart) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHttpMultiPart::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QHttpMultiPart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QHttpMultiPart::eventFilter")

	if ptr.Pointer() != nil {
		return C.QHttpMultiPart_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQHttpMultiPart_MetaObject
func callbackQHttpMultiPart_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QHttpMultiPart::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QHttpMultiPart::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQHttpMultiPartFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QHttpMultiPart) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QHttpMultiPart::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::metaObject", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QHttpMultiPart::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QHttpMultiPart::metaObject")
	}
}

func (ptr *QHttpMultiPart) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QHttpMultiPart::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHttpMultiPart_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QHttpMultiPart::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QHttpMultiPart_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QHttpPart struct {
	ptr unsafe.Pointer
}

type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (p *QHttpPart) QHttpPart_PTR() *QHttpPart {
	return p
}

func (p *QHttpPart) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QHttpPart) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQHttpPart(ptr QHttpPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = new(QHttpPart)
	n.SetPointer(ptr)
	return n
}
func NewQHttpPart() *QHttpPart {
	defer qt.Recovering("QHttpPart::QHttpPart")

	var tmpValue = NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart())
	runtime.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
	return tmpValue
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {
	defer qt.Recovering("QHttpPart::QHttpPart")

	var tmpValue = NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart2(PointerFromQHttpPart(other)))
	runtime.SetFinalizer(tmpValue, (*QHttpPart).DestroyQHttpPart)
	return tmpValue
}

func (ptr *QHttpPart) SetBody(body string) {
	defer qt.Recovering("QHttpPart::setBody")

	if ptr.Pointer() != nil {
		var bodyC = C.CString(hex.EncodeToString([]byte(body)))
		defer C.free(unsafe.Pointer(bodyC))
		C.QHttpPart_SetBody(ptr.Pointer(), bodyC)
	}
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QHttpPart::setBodyDevice")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetBodyDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QHttpPart::setHeader")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QHttpPart) SetRawHeader(headerName string, headerValue string) {
	defer qt.Recovering("QHttpPart::setRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		var headerValueC = C.CString(hex.EncodeToString([]byte(headerValue)))
		defer C.free(unsafe.Pointer(headerValueC))
		C.QHttpPart_SetRawHeader(ptr.Pointer(), headerNameC, headerValueC)
	}
}

func (ptr *QHttpPart) Swap(other QHttpPart_ITF) {
	defer qt.Recovering("QHttpPart::swap")

	if ptr.Pointer() != nil {
		C.QHttpPart_Swap(ptr.Pointer(), PointerFromQHttpPart(other))
	}
}

func (ptr *QHttpPart) DestroyQHttpPart() {
	defer qt.Recovering("QHttpPart::~QHttpPart")

	if ptr.Pointer() != nil {
		C.QHttpPart_DestroyQHttpPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption = QLocalServer__SocketOption(0x7)
)

type QLocalServer struct {
	core.QObject
}

type QLocalServer_ITF interface {
	core.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func (p *QLocalServer) QLocalServer_PTR() *QLocalServer {
	return p
}

func (p *QLocalServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QLocalServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQLocalServer(ptr QLocalServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServer_PTR().Pointer()
	}
	return nil
}

func NewQLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = new(QLocalServer)
	n.SetPointer(ptr)
	return n
}
func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	defer qt.Recovering("QLocalServer::setSocketOptions")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(ptr.Pointer(), C.longlong(options))
	}
}

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {
	defer qt.Recovering("QLocalServer::QLocalServer")

	var tmpValue = NewQLocalServerFromPointer(C.QLocalServer_NewQLocalServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLocalServer) Close() {
	defer qt.Recovering("QLocalServer::close")

	if ptr.Pointer() != nil {
		C.QLocalServer_Close(ptr.Pointer())
	}
}

func (ptr *QLocalServer) ErrorString() string {
	defer qt.Recovering("QLocalServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) FullServerName() string {
	defer qt.Recovering("QLocalServer::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_FullServerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQLocalServer_HasPendingConnections
func callbackQLocalServer_HasPendingConnections(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalServer::hasPendingConnections")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QLocalServer) ConnectHasPendingConnections(f func() bool) {
	defer qt.Recovering("connect QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::hasPendingConnections", f)
	}
}

func (ptr *QLocalServer) DisconnectHasPendingConnections() {
	defer qt.Recovering("disconnect QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::hasPendingConnections")
	}
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	defer qt.Recovering("QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) HasPendingConnectionsDefault() bool {
	defer qt.Recovering("QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnectionsDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalServer_IncomingConnection
func callbackQLocalServer_IncomingConnection(ptr unsafe.Pointer, socketDescriptor C.uintptr_t) {
	defer qt.Recovering("callback QLocalServer::incomingConnection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::incomingConnection"); signal != nil {
		signal.(func(uintptr))(uintptr(socketDescriptor))
	} else {
		NewQLocalServerFromPointer(ptr).IncomingConnectionDefault(uintptr(socketDescriptor))
	}
}

func (ptr *QLocalServer) ConnectIncomingConnection(f func(socketDescriptor uintptr)) {
	defer qt.Recovering("connect QLocalServer::incomingConnection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::incomingConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectIncomingConnection() {
	defer qt.Recovering("disconnect QLocalServer::incomingConnection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::incomingConnection")
	}
}

func (ptr *QLocalServer) IncomingConnection(socketDescriptor uintptr) {
	defer qt.Recovering("QLocalServer::incomingConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_IncomingConnection(ptr.Pointer(), C.uintptr_t(socketDescriptor))
	}
}

func (ptr *QLocalServer) IncomingConnectionDefault(socketDescriptor uintptr) {
	defer qt.Recovering("QLocalServer::incomingConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_IncomingConnectionDefault(ptr.Pointer(), C.uintptr_t(socketDescriptor))
	}
}

func (ptr *QLocalServer) IsListening() bool {
	defer qt.Recovering("QLocalServer::isListening")

	if ptr.Pointer() != nil {
		return C.QLocalServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	defer qt.Recovering("QLocalServer::listen")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QLocalServer_Listen(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QLocalServer) MaxPendingConnections() int {
	defer qt.Recovering("QLocalServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(int32(C.QLocalServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQLocalServer_NewConnection
func callbackQLocalServer_NewConnection(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::newConnection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::newConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::newConnection")
	}
}

func (ptr *QLocalServer) NewConnection() {
	defer qt.Recovering("QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQLocalServer_NextPendingConnection
func callbackQLocalServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QLocalServer::nextPendingConnection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::nextPendingConnection"); signal != nil {
		return PointerFromQLocalSocket(signal.(func() *QLocalSocket)())
	}

	return PointerFromQLocalSocket(NewQLocalServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QLocalServer) ConnectNextPendingConnection(f func() *QLocalSocket) {
	defer qt.Recovering("connect QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::nextPendingConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNextPendingConnection() {
	defer qt.Recovering("disconnect QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::nextPendingConnection")
	}
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	defer qt.Recovering("QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLocalServer) NextPendingConnectionDefault() *QLocalSocket {
	defer qt.Recovering("QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	defer qt.Recovering("QLocalServer::removeServer")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QLocalServer_QLocalServer_RemoveServer(nameC) != 0
}

func (ptr *QLocalServer) RemoveServer(name string) bool {
	defer qt.Recovering("QLocalServer::removeServer")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QLocalServer_QLocalServer_RemoveServer(nameC) != 0
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QLocalServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QLocalServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ServerName() string {
	defer qt.Recovering("QLocalServer::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QLocalServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	defer qt.Recovering("QLocalServer::socketOptions")

	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QLocalServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QLocalServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), C.char(int8(qt.GoBoolToInt(timedOut)))) != 0
	}
	return false
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	defer qt.Recovering("QLocalServer::~QLocalServer")

	if ptr.Pointer() != nil {
		C.QLocalServer_DestroyQLocalServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalServer_TimerEvent
func callbackQLocalServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLocalServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::timerEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLocalServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::timerEvent")
	}
}

func (ptr *QLocalServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLocalServer_ChildEvent
func callbackQLocalServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLocalServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::childEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLocalServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::childEvent")
	}
}

func (ptr *QLocalServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalServer::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalServer::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalServer_ConnectNotify
func callbackQLocalServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLocalServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::connectNotify", f)
	}
}

func (ptr *QLocalServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QLocalServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::connectNotify")
	}
}

func (ptr *QLocalServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_CustomEvent
func callbackQLocalServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLocalServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::customEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLocalServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::customEvent")
	}
}

func (ptr *QLocalServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalServer::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalServer::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalServer_DeleteLater
func callbackQLocalServer_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QLocalServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::deleteLater", f)
	}
}

func (ptr *QLocalServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QLocalServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::deleteLater")
	}
}

func (ptr *QLocalServer) DeleteLater() {
	defer qt.Recovering("QLocalServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QLocalServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalServer) DeleteLaterDefault() {
	defer qt.Recovering("QLocalServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QLocalServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalServer_DisconnectNotify
func callbackQLocalServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLocalServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::disconnectNotify", f)
	}
}

func (ptr *QLocalServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QLocalServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::disconnectNotify")
	}
}

func (ptr *QLocalServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalServer_Event
func callbackQLocalServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalServer::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QLocalServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::event", f)
	}
}

func (ptr *QLocalServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QLocalServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::event")
	}
}

func (ptr *QLocalServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalServer::event")

	if ptr.Pointer() != nil {
		return C.QLocalServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLocalServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalServer::event")

	if ptr.Pointer() != nil {
		return C.QLocalServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLocalServer_EventFilter
func callbackQLocalServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalServer::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QLocalServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::eventFilter", f)
	}
}

func (ptr *QLocalServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QLocalServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::eventFilter")
	}
}

func (ptr *QLocalServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLocalServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLocalServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLocalServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLocalServer_MetaObject
func callbackQLocalServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QLocalServer::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLocalServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QLocalServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::metaObject", f)
	}
}

func (ptr *QLocalServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QLocalServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalServer::metaObject")
	}
}

func (ptr *QLocalServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QLocalServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLocalServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QLocalServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

type QLocalSocket struct {
	core.QIODevice
}

type QLocalSocket_ITF interface {
	core.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func (p *QLocalSocket) QLocalSocket_PTR() *QLocalSocket {
	return p
}

func (p *QLocalSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QLocalSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQLocalSocket(ptr QLocalSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocket_PTR().Pointer()
	}
	return nil
}

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = new(QLocalSocket)
	n.SetPointer(ptr)
	return n
}

//export callbackQLocalSocket_Open
func callbackQLocalSocket_Open(ptr unsafe.Pointer, openMode C.longlong) C.char {
	defer qt.Recovering("callback QLocalSocket::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(openMode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(openMode)))))
}

func (ptr *QLocalSocket) ConnectOpen(f func(openMode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QLocalSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::open", f)
	}
}

func (ptr *QLocalSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QLocalSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::open")
	}
}

func (ptr *QLocalSocket) Open(openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QLocalSocket::open")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.longlong(openMode)) != 0
	}
	return false
}

func (ptr *QLocalSocket) OpenDefault(openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QLocalSocket::open")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_OpenDefault(ptr.Pointer(), C.longlong(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	defer qt.Recovering("QLocalSocket::QLocalSocket")

	var tmpValue = NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), nameC, C.longlong(openMode))
	}
}

//export callbackQLocalSocket_Connected
func callbackQLocalSocket_Connected(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::connected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connected")
	}
}

func (ptr *QLocalSocket) Connected() {
	defer qt.Recovering("QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Connected(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Disconnected
func callbackQLocalSocket_Disconnected(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::disconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnected")
	}
}

func (ptr *QLocalSocket) Disconnected() {
	defer qt.Recovering("QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQLocalSocket_Error2
func callbackQLocalSocket_Error2(ptr unsafe.Pointer, socketError C.longlong) {
	defer qt.Recovering("callback QLocalSocket::error")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::error2"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketError))(QLocalSocket__LocalSocketError(socketError))
	}

}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {
	defer qt.Recovering("connect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::error2", f)
	}
}

func (ptr *QLocalSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::error2")
	}
}

func (ptr *QLocalSocket) Error2(socketError QLocalSocket__LocalSocketError) {
	defer qt.Recovering("QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Error2(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QLocalSocket) FullServerName() string {
	defer qt.Recovering("QLocalSocket::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQLocalSocket_IsSequential
func callbackQLocalSocket_IsSequential(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::isSequential")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QLocalSocket) ConnectIsSequential(f func() bool) {
	defer qt.Recovering("connect QLocalSocket::isSequential")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::isSequential", f)
	}
}

func (ptr *QLocalSocket) DisconnectIsSequential() {
	defer qt.Recovering("disconnect QLocalSocket::isSequential")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::isSequential")
	}
}

func (ptr *QLocalSocket) IsSequential() bool {
	defer qt.Recovering("QLocalSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsSequentialDefault() bool {
	defer qt.Recovering("QLocalSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	defer qt.Recovering("QLocalSocket::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	defer qt.Recovering("QLocalSocket::setServerName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLocalSocket_SetServerName(ptr.Pointer(), nameC)
	}
}

func (ptr *QLocalSocket) State() QLocalSocket__LocalSocketState {
	defer qt.Recovering("QLocalSocket::state")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketState(C.QLocalSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_StateChanged
func callbackQLocalSocket_StateChanged(ptr unsafe.Pointer, socketState C.longlong) {
	defer qt.Recovering("callback QLocalSocket::stateChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::stateChanged"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
	}

}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	defer qt.Recovering("connect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::stateChanged")
	}
}

func (ptr *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {
	defer qt.Recovering("QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_StateChanged(ptr.Pointer(), C.longlong(socketState))
	}
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	defer qt.Recovering("QLocalSocket::~QLocalSocket")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	defer qt.Recovering("QLocalSocket::abort")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

//export callbackQLocalSocket_BytesAvailable
func callbackQLocalSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QLocalSocket::bytesAvailable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QLocalSocket) ConnectBytesAvailable(f func() int64) {
	defer qt.Recovering("connect QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesAvailable", f)
	}
}

func (ptr *QLocalSocket) DisconnectBytesAvailable() {
	defer qt.Recovering("disconnect QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesAvailable")
	}
}

func (ptr *QLocalSocket) BytesAvailable() int64 {
	defer qt.Recovering("QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesAvailableDefault() int64 {
	defer qt.Recovering("QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_BytesToWrite
func callbackQLocalSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QLocalSocket::bytesToWrite")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QLocalSocket) ConnectBytesToWrite(f func() int64) {
	defer qt.Recovering("connect QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesToWrite", f)
	}
}

func (ptr *QLocalSocket) DisconnectBytesToWrite() {
	defer qt.Recovering("disconnect QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::bytesToWrite")
	}
}

func (ptr *QLocalSocket) BytesToWrite() int64 {
	defer qt.Recovering("QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesToWriteDefault() int64 {
	defer qt.Recovering("QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_CanReadLine
func callbackQLocalSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::canReadLine")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QLocalSocket) ConnectCanReadLine(f func() bool) {
	defer qt.Recovering("connect QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::canReadLine", f)
	}
}

func (ptr *QLocalSocket) DisconnectCanReadLine() {
	defer qt.Recovering("disconnect QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::canReadLine")
	}
}

func (ptr *QLocalSocket) CanReadLine() bool {
	defer qt.Recovering("QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) CanReadLineDefault() bool {
	defer qt.Recovering("QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Close
func callbackQLocalSocket_Close(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::close"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QLocalSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::close", f)
	}
}

func (ptr *QLocalSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::close")
	}
}

func (ptr *QLocalSocket) Close() {
	defer qt.Recovering("QLocalSocket::close")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) CloseDefault() {
	defer qt.Recovering("QLocalSocket::close")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.longlong(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	defer qt.Recovering("QLocalSocket::disconnectFromServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	defer qt.Recovering("QLocalSocket::error")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	defer qt.Recovering("QLocalSocket::flush")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	defer qt.Recovering("QLocalSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QLocalSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QLocalSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQLocalSocket_WriteData
func callbackQLocalSocket_WriteData(ptr unsafe.Pointer, data *C.char, c C.longlong) C.longlong {
	defer qt.Recovering("callback QLocalSocket::writeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(c)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).WriteDataDefault(C.GoString(data), int64(c)))
}

func (ptr *QLocalSocket) ConnectWriteData(f func(data string, c int64) int64) {
	defer qt.Recovering("connect QLocalSocket::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::writeData", f)
	}
}

func (ptr *QLocalSocket) DisconnectWriteData() {
	defer qt.Recovering("disconnect QLocalSocket::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::writeData")
	}
}

func (ptr *QLocalSocket) WriteData(data string, c int64) int64 {
	defer qt.Recovering("QLocalSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_WriteData(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

func (ptr *QLocalSocket) WriteDataDefault(data string, c int64) int64 {
	defer qt.Recovering("QLocalSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(c)))
	}
	return 0
}

//export callbackQLocalSocket_AtEnd
func callbackQLocalSocket_AtEnd(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::atEnd")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QLocalSocket) ConnectAtEnd(f func() bool) {
	defer qt.Recovering("connect QLocalSocket::atEnd")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::atEnd", f)
	}
}

func (ptr *QLocalSocket) DisconnectAtEnd() {
	defer qt.Recovering("disconnect QLocalSocket::atEnd")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::atEnd")
	}
}

func (ptr *QLocalSocket) AtEnd() bool {
	defer qt.Recovering("QLocalSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) AtEndDefault() bool {
	defer qt.Recovering("QLocalSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Pos
func callbackQLocalSocket_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QLocalSocket::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).PosDefault())
}

func (ptr *QLocalSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QLocalSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::pos", f)
	}
}

func (ptr *QLocalSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QLocalSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::pos")
	}
}

func (ptr *QLocalSocket) Pos() int64 {
	defer qt.Recovering("QLocalSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) PosDefault() int64 {
	defer qt.Recovering("QLocalSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_ReadLineData
func callbackQLocalSocket_ReadLineData(ptr unsafe.Pointer, data *C.char, maxSize C.longlong) C.longlong {
	defer qt.Recovering("callback QLocalSocket::readLineData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(maxSize)))
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).ReadLineDataDefault(C.GoString(data), int64(maxSize)))
}

func (ptr *QLocalSocket) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	defer qt.Recovering("connect QLocalSocket::readLineData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readLineData", f)
	}
}

func (ptr *QLocalSocket) DisconnectReadLineData() {
	defer qt.Recovering("disconnect QLocalSocket::readLineData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::readLineData")
	}
}

func (ptr *QLocalSocket) ReadLineData(data string, maxSize int64) int64 {
	defer qt.Recovering("QLocalSocket::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QLocalSocket) ReadLineDataDefault(data string, maxSize int64) int64 {
	defer qt.Recovering("QLocalSocket::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QLocalSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQLocalSocket_Reset
func callbackQLocalSocket_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QLocalSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QLocalSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::reset", f)
	}
}

func (ptr *QLocalSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QLocalSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::reset")
	}
}

func (ptr *QLocalSocket) Reset() bool {
	defer qt.Recovering("QLocalSocket::reset")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ResetDefault() bool {
	defer qt.Recovering("QLocalSocket::reset")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQLocalSocket_Seek
func callbackQLocalSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QLocalSocket::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QLocalSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QLocalSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::seek", f)
	}
}

func (ptr *QLocalSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QLocalSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::seek")
	}
}

func (ptr *QLocalSocket) Seek(pos int64) bool {
	defer qt.Recovering("QLocalSocket::seek")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QLocalSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QLocalSocket::seek")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQLocalSocket_Size
func callbackQLocalSocket_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QLocalSocket::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQLocalSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QLocalSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QLocalSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::size", f)
	}
}

func (ptr *QLocalSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QLocalSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::size")
	}
}

func (ptr *QLocalSocket) Size() int64 {
	defer qt.Recovering("QLocalSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) SizeDefault() int64 {
	defer qt.Recovering("QLocalSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQLocalSocket_TimerEvent
func callbackQLocalSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::timerEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::timerEvent")
	}
}

func (ptr *QLocalSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLocalSocket_ChildEvent
func callbackQLocalSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLocalSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::childEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLocalSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::childEvent")
	}
}

func (ptr *QLocalSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLocalSocket_ConnectNotify
func callbackQLocalSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLocalSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connectNotify", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QLocalSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::connectNotify")
	}
}

func (ptr *QLocalSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_CustomEvent
func callbackQLocalSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLocalSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::customEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLocalSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::customEvent")
	}
}

func (ptr *QLocalSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLocalSocket_DeleteLater
func callbackQLocalSocket_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLocalSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QLocalSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::deleteLater", f)
	}
}

func (ptr *QLocalSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QLocalSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::deleteLater")
	}
}

func (ptr *QLocalSocket) DeleteLater() {
	defer qt.Recovering("QLocalSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) DeleteLaterDefault() {
	defer qt.Recovering("QLocalSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQLocalSocket_DisconnectNotify
func callbackQLocalSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLocalSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLocalSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLocalSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnectNotify", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QLocalSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::disconnectNotify")
	}
}

func (ptr *QLocalSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLocalSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLocalSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLocalSocket_Event
func callbackQLocalSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLocalSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QLocalSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::event", f)
	}
}

func (ptr *QLocalSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QLocalSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::event")
	}
}

func (ptr *QLocalSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalSocket::event")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLocalSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalSocket::event")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLocalSocket_EventFilter
func callbackQLocalSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QLocalSocket::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLocalSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLocalSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QLocalSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::eventFilter", f)
	}
}

func (ptr *QLocalSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QLocalSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::eventFilter")
	}
}

func (ptr *QLocalSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLocalSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLocalSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLocalSocket_MetaObject
func callbackQLocalSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QLocalSocket::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QLocalSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLocalSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLocalSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QLocalSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::metaObject", f)
	}
}

func (ptr *QLocalSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QLocalSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QLocalSocket::metaObject")
	}
}

func (ptr *QLocalSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QLocalSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLocalSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QLocalSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLocalSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           = QNetworkAccessManager__NetworkAccessibility(1)
)

//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
	QNetworkAccessManager__HeadOperation    = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation = QNetworkAccessManager__Operation(0)
)

type QNetworkAccessManager struct {
	core.QObject
}

type QNetworkAccessManager_ITF interface {
	core.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
}

func (p *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager {
	return p
}

func (p *QNetworkAccessManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkAccessManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkAccessManager(ptr QNetworkAccessManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAccessManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) *QNetworkAccessManager {
	var n = new(QNetworkAccessManager)
	n.SetPointer(ptr)
	return n
}
func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	defer qt.Recovering("QNetworkAccessManager::proxyFactory")

	if ptr.Pointer() != nil {
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	defer qt.Recovering("QNetworkAccessManager::QNetworkAccessManager")

	var tmpValue = NewQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkAccessManager::activeConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_ActiveConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_AuthenticationRequired
func callbackQNetworkAccessManager_AuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::authenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::authenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::authenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	defer qt.Recovering("QNetworkAccessManager::cache")

	if ptr.Pointer() != nil {
		var tmpValue = NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	defer qt.Recovering("QNetworkAccessManager::clearAccessCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) Configuration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkAccessManager::configuration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkAccessManager_Configuration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {
	defer qt.Recovering("QNetworkAccessManager::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkAccessManager_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port))
	}
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkAccessManager::connectToHostEncrypted")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkAccessManager_ConnectToHostEncrypted(ptr.Pointer(), hostNameC, C.ushort(port), PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	defer qt.Recovering("QNetworkAccessManager::cookieJar")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_CreateRequest
func callbackQNetworkAccessManager_CreateRequest(ptr unsafe.Pointer, op C.longlong, req unsafe.Pointer, outgoingData unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkAccessManager::createRequest")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::createRequest"); signal != nil {
		return PointerFromQNetworkReply(signal.(func(QNetworkAccessManager__Operation, *QNetworkRequest, *core.QIODevice) *QNetworkReply)(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(req), core.NewQIODeviceFromPointer(outgoingData)))
	}

	return PointerFromQNetworkReply(NewQNetworkAccessManagerFromPointer(ptr).CreateRequestDefault(QNetworkAccessManager__Operation(op), NewQNetworkRequestFromPointer(req), core.NewQIODeviceFromPointer(outgoingData)))
}

func (ptr *QNetworkAccessManager) ConnectCreateRequest(f func(op QNetworkAccessManager__Operation, req *QNetworkRequest, outgoingData *core.QIODevice) *QNetworkReply) {
	defer qt.Recovering("connect QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::createRequest", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCreateRequest() {
	defer qt.Recovering("disconnect QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::createRequest")
	}
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequest(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) CreateRequestDefault(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequestDefault(ptr.Pointer(), C.longlong(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::deleteResource")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_Encrypted
func callbackQNetworkAccessManager_Encrypted(ptr unsafe.Pointer, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::encrypted")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::encrypted"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::encrypted")
	}
}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Encrypted(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

//export callbackQNetworkAccessManager_Finished
func callbackQNetworkAccessManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::finished")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::finished"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::finished")
	}
}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Finished(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::get")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::head")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	defer qt.Recovering("QNetworkAccessManager::networkAccessible")

	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkAccessManager_NetworkAccessibleChanged
func callbackQNetworkAccessManager_NetworkAccessibleChanged(ptr unsafe.Pointer, accessible C.longlong) {
	defer qt.Recovering("callback QNetworkAccessManager::networkAccessibleChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::networkAccessibleChanged"); signal != nil {
		signal.(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
	}

}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	defer qt.Recovering("connect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	defer qt.Recovering("disconnect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::networkAccessibleChanged")
	}
}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_NetworkAccessibleChanged(ptr.Pointer(), C.longlong(accessible))
	}
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data string) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), dataC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Proxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkAccessManager::proxy")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkAccessManager_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkAccessManager_ProxyAuthenticationRequired
func callbackQNetworkAccessManager_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::proxyAuthenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkProxy, *QAuthenticator))(NewQNetworkProxyFromPointer(proxy), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::proxyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectProxyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::proxyAuthenticationRequired")
	}
}

func (ptr *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ProxyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkProxy(proxy), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, data string) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), dataC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb string, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::sendCustomRequest")

	if ptr.Pointer() != nil {
		var verbC = C.CString(hex.EncodeToString([]byte(verb)))
		defer C.free(unsafe.Pointer(verbC))
		var tmpValue = NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), verbC, core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCache(ptr.Pointer(), PointerFromQAbstractNetworkCache(cache))
	}
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCookieJar(ptr.Pointer(), PointerFromQNetworkCookieJar(cookieJar))
	}
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::setNetworkAccessible")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetNetworkAccessible(ptr.Pointer(), C.longlong(accessible))
	}
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxy")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(proxy))
	}
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxyFactory(ptr.Pointer(), PointerFromQNetworkProxyFactory(factory))
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	defer qt.Recovering("QNetworkAccessManager::supportedSchemes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQNetworkAccessManager_SupportedSchemesImplementation
func callbackQNetworkAccessManager_SupportedSchemesImplementation(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QNetworkAccessManager::supportedSchemesImplementation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::supportedSchemesImplementation"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(make([]string, 0), "|"))
}

func (ptr *QNetworkAccessManager) ConnectSupportedSchemesImplementation(f func() []string) {
	defer qt.Recovering("connect QNetworkAccessManager::supportedSchemesImplementation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::supportedSchemesImplementation", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectSupportedSchemesImplementation() {
	defer qt.Recovering("disconnect QNetworkAccessManager::supportedSchemesImplementation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::supportedSchemesImplementation")
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementation() []string {
	defer qt.Recovering("QNetworkAccessManager::supportedSchemesImplementation")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemesImplementation(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	defer qt.Recovering("QNetworkAccessManager::~QNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkAccessManager_TimerEvent
func callbackQNetworkAccessManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::timerEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::timerEvent")
	}
}

func (ptr *QNetworkAccessManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkAccessManager_ChildEvent
func callbackQNetworkAccessManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::childEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::childEvent")
	}
}

func (ptr *QNetworkAccessManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkAccessManager_ConnectNotify
func callbackQNetworkAccessManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkAccessManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::connectNotify", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkAccessManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::connectNotify")
	}
}

func (ptr *QNetworkAccessManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkAccessManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkAccessManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_CustomEvent
func callbackQNetworkAccessManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::customEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::customEvent")
	}
}

func (ptr *QNetworkAccessManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkAccessManager_DeleteLater
func callbackQNetworkAccessManager_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkAccessManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkAccessManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::deleteLater", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkAccessManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::deleteLater")
	}
}

func (ptr *QNetworkAccessManager) DeleteLater() {
	defer qt.Recovering("QNetworkAccessManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkAccessManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkAccessManager_DisconnectNotify
func callbackQNetworkAccessManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkAccessManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkAccessManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::disconnectNotify", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkAccessManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::disconnectNotify")
	}
}

func (ptr *QNetworkAccessManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkAccessManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkAccessManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkAccessManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkAccessManager_Event
func callbackQNetworkAccessManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkAccessManager::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkAccessManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkAccessManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::event", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::event")
	}
}

func (ptr *QNetworkAccessManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkAccessManager::event")

	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkAccessManager::event")

	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_EventFilter
func callbackQNetworkAccessManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkAccessManager::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkAccessManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkAccessManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkAccessManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::eventFilter", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkAccessManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::eventFilter")
	}
}

func (ptr *QNetworkAccessManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkAccessManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkAccessManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkAccessManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkAccessManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkAccessManager_MetaObject
func callbackQNetworkAccessManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkAccessManager::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkAccessManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkAccessManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkAccessManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkAccessManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::metaObject", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkAccessManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkAccessManager::metaObject")
	}
}

func (ptr *QNetworkAccessManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkAccessManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkAccessManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkAccessManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkAccessManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkAddressEntry struct {
	ptr unsafe.Pointer
}

type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (p *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry {
	return p
}

func (p *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkAddressEntry) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntry_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	var tmpValue = NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry())
	runtime.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	var tmpValue = NewQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(PointerFromQNetworkAddressEntry(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkAddressEntry).DestroyQNetworkAddressEntry)
	return tmpValue
}

func (ptr *QNetworkAddressEntry) Broadcast() *QHostAddress {
	defer qt.Recovering("QNetworkAddressEntry::broadcast")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Broadcast(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) Ip() *QHostAddress {
	defer qt.Recovering("QNetworkAddressEntry::ip")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Ip(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) Netmask() *QHostAddress {
	defer qt.Recovering("QNetworkAddressEntry::netmask")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QNetworkAddressEntry_Netmask(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {
	defer qt.Recovering("QNetworkAddressEntry::prefixLength")

	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkAddressEntry_PrefixLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setBroadcast")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(ptr.Pointer(), PointerFromQHostAddress(newBroadcast))
	}
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setIp")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetIp(ptr.Pointer(), PointerFromQHostAddress(newIp))
	}
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setNetmask")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetNetmask(ptr.Pointer(), PointerFromQHostAddress(newNetmask))
	}
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {
	defer qt.Recovering("QNetworkAddressEntry::setPrefixLength")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetPrefixLength(ptr.Pointer(), C.int(int32(length)))
	}
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::swap")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_Swap(ptr.Pointer(), PointerFromQNetworkAddressEntry(other))
	}
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {
	defer qt.Recovering("QNetworkAddressEntry::~QNetworkAddressEntry")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkCacheMetaData struct {
	ptr unsafe.Pointer
}

type QNetworkCacheMetaData_ITF interface {
	QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData
}

func (p *QNetworkCacheMetaData) QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData {
	return p
}

func (p *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkCacheMetaData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkCacheMetaData::QNetworkCacheMetaData")

	var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
	runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkCacheMetaData::QNetworkCacheMetaData")

	var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
	return tmpValue
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	defer qt.Recovering("QNetworkCacheMetaData::expirationDate")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	defer qt.Recovering("QNetworkCacheMetaData::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	defer qt.Recovering("QNetworkCacheMetaData::lastModified")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	defer qt.Recovering("QNetworkCacheMetaData::saveToDisk")

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setExpirationDate")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setLastModified")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	defer qt.Recovering("QNetworkCacheMetaData::setSaveToDisk")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allow))))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::swap")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(ptr.Pointer(), PointerFromQNetworkCacheMetaData(other))
	}
}

func (ptr *QNetworkCacheMetaData) Url() *core.QUrl {
	defer qt.Recovering("QNetworkCacheMetaData::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkCacheMetaData_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	defer qt.Recovering("QNetworkCacheMetaData::~QNetworkCacheMetaData")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
	QNetworkConfiguration__BearerUnknown   = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        = QNetworkConfiguration__BearerType(12)
)

//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose = QNetworkConfiguration__Purpose(3)
)

//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     = QNetworkConfiguration__StateFlag(0x000000e)
)

//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             = QNetworkConfiguration__Type(3)
)

type QNetworkConfiguration struct {
	ptr unsafe.Pointer
}

type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (p *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration {
	return p
}

func (p *QNetworkConfiguration) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkConfiguration) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration())
	runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
	return tmpValue
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(PointerFromQNetworkConfiguration(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
	return tmpValue
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerType")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeFamily")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerTypeFamily(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_BearerTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Identifier() string {
	defer qt.Recovering("QNetworkConfiguration::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	defer qt.Recovering("QNetworkConfiguration::isRoamingAvailable")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsRoamingAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	defer qt.Recovering("QNetworkConfiguration::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) Name() string {
	defer qt.Recovering("QNetworkConfiguration::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {
	defer qt.Recovering("QNetworkConfiguration::purpose")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Purpose(C.QNetworkConfiguration_Purpose(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) State() QNetworkConfiguration__StateFlag {
	defer qt.Recovering("QNetworkConfiguration::state")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__StateFlag(C.QNetworkConfiguration_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfiguration::swap")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_Swap(ptr.Pointer(), PointerFromQNetworkConfiguration(other))
	}
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {
	defer qt.Recovering("QNetworkConfiguration::type")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Type(C.QNetworkConfiguration_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {
	defer qt.Recovering("QNetworkConfiguration::~QNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    = QNetworkConfigurationManager__Capability(0x00000040)
)

type QNetworkConfigurationManager struct {
	core.QObject
}

type QNetworkConfigurationManager_ITF interface {
	core.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func (p *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return p
}

func (p *QNetworkConfigurationManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkConfigurationManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkConfigurationManager(parent core.QObject_ITF) *QNetworkConfigurationManager {
	defer qt.Recovering("QNetworkConfigurationManager::QNetworkConfigurationManager")

	var tmpValue = NewQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	defer qt.Recovering("QNetworkConfigurationManager::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkConfigurationManager_ConfigurationAdded
func callbackQNetworkConfigurationManager_ConfigurationAdded(ptr unsafe.Pointer, config unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::configurationAdded")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationAdded"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationAdded(f func(config *QNetworkConfiguration)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::configurationAdded")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationAdded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationAdded", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationAdded() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::configurationAdded")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationAdded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationAdded")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationAdded(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::configurationAdded")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationAdded(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

//export callbackQNetworkConfigurationManager_ConfigurationChanged
func callbackQNetworkConfigurationManager_ConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::configurationChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationChanged"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationChanged(f func(config *QNetworkConfiguration)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::configurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationChanged() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::configurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationChanged")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationChanged(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::configurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier string) *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfigurationManager::configurationFromIdentifier")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_ConfigurationFromIdentifier(ptr.Pointer(), identifierC))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkConfigurationManager_ConfigurationRemoved
func callbackQNetworkConfigurationManager_ConfigurationRemoved(ptr unsafe.Pointer, config unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::configurationRemoved")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::configurationRemoved"); signal != nil {
		signal.(func(*QNetworkConfiguration))(NewQNetworkConfigurationFromPointer(config))
	}

}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationRemoved(f func(config *QNetworkConfiguration)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::configurationRemoved")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectConfigurationRemoved(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationRemoved", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationRemoved() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::configurationRemoved")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectConfigurationRemoved(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::configurationRemoved")
	}
}

func (ptr *QNetworkConfigurationManager) ConfigurationRemoved(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::configurationRemoved")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConfigurationRemoved(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfigurationManager::defaultConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkConfigurationManager_DefaultConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	defer qt.Recovering("QNetworkConfigurationManager::isOnline")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_OnlineStateChanged
func callbackQNetworkConfigurationManager_OnlineStateChanged(ptr unsafe.Pointer, isOnline C.char) {
	defer qt.Recovering("callback QNetworkConfigurationManager::onlineStateChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::onlineStateChanged"); signal != nil {
		signal.(func(bool))(int8(isOnline) != 0)
	}

}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::onlineStateChanged")
	}
}

func (ptr *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	defer qt.Recovering("QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_OnlineStateChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isOnline))))
	}
}

//export callbackQNetworkConfigurationManager_UpdateCompleted
func callbackQNetworkConfigurationManager_UpdateCompleted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::updateCompleted")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::updateCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateCompleted")
	}
}

func (ptr *QNetworkConfigurationManager) UpdateCompleted() {
	defer qt.Recovering("QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateCompleted(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_UpdateConfigurations
func callbackQNetworkConfigurationManager_UpdateConfigurations(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::updateConfigurations")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::updateConfigurations"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkConfigurationManager) ConnectUpdateConfigurations(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::updateConfigurations")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateConfigurations", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateConfigurations() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::updateConfigurations")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::updateConfigurations")
	}
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	defer qt.Recovering("QNetworkConfigurationManager::updateConfigurations")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(ptr.Pointer())
	}
}

//export callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager
func callbackQNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::~QNetworkConfigurationManager"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DestroyQNetworkConfigurationManagerDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDestroyQNetworkConfigurationManager(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::~QNetworkConfigurationManager", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDestroyQNetworkConfigurationManager() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::~QNetworkConfigurationManager")
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	defer qt.Recovering("QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManagerDefault() {
	defer qt.Recovering("QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManagerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkConfigurationManager_TimerEvent
func callbackQNetworkConfigurationManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::timerEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::timerEvent")
	}
}

func (ptr *QNetworkConfigurationManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_ChildEvent
func callbackQNetworkConfigurationManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::childEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::childEvent")
	}
}

func (ptr *QNetworkConfigurationManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_ConnectNotify
func callbackQNetworkConfigurationManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::connectNotify", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::connectNotify")
	}
}

func (ptr *QNetworkConfigurationManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_CustomEvent
func callbackQNetworkConfigurationManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::customEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::customEvent")
	}
}

func (ptr *QNetworkConfigurationManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkConfigurationManager_DeleteLater
func callbackQNetworkConfigurationManager_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::deleteLater", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::deleteLater")
	}
}

func (ptr *QNetworkConfigurationManager) DeleteLater() {
	defer qt.Recovering("QNetworkConfigurationManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkConfigurationManager::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkConfigurationManager_DisconnectNotify
func callbackQNetworkConfigurationManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::disconnectNotify", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::disconnectNotify")
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkConfigurationManager_Event
func callbackQNetworkConfigurationManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkConfigurationManager::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkConfigurationManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkConfigurationManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::event", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::event")
	}
}

func (ptr *QNetworkConfigurationManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkConfigurationManager::event")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkConfigurationManager::event")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_EventFilter
func callbackQNetworkConfigurationManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkConfigurationManager::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkConfigurationManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkConfigurationManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkConfigurationManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::eventFilter", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::eventFilter")
	}
}

func (ptr *QNetworkConfigurationManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkConfigurationManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkConfigurationManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkConfigurationManager_MetaObject
func callbackQNetworkConfigurationManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkConfigurationManager::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkConfigurationManager::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkConfigurationManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkConfigurationManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkConfigurationManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::metaObject", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkConfigurationManager::metaObject")
	}
}

func (ptr *QNetworkConfigurationManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkConfigurationManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkConfigurationManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkConfigurationManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkConfigurationManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             = QNetworkCookie__RawForm(1)
)

type QNetworkCookie struct {
	ptr unsafe.Pointer
}

type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (p *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie {
	return p
}

func (p *QNetworkCookie) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkCookie) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkCookie(ptr QNetworkCookie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookie_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkCookie(name string, value string) *QNetworkCookie {
	defer qt.Recovering("QNetworkCookie::QNetworkCookie")

	var nameC = C.CString(hex.EncodeToString([]byte(name)))
	defer C.free(unsafe.Pointer(nameC))
	var valueC = C.CString(hex.EncodeToString([]byte(value)))
	defer C.free(unsafe.Pointer(valueC))
	var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie(nameC, valueC))
	runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
	return tmpValue
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {
	defer qt.Recovering("QNetworkCookie::QNetworkCookie")

	var tmpValue = NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie2(PointerFromQNetworkCookie(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkCookie).DestroyQNetworkCookie)
	return tmpValue
}

func (ptr *QNetworkCookie) Domain() string {
	defer qt.Recovering("QNetworkCookie::domain")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Domain(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) ExpirationDate() *core.QDateTime {
	defer qt.Recovering("QNetworkCookie::expirationDate")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QNetworkCookie_ExpirationDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookie::hasSameIdentifier")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_HasSameIdentifier(ptr.Pointer(), PointerFromQNetworkCookie(other)) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	defer qt.Recovering("QNetworkCookie::isHttpOnly")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsHttpOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	defer qt.Recovering("QNetworkCookie::isSecure")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSecure(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	defer qt.Recovering("QNetworkCookie::isSessionCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSessionCookie(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Name() string {
	defer qt.Recovering("QNetworkCookie::name")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNetworkCookie_Name(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkCookie) Normalize(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkCookie::normalize")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_Normalize(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCookie) Path() string {
	defer qt.Recovering("QNetworkCookie::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) SetDomain(domain string) {
	defer qt.Recovering("QNetworkCookie::setDomain")

	if ptr.Pointer() != nil {
		var domainC = C.CString(domain)
		defer C.free(unsafe.Pointer(domainC))
		C.QNetworkCookie_SetDomain(ptr.Pointer(), domainC)
	}
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCookie::setExpirationDate")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(date))
	}
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {
	defer qt.Recovering("QNetworkCookie::setHttpOnly")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetHttpOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QNetworkCookie) SetName(cookieName string) {
	defer qt.Recovering("QNetworkCookie::setName")

	if ptr.Pointer() != nil {
		var cookieNameC = C.CString(hex.EncodeToString([]byte(cookieName)))
		defer C.free(unsafe.Pointer(cookieNameC))
		C.QNetworkCookie_SetName(ptr.Pointer(), cookieNameC)
	}
}

func (ptr *QNetworkCookie) SetPath(path string) {
	defer qt.Recovering("QNetworkCookie::setPath")

	if ptr.Pointer() != nil {
		var pathC = C.CString(path)
		defer C.free(unsafe.Pointer(pathC))
		C.QNetworkCookie_SetPath(ptr.Pointer(), pathC)
	}
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {
	defer qt.Recovering("QNetworkCookie::setSecure")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetSecure(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(enable))))
	}
}

func (ptr *QNetworkCookie) SetValue(value string) {
	defer qt.Recovering("QNetworkCookie::setValue")

	if ptr.Pointer() != nil {
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QNetworkCookie_SetValue(ptr.Pointer(), valueC)
	}
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookie_ITF) {
	defer qt.Recovering("QNetworkCookie::swap")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_Swap(ptr.Pointer(), PointerFromQNetworkCookie(other))
	}
}

func (ptr *QNetworkCookie) ToRawForm(form QNetworkCookie__RawForm) string {
	defer qt.Recovering("QNetworkCookie::toRawForm")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNetworkCookie_ToRawForm(ptr.Pointer(), C.longlong(form))))
	}
	return ""
}

func (ptr *QNetworkCookie) Value() string {
	defer qt.Recovering("QNetworkCookie::value")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QNetworkCookie_Value(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	defer qt.Recovering("QNetworkCookie::~QNetworkCookie")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_DestroyQNetworkCookie(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func (p *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return p
}

func (p *QNetworkCookieJar) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkCookieJar) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	defer qt.Recovering("QNetworkCookieJar::QNetworkCookieJar")

	var tmpValue = NewQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkCookieJar_DeleteCookie
func callbackQNetworkCookieJar_DeleteCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::deleteCookie")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::deleteCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).DeleteCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectDeleteCookie(f func(cookie *QNetworkCookie) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDeleteCookie() {
	defer qt.Recovering("disconnect QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteCookie")
	}
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DeleteCookieDefault(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_InsertCookie
func callbackQNetworkCookieJar_InsertCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::insertCookie")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::insertCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).InsertCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectInsertCookie(f func(cookie *QNetworkCookie) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::insertCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectInsertCookie() {
	defer qt.Recovering("disconnect QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::insertCookie")
	}
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookieDefault(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_UpdateCookie
func callbackQNetworkCookieJar_UpdateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::updateCookie")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::updateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie) bool)(NewQNetworkCookieFromPointer(cookie)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).UpdateCookieDefault(NewQNetworkCookieFromPointer(cookie)))))
}

func (ptr *QNetworkCookieJar) ConnectUpdateCookie(f func(cookie *QNetworkCookie) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::updateCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectUpdateCookie() {
	defer qt.Recovering("disconnect QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::updateCookie")
	}
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookieDefault(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_ValidateCookie
func callbackQNetworkCookieJar_ValidateCookie(ptr unsafe.Pointer, cookie unsafe.Pointer, url unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::validateCookie")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::validateCookie"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QNetworkCookie, *core.QUrl) bool)(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).ValidateCookieDefault(NewQNetworkCookieFromPointer(cookie), core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkCookieJar) ConnectValidateCookie(f func(cookie *QNetworkCookie, url *core.QUrl) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::validateCookie", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectValidateCookie() {
	defer qt.Recovering("disconnect QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::validateCookie")
	}
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) ValidateCookieDefault(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookieDefault(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_DestroyQNetworkCookieJar
func callbackQNetworkCookieJar_DestroyQNetworkCookieJar(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::~QNetworkCookieJar")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::~QNetworkCookieJar"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DestroyQNetworkCookieJarDefault()
	}
}

func (ptr *QNetworkCookieJar) ConnectDestroyQNetworkCookieJar(f func()) {
	defer qt.Recovering("connect QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::~QNetworkCookieJar", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDestroyQNetworkCookieJar() {
	defer qt.Recovering("disconnect QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::~QNetworkCookieJar")
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	defer qt.Recovering("QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJarDefault() {
	defer qt.Recovering("QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJarDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkCookieJar_TimerEvent
func callbackQNetworkCookieJar_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::timerEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::timerEvent")
	}
}

func (ptr *QNetworkCookieJar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkCookieJar_ChildEvent
func callbackQNetworkCookieJar_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::childEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::childEvent")
	}
}

func (ptr *QNetworkCookieJar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkCookieJar_ConnectNotify
func callbackQNetworkCookieJar_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkCookieJar::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::connectNotify", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkCookieJar::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::connectNotify")
	}
}

func (ptr *QNetworkCookieJar) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkCookieJar::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkCookieJar::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_CustomEvent
func callbackQNetworkCookieJar_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::customEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::customEvent")
	}
}

func (ptr *QNetworkCookieJar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkCookieJar_DeleteLater
func callbackQNetworkCookieJar_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkCookieJar) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkCookieJar::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteLater", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkCookieJar::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::deleteLater")
	}
}

func (ptr *QNetworkCookieJar) DeleteLater() {
	defer qt.Recovering("QNetworkCookieJar::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkCookieJar::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkCookieJar_DisconnectNotify
func callbackQNetworkCookieJar_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkCookieJar) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkCookieJar::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::disconnectNotify", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkCookieJar::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::disconnectNotify")
	}
}

func (ptr *QNetworkCookieJar) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkCookieJar::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkCookieJar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkCookieJar::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkCookieJar_Event
func callbackQNetworkCookieJar_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkCookieJar) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::event", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::event")
	}
}

func (ptr *QNetworkCookieJar) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::event")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::event")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_EventFilter
func callbackQNetworkCookieJar_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkCookieJar::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkCookieJarFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkCookieJar) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkCookieJar::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::eventFilter", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkCookieJar::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::eventFilter")
	}
}

func (ptr *QNetworkCookieJar) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkCookieJar_MetaObject
func callbackQNetworkCookieJar_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkCookieJar::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkCookieJar::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkCookieJarFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkCookieJar) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkCookieJar::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::metaObject", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkCookieJar::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkCookieJar::metaObject")
	}
}

func (ptr *QNetworkCookieJar) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkCookieJar::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkCookieJar_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookieJar) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkCookieJar::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkCookieJar_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QNetworkDiskCache struct {
	QAbstractNetworkCache
}

type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func (p *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return p
}

func (p *QNetworkDiskCache) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkDiskCache) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractNetworkCache_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkDiskCache(ptr QNetworkDiskCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDiskCache_PTR().Pointer()
	}
	return nil
}

func NewQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) *QNetworkDiskCache {
	var n = new(QNetworkDiskCache)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	defer qt.Recovering("QNetworkDiskCache::QNetworkDiskCache")

	var tmpValue = NewQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	defer qt.Recovering("QNetworkDiskCache::cacheDirectory")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

//export callbackQNetworkDiskCache_CacheSize
func callbackQNetworkDiskCache_CacheSize(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkDiskCache::cacheSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::cacheSize"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).CacheSizeDefault())
}

func (ptr *QNetworkDiskCache) ConnectCacheSize(f func() int64) {
	defer qt.Recovering("connect QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::cacheSize", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCacheSize() {
	defer qt.Recovering("disconnect QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::cacheSize")
	}
}

func (ptr *QNetworkDiskCache) CacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) CacheSizeDefault() int64 {
	defer qt.Recovering("QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkDiskCache_Clear
func callbackQNetworkDiskCache_Clear(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::clear")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::clear"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {
	defer qt.Recovering("connect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::clear", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectClear() {
	defer qt.Recovering("disconnect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::clear")
	}
}

func (ptr *QNetworkDiskCache) Clear() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) ClearDefault() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ClearDefault(ptr.Pointer())
	}
}

//export callbackQNetworkDiskCache_Data
func callbackQNetworkDiskCache_Data(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkDiskCache::data")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::data"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*core.QUrl) *core.QIODevice)(core.NewQUrlFromPointer(url)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).DataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {
	defer qt.Recovering("connect QNetworkDiskCache::data")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::data", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectData() {
	defer qt.Recovering("disconnect QNetworkDiskCache::data")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::data")
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) DataDefault(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_DataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Expire
func callbackQNetworkDiskCache_Expire(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkDiskCache::expire")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::expire"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkDiskCacheFromPointer(ptr).ExpireDefault())
}

func (ptr *QNetworkDiskCache) ConnectExpire(f func() int64) {
	defer qt.Recovering("connect QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::expire", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectExpire() {
	defer qt.Recovering("disconnect QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::expire")
	}
}

func (ptr *QNetworkDiskCache) Expire() int64 {
	defer qt.Recovering("QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_Expire(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) ExpireDefault() int64 {
	defer qt.Recovering("QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_ExpireDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) FileMetaData(fileName string) *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkDiskCache::fileMetaData")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_FileMetaData(ptr.Pointer(), fileNameC))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Insert
func callbackQNetworkDiskCache_Insert(ptr unsafe.Pointer, device unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::insert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::insert"); signal != nil {
		signal.(func(*core.QIODevice))(core.NewQIODeviceFromPointer(device))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).InsertDefault(core.NewQIODeviceFromPointer(device))
	}
}

func (ptr *QNetworkDiskCache) ConnectInsert(f func(device *core.QIODevice)) {
	defer qt.Recovering("connect QNetworkDiskCache::insert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::insert", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectInsert() {
	defer qt.Recovering("disconnect QNetworkDiskCache::insert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::insert")
	}
}

func (ptr *QNetworkDiskCache) Insert(device core.QIODevice_ITF) {
	defer qt.Recovering("QNetworkDiskCache::insert")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Insert(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNetworkDiskCache) InsertDefault(device core.QIODevice_ITF) {
	defer qt.Recovering("QNetworkDiskCache::insert")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_InsertDefault(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QNetworkDiskCache) MaximumCacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::maximumCacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_MaximumCacheSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkDiskCache_MetaData
func callbackQNetworkDiskCache_MetaData(ptr unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkDiskCache::metaData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::metaData"); signal != nil {
		return PointerFromQNetworkCacheMetaData(signal.(func(*core.QUrl) *QNetworkCacheMetaData)(core.NewQUrlFromPointer(url)))
	}

	return PointerFromQNetworkCacheMetaData(NewQNetworkDiskCacheFromPointer(ptr).MetaDataDefault(core.NewQUrlFromPointer(url)))
}

func (ptr *QNetworkDiskCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {
	defer qt.Recovering("connect QNetworkDiskCache::metaData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaData", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectMetaData() {
	defer qt.Recovering("disconnect QNetworkDiskCache::metaData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaData")
	}
}

func (ptr *QNetworkDiskCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkDiskCache::metaData")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaData(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) MetaDataDefault(url core.QUrl_ITF) *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkDiskCache::metaData")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkCacheMetaDataFromPointer(C.QNetworkDiskCache_MetaDataDefault(ptr.Pointer(), core.PointerFromQUrl(url)))
		runtime.SetFinalizer(tmpValue, (*QNetworkCacheMetaData).DestroyQNetworkCacheMetaData)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Prepare
func callbackQNetworkDiskCache_Prepare(ptr unsafe.Pointer, metaData unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkDiskCache::prepare")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::prepare"); signal != nil {
		return core.PointerFromQIODevice(signal.(func(*QNetworkCacheMetaData) *core.QIODevice)(NewQNetworkCacheMetaDataFromPointer(metaData)))
	}

	return core.PointerFromQIODevice(NewQNetworkDiskCacheFromPointer(ptr).PrepareDefault(NewQNetworkCacheMetaDataFromPointer(metaData)))
}

func (ptr *QNetworkDiskCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {
	defer qt.Recovering("connect QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::prepare", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectPrepare() {
	defer qt.Recovering("disconnect QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::prepare")
	}
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkDiskCache) PrepareDefault(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QNetworkDiskCache_PrepareDefault(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkDiskCache_Remove
func callbackQNetworkDiskCache_Remove(ptr unsafe.Pointer, url unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkDiskCache::remove")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::remove"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QUrl) bool)(core.NewQUrlFromPointer(url)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).RemoveDefault(core.NewQUrlFromPointer(url)))))
}

func (ptr *QNetworkDiskCache) ConnectRemove(f func(url *core.QUrl) bool) {
	defer qt.Recovering("connect QNetworkDiskCache::remove")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::remove", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectRemove() {
	defer qt.Recovering("disconnect QNetworkDiskCache::remove")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::remove")
	}
}

func (ptr *QNetworkDiskCache) Remove(url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::remove")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_Remove(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) RemoveDefault(url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::remove")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_RemoveDefault(ptr.Pointer(), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	defer qt.Recovering("QNetworkDiskCache::setCacheDirectory")

	if ptr.Pointer() != nil {
		var cacheDirC = C.CString(cacheDir)
		defer C.free(unsafe.Pointer(cacheDirC))
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), cacheDirC)
	}
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	defer qt.Recovering("QNetworkDiskCache::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetMaximumCacheSize(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQNetworkDiskCache_UpdateMetaData
func callbackQNetworkDiskCache_UpdateMetaData(ptr unsafe.Pointer, metaData unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::updateMetaData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::updateMetaData"); signal != nil {
		signal.(func(*QNetworkCacheMetaData))(NewQNetworkCacheMetaDataFromPointer(metaData))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).UpdateMetaDataDefault(NewQNetworkCacheMetaDataFromPointer(metaData))
	}
}

func (ptr *QNetworkDiskCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {
	defer qt.Recovering("connect QNetworkDiskCache::updateMetaData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::updateMetaData", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectUpdateMetaData() {
	defer qt.Recovering("disconnect QNetworkDiskCache::updateMetaData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::updateMetaData")
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QNetworkDiskCache::updateMetaData")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) UpdateMetaDataDefault(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QNetworkDiskCache::updateMetaData")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_UpdateMetaDataDefault(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	defer qt.Recovering("QNetworkDiskCache::~QNetworkDiskCache")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkDiskCache_TimerEvent
func callbackQNetworkDiskCache_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::timerEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::timerEvent")
	}
}

func (ptr *QNetworkDiskCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkDiskCache_ChildEvent
func callbackQNetworkDiskCache_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::childEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::childEvent")
	}
}

func (ptr *QNetworkDiskCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkDiskCache_ConnectNotify
func callbackQNetworkDiskCache_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkDiskCache::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::connectNotify", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkDiskCache::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::connectNotify")
	}
}

func (ptr *QNetworkDiskCache) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkDiskCache::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkDiskCache::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkDiskCache_CustomEvent
func callbackQNetworkDiskCache_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::customEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::customEvent")
	}
}

func (ptr *QNetworkDiskCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkDiskCache_DeleteLater
func callbackQNetworkDiskCache_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkDiskCache) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkDiskCache::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::deleteLater", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkDiskCache::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::deleteLater")
	}
}

func (ptr *QNetworkDiskCache) DeleteLater() {
	defer qt.Recovering("QNetworkDiskCache::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkDiskCache) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkDiskCache::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkDiskCache_DisconnectNotify
func callbackQNetworkDiskCache_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkDiskCache) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkDiskCache::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::disconnectNotify", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkDiskCache::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::disconnectNotify")
	}
}

func (ptr *QNetworkDiskCache) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkDiskCache::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkDiskCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkDiskCache::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkDiskCache_Event
func callbackQNetworkDiskCache_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkDiskCache::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkDiskCache) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkDiskCache::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::event", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::event")
	}
}

func (ptr *QNetworkDiskCache) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::event")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::event")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkDiskCache_EventFilter
func callbackQNetworkDiskCache_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkDiskCache::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkDiskCacheFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkDiskCache) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkDiskCache::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::eventFilter", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkDiskCache::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::eventFilter")
	}
}

func (ptr *QNetworkDiskCache) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkDiskCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkDiskCache::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkDiskCache_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkDiskCache_MetaObject
func callbackQNetworkDiskCache_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkDiskCache::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkDiskCache::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkDiskCacheFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkDiskCache) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkDiskCache::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaObject", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkDiskCache::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkDiskCache::metaObject")
	}
}

func (ptr *QNetworkDiskCache) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkDiskCache::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkDiskCache_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkDiskCache) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkDiskCache::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkDiskCache_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   = QNetworkInterface__InterfaceFlag(0x20)
)

type QNetworkInterface struct {
	ptr unsafe.Pointer
}

type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (p *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface {
	return p
}

func (p *QNetworkInterface) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkInterface) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkInterface(ptr QNetworkInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterface_PTR().Pointer()
	}
	return nil
}

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkInterface() *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::QNetworkInterface")

	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::QNetworkInterface")

	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	defer qt.Recovering("QNetworkInterface::flags")

	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	defer qt.Recovering("QNetworkInterface::hardwareAddress")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HardwareAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	defer qt.Recovering("QNetworkInterface::humanReadableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HumanReadableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	defer qt.Recovering("QNetworkInterface::index")

	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkInterface_Index(ptr.Pointer())))
	}
	return 0
}

func QNetworkInterface_InterfaceFromIndex(index int) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::interfaceFromIndex")

	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::interfaceFromIndex")

	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromIndex(C.int(int32(index))))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::interfaceFromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func (ptr *QNetworkInterface) InterfaceFromName(name string) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::interfaceFromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkInterface_QNetworkInterface_InterfaceFromName(nameC))
	runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
	return tmpValue
}

func QNetworkInterface_InterfaceIndexFromName(name string) int {
	defer qt.Recovering("QNetworkInterface::interfaceIndexFromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(nameC)))
}

func (ptr *QNetworkInterface) InterfaceIndexFromName(name string) int {
	defer qt.Recovering("QNetworkInterface::interfaceIndexFromName")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return int(int32(C.QNetworkInterface_QNetworkInterface_InterfaceIndexFromName(nameC)))
}

func QNetworkInterface_InterfaceNameFromIndex(index int) string {
	defer qt.Recovering("QNetworkInterface::interfaceNameFromIndex")

	return C.GoString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) InterfaceNameFromIndex(index int) string {
	defer qt.Recovering("QNetworkInterface::interfaceNameFromIndex")

	return C.GoString(C.QNetworkInterface_QNetworkInterface_InterfaceNameFromIndex(C.int(int32(index))))
}

func (ptr *QNetworkInterface) IsValid() bool {
	defer qt.Recovering("QNetworkInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	defer qt.Recovering("QNetworkInterface::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	defer qt.Recovering("QNetworkInterface::swap")

	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(ptr.Pointer(), PointerFromQNetworkInterface(other))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	defer qt.Recovering("QNetworkInterface::~QNetworkInterface")

	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability = QNetworkProxy__Capability(0x0010)
)

//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  = QNetworkProxy__ProxyType(5)
)

type QNetworkProxy struct {
	ptr unsafe.Pointer
}

type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (p *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy {
	return p
}

func (p *QNetworkProxy) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkProxy) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkProxy(ptr QNetworkProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxy_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkProxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy2(ty QNetworkProxy__ProxyType, hostName string, port uint16, user string, password string) *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	var hostNameC = C.CString(hostName)
	defer C.free(unsafe.Pointer(hostNameC))
	var userC = C.CString(user)
	defer C.free(unsafe.Pointer(userC))
	var passwordC = C.CString(password)
	defer C.free(unsafe.Pointer(passwordC))
	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy2(C.longlong(ty), hostNameC, C.ushort(port), userC, passwordC))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy3(PointerFromQNetworkProxy(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func QNetworkProxy_ApplicationProxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::applicationProxy")

	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func (ptr *QNetworkProxy) ApplicationProxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::applicationProxy")

	var tmpValue = NewQNetworkProxyFromPointer(C.QNetworkProxy_QNetworkProxy_ApplicationProxy())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
	return tmpValue
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {
	defer qt.Recovering("QNetworkProxy::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkProxy__Capability(C.QNetworkProxy_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) HasRawHeader(headerName string) bool {
	defer qt.Recovering("QNetworkProxy::hasRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return C.QNetworkProxy_HasRawHeader(ptr.Pointer(), headerNameC) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkProxy::header")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkProxy_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxy) HostName() string {
	defer qt.Recovering("QNetworkProxy::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_HostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {
	defer qt.Recovering("QNetworkProxy::isCachingProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsCachingProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	defer qt.Recovering("QNetworkProxy::isTransparentProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsTransparentProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Password() string {
	defer qt.Recovering("QNetworkProxy::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) Port() uint16 {
	defer qt.Recovering("QNetworkProxy::port")

	if ptr.Pointer() != nil {
		return uint16(C.QNetworkProxy_Port(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) RawHeader(headerName string) string {
	defer qt.Recovering("QNetworkProxy::rawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return qt.HexDecodeToString(C.GoString(C.QNetworkProxy_RawHeader(ptr.Pointer(), headerNameC)))
	}
	return ""
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::setApplicationProxy")

	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::setApplicationProxy")

	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {
	defer qt.Recovering("QNetworkProxy::setCapabilities")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetCapabilities(ptr.Pointer(), C.longlong(capabilities))
	}
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkProxy::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {
	defer qt.Recovering("QNetworkProxy::setHostName")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QNetworkProxy_SetHostName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	defer qt.Recovering("QNetworkProxy::setPassword")

	if ptr.Pointer() != nil {
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		C.QNetworkProxy_SetPassword(ptr.Pointer(), passwordC)
	}
}

func (ptr *QNetworkProxy) SetPort(port uint16) {
	defer qt.Recovering("QNetworkProxy::setPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetPort(ptr.Pointer(), C.ushort(port))
	}
}

func (ptr *QNetworkProxy) SetRawHeader(headerName string, headerValue string) {
	defer qt.Recovering("QNetworkProxy::setRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		var headerValueC = C.CString(hex.EncodeToString([]byte(headerValue)))
		defer C.free(unsafe.Pointer(headerValueC))
		C.QNetworkProxy_SetRawHeader(ptr.Pointer(), headerNameC, headerValueC)
	}
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {
	defer qt.Recovering("QNetworkProxy::setType")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QNetworkProxy) SetUser(user string) {
	defer qt.Recovering("QNetworkProxy::setUser")

	if ptr.Pointer() != nil {
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		C.QNetworkProxy_SetUser(ptr.Pointer(), userC)
	}
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_Swap(ptr.Pointer(), PointerFromQNetworkProxy(other))
	}
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	defer qt.Recovering("QNetworkProxy::type")

	if ptr.Pointer() != nil {
		return QNetworkProxy__ProxyType(C.QNetworkProxy_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) User() string {
	defer qt.Recovering("QNetworkProxy::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {
	defer qt.Recovering("QNetworkProxy::~QNetworkProxy")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_DestroyQNetworkProxy(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QNetworkProxyFactory struct {
	ptr unsafe.Pointer
}

type QNetworkProxyFactory_ITF interface {
	QNetworkProxyFactory_PTR() *QNetworkProxyFactory
}

func (p *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return p
}

func (p *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkProxyFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactory_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	return n
}
func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkProxyFactory::setApplicationProxyFactory")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func (ptr *QNetworkProxyFactory) SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkProxyFactory::setApplicationProxyFactory")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	defer qt.Recovering("QNetworkProxyFactory::setUseSystemConfiguration")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.char(int8(qt.GoBoolToInt(enable))))
}

func (ptr *QNetworkProxyFactory) SetUseSystemConfiguration(enable bool) {
	defer qt.Recovering("QNetworkProxyFactory::setUseSystemConfiguration")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.char(int8(qt.GoBoolToInt(enable))))
}

//export callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory
func callbackQNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkProxyFactory::~QNetworkProxyFactory")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkProxyFactory::~QNetworkProxyFactory"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkProxyFactoryFromPointer(ptr).DestroyQNetworkProxyFactoryDefault()
	}
}

func (ptr *QNetworkProxyFactory) ConnectDestroyQNetworkProxyFactory(f func()) {
	defer qt.Recovering("connect QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::~QNetworkProxyFactory", f)
	}
}

func (ptr *QNetworkProxyFactory) DisconnectDestroyQNetworkProxyFactory() {
	defer qt.Recovering("disconnect QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkProxyFactory::~QNetworkProxyFactory")
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	defer qt.Recovering("QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactoryDefault() {
	defer qt.Recovering("QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactoryDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__TcpServer  = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest = QNetworkProxyQuery__QueryType(101)
)

type QNetworkProxyQuery struct {
	ptr unsafe.Pointer
}

type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (p *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery {
	return p
}

func (p *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkProxyQuery) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQuery_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery())
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery6(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var hostnameC = C.CString(hostname)
	defer C.free(unsafe.Pointer(hostnameC))
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery6(PointerFromQNetworkConfiguration(networkConfiguration), hostnameC, C.int(int32(port)), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery5(networkConfiguration QNetworkConfiguration_ITF, requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery5(PointerFromQNetworkConfiguration(networkConfiguration), core.PointerFromQUrl(requestUrl), C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery7(networkConfiguration QNetworkConfiguration_ITF, bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery7(PointerFromQNetworkConfiguration(networkConfiguration), C.ushort(bindPort), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery8(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery8(PointerFromQNetworkProxyQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var hostnameC = C.CString(hostname)
	defer C.free(unsafe.Pointer(hostnameC))
	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(hostnameC, C.int(int32(port)), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(core.PointerFromQUrl(requestUrl), C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func NewQNetworkProxyQuery4(bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	var protocolTagC = C.CString(protocolTag)
	defer C.free(unsafe.Pointer(protocolTagC))
	var tmpValue = NewQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery4(C.ushort(bindPort), protocolTagC, C.longlong(queryType)))
	runtime.SetFinalizer(tmpValue, (*QNetworkProxyQuery).DestroyQNetworkProxyQuery)
	return tmpValue
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	defer qt.Recovering("QNetworkProxyQuery::localPort")

	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkProxyQuery_LocalPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) NetworkConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkProxyQuery::networkConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkProxyQuery_NetworkConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {
	defer qt.Recovering("QNetworkProxyQuery::peerHostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_PeerHostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) PeerPort() int {
	defer qt.Recovering("QNetworkProxyQuery::peerPort")

	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkProxyQuery_PeerPort(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {
	defer qt.Recovering("QNetworkProxyQuery::protocolTag")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_ProtocolTag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {
	defer qt.Recovering("QNetworkProxyQuery::queryType")

	if ptr.Pointer() != nil {
		return QNetworkProxyQuery__QueryType(C.QNetworkProxyQuery_QueryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setLocalPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetLocalPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetNetworkConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(networkConfiguration))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerHostName")

	if ptr.Pointer() != nil {
		var hostnameC = C.CString(hostname)
		defer C.free(unsafe.Pointer(hostnameC))
		C.QNetworkProxyQuery_SetPeerHostName(ptr.Pointer(), hostnameC)
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	defer qt.Recovering("QNetworkProxyQuery::setProtocolTag")

	if ptr.Pointer() != nil {
		var protocolTagC = C.CString(protocolTag)
		defer C.free(unsafe.Pointer(protocolTagC))
		C.QNetworkProxyQuery_SetProtocolTag(ptr.Pointer(), protocolTagC)
	}
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {
	defer qt.Recovering("QNetworkProxyQuery::setQueryType")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetQueryType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QNetworkProxyQuery) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_Swap(ptr.Pointer(), PointerFromQNetworkProxyQuery(other))
	}
}

func (ptr *QNetworkProxyQuery) Url() *core.QUrl {
	defer qt.Recovering("QNetworkProxyQuery::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkProxyQuery_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	defer qt.Recovering("QNetworkProxyQuery::~QNetworkProxyQuery")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int64

const (
	QNetworkReply__NoError                           = QNetworkReply__NetworkError(0)
	QNetworkReply__ConnectionRefusedError            = QNetworkReply__NetworkError(1)
	QNetworkReply__RemoteHostClosedError             = QNetworkReply__NetworkError(2)
	QNetworkReply__HostNotFoundError                 = QNetworkReply__NetworkError(3)
	QNetworkReply__TimeoutError                      = QNetworkReply__NetworkError(4)
	QNetworkReply__OperationCanceledError            = QNetworkReply__NetworkError(5)
	QNetworkReply__SslHandshakeFailedError           = QNetworkReply__NetworkError(6)
	QNetworkReply__TemporaryNetworkFailureError      = QNetworkReply__NetworkError(7)
	QNetworkReply__NetworkSessionFailedError         = QNetworkReply__NetworkError(8)
	QNetworkReply__BackgroundRequestNotAllowedError  = QNetworkReply__NetworkError(9)
	QNetworkReply__TooManyRedirectsError             = QNetworkReply__NetworkError(10)
	QNetworkReply__InsecureRedirectError             = QNetworkReply__NetworkError(11)
	QNetworkReply__UnknownNetworkError               = QNetworkReply__NetworkError(99)
	QNetworkReply__ProxyConnectionRefusedError       = QNetworkReply__NetworkError(101)
	QNetworkReply__ProxyConnectionClosedError        = QNetworkReply__NetworkError(102)
	QNetworkReply__ProxyNotFoundError                = QNetworkReply__NetworkError(103)
	QNetworkReply__ProxyTimeoutError                 = QNetworkReply__NetworkError(104)
	QNetworkReply__ProxyAuthenticationRequiredError  = QNetworkReply__NetworkError(105)
	QNetworkReply__UnknownProxyError                 = QNetworkReply__NetworkError(199)
	QNetworkReply__ContentAccessDenied               = QNetworkReply__NetworkError(201)
	QNetworkReply__ContentOperationNotPermittedError = QNetworkReply__NetworkError(202)
	QNetworkReply__ContentNotFoundError              = QNetworkReply__NetworkError(203)
	QNetworkReply__AuthenticationRequiredError       = QNetworkReply__NetworkError(204)
	QNetworkReply__ContentReSendError                = QNetworkReply__NetworkError(205)
	QNetworkReply__ContentConflictError              = QNetworkReply__NetworkError(206)
	QNetworkReply__ContentGoneError                  = QNetworkReply__NetworkError(207)
	QNetworkReply__UnknownContentError               = QNetworkReply__NetworkError(299)
	QNetworkReply__ProtocolUnknownError              = QNetworkReply__NetworkError(301)
	QNetworkReply__ProtocolInvalidOperationError     = QNetworkReply__NetworkError(302)
	QNetworkReply__ProtocolFailure                   = QNetworkReply__NetworkError(399)
	QNetworkReply__InternalServerError               = QNetworkReply__NetworkError(401)
	QNetworkReply__OperationNotImplementedError      = QNetworkReply__NetworkError(402)
	QNetworkReply__ServiceUnavailableError           = QNetworkReply__NetworkError(403)
	QNetworkReply__UnknownServerError                = QNetworkReply__NetworkError(499)
)

type QNetworkReply struct {
	core.QIODevice
}

type QNetworkReply_ITF interface {
	core.QIODevice_ITF
	QNetworkReply_PTR() *QNetworkReply
}

func (p *QNetworkReply) QNetworkReply_PTR() *QNetworkReply {
	return p
}

func (p *QNetworkReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkReply(ptr QNetworkReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkReply_PTR().Pointer()
	}
	return nil
}

func NewQNetworkReplyFromPointer(ptr unsafe.Pointer) *QNetworkReply {
	var n = new(QNetworkReply)
	n.SetPointer(ptr)
	return n
}

//export callbackQNetworkReply_SetSslConfigurationImplementation
func callbackQNetworkReply_SetSslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::setSslConfigurationImplementation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::setSslConfigurationImplementation"); signal != nil {
		signal.(func(*QSslConfiguration))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetSslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSetSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	defer qt.Recovering("connect QNetworkReply::setSslConfigurationImplementation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setSslConfigurationImplementation", f)
	}
}

func (ptr *QNetworkReply) DisconnectSetSslConfigurationImplementation() {
	defer qt.Recovering("disconnect QNetworkReply::setSslConfigurationImplementation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setSslConfigurationImplementation")
	}
}

func (ptr *QNetworkReply) SetSslConfigurationImplementation(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::setSslConfigurationImplementation")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfigurationImplementation(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QNetworkReply) SetSslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::setSslConfigurationImplementation")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfigurationImplementationDefault(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

//export callbackQNetworkReply_SslConfigurationImplementation
func callbackQNetworkReply_SslConfigurationImplementation(ptr unsafe.Pointer, configuration unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::sslConfigurationImplementation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::sslConfigurationImplementation"); signal != nil {
		signal.(func(*QSslConfiguration))(NewQSslConfigurationFromPointer(configuration))
	} else {
		NewQNetworkReplyFromPointer(ptr).SslConfigurationImplementationDefault(NewQSslConfigurationFromPointer(configuration))
	}
}

func (ptr *QNetworkReply) ConnectSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {
	defer qt.Recovering("connect QNetworkReply::sslConfigurationImplementation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::sslConfigurationImplementation", f)
	}
}

func (ptr *QNetworkReply) DisconnectSslConfigurationImplementation() {
	defer qt.Recovering("disconnect QNetworkReply::sslConfigurationImplementation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::sslConfigurationImplementation")
	}
}

func (ptr *QNetworkReply) SslConfigurationImplementation(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::sslConfigurationImplementation")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SslConfigurationImplementation(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QNetworkReply) SslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::sslConfigurationImplementation")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SslConfigurationImplementationDefault(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

//export callbackQNetworkReply_Abort
func callbackQNetworkReply_Abort(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::abort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::abort"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectAbort(f func()) {
	defer qt.Recovering("connect QNetworkReply::abort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::abort", f)
	}
}

func (ptr *QNetworkReply) DisconnectAbort() {
	defer qt.Recovering("disconnect QNetworkReply::abort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::abort")
	}
}

func (ptr *QNetworkReply) Abort() {
	defer qt.Recovering("QNetworkReply::abort")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Abort(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *core.QVariant {
	defer qt.Recovering("QNetworkReply::attribute")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkReply_Attribute(ptr.Pointer(), C.longlong(code)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_Close
func callbackQNetworkReply_Close(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::close"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkReply) ConnectClose(f func()) {
	defer qt.Recovering("connect QNetworkReply::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::close", f)
	}
}

func (ptr *QNetworkReply) DisconnectClose() {
	defer qt.Recovering("disconnect QNetworkReply::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::close")
	}
}

func (ptr *QNetworkReply) Close() {
	defer qt.Recovering("QNetworkReply::close")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) CloseDefault() {
	defer qt.Recovering("QNetworkReply::close")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CloseDefault(ptr.Pointer())
	}
}

//export callbackQNetworkReply_DownloadProgress
func callbackQNetworkReply_DownloadProgress(ptr unsafe.Pointer, bytesReceived C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::downloadProgress")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	defer qt.Recovering("connect QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectDownloadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::downloadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectDownloadProgress() {
	defer qt.Recovering("disconnect QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::downloadProgress")
	}
}

func (ptr *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	defer qt.Recovering("QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

//export callbackQNetworkReply_Encrypted
func callbackQNetworkReply_Encrypted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::encrypted")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {
	defer qt.Recovering("connect QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::encrypted", f)
	}
}

func (ptr *QNetworkReply) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::encrypted")
	}
}

func (ptr *QNetworkReply) Encrypted() {
	defer qt.Recovering("QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Encrypted(ptr.Pointer())
	}
}

//export callbackQNetworkReply_Error2
func callbackQNetworkReply_Error2(ptr unsafe.Pointer, code C.longlong) {
	defer qt.Recovering("callback QNetworkReply::error")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::error2"); signal != nil {
		signal.(func(QNetworkReply__NetworkError))(QNetworkReply__NetworkError(code))
	}

}

func (ptr *QNetworkReply) ConnectError2(f func(code QNetworkReply__NetworkError)) {
	defer qt.Recovering("connect QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::error2", f)
	}
}

func (ptr *QNetworkReply) DisconnectError2() {
	defer qt.Recovering("disconnect QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::error2")
	}
}

func (ptr *QNetworkReply) Error2(code QNetworkReply__NetworkError) {
	defer qt.Recovering("QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Error2(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {
	defer qt.Recovering("QNetworkReply::error")

	if ptr.Pointer() != nil {
		return QNetworkReply__NetworkError(C.QNetworkReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_Finished
func callbackQNetworkReply_Finished(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::finished")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectFinished(f func()) {
	defer qt.Recovering("connect QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::finished", f)
	}
}

func (ptr *QNetworkReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::finished")
	}
}

func (ptr *QNetworkReply) Finished() {
	defer qt.Recovering("QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Finished(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) HasRawHeader(headerName string) bool {
	defer qt.Recovering("QNetworkReply::hasRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return C.QNetworkReply_HasRawHeader(ptr.Pointer(), headerNameC) != 0
	}
	return false
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkReply::header")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkReply_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_IgnoreSslErrors
func callbackQNetworkReply_IgnoreSslErrors(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::ignoreSslErrors")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::ignoreSslErrors"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).IgnoreSslErrorsDefault()
	}
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrors(f func()) {
	defer qt.Recovering("connect QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::ignoreSslErrors", f)
	}
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrors() {
	defer qt.Recovering("disconnect QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::ignoreSslErrors")
	}
}

func (ptr *QNetworkReply) IgnoreSslErrors() {
	defer qt.Recovering("QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IgnoreSslErrorsDefault() {
	defer qt.Recovering("QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrorsDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IsFinished() bool {
	defer qt.Recovering("QNetworkReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsRunning() bool {
	defer qt.Recovering("QNetworkReply::isRunning")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {
	defer qt.Recovering("QNetworkReply::manager")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkAccessManagerFromPointer(C.QNetworkReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_MetaDataChanged
func callbackQNetworkReply_MetaDataChanged(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::metaDataChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaDataChanged", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaDataChanged")
	}
}

func (ptr *QNetworkReply) MetaDataChanged() {
	defer qt.Recovering("QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Operation() QNetworkAccessManager__Operation {
	defer qt.Recovering("QNetworkReply::operation")

	if ptr.Pointer() != nil {
		return QNetworkAccessManager__Operation(C.QNetworkReply_Operation(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_PreSharedKeyAuthenticationRequired
func callbackQNetworkReply_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkReply) RawHeader(headerName string) string {
	defer qt.Recovering("QNetworkReply::rawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return qt.HexDecodeToString(C.GoString(C.QNetworkReply_RawHeader(ptr.Pointer(), headerNameC)))
	}
	return ""
}

func (ptr *QNetworkReply) ReadBufferSize() int64 {
	defer qt.Recovering("QNetworkReply::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_Redirected
func callbackQNetworkReply_Redirected(ptr unsafe.Pointer, url unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::redirected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::redirected"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QNetworkReply) ConnectRedirected(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QNetworkReply::redirected")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectRedirected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::redirected", f)
	}
}

func (ptr *QNetworkReply) DisconnectRedirected() {
	defer qt.Recovering("disconnect QNetworkReply::redirected")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectRedirected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::redirected")
	}
}

func (ptr *QNetworkReply) Redirected(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkReply::redirected")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Redirected(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) Request() *QNetworkRequest {
	defer qt.Recovering("QNetworkReply::request")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkReply_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkReply::setAttribute")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkReply) SetError(errorCode QNetworkReply__NetworkError, errorString string) {
	defer qt.Recovering("QNetworkReply::setError")

	if ptr.Pointer() != nil {
		var errorStringC = C.CString(errorString)
		defer C.free(unsafe.Pointer(errorStringC))
		C.QNetworkReply_SetError(ptr.Pointer(), C.longlong(errorCode), errorStringC)
	}
}

func (ptr *QNetworkReply) SetFinished(finished bool) {
	defer qt.Recovering("QNetworkReply::setFinished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetFinished(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(finished))))
	}
}

func (ptr *QNetworkReply) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkReply::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkReply) SetOperation(operation QNetworkAccessManager__Operation) {
	defer qt.Recovering("QNetworkReply::setOperation")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetOperation(ptr.Pointer(), C.longlong(operation))
	}
}

func (ptr *QNetworkReply) SetRawHeader(headerName string, value string) {
	defer qt.Recovering("QNetworkReply::setRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QNetworkReply_SetRawHeader(ptr.Pointer(), headerNameC, valueC)
	}
}

//export callbackQNetworkReply_SetReadBufferSize
func callbackQNetworkReply_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QNetworkReply::setReadBufferSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QNetworkReply) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setReadBufferSize", f)
	}
}

func (ptr *QNetworkReply) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::setReadBufferSize")
	}
}

func (ptr *QNetworkReply) SetReadBufferSize(size int64) {
	defer qt.Recovering("QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetRequest(request QNetworkRequest_ITF) {
	defer qt.Recovering("QNetworkReply::setRequest")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetRequest(ptr.Pointer(), PointerFromQNetworkRequest(request))
	}
}

func (ptr *QNetworkReply) SetSslConfiguration(config QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkReply) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkReply::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkReply) SslConfiguration() *QSslConfiguration {
	defer qt.Recovering("QNetworkReply::sslConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QNetworkReply_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkReply_UploadProgress
func callbackQNetworkReply_UploadProgress(ptr unsafe.Pointer, bytesSent C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::uploadProgress")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::uploadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesSent), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) ConnectUploadProgress(f func(bytesSent int64, bytesTotal int64)) {
	defer qt.Recovering("connect QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectUploadProgress(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::uploadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectUploadProgress() {
	defer qt.Recovering("disconnect QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectUploadProgress(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::uploadProgress")
	}
}

func (ptr *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	defer qt.Recovering("QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_UploadProgress(ptr.Pointer(), C.longlong(bytesSent), C.longlong(bytesTotal))
	}
}

func (ptr *QNetworkReply) Url() *core.QUrl {
	defer qt.Recovering("QNetworkReply::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkReply_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {
	defer qt.Recovering("QNetworkReply::~QNetworkReply")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DestroyQNetworkReply(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkReply_AtEnd
func callbackQNetworkReply_AtEnd(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::atEnd")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).AtEndDefault())))
}

func (ptr *QNetworkReply) ConnectAtEnd(f func() bool) {
	defer qt.Recovering("connect QNetworkReply::atEnd")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::atEnd", f)
	}
}

func (ptr *QNetworkReply) DisconnectAtEnd() {
	defer qt.Recovering("disconnect QNetworkReply::atEnd")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::atEnd")
	}
}

func (ptr *QNetworkReply) AtEnd() bool {
	defer qt.Recovering("QNetworkReply::atEnd")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) AtEndDefault() bool {
	defer qt.Recovering("QNetworkReply::atEnd")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_BytesAvailable
func callbackQNetworkReply_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkReply::bytesAvailable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QNetworkReply) ConnectBytesAvailable(f func() int64) {
	defer qt.Recovering("connect QNetworkReply::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesAvailable", f)
	}
}

func (ptr *QNetworkReply) DisconnectBytesAvailable() {
	defer qt.Recovering("disconnect QNetworkReply::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesAvailable")
	}
}

func (ptr *QNetworkReply) BytesAvailable() int64 {
	defer qt.Recovering("QNetworkReply::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) BytesAvailableDefault() int64 {
	defer qt.Recovering("QNetworkReply::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_BytesToWrite
func callbackQNetworkReply_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkReply::bytesToWrite")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QNetworkReply) ConnectBytesToWrite(f func() int64) {
	defer qt.Recovering("connect QNetworkReply::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesToWrite", f)
	}
}

func (ptr *QNetworkReply) DisconnectBytesToWrite() {
	defer qt.Recovering("disconnect QNetworkReply::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::bytesToWrite")
	}
}

func (ptr *QNetworkReply) BytesToWrite() int64 {
	defer qt.Recovering("QNetworkReply::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) BytesToWriteDefault() int64 {
	defer qt.Recovering("QNetworkReply::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_CanReadLine
func callbackQNetworkReply_CanReadLine(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::canReadLine")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QNetworkReply) ConnectCanReadLine(f func() bool) {
	defer qt.Recovering("connect QNetworkReply::canReadLine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::canReadLine", f)
	}
}

func (ptr *QNetworkReply) DisconnectCanReadLine() {
	defer qt.Recovering("disconnect QNetworkReply::canReadLine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::canReadLine")
	}
}

func (ptr *QNetworkReply) CanReadLine() bool {
	defer qt.Recovering("QNetworkReply::canReadLine")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) CanReadLineDefault() bool {
	defer qt.Recovering("QNetworkReply::canReadLine")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_IsSequential
func callbackQNetworkReply_IsSequential(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::isSequential")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QNetworkReply) ConnectIsSequential(f func() bool) {
	defer qt.Recovering("connect QNetworkReply::isSequential")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::isSequential", f)
	}
}

func (ptr *QNetworkReply) DisconnectIsSequential() {
	defer qt.Recovering("disconnect QNetworkReply::isSequential")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::isSequential")
	}
}

func (ptr *QNetworkReply) IsSequential() bool {
	defer qt.Recovering("QNetworkReply::isSequential")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsSequentialDefault() bool {
	defer qt.Recovering("QNetworkReply::isSequential")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_Open
func callbackQNetworkReply_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QNetworkReply::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QNetworkReply) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QNetworkReply::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::open", f)
	}
}

func (ptr *QNetworkReply) DisconnectOpen() {
	defer qt.Recovering("disconnect QNetworkReply::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::open")
	}
}

func (ptr *QNetworkReply) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QNetworkReply::open")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QNetworkReply) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QNetworkReply::open")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQNetworkReply_Pos
func callbackQNetworkReply_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkReply::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).PosDefault())
}

func (ptr *QNetworkReply) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QNetworkReply::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::pos", f)
	}
}

func (ptr *QNetworkReply) DisconnectPos() {
	defer qt.Recovering("disconnect QNetworkReply::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::pos")
	}
}

func (ptr *QNetworkReply) Pos() int64 {
	defer qt.Recovering("QNetworkReply::pos")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) PosDefault() int64 {
	defer qt.Recovering("QNetworkReply::pos")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_ReadLineData
func callbackQNetworkReply_ReadLineData(ptr unsafe.Pointer, data *C.char, maxSize C.longlong) C.longlong {
	defer qt.Recovering("callback QNetworkReply::readLineData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(maxSize)))
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).ReadLineDataDefault(C.GoString(data), int64(maxSize)))
}

func (ptr *QNetworkReply) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	defer qt.Recovering("connect QNetworkReply::readLineData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readLineData", f)
	}
}

func (ptr *QNetworkReply) DisconnectReadLineData() {
	defer qt.Recovering("disconnect QNetworkReply::readLineData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::readLineData")
	}
}

func (ptr *QNetworkReply) ReadLineData(data string, maxSize int64) int64 {
	defer qt.Recovering("QNetworkReply::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QNetworkReply) ReadLineDataDefault(data string, maxSize int64) int64 {
	defer qt.Recovering("QNetworkReply::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_Reset
func callbackQNetworkReply_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).ResetDefault())))
}

func (ptr *QNetworkReply) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QNetworkReply::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::reset", f)
	}
}

func (ptr *QNetworkReply) DisconnectReset() {
	defer qt.Recovering("disconnect QNetworkReply::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::reset")
	}
}

func (ptr *QNetworkReply) Reset() bool {
	defer qt.Recovering("QNetworkReply::reset")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) ResetDefault() bool {
	defer qt.Recovering("QNetworkReply::reset")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkReply_Seek
func callbackQNetworkReply_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QNetworkReply::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QNetworkReply) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QNetworkReply::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::seek", f)
	}
}

func (ptr *QNetworkReply) DisconnectSeek() {
	defer qt.Recovering("disconnect QNetworkReply::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::seek")
	}
}

func (ptr *QNetworkReply) Seek(pos int64) bool {
	defer qt.Recovering("QNetworkReply::seek")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QNetworkReply) SeekDefault(pos int64) bool {
	defer qt.Recovering("QNetworkReply::seek")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQNetworkReply_Size
func callbackQNetworkReply_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QNetworkReply::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).SizeDefault())
}

func (ptr *QNetworkReply) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QNetworkReply::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::size", f)
	}
}

func (ptr *QNetworkReply) DisconnectSize() {
	defer qt.Recovering("disconnect QNetworkReply::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::size")
	}
}

func (ptr *QNetworkReply) Size() int64 {
	defer qt.Recovering("QNetworkReply::size")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) SizeDefault() int64 {
	defer qt.Recovering("QNetworkReply::size")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkReply_WaitForBytesWritten
func callbackQNetworkReply_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QNetworkReply::waitForBytesWritten")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	defer qt.Recovering("connect QNetworkReply::waitForBytesWritten")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForBytesWritten", f)
	}
}

func (ptr *QNetworkReply) DisconnectWaitForBytesWritten() {
	defer qt.Recovering("disconnect QNetworkReply::waitForBytesWritten")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForBytesWritten")
	}
}

func (ptr *QNetworkReply) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QNetworkReply::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QNetworkReply) WaitForBytesWrittenDefault(msecs int) bool {
	defer qt.Recovering("QNetworkReply::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkReply_WaitForReadyRead
func callbackQNetworkReply_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QNetworkReply::waitForReadyRead")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QNetworkReply) ConnectWaitForReadyRead(f func(msecs int) bool) {
	defer qt.Recovering("connect QNetworkReply::waitForReadyRead")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForReadyRead", f)
	}
}

func (ptr *QNetworkReply) DisconnectWaitForReadyRead() {
	defer qt.Recovering("disconnect QNetworkReply::waitForReadyRead")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::waitForReadyRead")
	}
}

func (ptr *QNetworkReply) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QNetworkReply::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QNetworkReply) WaitForReadyReadDefault(msecs int) bool {
	defer qt.Recovering("QNetworkReply::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkReply_WriteData
func callbackQNetworkReply_WriteData(ptr unsafe.Pointer, data *C.char, maxSize C.longlong) C.longlong {
	defer qt.Recovering("callback QNetworkReply::writeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(maxSize)))
	}

	return C.longlong(NewQNetworkReplyFromPointer(ptr).WriteDataDefault(C.GoString(data), int64(maxSize)))
}

func (ptr *QNetworkReply) ConnectWriteData(f func(data string, maxSize int64) int64) {
	defer qt.Recovering("connect QNetworkReply::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::writeData", f)
	}
}

func (ptr *QNetworkReply) DisconnectWriteData() {
	defer qt.Recovering("disconnect QNetworkReply::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::writeData")
	}
}

func (ptr *QNetworkReply) WriteData(data string, maxSize int64) int64 {
	defer qt.Recovering("QNetworkReply::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QNetworkReply) WriteDataDefault(data string, maxSize int64) int64 {
	defer qt.Recovering("QNetworkReply::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QNetworkReply_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQNetworkReply_TimerEvent
func callbackQNetworkReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::timerEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::timerEvent")
	}
}

func (ptr *QNetworkReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkReply_ChildEvent
func callbackQNetworkReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::childEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::childEvent")
	}
}

func (ptr *QNetworkReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkReply::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkReply::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkReply_ConnectNotify
func callbackQNetworkReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::connectNotify", f)
	}
}

func (ptr *QNetworkReply) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::connectNotify")
	}
}

func (ptr *QNetworkReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_CustomEvent
func callbackQNetworkReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::customEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::customEvent")
	}
}

func (ptr *QNetworkReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkReply::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkReply::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkReply_DeleteLater
func callbackQNetworkReply_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkReply) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::deleteLater", f)
	}
}

func (ptr *QNetworkReply) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::deleteLater")
	}
}

func (ptr *QNetworkReply) DeleteLater() {
	defer qt.Recovering("QNetworkReply::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkReply::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkReply_DisconnectNotify
func callbackQNetworkReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::disconnectNotify", f)
	}
}

func (ptr *QNetworkReply) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::disconnectNotify")
	}
}

func (ptr *QNetworkReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkReply_Event
func callbackQNetworkReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkReply) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkReply::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::event", f)
	}
}

func (ptr *QNetworkReply) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkReply::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::event")
	}
}

func (ptr *QNetworkReply) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkReply::event")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkReply) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkReply::event")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkReply_EventFilter
func callbackQNetworkReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkReply::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::eventFilter", f)
	}
}

func (ptr *QNetworkReply) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::eventFilter")
	}
}

func (ptr *QNetworkReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkReply_MetaObject
func callbackQNetworkReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkReply::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkReply::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkReply) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkReply::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaObject", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkReply::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkReply::metaObject")
	}
}

func (ptr *QNetworkReply) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkReply) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int64

const (
	QNetworkRequest__HttpStatusCodeAttribute               = QNetworkRequest__Attribute(0)
	QNetworkRequest__HttpReasonPhraseAttribute             = QNetworkRequest__Attribute(1)
	QNetworkRequest__RedirectionTargetAttribute            = QNetworkRequest__Attribute(2)
	QNetworkRequest__ConnectionEncryptedAttribute          = QNetworkRequest__Attribute(3)
	QNetworkRequest__CacheLoadControlAttribute             = QNetworkRequest__Attribute(4)
	QNetworkRequest__CacheSaveControlAttribute             = QNetworkRequest__Attribute(5)
	QNetworkRequest__SourceIsFromCacheAttribute            = QNetworkRequest__Attribute(6)
	QNetworkRequest__DoNotBufferUploadDataAttribute        = QNetworkRequest__Attribute(7)
	QNetworkRequest__HttpPipeliningAllowedAttribute        = QNetworkRequest__Attribute(8)
	QNetworkRequest__HttpPipeliningWasUsedAttribute        = QNetworkRequest__Attribute(9)
	QNetworkRequest__CustomVerbAttribute                   = QNetworkRequest__Attribute(10)
	QNetworkRequest__CookieLoadControlAttribute            = QNetworkRequest__Attribute(11)
	QNetworkRequest__AuthenticationReuseAttribute          = QNetworkRequest__Attribute(12)
	QNetworkRequest__CookieSaveControlAttribute            = QNetworkRequest__Attribute(13)
	QNetworkRequest__MaximumDownloadBufferSizeAttribute    = QNetworkRequest__Attribute(14)
	QNetworkRequest__DownloadBufferAttribute               = QNetworkRequest__Attribute(15)
	QNetworkRequest__SynchronousRequestAttribute           = QNetworkRequest__Attribute(16)
	QNetworkRequest__BackgroundRequestAttribute            = QNetworkRequest__Attribute(17)
	QNetworkRequest__SpdyAllowedAttribute                  = QNetworkRequest__Attribute(18)
	QNetworkRequest__SpdyWasUsedAttribute                  = QNetworkRequest__Attribute(19)
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute = QNetworkRequest__Attribute(20)
	QNetworkRequest__FollowRedirectsAttribute              = QNetworkRequest__Attribute(21)
	QNetworkRequest__User                                  = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               = QNetworkRequest__Attribute(32767)
)

//QNetworkRequest::CacheLoadControl
type QNetworkRequest__CacheLoadControl int64

const (
	QNetworkRequest__AlwaysNetwork = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   = QNetworkRequest__CacheLoadControl(3)
)

//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int64

const (
	QNetworkRequest__ContentTypeHeader        = QNetworkRequest__KnownHeaders(0)
	QNetworkRequest__ContentLengthHeader      = QNetworkRequest__KnownHeaders(1)
	QNetworkRequest__LocationHeader           = QNetworkRequest__KnownHeaders(2)
	QNetworkRequest__LastModifiedHeader       = QNetworkRequest__KnownHeaders(3)
	QNetworkRequest__CookieHeader             = QNetworkRequest__KnownHeaders(4)
	QNetworkRequest__SetCookieHeader          = QNetworkRequest__KnownHeaders(5)
	QNetworkRequest__ContentDispositionHeader = QNetworkRequest__KnownHeaders(6)
	QNetworkRequest__UserAgentHeader          = QNetworkRequest__KnownHeaders(7)
	QNetworkRequest__ServerHeader             = QNetworkRequest__KnownHeaders(8)
)

//QNetworkRequest::LoadControl
type QNetworkRequest__LoadControl int64

const (
	QNetworkRequest__Automatic = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    = QNetworkRequest__LoadControl(1)
)

//QNetworkRequest::Priority
type QNetworkRequest__Priority int64

const (
	QNetworkRequest__HighPriority   = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    = QNetworkRequest__Priority(5)
)

type QNetworkRequest struct {
	ptr unsafe.Pointer
}

type QNetworkRequest_ITF interface {
	QNetworkRequest_PTR() *QNetworkRequest
}

func (p *QNetworkRequest) QNetworkRequest_PTR() *QNetworkRequest {
	return p
}

func (p *QNetworkRequest) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QNetworkRequest) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQNetworkRequest(ptr QNetworkRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkRequest_PTR().Pointer()
	}
	return nil
}

func NewQNetworkRequestFromPointer(ptr unsafe.Pointer) *QNetworkRequest {
	var n = new(QNetworkRequest)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkRequest2(other QNetworkRequest_ITF) *QNetworkRequest {
	defer qt.Recovering("QNetworkRequest::QNetworkRequest")

	var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest2(PointerFromQNetworkRequest(other)))
	runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {
	defer qt.Recovering("QNetworkRequest::QNetworkRequest")

	var tmpValue = NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest(core.PointerFromQUrl(url)))
	runtime.SetFinalizer(tmpValue, (*QNetworkRequest).DestroyQNetworkRequest)
	return tmpValue
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QNetworkRequest::attribute")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) HasRawHeader(headerName string) bool {
	defer qt.Recovering("QNetworkRequest::hasRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return C.QNetworkRequest_HasRawHeader(ptr.Pointer(), headerNameC) != 0
	}
	return false
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkRequest::header")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkRequest_Header(ptr.Pointer(), C.longlong(header)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) MaximumRedirectsAllowed() int {
	defer qt.Recovering("QNetworkRequest::maximumRedirectsAllowed")

	if ptr.Pointer() != nil {
		return int(int32(C.QNetworkRequest_MaximumRedirectsAllowed(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkRequest) OriginatingObject() *core.QObject {
	defer qt.Recovering("QNetworkRequest::originatingObject")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QNetworkRequest_OriginatingObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {
	defer qt.Recovering("QNetworkRequest::priority")

	if ptr.Pointer() != nil {
		return QNetworkRequest__Priority(C.QNetworkRequest_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkRequest) RawHeader(headerName string) string {
	defer qt.Recovering("QNetworkRequest::rawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		return qt.HexDecodeToString(C.GoString(C.QNetworkRequest_RawHeader(ptr.Pointer(), headerNameC)))
	}
	return ""
}

func (ptr *QNetworkRequest) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkRequest::setAttribute")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkRequest::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetHeader(ptr.Pointer(), C.longlong(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetMaximumRedirectsAllowed(maxRedirectsAllowed int) {
	defer qt.Recovering("QNetworkRequest::setMaximumRedirectsAllowed")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetMaximumRedirectsAllowed(ptr.Pointer(), C.int(int32(maxRedirectsAllowed)))
	}
}

func (ptr *QNetworkRequest) SetOriginatingObject(object core.QObject_ITF) {
	defer qt.Recovering("QNetworkRequest::setOriginatingObject")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetOriginatingObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QNetworkRequest) SetPriority(priority QNetworkRequest__Priority) {
	defer qt.Recovering("QNetworkRequest::setPriority")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetPriority(ptr.Pointer(), C.longlong(priority))
	}
}

func (ptr *QNetworkRequest) SetRawHeader(headerName string, headerValue string) {
	defer qt.Recovering("QNetworkRequest::setRawHeader")

	if ptr.Pointer() != nil {
		var headerNameC = C.CString(hex.EncodeToString([]byte(headerName)))
		defer C.free(unsafe.Pointer(headerNameC))
		var headerValueC = C.CString(hex.EncodeToString([]byte(headerValue)))
		defer C.free(unsafe.Pointer(headerValueC))
		C.QNetworkRequest_SetRawHeader(ptr.Pointer(), headerNameC, headerValueC)
	}
}

func (ptr *QNetworkRequest) SetSslConfiguration(config QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkRequest::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkRequest) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkRequest::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkRequest) SslConfiguration() *QSslConfiguration {
	defer qt.Recovering("QNetworkRequest::sslConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QNetworkRequest_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) Swap(other QNetworkRequest_ITF) {
	defer qt.Recovering("QNetworkRequest::swap")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_Swap(ptr.Pointer(), PointerFromQNetworkRequest(other))
	}
}

func (ptr *QNetworkRequest) Url() *core.QUrl {
	defer qt.Recovering("QNetworkRequest::url")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QNetworkRequest_Url(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {
	defer qt.Recovering("QNetworkRequest::~QNetworkRequest")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_DestroyQNetworkRequest(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  = QNetworkSession__SessionError(4)
)

//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      = QNetworkSession__State(0)
	QNetworkSession__NotAvailable = QNetworkSession__State(1)
	QNetworkSession__Connecting   = QNetworkSession__State(2)
	QNetworkSession__Connected    = QNetworkSession__State(3)
	QNetworkSession__Closing      = QNetworkSession__State(4)
	QNetworkSession__Disconnected = QNetworkSession__State(5)
	QNetworkSession__Roaming      = QNetworkSession__State(6)
)

//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy = QNetworkSession__UsagePolicy(1)
)

type QNetworkSession struct {
	core.QObject
}

type QNetworkSession_ITF interface {
	core.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func (p *QNetworkSession) QNetworkSession_PTR() *QNetworkSession {
	return p
}

func (p *QNetworkSession) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QNetworkSession) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQNetworkSession(ptr QNetworkSession_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSession_PTR().Pointer()
	}
	return nil
}

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = new(QNetworkSession)
	n.SetPointer(ptr)
	return n
}
func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {
	defer qt.Recovering("QNetworkSession::QNetworkSession")

	var tmpValue = NewQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQNetworkSession_Accept
func callbackQNetworkSession_Accept(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::accept")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::accept"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectAccept(f func()) {
	defer qt.Recovering("connect QNetworkSession::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::accept", f)
	}
}

func (ptr *QNetworkSession) DisconnectAccept() {
	defer qt.Recovering("disconnect QNetworkSession::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::accept")
	}
}

func (ptr *QNetworkSession) Accept() {
	defer qt.Recovering("QNetworkSession::accept")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ActiveTime() uint64 {
	defer qt.Recovering("QNetworkSession::activeTime")

	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_ActiveTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) BytesReceived() uint64 {
	defer qt.Recovering("QNetworkSession::bytesReceived")

	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_BytesReceived(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) BytesWritten() uint64 {
	defer qt.Recovering("QNetworkSession::bytesWritten")

	if ptr.Pointer() != nil {
		return uint64(C.QNetworkSession_BytesWritten(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_Close
func callbackQNetworkSession_Close(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectClose(f func()) {
	defer qt.Recovering("connect QNetworkSession::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::close", f)
	}
}

func (ptr *QNetworkSession) DisconnectClose() {
	defer qt.Recovering("disconnect QNetworkSession::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::close")
	}
}

func (ptr *QNetworkSession) Close() {
	defer qt.Recovering("QNetworkSession::close")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Closed
func callbackQNetworkSession_Closed(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::closed")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	defer qt.Recovering("connect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	defer qt.Recovering("disconnect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::closed")
	}
}

func (ptr *QNetworkSession) Closed() {
	defer qt.Recovering("QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Closed(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Configuration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkSession::configuration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkConfigurationFromPointer(C.QNetworkSession_Configuration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkConfiguration).DestroyQNetworkConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQNetworkSession_Error2
func callbackQNetworkSession_Error2(ptr unsafe.Pointer, error C.longlong) {
	defer qt.Recovering("callback QNetworkSession::error")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::error2"); signal != nil {
		signal.(func(QNetworkSession__SessionError))(QNetworkSession__SessionError(error))
	}

}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {
	defer qt.Recovering("connect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::error2", f)
	}
}

func (ptr *QNetworkSession) DisconnectError2() {
	defer qt.Recovering("disconnect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::error2")
	}
}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Error2(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	defer qt.Recovering("QNetworkSession::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQNetworkSession_Ignore
func callbackQNetworkSession_Ignore(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::ignore")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::ignore"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectIgnore(f func()) {
	defer qt.Recovering("connect QNetworkSession::ignore")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::ignore", f)
	}
}

func (ptr *QNetworkSession) DisconnectIgnore() {
	defer qt.Recovering("disconnect QNetworkSession::ignore")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::ignore")
	}
}

func (ptr *QNetworkSession) Ignore() {
	defer qt.Recovering("QNetworkSession::ignore")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Interface() *QNetworkInterface {
	defer qt.Recovering("QNetworkSession::interface")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkInterfaceFromPointer(C.QNetworkSession_Interface(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) IsOpen() bool {
	defer qt.Recovering("QNetworkSession::isOpen")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQNetworkSession_Migrate
func callbackQNetworkSession_Migrate(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::migrate")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::migrate"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectMigrate(f func()) {
	defer qt.Recovering("connect QNetworkSession::migrate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::migrate", f)
	}
}

func (ptr *QNetworkSession) DisconnectMigrate() {
	defer qt.Recovering("disconnect QNetworkSession::migrate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::migrate")
	}
}

func (ptr *QNetworkSession) Migrate() {
	defer qt.Recovering("QNetworkSession::migrate")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(ptr.Pointer())
	}
}

//export callbackQNetworkSession_NewConfigurationActivated
func callbackQNetworkSession_NewConfigurationActivated(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::newConfigurationActivated")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::newConfigurationActivated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	defer qt.Recovering("connect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	defer qt.Recovering("disconnect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::newConfigurationActivated")
	}
}

func (ptr *QNetworkSession) NewConfigurationActivated() {
	defer qt.Recovering("QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_NewConfigurationActivated(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Open
func callbackQNetworkSession_Open(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::open"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectOpen(f func()) {
	defer qt.Recovering("connect QNetworkSession::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::open", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpen() {
	defer qt.Recovering("disconnect QNetworkSession::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::open")
	}
}

func (ptr *QNetworkSession) Open() {
	defer qt.Recovering("QNetworkSession::open")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(ptr.Pointer())
	}
}

//export callbackQNetworkSession_Opened
func callbackQNetworkSession_Opened(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::opened")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::opened"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	defer qt.Recovering("connect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	defer qt.Recovering("disconnect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::opened")
	}
}

func (ptr *QNetworkSession) Opened() {
	defer qt.Recovering("QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Opened(ptr.Pointer())
	}
}

//export callbackQNetworkSession_PreferredConfigurationChanged
func callbackQNetworkSession_PreferredConfigurationChanged(ptr unsafe.Pointer, config unsafe.Pointer, isSeamless C.char) {
	defer qt.Recovering("callback QNetworkSession::preferredConfigurationChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::preferredConfigurationChanged"); signal != nil {
		signal.(func(*QNetworkConfiguration, bool))(NewQNetworkConfigurationFromPointer(config), int8(isSeamless) != 0)
	}

}

func (ptr *QNetworkSession) ConnectPreferredConfigurationChanged(f func(config *QNetworkConfiguration, isSeamless bool)) {
	defer qt.Recovering("connect QNetworkSession::preferredConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectPreferredConfigurationChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::preferredConfigurationChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectPreferredConfigurationChanged() {
	defer qt.Recovering("disconnect QNetworkSession::preferredConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectPreferredConfigurationChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::preferredConfigurationChanged")
	}
}

func (ptr *QNetworkSession) PreferredConfigurationChanged(config QNetworkConfiguration_ITF, isSeamless bool) {
	defer qt.Recovering("QNetworkSession::preferredConfigurationChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_PreferredConfigurationChanged(ptr.Pointer(), PointerFromQNetworkConfiguration(config), C.char(int8(qt.GoBoolToInt(isSeamless))))
	}
}

//export callbackQNetworkSession_Reject
func callbackQNetworkSession_Reject(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::reject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::reject"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectReject(f func()) {
	defer qt.Recovering("connect QNetworkSession::reject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::reject", f)
	}
}

func (ptr *QNetworkSession) DisconnectReject() {
	defer qt.Recovering("disconnect QNetworkSession::reject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::reject")
	}
}

func (ptr *QNetworkSession) Reject() {
	defer qt.Recovering("QNetworkSession::reject")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {
	defer qt.Recovering("QNetworkSession::sessionProperty")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), keyC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkSession::setSessionProperty")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), keyC, core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkSession) State() QNetworkSession__State {
	defer qt.Recovering("QNetworkSession::state")

	if ptr.Pointer() != nil {
		return QNetworkSession__State(C.QNetworkSession_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_StateChanged
func callbackQNetworkSession_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	defer qt.Recovering("callback QNetworkSession::stateChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::stateChanged"); signal != nil {
		signal.(func(QNetworkSession__State))(QNetworkSession__State(state))
	}

}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	defer qt.Recovering("connect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stateChanged")
	}
}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {
	defer qt.Recovering("QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQNetworkSession_Stop
func callbackQNetworkSession_Stop(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::stop")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) ConnectStop(f func()) {
	defer qt.Recovering("connect QNetworkSession::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stop", f)
	}
}

func (ptr *QNetworkSession) DisconnectStop() {
	defer qt.Recovering("disconnect QNetworkSession::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::stop")
	}
}

func (ptr *QNetworkSession) Stop() {
	defer qt.Recovering("QNetworkSession::stop")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	defer qt.Recovering("QNetworkSession::usagePolicies")

	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(ptr.Pointer()))
	}
	return 0
}

//export callbackQNetworkSession_UsagePoliciesChanged
func callbackQNetworkSession_UsagePoliciesChanged(ptr unsafe.Pointer, usagePolicies C.longlong) {
	defer qt.Recovering("callback QNetworkSession::usagePoliciesChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::usagePoliciesChanged"); signal != nil {
		signal.(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
	}

}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	defer qt.Recovering("connect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	defer qt.Recovering("disconnect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::usagePoliciesChanged")
	}
}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	defer qt.Recovering("QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_UsagePoliciesChanged(ptr.Pointer(), C.longlong(usagePolicies))
	}
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	defer qt.Recovering("QNetworkSession::waitForOpened")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQNetworkSession_DestroyQNetworkSession
func callbackQNetworkSession_DestroyQNetworkSession(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::~QNetworkSession")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::~QNetworkSession"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).DestroyQNetworkSessionDefault()
	}
}

func (ptr *QNetworkSession) ConnectDestroyQNetworkSession(f func()) {
	defer qt.Recovering("connect QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::~QNetworkSession", f)
	}
}

func (ptr *QNetworkSession) DisconnectDestroyQNetworkSession() {
	defer qt.Recovering("disconnect QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::~QNetworkSession")
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	defer qt.Recovering("QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) DestroyQNetworkSessionDefault() {
	defer qt.Recovering("QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSessionDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkSession_TimerEvent
func callbackQNetworkSession_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::timerEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::timerEvent")
	}
}

func (ptr *QNetworkSession) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQNetworkSession_ChildEvent
func callbackQNetworkSession_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::childEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::childEvent")
	}
}

func (ptr *QNetworkSession) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQNetworkSession_ConnectNotify
func callbackQNetworkSession_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkSession::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::connectNotify", f)
	}
}

func (ptr *QNetworkSession) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QNetworkSession::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::connectNotify")
	}
}

func (ptr *QNetworkSession) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkSession::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkSession) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkSession::connectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_CustomEvent
func callbackQNetworkSession_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::customEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::customEvent")
	}
}

func (ptr *QNetworkSession) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQNetworkSession_DeleteLater
func callbackQNetworkSession_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkSessionFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QNetworkSession) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QNetworkSession::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::deleteLater", f)
	}
}

func (ptr *QNetworkSession) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QNetworkSession::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::deleteLater")
	}
}

func (ptr *QNetworkSession) DeleteLater() {
	defer qt.Recovering("QNetworkSession::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) DeleteLaterDefault() {
	defer qt.Recovering("QNetworkSession::deleteLater")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQNetworkSession_DisconnectNotify
func callbackQNetworkSession_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQNetworkSessionFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QNetworkSession) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QNetworkSession::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::disconnectNotify", f)
	}
}

func (ptr *QNetworkSession) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QNetworkSession::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::disconnectNotify")
	}
}

func (ptr *QNetworkSession) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkSession::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QNetworkSession) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QNetworkSession::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQNetworkSession_Event
func callbackQNetworkSession_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkSession::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QNetworkSession) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkSession::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::event", f)
	}
}

func (ptr *QNetworkSession) DisconnectEvent() {
	defer qt.Recovering("disconnect QNetworkSession::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::event")
	}
}

func (ptr *QNetworkSession) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkSession::event")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QNetworkSession) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkSession::event")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQNetworkSession_EventFilter
func callbackQNetworkSession_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QNetworkSession::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQNetworkSessionFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QNetworkSession) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QNetworkSession::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::eventFilter", f)
	}
}

func (ptr *QNetworkSession) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QNetworkSession::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::eventFilter")
	}
}

func (ptr *QNetworkSession) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkSession::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QNetworkSession) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QNetworkSession::eventFilter")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQNetworkSession_MetaObject
func callbackQNetworkSession_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QNetworkSession::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QNetworkSession::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQNetworkSessionFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QNetworkSession) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QNetworkSession::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::metaObject", f)
	}
}

func (ptr *QNetworkSession) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QNetworkSession::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QNetworkSession::metaObject")
	}
}

func (ptr *QNetworkSession) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QNetworkSession::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkSession_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkSession) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QNetworkSession::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QNetworkSession_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int64

const (
	QSslCertificate__Organization               = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               = QSslCertificate__SubjectInfo(8)
)

type QSslCertificate struct {
	ptr unsafe.Pointer
}

type QSslCertificate_ITF interface {
	QSslCertificate_PTR() *QSslCertificate
}

func (p *QSslCertificate) QSslCertificate_PTR() *QSslCertificate {
	return p
}

func (p *QSslCertificate) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslCertificate) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslCertificate(ptr QSslCertificate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificate_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateFromPointer(ptr unsafe.Pointer) *QSslCertificate {
	var n = new(QSslCertificate)
	n.SetPointer(ptr)
	return n
}
func NewQSslCertificate3(other QSslCertificate_ITF) *QSslCertificate {
	defer qt.Recovering("QSslCertificate::QSslCertificate")

	var tmpValue = NewQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate3(PointerFromQSslCertificate(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
	return tmpValue
}

func (ptr *QSslCertificate) Clear() {
	defer qt.Recovering("QSslCertificate::clear")

	if ptr.Pointer() != nil {
		C.QSslCertificate_Clear(ptr.Pointer())
	}
}

func (ptr *QSslCertificate) Digest(algorithm core.QCryptographicHash__Algorithm) string {
	defer qt.Recovering("QSslCertificate::digest")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslCertificate_Digest(ptr.Pointer(), C.longlong(algorithm))))
	}
	return ""
}

func (ptr *QSslCertificate) IsBlacklisted() bool {
	defer qt.Recovering("QSslCertificate::isBlacklisted")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsBlacklisted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) Swap(other QSslCertificate_ITF) {
	defer qt.Recovering("QSslCertificate::swap")

	if ptr.Pointer() != nil {
		C.QSslCertificate_Swap(ptr.Pointer(), PointerFromQSslCertificate(other))
	}
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {
	defer qt.Recovering("QSslCertificate::~QSslCertificate")

	if ptr.Pointer() != nil {
		C.QSslCertificate_DestroyQSslCertificate(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslCertificate) EffectiveDate() *core.QDateTime {
	defer qt.Recovering("QSslCertificate::effectiveDate")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QSslCertificate_EffectiveDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) ExpiryDate() *core.QDateTime {
	defer qt.Recovering("QSslCertificate::expiryDate")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQDateTimeFromPointer(C.QSslCertificate_ExpiryDate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QDateTime).DestroyQDateTime)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) IsNull() bool {
	defer qt.Recovering("QSslCertificate::isNull")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsSelfSigned() bool {
	defer qt.Recovering("QSslCertificate::isSelfSigned")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsSelfSigned(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {
	defer qt.Recovering("QSslCertificate::issuerInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo(ptr.Pointer(), C.longlong(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute string) []string {
	defer qt.Recovering("QSslCertificate::issuerInfo")

	if ptr.Pointer() != nil {
		var attributeC = C.CString(hex.EncodeToString([]byte(attribute)))
		defer C.free(unsafe.Pointer(attributeC))
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo2(ptr.Pointer(), attributeC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) PublicKey() *QSslKey {
	defer qt.Recovering("QSslCertificate::publicKey")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslCertificate_PublicKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificate) SerialNumber() string {
	defer qt.Recovering("QSslCertificate::serialNumber")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslCertificate_SerialNumber(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {
	defer qt.Recovering("QSslCertificate::subjectInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo(ptr.Pointer(), C.longlong(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute string) []string {
	defer qt.Recovering("QSslCertificate::subjectInfo")

	if ptr.Pointer() != nil {
		var attributeC = C.CString(hex.EncodeToString([]byte(attribute)))
		defer C.free(unsafe.Pointer(attributeC))
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo2(ptr.Pointer(), attributeC)), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) ToDer() string {
	defer qt.Recovering("QSslCertificate::toDer")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslCertificate_ToDer(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificate) ToPem() string {
	defer qt.Recovering("QSslCertificate::toPem")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslCertificate_ToPem(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCertificate) ToText() string {
	defer qt.Recovering("QSslCertificate::toText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificate_ToText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificate) Version() string {
	defer qt.Recovering("QSslCertificate::version")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslCertificate_Version(ptr.Pointer())))
	}
	return ""
}

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (p *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return p
}

func (p *QSslCertificateExtension) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslCertificateExtension) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}
func NewQSslCertificateExtension() *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	var tmpValue = NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
	runtime.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	var tmpValue = NewQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCertificateExtension).DestroyQSslCertificateExtension)
	return tmpValue
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	defer qt.Recovering("QSslCertificateExtension::isCritical")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	defer qt.Recovering("QSslCertificateExtension::isSupported")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	defer qt.Recovering("QSslCertificateExtension::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	defer qt.Recovering("QSslCertificateExtension::oid")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Oid(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	defer qt.Recovering("QSslCertificateExtension::swap")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(ptr.Pointer(), PointerFromQSslCertificateExtension(other))
	}
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {
	defer qt.Recovering("QSslCertificateExtension::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	defer qt.Recovering("QSslCertificateExtension::~QSslCertificateExtension")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslCipher struct {
	ptr unsafe.Pointer
}

type QSslCipher_ITF interface {
	QSslCipher_PTR() *QSslCipher
}

func (p *QSslCipher) QSslCipher_PTR() *QSslCipher {
	return p
}

func (p *QSslCipher) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslCipher) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslCipher(ptr QSslCipher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCipher_PTR().Pointer()
	}
	return nil
}

func NewQSslCipherFromPointer(ptr unsafe.Pointer) *QSslCipher {
	var n = new(QSslCipher)
	n.SetPointer(ptr)
	return n
}
func NewQSslCipher() *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher())
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher4(other QSslCipher_ITF) *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher4(PointerFromQSslCipher(other)))
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func NewQSslCipher2(name string) *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSslCipherFromPointer(C.QSslCipher_NewQSslCipher2(nameC))
	runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
	return tmpValue
}

func (ptr *QSslCipher) AuthenticationMethod() string {
	defer qt.Recovering("QSslCipher::authenticationMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_AuthenticationMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) EncryptionMethod() string {
	defer qt.Recovering("QSslCipher::encryptionMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_EncryptionMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) IsNull() bool {
	defer qt.Recovering("QSslCipher::isNull")

	if ptr.Pointer() != nil {
		return C.QSslCipher_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCipher) KeyExchangeMethod() string {
	defer qt.Recovering("QSslCipher::keyExchangeMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_KeyExchangeMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) Name() string {
	defer qt.Recovering("QSslCipher::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) ProtocolString() string {
	defer qt.Recovering("QSslCipher::protocolString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_ProtocolString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) SupportedBits() int {
	defer qt.Recovering("QSslCipher::supportedBits")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslCipher_SupportedBits(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) Swap(other QSslCipher_ITF) {
	defer qt.Recovering("QSslCipher::swap")

	if ptr.Pointer() != nil {
		C.QSslCipher_Swap(ptr.Pointer(), PointerFromQSslCipher(other))
	}
}

func (ptr *QSslCipher) UsedBits() int {
	defer qt.Recovering("QSslCipher::usedBits")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslCipher_UsedBits(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) DestroyQSslCipher() {
	defer qt.Recovering("QSslCipher::~QSslCipher")

	if ptr.Pointer() != nil {
		C.QSslCipher_DestroyQSslCipher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

type QSslConfiguration struct {
	ptr unsafe.Pointer
}

type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (p *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration {
	return p
}

func (p *QSslConfiguration) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslConfiguration) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslConfiguration(ptr QSslConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return n
}
func NewQSslConfiguration() *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::QSslConfiguration")

	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::QSslConfiguration")

	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration2(PointerFromQSslConfiguration(other)))
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func QSslConfiguration_DefaultConfiguration() *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::defaultConfiguration")

	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) DefaultConfiguration() *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::defaultConfiguration")

	var tmpValue = NewQSslConfigurationFromPointer(C.QSslConfiguration_QSslConfiguration_DefaultConfiguration())
	runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
	return tmpValue
}

func (ptr *QSslConfiguration) EphemeralServerKey() *QSslKey {
	defer qt.Recovering("QSslConfiguration::ephemeralServerKey")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslConfiguration_EphemeralServerKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) IsNull() bool {
	defer qt.Recovering("QSslConfiguration::isNull")

	if ptr.Pointer() != nil {
		return C.QSslConfiguration_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslConfiguration) LocalCertificate() *QSslCertificate {
	defer qt.Recovering("QSslConfiguration::localCertificate")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_LocalCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() string {
	defer qt.Recovering("QSslConfiguration::nextNegotiatedProtocol")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslConfiguration_NextNegotiatedProtocol(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {
	defer qt.Recovering("QSslConfiguration::nextProtocolNegotiationStatus")

	if ptr.Pointer() != nil {
		return QSslConfiguration__NextProtocolNegotiationStatus(C.QSslConfiguration_NextProtocolNegotiationStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerCertificate() *QSslCertificate {
	defer qt.Recovering("QSslConfiguration::peerCertificate")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslConfiguration_PeerCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {
	defer qt.Recovering("QSslConfiguration::peerVerifyDepth")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslConfiguration_PeerVerifyDepth(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	defer qt.Recovering("QSslConfiguration::peerVerifyMode")

	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslConfiguration_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PrivateKey() *QSslKey {
	defer qt.Recovering("QSslConfiguration::privateKey")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslConfiguration_PrivateKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) SessionCipher() *QSslCipher {
	defer qt.Recovering("QSslConfiguration::sessionCipher")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslConfiguration_SessionCipher(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslConfiguration) SessionTicket() string {
	defer qt.Recovering("QSslConfiguration::sessionTicket")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslConfiguration_SessionTicket(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {
	defer qt.Recovering("QSslConfiguration::sessionTicketLifeTimeHint")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslConfiguration_SessionTicketLifeTimeHint(ptr.Pointer())))
	}
	return 0
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QSslConfiguration::setDefaultConfiguration")

	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QSslConfiguration::setDefaultConfiguration")

	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslConfiguration::setLocalCertificate")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	defer qt.Recovering("QSslConfiguration::setPeerVerifyDepth")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	defer qt.Recovering("QSslConfiguration::setPeerVerifyMode")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	defer qt.Recovering("QSslConfiguration::setPrivateKey")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket string) {
	defer qt.Recovering("QSslConfiguration::setSessionTicket")

	if ptr.Pointer() != nil {
		var sessionTicketC = C.CString(hex.EncodeToString([]byte(sessionTicket)))
		defer C.free(unsafe.Pointer(sessionTicketC))
		C.QSslConfiguration_SetSessionTicket(ptr.Pointer(), sessionTicketC)
	}
}

func (ptr *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	defer qt.Recovering("QSslConfiguration::swap")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_Swap(ptr.Pointer(), PointerFromQSslConfiguration(other))
	}
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	defer qt.Recovering("QSslConfiguration::~QSslConfiguration")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_DestroyQSslConfiguration(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslEllipticCurve struct {
	ptr unsafe.Pointer
}

type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (p *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve {
	return p
}

func (p *QSslEllipticCurve) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslEllipticCurve) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurve_PTR().Pointer()
	}
	return nil
}

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslEllipticCurve) DestroyQSslEllipticCurve() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSslEllipticCurve() *QSslEllipticCurve {
	defer qt.Recovering("QSslEllipticCurve::QSslEllipticCurve")

	var tmpValue = NewQSslEllipticCurveFromPointer(C.QSslEllipticCurve_NewQSslEllipticCurve())
	runtime.SetFinalizer(tmpValue, (*QSslEllipticCurve).DestroyQSslEllipticCurve)
	return tmpValue
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	defer qt.Recovering("QSslEllipticCurve::isValid")

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	defer qt.Recovering("QSslEllipticCurve::isTlsNamedCurve")

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsTlsNamedCurve(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) LongName() string {
	defer qt.Recovering("QSslEllipticCurve::longName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_LongName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslEllipticCurve) ShortName() string {
	defer qt.Recovering("QSslEllipticCurve::shortName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_ShortName(ptr.Pointer()))
	}
	return ""
}

//QSslError::SslError
type QSslError__SslError int64

const (
	QSslError__NoError                             = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              = QSslError__SslError(5)
	QSslError__CertificateExpired                  = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      = QSslError__SslError(12)
	QSslError__CertificateRevoked                  = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  = QSslError__SslError(15)
	QSslError__InvalidPurpose                      = QSslError__SslError(16)
	QSslError__CertificateUntrusted                = QSslError__SslError(17)
	QSslError__CertificateRejected                 = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   = QSslError__SslError(21)
	QSslError__HostNameMismatch                    = QSslError__SslError(22)
	QSslError__NoSslSupport                        = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              = QSslError__SslError(24)
	QSslError__UnspecifiedError                    = QSslError__SslError(-1)
)

type QSslError struct {
	ptr unsafe.Pointer
}

type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (p *QSslError) QSslError_PTR() *QSslError {
	return p
}

func (p *QSslError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslError(ptr QSslError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslError_PTR().Pointer()
	}
	return nil
}

func NewQSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = new(QSslError)
	n.SetPointer(ptr)
	return n
}
func NewQSslError() *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError())
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError2(C.longlong(error)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError3(C.longlong(error), PointerFromQSslCertificate(certificate)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func NewQSslError4(other QSslError_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	var tmpValue = NewQSslErrorFromPointer(C.QSslError_NewQSslError4(PointerFromQSslError(other)))
	runtime.SetFinalizer(tmpValue, (*QSslError).DestroyQSslError)
	return tmpValue
}

func (ptr *QSslError) Certificate() *QSslCertificate {
	defer qt.Recovering("QSslError::certificate")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslError_Certificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslError) Error() QSslError__SslError {
	defer qt.Recovering("QSslError::error")

	if ptr.Pointer() != nil {
		return QSslError__SslError(C.QSslError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslError) ErrorString() string {
	defer qt.Recovering("QSslError::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslError_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslError) Swap(other QSslError_ITF) {
	defer qt.Recovering("QSslError::swap")

	if ptr.Pointer() != nil {
		C.QSslError_Swap(ptr.Pointer(), PointerFromQSslError(other))
	}
}

func (ptr *QSslError) DestroyQSslError() {
	defer qt.Recovering("QSslError::~QSslError")

	if ptr.Pointer() != nil {
		C.QSslError_DestroyQSslError(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslKey struct {
	ptr unsafe.Pointer
}

type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (p *QSslKey) QSslKey_PTR() *QSslKey {
	return p
}

func (p *QSslKey) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslKey) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslKey(ptr QSslKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKey_PTR().Pointer()
	}
	return nil
}

func NewQSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = new(QSslKey)
	n.SetPointer(ptr)
	return n
}
func NewQSslKey() *QSslKey {
	defer qt.Recovering("QSslKey::QSslKey")

	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey())
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	defer qt.Recovering("QSslKey::QSslKey")

	var tmpValue = NewQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
	runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
	return tmpValue
}

func (ptr *QSslKey) Clear() {
	defer qt.Recovering("QSslKey::clear")

	if ptr.Pointer() != nil {
		C.QSslKey_Clear(ptr.Pointer())
	}
}

func (ptr *QSslKey) IsNull() bool {
	defer qt.Recovering("QSslKey::isNull")

	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	defer qt.Recovering("QSslKey::length")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslKey_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {
	defer qt.Recovering("QSslKey::swap")

	if ptr.Pointer() != nil {
		C.QSslKey_Swap(ptr.Pointer(), PointerFromQSslKey(other))
	}
}

func (ptr *QSslKey) ToDer(passPhrase string) string {
	defer qt.Recovering("QSslKey::toDer")

	if ptr.Pointer() != nil {
		var passPhraseC = C.CString(hex.EncodeToString([]byte(passPhrase)))
		defer C.free(unsafe.Pointer(passPhraseC))
		return qt.HexDecodeToString(C.GoString(C.QSslKey_ToDer(ptr.Pointer(), passPhraseC)))
	}
	return ""
}

func (ptr *QSslKey) ToPem(passPhrase string) string {
	defer qt.Recovering("QSslKey::toPem")

	if ptr.Pointer() != nil {
		var passPhraseC = C.CString(hex.EncodeToString([]byte(passPhrase)))
		defer C.free(unsafe.Pointer(passPhraseC))
		return qt.HexDecodeToString(C.GoString(C.QSslKey_ToPem(ptr.Pointer(), passPhraseC)))
	}
	return ""
}

func (ptr *QSslKey) DestroyQSslKey() {
	defer qt.Recovering("QSslKey::~QSslKey")

	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSslPreSharedKeyAuthenticator struct {
	ptr unsafe.Pointer
}

type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (p *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return p
}

func (p *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSslPreSharedKeyAuthenticator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return n
}
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::QSslPreSharedKeyAuthenticator")

	var tmpValue = NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator())
	runtime.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::QSslPreSharedKeyAuthenticator")

	var tmpValue = NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
	runtime.SetFinalizer(tmpValue, (*QSslPreSharedKeyAuthenticator).DestroyQSslPreSharedKeyAuthenticator)
	return tmpValue
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() string {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::identity")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslPreSharedKeyAuthenticator_Identity(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() string {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::identityHint")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslPreSharedKeyAuthenticator_IdentityHint(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::maximumIdentityLength")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslPreSharedKeyAuthenticator_MaximumIdentityLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::maximumPreSharedKeyLength")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) PreSharedKey() string {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::preSharedKey")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QSslPreSharedKeyAuthenticator_PreSharedKey(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity string) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::setIdentity")

	if ptr.Pointer() != nil {
		var identityC = C.CString(hex.EncodeToString([]byte(identity)))
		defer C.free(unsafe.Pointer(identityC))
		C.QSslPreSharedKeyAuthenticator_SetIdentity(ptr.Pointer(), identityC)
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey string) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::setPreSharedKey")

	if ptr.Pointer() != nil {
		var preSharedKeyC = C.CString(hex.EncodeToString([]byte(preSharedKey)))
		defer C.free(unsafe.Pointer(preSharedKeyC))
		C.QSslPreSharedKeyAuthenticator_SetPreSharedKey(ptr.Pointer(), preSharedKeyC)
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::swap")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_Swap(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::~QSslPreSharedKeyAuthenticator")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer = QSslSocket__PeerVerifyMode(3)
)

//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   = QSslSocket__SslMode(2)
)

type QSslSocket struct {
	QTcpSocket
}

type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func (p *QSslSocket) QSslSocket_PTR() *QSslSocket {
	return p
}

func (p *QSslSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (p *QSslSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QTcpSocket_PTR().SetPointer(ptr)
	}
}

func PointerFromQSslSocket(ptr QSslSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslSocket_PTR().Pointer()
	}
	return nil
}

func NewQSslSocketFromPointer(ptr unsafe.Pointer) *QSslSocket {
	var n = new(QSslSocket)
	n.SetPointer(ptr)
	return n
}
func NewQSslSocket(parent core.QObject_ITF) *QSslSocket {
	defer qt.Recovering("QSslSocket::QSslSocket")

	var tmpValue = NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSslSocket) Abort() {
	defer qt.Recovering("QSslSocket::abort")

	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::addCaCertificate")

	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::addDefaultCaCertificate")

	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::addDefaultCaCertificate")

	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AtEnd() bool {
	defer qt.Recovering("QSslSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) BytesAvailable() int64 {
	defer qt.Recovering("QSslSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) BytesToWrite() int64 {
	defer qt.Recovering("QSslSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) CanReadLine() bool {
	defer qt.Recovering("QSslSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Close() {
	defer qt.Recovering("QSslSocket::close")

	if ptr.Pointer() != nil {
		C.QSslSocket_Close(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QSslSocket::connectToHostEncrypted")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHostEncrypted(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(mode), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) ConnectToHostEncrypted2(hostName string, port uint16, sslPeerName string, mode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QSslSocket::connectToHostEncrypted")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		var sslPeerNameC = C.CString(sslPeerName)
		defer C.free(unsafe.Pointer(sslPeerNameC))
		C.QSslSocket_ConnectToHostEncrypted2(ptr.Pointer(), hostNameC, C.ushort(port), sslPeerNameC, C.longlong(mode), C.longlong(protocol))
	}
}

//export callbackQSslSocket_Encrypted
func callbackQSslSocket_Encrypted(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::encrypted")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	defer qt.Recovering("connect QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encrypted")
	}
}

func (ptr *QSslSocket) Encrypted() {
	defer qt.Recovering("QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_Encrypted(ptr.Pointer())
	}
}

func (ptr *QSslSocket) EncryptedBytesAvailable() int64 {
	defer qt.Recovering("QSslSocket::encryptedBytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) EncryptedBytesToWrite() int64 {
	defer qt.Recovering("QSslSocket::encryptedBytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesToWrite(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_EncryptedBytesWritten
func callbackQSslSocket_EncryptedBytesWritten(ptr unsafe.Pointer, written C.longlong) {
	defer qt.Recovering("callback QSslSocket::encryptedBytesWritten")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::encryptedBytesWritten"); signal != nil {
		signal.(func(int64))(int64(written))
	}

}

func (ptr *QSslSocket) ConnectEncryptedBytesWritten(f func(written int64)) {
	defer qt.Recovering("connect QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncryptedBytesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encryptedBytesWritten", f)
	}
}

func (ptr *QSslSocket) DisconnectEncryptedBytesWritten() {
	defer qt.Recovering("disconnect QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncryptedBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::encryptedBytesWritten")
	}
}

func (ptr *QSslSocket) EncryptedBytesWritten(written int64) {
	defer qt.Recovering("QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_EncryptedBytesWritten(ptr.Pointer(), C.longlong(written))
	}
}

func (ptr *QSslSocket) Flush() bool {
	defer qt.Recovering("QSslSocket::flush")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_IgnoreSslErrors
func callbackQSslSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::ignoreSslErrors")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::ignoreSslErrors"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) ConnectIgnoreSslErrors(f func()) {
	defer qt.Recovering("connect QSslSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::ignoreSslErrors", f)
	}
}

func (ptr *QSslSocket) DisconnectIgnoreSslErrors() {
	defer qt.Recovering("disconnect QSslSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::ignoreSslErrors")
	}
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	defer qt.Recovering("QSslSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	defer qt.Recovering("QSslSocket::isEncrypted")

	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) LocalCertificate() *QSslCertificate {
	defer qt.Recovering("QSslSocket::localCertificate")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_LocalCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	defer qt.Recovering("QSslSocket::mode")

	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_ModeChanged
func callbackQSslSocket_ModeChanged(ptr unsafe.Pointer, mode C.longlong) {
	defer qt.Recovering("callback QSslSocket::modeChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::modeChanged"); signal != nil {
		signal.(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
	}

}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	defer qt.Recovering("connect QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	defer qt.Recovering("disconnect QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::modeChanged")
	}
}

func (ptr *QSslSocket) ModeChanged(mode QSslSocket__SslMode) {
	defer qt.Recovering("QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_ModeChanged(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslSocket) PeerCertificate() *QSslCertificate {
	defer qt.Recovering("QSslSocket::peerCertificate")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCertificateFromPointer(C.QSslSocket_PeerCertificate(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCertificate).DestroyQSslCertificate)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	defer qt.Recovering("QSslSocket::peerVerifyDepth")

	if ptr.Pointer() != nil {
		return int(int32(C.QSslSocket_PeerVerifyDepth(ptr.Pointer())))
	}
	return 0
}

//export callbackQSslSocket_PeerVerifyError
func callbackQSslSocket_PeerVerifyError(ptr unsafe.Pointer, error unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::peerVerifyError")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::peerVerifyError"); signal != nil {
		signal.(func(*QSslError))(NewQSslErrorFromPointer(error))
	}

}

func (ptr *QSslSocket) ConnectPeerVerifyError(f func(error *QSslError)) {
	defer qt.Recovering("connect QSslSocket::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPeerVerifyError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::peerVerifyError", f)
	}
}

func (ptr *QSslSocket) DisconnectPeerVerifyError() {
	defer qt.Recovering("disconnect QSslSocket::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::peerVerifyError")
	}
}

func (ptr *QSslSocket) PeerVerifyError(error QSslError_ITF) {
	defer qt.Recovering("QSslSocket::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QSslSocket_PeerVerifyError(ptr.Pointer(), PointerFromQSslError(error))
	}
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	defer qt.Recovering("QSslSocket::peerVerifyMode")

	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	defer qt.Recovering("QSslSocket::peerVerifyName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

//export callbackQSslSocket_PreSharedKeyAuthenticationRequired
func callbackQSslSocket_PreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::preSharedKeyAuthenticationRequired")
	}
}

func (ptr *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslSocket) PrivateKey() *QSslKey {
	defer qt.Recovering("QSslSocket::privateKey")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslKeyFromPointer(C.QSslSocket_PrivateKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslKey).DestroyQSslKey)
		return tmpValue
	}
	return nil
}

//export callbackQSslSocket_Resume
func callbackQSslSocket_Resume(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::resume")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QSslSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QSslSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::resume", f)
	}
}

func (ptr *QSslSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QSslSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::resume")
	}
}

func (ptr *QSslSocket) Resume() {
	defer qt.Recovering("QSslSocket::resume")

	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ResumeDefault() {
	defer qt.Recovering("QSslSocket::resume")

	if ptr.Pointer() != nil {
		C.QSslSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) SessionCipher() *QSslCipher {
	defer qt.Recovering("QSslSocket::sessionCipher")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslCipherFromPointer(C.QSslSocket_SessionCipher(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslCipher).DestroyQSslCipher)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::setLocalCertificate")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	defer qt.Recovering("QSslSocket::setPeerVerifyDepth")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(ptr.Pointer(), C.int(int32(depth)))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	defer qt.Recovering("QSslSocket::setPeerVerifyMode")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	defer qt.Recovering("QSslSocket::setPeerVerifyName")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), hostNameC)
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	defer qt.Recovering("QSslSocket::setPrivateKey")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

//export callbackQSslSocket_SetReadBufferSize
func callbackQSslSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QSslSocket::setReadBufferSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQSslSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QSslSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setReadBufferSize", f)
	}
}

func (ptr *QSslSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setReadBufferSize")
	}
}

func (ptr *QSslSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QSslSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQSslSocket_SetSocketOption
func callbackQSslSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::setSocketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQSslSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QSslSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setSocketOption", f)
	}
}

func (ptr *QSslSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::setSocketOption")
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QSslSocket::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

//export callbackQSslSocket_SocketOption
func callbackQSslSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QSslSocket::socketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQSslSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QSslSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	defer qt.Recovering("connect QSslSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::socketOption", f)
	}
}

func (ptr *QSslSocket) DisconnectSocketOption() {
	defer qt.Recovering("disconnect QSslSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::socketOption")
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QSslSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QSslSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSslSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSslSocket) SslConfiguration() *QSslConfiguration {
	defer qt.Recovering("QSslSocket::sslConfiguration")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSslConfigurationFromPointer(C.QSslSocket_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func QSslSocket_SslLibraryBuildVersionNumber() int {
	defer qt.Recovering("QSslSocket::sslLibraryBuildVersionNumber")

	return int(int32(C.QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()))
}

func (ptr *QSslSocket) SslLibraryBuildVersionNumber() int {
	defer qt.Recovering("QSslSocket::sslLibraryBuildVersionNumber")

	return int(int32(C.QSslSocket_QSslSocket_SslLibraryBuildVersionNumber()))
}

func QSslSocket_SslLibraryBuildVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryBuildVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func (ptr *QSslSocket) SslLibraryBuildVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryBuildVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionNumber() int {
	defer qt.Recovering("QSslSocket::sslLibraryVersionNumber")

	return int(int32(C.QSslSocket_QSslSocket_SslLibraryVersionNumber()))
}

func (ptr *QSslSocket) SslLibraryVersionNumber() int {
	defer qt.Recovering("QSslSocket::sslLibraryVersionNumber")

	return int(int32(C.QSslSocket_QSslSocket_SslLibraryVersionNumber()))
}

func QSslSocket_SslLibraryVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) SslLibraryVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

//export callbackQSslSocket_StartClientEncryption
func callbackQSslSocket_StartClientEncryption(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::startClientEncryption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::startClientEncryption"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) ConnectStartClientEncryption(f func()) {
	defer qt.Recovering("connect QSslSocket::startClientEncryption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startClientEncryption", f)
	}
}

func (ptr *QSslSocket) DisconnectStartClientEncryption() {
	defer qt.Recovering("disconnect QSslSocket::startClientEncryption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startClientEncryption")
	}
}

func (ptr *QSslSocket) StartClientEncryption() {
	defer qt.Recovering("QSslSocket::startClientEncryption")

	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(ptr.Pointer())
	}
}

//export callbackQSslSocket_StartServerEncryption
func callbackQSslSocket_StartServerEncryption(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::startServerEncryption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::startServerEncryption"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) ConnectStartServerEncryption(f func()) {
	defer qt.Recovering("connect QSslSocket::startServerEncryption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startServerEncryption", f)
	}
}

func (ptr *QSslSocket) DisconnectStartServerEncryption() {
	defer qt.Recovering("disconnect QSslSocket::startServerEncryption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::startServerEncryption")
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	defer qt.Recovering("QSslSocket::startServerEncryption")

	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(ptr.Pointer())
	}
}

func QSslSocket_SupportsSsl() bool {
	defer qt.Recovering("QSslSocket::supportsSsl")

	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) SupportsSsl() bool {
	defer qt.Recovering("QSslSocket::supportsSsl")

	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WaitForConnected
func callbackQSslSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QSslSocket::waitForConnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QSslSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForConnected", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForConnected() {
	defer qt.Recovering("disconnect QSslSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForConnected")
	}
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnectedDefault(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQSslSocket_WaitForDisconnected
func callbackQSslSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QSslSocket::waitForDisconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QSslSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QSslSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForDisconnected", f)
	}
}

func (ptr *QSslSocket) DisconnectWaitForDisconnected() {
	defer qt.Recovering("disconnect QSslSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::waitForDisconnected")
	}
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnectedDefault(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForEncrypted")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QSslSocket) WriteData(data string, len int64) int64 {
	defer qt.Recovering("QSslSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QSslSocket_WriteData(ptr.Pointer(), dataC, C.longlong(len)))
	}
	return 0
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	defer qt.Recovering("QSslSocket::~QSslSocket")

	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSslSocket_ConnectToHost2
func callbackQSslSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	defer qt.Recovering("callback QSslSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QSslSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	defer qt.Recovering("connect QSslSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost2", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectToHost2() {
	defer qt.Recovering("disconnect QSslSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost2")
	}
}

func (ptr *QSslSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QSslSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QSslSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QSslSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQSslSocket_ConnectToHost
func callbackQSslSocket_ConnectToHost(ptr unsafe.Pointer, hostName *C.char, port C.ushort, openMode C.longlong, protocol C.longlong) {
	defer qt.Recovering("callback QSslSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectToHostDefault(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QSslSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	defer qt.Recovering("connect QSslSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectToHost() {
	defer qt.Recovering("disconnect QSslSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectToHost")
	}
}

func (ptr *QSslSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QSslSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QSslSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QSslSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QSslSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQSslSocket_DisconnectFromHost
func callbackQSslSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::disconnectFromHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QSslSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectFromHost", f)
	}
}

func (ptr *QSslSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectFromHost")
	}
}

func (ptr *QSslSocket) DisconnectFromHost() {
	defer qt.Recovering("QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QSslSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQSslSocket_Open
func callbackQSslSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QSslSocket::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QSslSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QSslSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::open", f)
	}
}

func (ptr *QSslSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QSslSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::open")
	}
}

func (ptr *QSslSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QSslSocket::open")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QSslSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QSslSocket::open")

	if ptr.Pointer() != nil {
		return C.QSslSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQSslSocket_Pos
func callbackQSslSocket_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSslSocket::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).PosDefault())
}

func (ptr *QSslSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QSslSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::pos", f)
	}
}

func (ptr *QSslSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QSslSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::pos")
	}
}

func (ptr *QSslSocket) Pos() int64 {
	defer qt.Recovering("QSslSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PosDefault() int64 {
	defer qt.Recovering("QSslSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_Reset
func callbackQSslSocket_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSslSocket::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QSslSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QSslSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::reset", f)
	}
}

func (ptr *QSslSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QSslSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::reset")
	}
}

func (ptr *QSslSocket) Reset() bool {
	defer qt.Recovering("QSslSocket::reset")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) ResetDefault() bool {
	defer qt.Recovering("QSslSocket::reset")

	if ptr.Pointer() != nil {
		return C.QSslSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSslSocket_Seek
func callbackQSslSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QSslSocket::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QSslSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QSslSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::seek", f)
	}
}

func (ptr *QSslSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QSslSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::seek")
	}
}

func (ptr *QSslSocket) Seek(pos int64) bool {
	defer qt.Recovering("QSslSocket::seek")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QSslSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QSslSocket::seek")

	if ptr.Pointer() != nil {
		return C.QSslSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQSslSocket_Size
func callbackQSslSocket_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSslSocket::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQSslSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QSslSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QSslSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::size", f)
	}
}

func (ptr *QSslSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QSslSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::size")
	}
}

func (ptr *QSslSocket) Size() int64 {
	defer qt.Recovering("QSslSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) SizeDefault() int64 {
	defer qt.Recovering("QSslSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSslSocket_TimerEvent
func callbackQSslSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSslSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::timerEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSslSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::timerEvent")
	}
}

func (ptr *QSslSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSslSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSslSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSslSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSslSocket_ChildEvent
func callbackQSslSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSslSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::childEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSslSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::childEvent")
	}
}

func (ptr *QSslSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSslSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSslSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSslSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSslSocket_ConnectNotify
func callbackQSslSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSslSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSslSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSslSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectNotify", f)
	}
}

func (ptr *QSslSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSslSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::connectNotify")
	}
}

func (ptr *QSslSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSslSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSslSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSslSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSslSocket_CustomEvent
func callbackQSslSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSslSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::customEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSslSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::customEvent")
	}
}

func (ptr *QSslSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSslSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSslSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSslSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSslSocket_DeleteLater
func callbackQSslSocket_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSslSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSslSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::deleteLater", f)
	}
}

func (ptr *QSslSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSslSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::deleteLater")
	}
}

func (ptr *QSslSocket) DeleteLater() {
	defer qt.Recovering("QSslSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QSslSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) DeleteLaterDefault() {
	defer qt.Recovering("QSslSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QSslSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSslSocket_DisconnectNotify
func callbackQSslSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSslSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSslSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSslSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectNotify", f)
	}
}

func (ptr *QSslSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSslSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::disconnectNotify")
	}
}

func (ptr *QSslSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSslSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSslSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSslSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSslSocket_Event
func callbackQSslSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSslSocket::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSslSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSslSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::event", f)
	}
}

func (ptr *QSslSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QSslSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::event")
	}
}

func (ptr *QSslSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSslSocket::event")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSslSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSslSocket::event")

	if ptr.Pointer() != nil {
		return C.QSslSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSslSocket_EventFilter
func callbackQSslSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSslSocket::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSslSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSslSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSslSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::eventFilter", f)
	}
}

func (ptr *QSslSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSslSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::eventFilter")
	}
}

func (ptr *QSslSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSslSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSslSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSslSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSslSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSslSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSslSocket_MetaObject
func callbackQSslSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSslSocket::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSslSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSslSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSslSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSslSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::metaObject", f)
	}
}

func (ptr *QSslSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSslSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSslSocket::metaObject")
	}
}

func (ptr *QSslSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSslSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSslSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSslSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSslSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTcpServer struct {
	core.QObject
}

type QTcpServer_ITF interface {
	core.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func (p *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return p
}

func (p *QTcpServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QTcpServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQTcpServer(ptr QTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func NewQTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = new(QTcpServer)
	n.SetPointer(ptr)
	return n
}
func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	defer qt.Recovering("QTcpServer::QTcpServer")

	var tmpValue = NewQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpServer_AcceptError
func callbackQTcpServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {
	defer qt.Recovering("callback QTcpServer::acceptError")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::acceptError"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	defer qt.Recovering("disconnect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::acceptError")
	}
}

func (ptr *QTcpServer) AcceptError(socketError QAbstractSocket__SocketError) {
	defer qt.Recovering("QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_AcceptError(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QTcpServer) AddPendingConnection(socket QTcpSocket_ITF) {
	defer qt.Recovering("QTcpServer::addPendingConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_AddPendingConnection(ptr.Pointer(), PointerFromQTcpSocket(socket))
	}
}

func (ptr *QTcpServer) Close() {
	defer qt.Recovering("QTcpServer::close")

	if ptr.Pointer() != nil {
		C.QTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ErrorString() string {
	defer qt.Recovering("QTcpServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTcpServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQTcpServer_HasPendingConnections
func callbackQTcpServer_HasPendingConnections(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpServer::hasPendingConnections")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::hasPendingConnections"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).HasPendingConnectionsDefault())))
}

func (ptr *QTcpServer) ConnectHasPendingConnections(f func() bool) {
	defer qt.Recovering("connect QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::hasPendingConnections", f)
	}
}

func (ptr *QTcpServer) DisconnectHasPendingConnections() {
	defer qt.Recovering("disconnect QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::hasPendingConnections")
	}
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	defer qt.Recovering("QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) HasPendingConnectionsDefault() bool {
	defer qt.Recovering("QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnectionsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	defer qt.Recovering("QTcpServer::isListening")

	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) Listen(address QHostAddress_ITF, port uint16) bool {
	defer qt.Recovering("QTcpServer::listen")

	if ptr.Pointer() != nil {
		return C.QTcpServer_Listen(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	defer qt.Recovering("QTcpServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(int32(C.QTcpServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQTcpServer_NewConnection
func callbackQTcpServer_NewConnection(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::newConnection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::newConnection")
	}
}

func (ptr *QTcpServer) NewConnection() {
	defer qt.Recovering("QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQTcpServer_NextPendingConnection
func callbackQTcpServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QTcpServer::nextPendingConnection")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::nextPendingConnection"); signal != nil {
		return PointerFromQTcpSocket(signal.(func() *QTcpSocket)())
	}

	return PointerFromQTcpSocket(NewQTcpServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QTcpServer) ConnectNextPendingConnection(f func() *QTcpSocket) {
	defer qt.Recovering("connect QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::nextPendingConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNextPendingConnection() {
	defer qt.Recovering("disconnect QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::nextPendingConnection")
	}
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	defer qt.Recovering("QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) NextPendingConnectionDefault() *QTcpSocket {
	defer qt.Recovering("QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		var tmpValue = NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	defer qt.Recovering("QTcpServer::pauseAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) Proxy() *QNetworkProxy {
	defer qt.Recovering("QTcpServer::proxy")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkProxyFromPointer(C.QTcpServer_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) ResumeAccepting() {
	defer qt.Recovering("QTcpServer::resumeAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ServerAddress() *QHostAddress {
	defer qt.Recovering("QTcpServer::serverAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQHostAddressFromPointer(C.QTcpServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QTcpServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) ServerPort() uint16 {
	defer qt.Recovering("QTcpServer::serverPort")

	if ptr.Pointer() != nil {
		return uint16(C.QTcpServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QTcpServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QTcpServer::setProxy")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QTcpServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(int32(msec)), C.char(int8(qt.GoBoolToInt(timedOut)))) != 0
	}
	return false
}

//export callbackQTcpServer_DestroyQTcpServer
func callbackQTcpServer_DestroyQTcpServer(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::~QTcpServer")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::~QTcpServer"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpServerFromPointer(ptr).DestroyQTcpServerDefault()
	}
}

func (ptr *QTcpServer) ConnectDestroyQTcpServer(f func()) {
	defer qt.Recovering("connect QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::~QTcpServer", f)
	}
}

func (ptr *QTcpServer) DisconnectDestroyQTcpServer() {
	defer qt.Recovering("disconnect QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::~QTcpServer")
	}
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	defer qt.Recovering("QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) DestroyQTcpServerDefault() {
	defer qt.Recovering("QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpServer_TimerEvent
func callbackQTcpServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::timerEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::timerEvent")
	}
}

func (ptr *QTcpServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTcpServer_ChildEvent
func callbackQTcpServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::childEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::childEvent")
	}
}

func (ptr *QTcpServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTcpServer_ConnectNotify
func callbackQTcpServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QTcpServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::connectNotify", f)
	}
}

func (ptr *QTcpServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QTcpServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::connectNotify")
	}
}

func (ptr *QTcpServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_CustomEvent
func callbackQTcpServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::customEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::customEvent")
	}
}

func (ptr *QTcpServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTcpServer_DeleteLater
func callbackQTcpServer_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTcpServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QTcpServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::deleteLater", f)
	}
}

func (ptr *QTcpServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QTcpServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::deleteLater")
	}
}

func (ptr *QTcpServer) DeleteLater() {
	defer qt.Recovering("QTcpServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QTcpServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) DeleteLaterDefault() {
	defer qt.Recovering("QTcpServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QTcpServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpServer_DisconnectNotify
func callbackQTcpServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::disconnectNotify", f)
	}
}

func (ptr *QTcpServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::disconnectNotify")
	}
}

func (ptr *QTcpServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpServer_Event
func callbackQTcpServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpServer::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTcpServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QTcpServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::event", f)
	}
}

func (ptr *QTcpServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QTcpServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::event")
	}
}

func (ptr *QTcpServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpServer::event")

	if ptr.Pointer() != nil {
		return C.QTcpServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTcpServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpServer::event")

	if ptr.Pointer() != nil {
		return C.QTcpServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTcpServer_EventFilter
func callbackQTcpServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpServer::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTcpServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QTcpServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::eventFilter", f)
	}
}

func (ptr *QTcpServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QTcpServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::eventFilter")
	}
}

func (ptr *QTcpServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QTcpServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QTcpServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTcpServer_MetaObject
func callbackQTcpServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QTcpServer::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTcpServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTcpServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QTcpServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::metaObject", f)
	}
}

func (ptr *QTcpServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QTcpServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpServer::metaObject")
	}
}

func (ptr *QTcpServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QTcpServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QTcpServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func (p *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return p
}

func (p *QTcpSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (p *QTcpSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractSocket_PTR().SetPointer(ptr)
	}
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	return n
}
func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	defer qt.Recovering("QTcpSocket::QTcpSocket")

	var tmpValue = NewQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTcpSocket_DestroyQTcpSocket
func callbackQTcpSocket_DestroyQTcpSocket(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::~QTcpSocket")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::~QTcpSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DestroyQTcpSocketDefault()
	}
}

func (ptr *QTcpSocket) ConnectDestroyQTcpSocket(f func()) {
	defer qt.Recovering("connect QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::~QTcpSocket", f)
	}
}

func (ptr *QTcpSocket) DisconnectDestroyQTcpSocket() {
	defer qt.Recovering("disconnect QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::~QTcpSocket")
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	defer qt.Recovering("QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) DestroyQTcpSocketDefault() {
	defer qt.Recovering("QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpSocket_ConnectToHost2
func callbackQTcpSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	defer qt.Recovering("callback QTcpSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QTcpSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	defer qt.Recovering("connect QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost2", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectToHost2() {
	defer qt.Recovering("disconnect QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost2")
	}
}

func (ptr *QTcpSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QTcpSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQTcpSocket_ConnectToHost
func callbackQTcpSocket_ConnectToHost(ptr unsafe.Pointer, hostName *C.char, port C.ushort, openMode C.longlong, protocol C.longlong) {
	defer qt.Recovering("callback QTcpSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectToHostDefault(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QTcpSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	defer qt.Recovering("connect QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectToHost() {
	defer qt.Recovering("disconnect QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectToHost")
	}
}

func (ptr *QTcpSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QTcpSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QTcpSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QTcpSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QTcpSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQTcpSocket_DisconnectFromHost
func callbackQTcpSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::disconnectFromHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QTcpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectFromHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectFromHost")
	}
}

func (ptr *QTcpSocket) DisconnectFromHost() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQTcpSocket_Resume
func callbackQTcpSocket_Resume(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::resume")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTcpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::resume", f)
	}
}

func (ptr *QTcpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::resume")
	}
}

func (ptr *QTcpSocket) Resume() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ResumeDefault() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQTcpSocket_SetReadBufferSize
func callbackQTcpSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QTcpSocket::setReadBufferSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQTcpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QTcpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setReadBufferSize", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setReadBufferSize")
	}
}

func (ptr *QTcpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQTcpSocket_SetSocketOption
func callbackQTcpSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::setSocketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQTcpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QTcpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setSocketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::setSocketOption")
	}
}

func (ptr *QTcpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

//export callbackQTcpSocket_SocketOption
func callbackQTcpSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QTcpSocket::socketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQTcpSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QTcpSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	defer qt.Recovering("connect QTcpSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::socketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSocketOption() {
	defer qt.Recovering("disconnect QTcpSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::socketOption")
	}
}

func (ptr *QTcpSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QTcpSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QTcpSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTcpSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QTcpSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QTcpSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQTcpSocket_WaitForConnected
func callbackQTcpSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QTcpSocket::waitForConnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QTcpSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForConnected", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForConnected() {
	defer qt.Recovering("disconnect QTcpSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForConnected")
	}
}

func (ptr *QTcpSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QTcpSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForConnectedDefault(msecs int) bool {
	defer qt.Recovering("QTcpSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_WaitForDisconnected
func callbackQTcpSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QTcpSocket::waitForDisconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QTcpSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QTcpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForDisconnected", f)
	}
}

func (ptr *QTcpSocket) DisconnectWaitForDisconnected() {
	defer qt.Recovering("disconnect QTcpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::waitForDisconnected")
	}
}

func (ptr *QTcpSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QTcpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QTcpSocket) WaitForDisconnectedDefault(msecs int) bool {
	defer qt.Recovering("QTcpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQTcpSocket_Open
func callbackQTcpSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QTcpSocket::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QTcpSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QTcpSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::open", f)
	}
}

func (ptr *QTcpSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QTcpSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::open")
	}
}

func (ptr *QTcpSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QTcpSocket::open")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QTcpSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QTcpSocket::open")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQTcpSocket_Pos
func callbackQTcpSocket_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QTcpSocket::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).PosDefault())
}

func (ptr *QTcpSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QTcpSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::pos", f)
	}
}

func (ptr *QTcpSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QTcpSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::pos")
	}
}

func (ptr *QTcpSocket) Pos() int64 {
	defer qt.Recovering("QTcpSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) PosDefault() int64 {
	defer qt.Recovering("QTcpSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_Reset
func callbackQTcpSocket_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpSocket::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QTcpSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QTcpSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::reset", f)
	}
}

func (ptr *QTcpSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QTcpSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::reset")
	}
}

func (ptr *QTcpSocket) Reset() bool {
	defer qt.Recovering("QTcpSocket::reset")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpSocket) ResetDefault() bool {
	defer qt.Recovering("QTcpSocket::reset")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQTcpSocket_Seek
func callbackQTcpSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QTcpSocket::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QTcpSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QTcpSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::seek", f)
	}
}

func (ptr *QTcpSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QTcpSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::seek")
	}
}

func (ptr *QTcpSocket) Seek(pos int64) bool {
	defer qt.Recovering("QTcpSocket::seek")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QTcpSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QTcpSocket::seek")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQTcpSocket_Size
func callbackQTcpSocket_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QTcpSocket::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQTcpSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QTcpSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QTcpSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::size", f)
	}
}

func (ptr *QTcpSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QTcpSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::size")
	}
}

func (ptr *QTcpSocket) Size() int64 {
	defer qt.Recovering("QTcpSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpSocket) SizeDefault() int64 {
	defer qt.Recovering("QTcpSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QTcpSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQTcpSocket_TimerEvent
func callbackQTcpSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::timerEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::timerEvent")
	}
}

func (ptr *QTcpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQTcpSocket_ChildEvent
func callbackQTcpSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::childEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::childEvent")
	}
}

func (ptr *QTcpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTcpSocket_ConnectNotify
func callbackQTcpSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QTcpSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectNotify", f)
	}
}

func (ptr *QTcpSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QTcpSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::connectNotify")
	}
}

func (ptr *QTcpSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpSocket_CustomEvent
func callbackQTcpSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::customEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::customEvent")
	}
}

func (ptr *QTcpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTcpSocket_DeleteLater
func callbackQTcpSocket_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTcpSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QTcpSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::deleteLater", f)
	}
}

func (ptr *QTcpSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QTcpSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::deleteLater")
	}
}

func (ptr *QTcpSocket) DeleteLater() {
	defer qt.Recovering("QTcpSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) DeleteLaterDefault() {
	defer qt.Recovering("QTcpSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQTcpSocket_DisconnectNotify
func callbackQTcpSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTcpSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QTcpSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectNotify", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QTcpSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::disconnectNotify")
	}
}

func (ptr *QTcpSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QTcpSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QTcpSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTcpSocket_Event
func callbackQTcpSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpSocket::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTcpSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QTcpSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::event", f)
	}
}

func (ptr *QTcpSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QTcpSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::event")
	}
}

func (ptr *QTcpSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpSocket::event")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QTcpSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpSocket::event")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQTcpSocket_EventFilter
func callbackQTcpSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QTcpSocket::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTcpSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTcpSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QTcpSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::eventFilter", f)
	}
}

func (ptr *QTcpSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QTcpSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::eventFilter")
	}
}

func (ptr *QTcpSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QTcpSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QTcpSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QTcpSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQTcpSocket_MetaObject
func callbackQTcpSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QTcpSocket::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QTcpSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTcpSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTcpSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QTcpSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::metaObject", f)
	}
}

func (ptr *QTcpSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QTcpSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QTcpSocket::metaObject")
	}
}

func (ptr *QTcpSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QTcpSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QTcpSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTcpSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func (p *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return p
}

func (p *QUdpSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (p *QUdpSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractSocket_PTR().SetPointer(ptr)
	}
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	return n
}
func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	defer qt.Recovering("QUdpSocket::QUdpSocket")

	var tmpValue = NewQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	defer qt.Recovering("QUdpSocket::hasPendingDatagrams")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) MulticastInterface() *QNetworkInterface {
	defer qt.Recovering("QUdpSocket::multicastInterface")

	if ptr.Pointer() != nil {
		var tmpValue = NewQNetworkInterfaceFromPointer(C.QUdpSocket_MulticastInterface(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QNetworkInterface).DestroyQNetworkInterface)
		return tmpValue
	}
	return nil
}

func (ptr *QUdpSocket) PendingDatagramSize() int64 {
	defer qt.Recovering("QUdpSocket::pendingDatagramSize")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PendingDatagramSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) ReadDatagram(data string, maxSize int64, address QHostAddress_ITF, port uint16) int64 {
	defer qt.Recovering("QUdpSocket::readDatagram")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_ReadDatagram(ptr.Pointer(), dataC, C.longlong(maxSize), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	defer qt.Recovering("QUdpSocket::setMulticastInterface")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) WriteDatagram2(datagram string, host QHostAddress_ITF, port uint16) int64 {
	defer qt.Recovering("QUdpSocket::writeDatagram")

	if ptr.Pointer() != nil {
		var datagramC = C.CString(hex.EncodeToString([]byte(datagram)))
		defer C.free(unsafe.Pointer(datagramC))
		return int64(C.QUdpSocket_WriteDatagram2(ptr.Pointer(), datagramC, PointerFromQHostAddress(host), C.ushort(port)))
	}
	return 0
}

func (ptr *QUdpSocket) WriteDatagram(data string, size int64, address QHostAddress_ITF, port uint16) int64 {
	defer qt.Recovering("QUdpSocket::writeDatagram")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QUdpSocket_WriteDatagram(ptr.Pointer(), dataC, C.longlong(size), PointerFromQHostAddress(address), C.ushort(port)))
	}
	return 0
}

//export callbackQUdpSocket_DestroyQUdpSocket
func callbackQUdpSocket_DestroyQUdpSocket(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::~QUdpSocket")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::~QUdpSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DestroyQUdpSocketDefault()
	}
}

func (ptr *QUdpSocket) ConnectDestroyQUdpSocket(f func()) {
	defer qt.Recovering("connect QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::~QUdpSocket", f)
	}
}

func (ptr *QUdpSocket) DisconnectDestroyQUdpSocket() {
	defer qt.Recovering("disconnect QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::~QUdpSocket")
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	defer qt.Recovering("QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocketDefault() {
	defer qt.Recovering("QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUdpSocket_ConnectToHost2
func callbackQUdpSocket_ConnectToHost2(ptr unsafe.Pointer, address unsafe.Pointer, port C.ushort, openMode C.longlong) {
	defer qt.Recovering("callback QUdpSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectToHost2"); signal != nil {
		signal.(func(*QHostAddress, uint16, core.QIODevice__OpenModeFlag))(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectToHost2Default(NewQHostAddressFromPointer(address), uint16(port), core.QIODevice__OpenModeFlag(openMode))
	}
}

func (ptr *QUdpSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {
	defer qt.Recovering("connect QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost2", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectToHost2() {
	defer qt.Recovering("disconnect QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost2")
	}
}

func (ptr *QUdpSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectToHost2(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QUdpSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectToHost2Default(ptr.Pointer(), PointerFromQHostAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

//export callbackQUdpSocket_ConnectToHost
func callbackQUdpSocket_ConnectToHost(ptr unsafe.Pointer, hostName *C.char, port C.ushort, openMode C.longlong, protocol C.longlong) {
	defer qt.Recovering("callback QUdpSocket::connectToHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectToHost"); signal != nil {
		signal.(func(string, uint16, core.QIODevice__OpenModeFlag, QAbstractSocket__NetworkLayerProtocol))(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectToHostDefault(C.GoString(hostName), uint16(port), core.QIODevice__OpenModeFlag(openMode), QAbstractSocket__NetworkLayerProtocol(protocol))
	}
}

func (ptr *QUdpSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol)) {
	defer qt.Recovering("connect QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectToHost() {
	defer qt.Recovering("disconnect QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectToHost")
	}
}

func (ptr *QUdpSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QUdpSocket_ConnectToHost(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

func (ptr *QUdpSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protocol QAbstractSocket__NetworkLayerProtocol) {
	defer qt.Recovering("QUdpSocket::connectToHost")

	if ptr.Pointer() != nil {
		var hostNameC = C.CString(hostName)
		defer C.free(unsafe.Pointer(hostNameC))
		C.QUdpSocket_ConnectToHostDefault(ptr.Pointer(), hostNameC, C.ushort(port), C.longlong(openMode), C.longlong(protocol))
	}
}

//export callbackQUdpSocket_DisconnectFromHost
func callbackQUdpSocket_DisconnectFromHost(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::disconnectFromHost")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QUdpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectFromHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectFromHost")
	}
}

func (ptr *QUdpSocket) DisconnectFromHost() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

//export callbackQUdpSocket_Resume
func callbackQUdpSocket_Resume(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::resume")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::resume"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QUdpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::resume", f)
	}
}

func (ptr *QUdpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::resume")
	}
}

func (ptr *QUdpSocket) Resume() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ResumeDefault() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQUdpSocket_SetReadBufferSize
func callbackQUdpSocket_SetReadBufferSize(ptr unsafe.Pointer, size C.longlong) {
	defer qt.Recovering("callback QUdpSocket::setReadBufferSize")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQUdpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QUdpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setReadBufferSize", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setReadBufferSize")
	}
}

func (ptr *QUdpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

//export callbackQUdpSocket_SetSocketOption
func callbackQUdpSocket_SetSocketOption(ptr unsafe.Pointer, option C.longlong, value unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::setSocketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQUdpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QUdpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setSocketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::setSocketOption")
	}
}

func (ptr *QUdpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOption(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOptionDefault(ptr.Pointer(), C.longlong(option), core.PointerFromQVariant(value))
	}
}

//export callbackQUdpSocket_SocketOption
func callbackQUdpSocket_SocketOption(ptr unsafe.Pointer, option C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QUdpSocket::socketOption")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::socketOption"); signal != nil {
		return core.PointerFromQVariant(signal.(func(QAbstractSocket__SocketOption) *core.QVariant)(QAbstractSocket__SocketOption(option)))
	}

	return core.PointerFromQVariant(NewQUdpSocketFromPointer(ptr).SocketOptionDefault(QAbstractSocket__SocketOption(option)))
}

func (ptr *QUdpSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {
	defer qt.Recovering("connect QUdpSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::socketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSocketOption() {
	defer qt.Recovering("disconnect QUdpSocket::socketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::socketOption")
	}
}

func (ptr *QUdpSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QUdpSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QUdpSocket_SocketOption(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QUdpSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QUdpSocket::socketOption")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QUdpSocket_SocketOptionDefault(ptr.Pointer(), C.longlong(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQUdpSocket_WaitForConnected
func callbackQUdpSocket_WaitForConnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QUdpSocket::waitForConnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForConnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForConnectedDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForConnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QUdpSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForConnected", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForConnected() {
	defer qt.Recovering("disconnect QUdpSocket::waitForConnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForConnected")
	}
}

func (ptr *QUdpSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QUdpSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForConnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForConnectedDefault(msecs int) bool {
	defer qt.Recovering("QUdpSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForConnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_WaitForDisconnected
func callbackQUdpSocket_WaitForDisconnected(ptr unsafe.Pointer, msecs C.int) C.char {
	defer qt.Recovering("callback QUdpSocket::waitForDisconnected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::waitForDisconnected"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).WaitForDisconnectedDefault(int(int32(msecs))))))
}

func (ptr *QUdpSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {
	defer qt.Recovering("connect QUdpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForDisconnected", f)
	}
}

func (ptr *QUdpSocket) DisconnectWaitForDisconnected() {
	defer qt.Recovering("disconnect QUdpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::waitForDisconnected")
	}
}

func (ptr *QUdpSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QUdpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForDisconnected(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QUdpSocket) WaitForDisconnectedDefault(msecs int) bool {
	defer qt.Recovering("QUdpSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_WaitForDisconnectedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQUdpSocket_Open
func callbackQUdpSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QUdpSocket::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QUdpSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QUdpSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::open", f)
	}
}

func (ptr *QUdpSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QUdpSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::open")
	}
}

func (ptr *QUdpSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QUdpSocket::open")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_Open(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QUdpSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QUdpSocket::open")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQUdpSocket_Pos
func callbackQUdpSocket_Pos(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QUdpSocket::pos")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).PosDefault())
}

func (ptr *QUdpSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QUdpSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::pos", f)
	}
}

func (ptr *QUdpSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QUdpSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::pos")
	}
}

func (ptr *QUdpSocket) Pos() int64 {
	defer qt.Recovering("QUdpSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) PosDefault() int64 {
	defer qt.Recovering("QUdpSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_Reset
func callbackQUdpSocket_Reset(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QUdpSocket::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QUdpSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QUdpSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::reset", f)
	}
}

func (ptr *QUdpSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QUdpSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::reset")
	}
}

func (ptr *QUdpSocket) Reset() bool {
	defer qt.Recovering("QUdpSocket::reset")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) ResetDefault() bool {
	defer qt.Recovering("QUdpSocket::reset")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQUdpSocket_Seek
func callbackQUdpSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	defer qt.Recovering("callback QUdpSocket::seek")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QUdpSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QUdpSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::seek", f)
	}
}

func (ptr *QUdpSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QUdpSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::seek")
	}
}

func (ptr *QUdpSocket) Seek(pos int64) bool {
	defer qt.Recovering("QUdpSocket::seek")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QUdpSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QUdpSocket::seek")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQUdpSocket_Size
func callbackQUdpSocket_Size(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QUdpSocket::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQUdpSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QUdpSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QUdpSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::size", f)
	}
}

func (ptr *QUdpSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QUdpSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::size")
	}
}

func (ptr *QUdpSocket) Size() int64 {
	defer qt.Recovering("QUdpSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) SizeDefault() int64 {
	defer qt.Recovering("QUdpSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQUdpSocket_TimerEvent
func callbackQUdpSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::timerEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::timerEvent")
	}
}

func (ptr *QUdpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQUdpSocket_ChildEvent
func callbackQUdpSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::childEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::childEvent")
	}
}

func (ptr *QUdpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQUdpSocket_ConnectNotify
func callbackQUdpSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUdpSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUdpSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QUdpSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectNotify", f)
	}
}

func (ptr *QUdpSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QUdpSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::connectNotify")
	}
}

func (ptr *QUdpSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUdpSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUdpSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUdpSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUdpSocket_CustomEvent
func callbackQUdpSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::customEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::customEvent")
	}
}

func (ptr *QUdpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUdpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQUdpSocket_DeleteLater
func callbackQUdpSocket_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QUdpSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QUdpSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::deleteLater", f)
	}
}

func (ptr *QUdpSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QUdpSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::deleteLater")
	}
}

func (ptr *QUdpSocket) DeleteLater() {
	defer qt.Recovering("QUdpSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) DeleteLaterDefault() {
	defer qt.Recovering("QUdpSocket::deleteLater")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQUdpSocket_DisconnectNotify
func callbackQUdpSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QUdpSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QUdpSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectNotify", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QUdpSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::disconnectNotify")
	}
}

func (ptr *QUdpSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUdpSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QUdpSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QUdpSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQUdpSocket_Event
func callbackQUdpSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QUdpSocket::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QUdpSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QUdpSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::event", f)
	}
}

func (ptr *QUdpSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QUdpSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::event")
	}
}

func (ptr *QUdpSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QUdpSocket::event")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QUdpSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QUdpSocket::event")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQUdpSocket_EventFilter
func callbackQUdpSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QUdpSocket::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQUdpSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QUdpSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QUdpSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::eventFilter", f)
	}
}

func (ptr *QUdpSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QUdpSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::eventFilter")
	}
}

func (ptr *QUdpSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QUdpSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QUdpSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QUdpSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQUdpSocket_MetaObject
func callbackQUdpSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QUdpSocket::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QUdpSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQUdpSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QUdpSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QUdpSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::metaObject", f)
	}
}

func (ptr *QUdpSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QUdpSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QUdpSocket::metaObject")
	}
}

func (ptr *QUdpSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QUdpSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUdpSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUdpSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QUdpSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QUdpSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
