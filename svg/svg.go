// +build !minimal

package svg

//#include <stdint.h>
//#include <stdlib.h>
//#include "svg.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"unsafe"
)

type QGraphicsSvgItem struct {
	widgets.QGraphicsObject
}

type QGraphicsSvgItem_ITF interface {
	widgets.QGraphicsObject_ITF
	QGraphicsSvgItem_PTR() *QGraphicsSvgItem
}

func (p *QGraphicsSvgItem) QGraphicsSvgItem_PTR() *QGraphicsSvgItem {
	return p
}

func (p *QGraphicsSvgItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QGraphicsObject_PTR().Pointer()
	}
	return nil
}

func (p *QGraphicsSvgItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QGraphicsObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQGraphicsSvgItem(ptr QGraphicsSvgItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSvgItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSvgItemFromPointer(ptr unsafe.Pointer) *QGraphicsSvgItem {
	var n = new(QGraphicsSvgItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSvgItem) DestroyQGraphicsSvgItem() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

func NewQGraphicsSvgItem(parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {
	defer qt.Recovering("QGraphicsSvgItem::QGraphicsSvgItem")

	var tmpValue = NewQGraphicsSvgItemFromPointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem(widgets.PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGraphicsSvgItem2(fileName string, parent widgets.QGraphicsItem_ITF) *QGraphicsSvgItem {
	defer qt.Recovering("QGraphicsSvgItem::QGraphicsSvgItem")

	var fileNameC = C.CString(fileName)
	defer C.free(unsafe.Pointer(fileNameC))
	var tmpValue = NewQGraphicsSvgItemFromPointer(C.QGraphicsSvgItem_NewQGraphicsSvgItem2(fileNameC, widgets.PointerFromQGraphicsItem(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQGraphicsSvgItem_BoundingRect
func callbackQGraphicsSvgItem_BoundingRect(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::boundingRect")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::boundingRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQGraphicsSvgItemFromPointer(ptr).BoundingRectDefault())
}

func (ptr *QGraphicsSvgItem) ConnectBoundingRect(f func() *core.QRectF) {
	defer qt.Recovering("connect QGraphicsSvgItem::boundingRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::boundingRect", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectBoundingRect() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::boundingRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::boundingRect")
	}
}

func (ptr *QGraphicsSvgItem) BoundingRect() *core.QRectF {
	defer qt.Recovering("QGraphicsSvgItem::boundingRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QGraphicsSvgItem_BoundingRect(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) BoundingRectDefault() *core.QRectF {
	defer qt.Recovering("QGraphicsSvgItem::boundingRect")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QGraphicsSvgItem_BoundingRectDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ElementId() string {
	defer qt.Recovering("QGraphicsSvgItem::elementId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsSvgItem_ElementId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsSvgItem) MaximumCacheSize() *core.QSize {
	defer qt.Recovering("QGraphicsSvgItem::maximumCacheSize")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QGraphicsSvgItem_MaximumCacheSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_Paint
func callbackQGraphicsSvgItem_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::paint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionGraphicsItem, *widgets.QWidget))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionGraphicsItemFromPointer(option), widgets.NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsSvgItem) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionGraphicsItem, widget *widgets.QWidget)) {
	defer qt.Recovering("connect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::paint", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::paint")
	}
}

func (ptr *QGraphicsSvgItem) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionGraphicsItem_ITF, widget widgets.QWidget_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionGraphicsItem(option), widgets.PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	defer qt.Recovering("QGraphicsSvgItem::renderer")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSvgRendererFromPointer(C.QGraphicsSvgItem_Renderer(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) SetElementId(id string) {
	defer qt.Recovering("QGraphicsSvgItem::setElementId")

	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QGraphicsSvgItem_SetElementId(ptr.Pointer(), idC)
	}
}

func (ptr *QGraphicsSvgItem) SetMaximumCacheSize(size core.QSize_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetMaximumCacheSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QGraphicsSvgItem) SetSharedRenderer(renderer QSvgRenderer_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::setSharedRenderer")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_SetSharedRenderer(ptr.Pointer(), PointerFromQSvgRenderer(renderer))
	}
}

//export callbackQGraphicsSvgItem_Type
func callbackQGraphicsSvgItem_Type(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QGraphicsSvgItem::type")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::type"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQGraphicsSvgItemFromPointer(ptr).TypeDefault()))
}

func (ptr *QGraphicsSvgItem) ConnectType(f func() int) {
	defer qt.Recovering("connect QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::type", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectType() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::type")
	}
}

func (ptr *QGraphicsSvgItem) Type() int {
	defer qt.Recovering("QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsSvgItem_Type(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSvgItem) TypeDefault() int {
	defer qt.Recovering("QGraphicsSvgItem::type")

	if ptr.Pointer() != nil {
		return int(int32(C.QGraphicsSvgItem_TypeDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQGraphicsSvgItem_UpdateMicroFocus
func callbackQGraphicsSvgItem_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QGraphicsSvgItem) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QGraphicsSvgItem::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::updateMicroFocus", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::updateMicroFocus")
	}
}

func (ptr *QGraphicsSvgItem) UpdateMicroFocus() {
	defer qt.Recovering("QGraphicsSvgItem::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QGraphicsSvgItem) UpdateMicroFocusDefault() {
	defer qt.Recovering("QGraphicsSvgItem::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQGraphicsSvgItem_TimerEvent
func callbackQGraphicsSvgItem_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::timerEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::timerEvent")
	}
}

func (ptr *QGraphicsSvgItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGraphicsSvgItem_ChildEvent
func callbackQGraphicsSvgItem_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::childEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::childEvent")
	}
}

func (ptr *QGraphicsSvgItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQGraphicsSvgItem_ConnectNotify
func callbackQGraphicsSvgItem_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsSvgItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsSvgItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::connectNotify", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::connectNotify")
	}
}

func (ptr *QGraphicsSvgItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsSvgItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsSvgItem_CustomEvent
func callbackQGraphicsSvgItem_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::customEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::customEvent")
	}
}

func (ptr *QGraphicsSvgItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DeleteLater
func callbackQGraphicsSvgItem_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGraphicsSvgItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QGraphicsSvgItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::deleteLater", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::deleteLater")
	}
}

func (ptr *QGraphicsSvgItem) DeleteLater() {
	defer qt.Recovering("QGraphicsSvgItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsSvgItem) DeleteLaterDefault() {
	defer qt.Recovering("QGraphicsSvgItem::deleteLater")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQGraphicsSvgItem_DisconnectNotify
func callbackQGraphicsSvgItem_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGraphicsSvgItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QGraphicsSvgItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::disconnectNotify", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::disconnectNotify")
	}
}

func (ptr *QGraphicsSvgItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QGraphicsSvgItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGraphicsSvgItem_EventFilter
func callbackQGraphicsSvgItem_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::eventFilter", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::eventFilter")
	}
}

func (ptr *QGraphicsSvgItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_MetaObject
func callbackQGraphicsSvgItem_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQGraphicsSvgItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QGraphicsSvgItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QGraphicsSvgItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::metaObject", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::metaObject")
	}
}

func (ptr *QGraphicsSvgItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QGraphicsSvgItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsSvgItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSvgItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QGraphicsSvgItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QGraphicsSvgItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGraphicsSvgItem_Advance
func callbackQGraphicsSvgItem_Advance(ptr unsafe.Pointer, phase C.int) {
	defer qt.Recovering("callback QGraphicsSvgItem::advance")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::advance"); signal != nil {
		signal.(func(int))(int(int32(phase)))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).AdvanceDefault(int(int32(phase)))
	}
}

func (ptr *QGraphicsSvgItem) ConnectAdvance(f func(phase int)) {
	defer qt.Recovering("connect QGraphicsSvgItem::advance")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::advance", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectAdvance() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::advance")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::advance")
	}
}

func (ptr *QGraphicsSvgItem) Advance(phase int) {
	defer qt.Recovering("QGraphicsSvgItem::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_Advance(ptr.Pointer(), C.int(int32(phase)))
	}
}

func (ptr *QGraphicsSvgItem) AdvanceDefault(phase int) {
	defer qt.Recovering("QGraphicsSvgItem::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_AdvanceDefault(ptr.Pointer(), C.int(int32(phase)))
	}
}

//export callbackQGraphicsSvgItem_CollidesWithItem
func callbackQGraphicsSvgItem_CollidesWithItem(ptr unsafe.Pointer, other unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::collidesWithItem")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::collidesWithItem"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, core.Qt__ItemSelectionMode) bool)(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).CollidesWithItemDefault(widgets.NewQGraphicsItemFromPointer(other), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsSvgItem) ConnectCollidesWithItem(f func(other *widgets.QGraphicsItem, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::collidesWithItem", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectCollidesWithItem() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::collidesWithItem")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::collidesWithItem")
	}
}

func (ptr *QGraphicsSvgItem) CollidesWithItem(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsSvgItem::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_CollidesWithItem(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) CollidesWithItemDefault(other widgets.QGraphicsItem_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsSvgItem::collidesWithItem")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_CollidesWithItemDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(other), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_CollidesWithPath
func callbackQGraphicsSvgItem_CollidesWithPath(ptr unsafe.Pointer, path unsafe.Pointer, mode C.longlong) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::collidesWithPath")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::collidesWithPath"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QPainterPath, core.Qt__ItemSelectionMode) bool)(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).CollidesWithPathDefault(gui.NewQPainterPathFromPointer(path), core.Qt__ItemSelectionMode(mode)))))
}

func (ptr *QGraphicsSvgItem) ConnectCollidesWithPath(f func(path *gui.QPainterPath, mode core.Qt__ItemSelectionMode) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::collidesWithPath", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectCollidesWithPath() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::collidesWithPath")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::collidesWithPath")
	}
}

func (ptr *QGraphicsSvgItem) CollidesWithPath(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsSvgItem::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_CollidesWithPath(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) CollidesWithPathDefault(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode) bool {
	defer qt.Recovering("QGraphicsSvgItem::collidesWithPath")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_CollidesWithPathDefault(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_Contains
func callbackQGraphicsSvgItem_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::contains")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point)))))
}

func (ptr *QGraphicsSvgItem) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::contains", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectContains() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::contains")
	}
}

func (ptr *QGraphicsSvgItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_ContextMenuEvent
func callbackQGraphicsSvgItem_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::contextMenuEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneContextMenuEvent))(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).ContextMenuEventDefault(widgets.NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectContextMenuEvent(f func(event *widgets.QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::contextMenuEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::contextMenuEvent")
	}
}

func (ptr *QGraphicsSvgItem) ContextMenuEvent(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ContextMenuEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) ContextMenuEventDefault(event widgets.QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_ContextMenuEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragEnterEvent
func callbackQGraphicsSvgItem_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::dragEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragEnterEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectDragEnterEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragEnterEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragEnterEvent")
	}
}

func (ptr *QGraphicsSvgItem) DragEnterEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) DragEnterEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragLeaveEvent
func callbackQGraphicsSvgItem_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::dragLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragLeaveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectDragLeaveEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragLeaveEvent")
	}
}

func (ptr *QGraphicsSvgItem) DragLeaveEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) DragLeaveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DragMoveEvent
func callbackQGraphicsSvgItem_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::dragMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DragMoveEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectDragMoveEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragMoveEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dragMoveEvent")
	}
}

func (ptr *QGraphicsSvgItem) DragMoveEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) DragMoveEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DragMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_DropEvent
func callbackQGraphicsSvgItem_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::dropEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneDragDropEvent))(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).DropEventDefault(widgets.NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectDropEvent(f func(event *widgets.QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dropEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::dropEvent")
	}
}

func (ptr *QGraphicsSvgItem) DropEvent(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DropEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) DropEventDefault(event widgets.QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_DropEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

//export callbackQGraphicsSvgItem_FocusInEvent
func callbackQGraphicsSvgItem_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::focusInEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::focusInEvent")
	}
}

func (ptr *QGraphicsSvgItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsSvgItem_FocusOutEvent
func callbackQGraphicsSvgItem_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::focusOutEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::focusOutEvent")
	}
}

func (ptr *QGraphicsSvgItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverEnterEvent
func callbackQGraphicsSvgItem_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::hoverEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::hoverEnterEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverEnterEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectHoverEnterEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverEnterEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverEnterEvent")
	}
}

func (ptr *QGraphicsSvgItem) HoverEnterEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverEnterEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverEnterEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverEnterEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverLeaveEvent
func callbackQGraphicsSvgItem_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::hoverLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::hoverLeaveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverLeaveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectHoverLeaveEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverLeaveEvent")
	}
}

func (ptr *QGraphicsSvgItem) HoverLeaveEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverLeaveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverLeaveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverLeaveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_HoverMoveEvent
func callbackQGraphicsSvgItem_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::hoverMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::hoverMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneHoverEvent))(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).HoverMoveEventDefault(widgets.NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectHoverMoveEvent(f func(event *widgets.QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::hoverMoveEvent")
	}
}

func (ptr *QGraphicsSvgItem) HoverMoveEvent(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) HoverMoveEventDefault(event widgets.QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_HoverMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneHoverEvent(event))
	}
}

//export callbackQGraphicsSvgItem_InputMethodEvent
func callbackQGraphicsSvgItem_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::inputMethodEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::inputMethodEvent")
	}
}

func (ptr *QGraphicsSvgItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQGraphicsSvgItem_InputMethodQuery
func callbackQGraphicsSvgItem_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQGraphicsSvgItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QGraphicsSvgItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsSvgItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::inputMethodQuery", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::inputMethodQuery")
	}
}

func (ptr *QGraphicsSvgItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsSvgItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsSvgItem_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsSvgItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsSvgItem_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_IsObscuredBy
func callbackQGraphicsSvgItem_IsObscuredBy(ptr unsafe.Pointer, item unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::isObscuredBy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::isObscuredBy"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem) bool)(widgets.NewQGraphicsItemFromPointer(item)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).IsObscuredByDefault(widgets.NewQGraphicsItemFromPointer(item)))))
}

func (ptr *QGraphicsSvgItem) ConnectIsObscuredBy(f func(item *widgets.QGraphicsItem) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::isObscuredBy", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectIsObscuredBy() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::isObscuredBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::isObscuredBy")
	}
}

func (ptr *QGraphicsSvgItem) IsObscuredBy(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_IsObscuredBy(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) IsObscuredByDefault(item widgets.QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_IsObscuredByDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_ItemChange
func callbackQGraphicsSvgItem_ItemChange(ptr unsafe.Pointer, change C.longlong, value unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::itemChange")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::itemChange"); signal != nil {
		return core.PointerFromQVariant(signal.(func(widgets.QGraphicsItem__GraphicsItemChange, *core.QVariant) *core.QVariant)(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
	}

	return core.PointerFromQVariant(NewQGraphicsSvgItemFromPointer(ptr).ItemChangeDefault(widgets.QGraphicsItem__GraphicsItemChange(change), core.NewQVariantFromPointer(value)))
}

func (ptr *QGraphicsSvgItem) ConnectItemChange(f func(change widgets.QGraphicsItem__GraphicsItemChange, value *core.QVariant) *core.QVariant) {
	defer qt.Recovering("connect QGraphicsSvgItem::itemChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::itemChange", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectItemChange() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::itemChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::itemChange")
	}
}

func (ptr *QGraphicsSvgItem) ItemChange(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsSvgItem::itemChange")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsSvgItem_ItemChange(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ItemChangeDefault(change widgets.QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsSvgItem::itemChange")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QGraphicsSvgItem_ItemChangeDefault(ptr.Pointer(), C.longlong(change), core.PointerFromQVariant(value)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_KeyPressEvent
func callbackQGraphicsSvgItem_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::keyPressEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::keyPressEvent")
	}
}

func (ptr *QGraphicsSvgItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsSvgItem_KeyReleaseEvent
func callbackQGraphicsSvgItem_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::keyReleaseEvent")
	}
}

func (ptr *QGraphicsSvgItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseDoubleClickEvent
func callbackQGraphicsSvgItem_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseDoubleClickEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectMouseDoubleClickEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseDoubleClickEvent")
	}
}

func (ptr *QGraphicsSvgItem) MouseDoubleClickEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseDoubleClickEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseDoubleClickEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseDoubleClickEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseMoveEvent
func callbackQGraphicsSvgItem_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::mouseMoveEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseMoveEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectMouseMoveEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseMoveEvent")
	}
}

func (ptr *QGraphicsSvgItem) MouseMoveEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseMoveEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseMoveEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseMoveEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MousePressEvent
func callbackQGraphicsSvgItem_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::mousePressEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MousePressEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectMousePressEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mousePressEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mousePressEvent")
	}
}

func (ptr *QGraphicsSvgItem) MousePressEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MousePressEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) MousePressEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MousePressEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_MouseReleaseEvent
func callbackQGraphicsSvgItem_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::mouseReleaseEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneMouseEvent))(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).MouseReleaseEventDefault(widgets.NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectMouseReleaseEvent(f func(event *widgets.QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::mouseReleaseEvent")
	}
}

func (ptr *QGraphicsSvgItem) MouseReleaseEvent(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseReleaseEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) MouseReleaseEventDefault(event widgets.QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_MouseReleaseEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneMouseEvent(event))
	}
}

//export callbackQGraphicsSvgItem_OpaqueArea
func callbackQGraphicsSvgItem_OpaqueArea(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::opaqueArea")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::opaqueArea"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsSvgItemFromPointer(ptr).OpaqueAreaDefault())
}

func (ptr *QGraphicsSvgItem) ConnectOpaqueArea(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsSvgItem::opaqueArea")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::opaqueArea", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectOpaqueArea() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::opaqueArea")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::opaqueArea")
	}
}

func (ptr *QGraphicsSvgItem) OpaqueArea() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsSvgItem::opaqueArea")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_OpaqueArea(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) OpaqueAreaDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsSvgItem::opaqueArea")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_OpaqueAreaDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_SceneEvent
func callbackQGraphicsSvgItem_SceneEvent(ptr unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::sceneEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::sceneEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).SceneEventDefault(core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) ConnectSceneEvent(f func(event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::sceneEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::sceneEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectSceneEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::sceneEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::sceneEvent")
	}
}

func (ptr *QGraphicsSvgItem) SceneEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_SceneEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) SceneEventDefault(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_SceneEventDefault(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_SceneEventFilter
func callbackQGraphicsSvgItem_SceneEventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QGraphicsSvgItem::sceneEventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::sceneEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*widgets.QGraphicsItem, *core.QEvent) bool)(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGraphicsSvgItemFromPointer(ptr).SceneEventFilterDefault(widgets.NewQGraphicsItemFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QGraphicsSvgItem) ConnectSceneEventFilter(f func(watched *widgets.QGraphicsItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QGraphicsSvgItem::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::sceneEventFilter", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectSceneEventFilter() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::sceneEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::sceneEventFilter")
	}
}

func (ptr *QGraphicsSvgItem) SceneEventFilter(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_SceneEventFilter(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsSvgItem) SceneEventFilterDefault(watched widgets.QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsSvgItem::sceneEventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsSvgItem_SceneEventFilterDefault(ptr.Pointer(), widgets.PointerFromQGraphicsItem(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGraphicsSvgItem_Shape
func callbackQGraphicsSvgItem_Shape(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QGraphicsSvgItem::shape")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::shape"); signal != nil {
		return gui.PointerFromQPainterPath(signal.(func() *gui.QPainterPath)())
	}

	return gui.PointerFromQPainterPath(NewQGraphicsSvgItemFromPointer(ptr).ShapeDefault())
}

func (ptr *QGraphicsSvgItem) ConnectShape(f func() *gui.QPainterPath) {
	defer qt.Recovering("connect QGraphicsSvgItem::shape")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::shape", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectShape() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::shape")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::shape")
	}
}

func (ptr *QGraphicsSvgItem) Shape() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsSvgItem::shape")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_Shape(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

func (ptr *QGraphicsSvgItem) ShapeDefault() *gui.QPainterPath {
	defer qt.Recovering("QGraphicsSvgItem::shape")

	if ptr.Pointer() != nil {
		var tmpValue = gui.NewQPainterPathFromPointer(C.QGraphicsSvgItem_ShapeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*gui.QPainterPath).DestroyQPainterPath)
		return tmpValue
	}
	return nil
}

//export callbackQGraphicsSvgItem_WheelEvent
func callbackQGraphicsSvgItem_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsSvgItem::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QGraphicsSvgItem::wheelEvent"); signal != nil {
		signal.(func(*widgets.QGraphicsSceneWheelEvent))(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQGraphicsSvgItemFromPointer(ptr).WheelEventDefault(widgets.NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QGraphicsSvgItem) ConnectWheelEvent(f func(event *widgets.QGraphicsSceneWheelEvent)) {
	defer qt.Recovering("connect QGraphicsSvgItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::wheelEvent", f)
	}
}

func (ptr *QGraphicsSvgItem) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsSvgItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QGraphicsSvgItem::wheelEvent")
	}
}

func (ptr *QGraphicsSvgItem) WheelEvent(event widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_WheelEvent(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

func (ptr *QGraphicsSvgItem) WheelEventDefault(event widgets.QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsSvgItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSvgItem_WheelEventDefault(ptr.Pointer(), widgets.PointerFromQGraphicsSceneWheelEvent(event))
	}
}

type QSvgGenerator struct {
	gui.QPaintDevice
}

type QSvgGenerator_ITF interface {
	gui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func (p *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator {
	return p
}

func (p *QSvgGenerator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QPaintDevice_PTR().Pointer()
	}
	return nil
}

func (p *QSvgGenerator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QPaintDevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQSvgGenerator(ptr QSvgGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgGenerator_PTR().Pointer()
	}
	return nil
}

func NewQSvgGeneratorFromPointer(ptr unsafe.Pointer) *QSvgGenerator {
	var n = new(QSvgGenerator)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSvgGenerator) Description() string {
	defer qt.Recovering("QSvgGenerator::description")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) FileName() string {
	defer qt.Recovering("QSvgGenerator::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) OutputDevice() *core.QIODevice {
	defer qt.Recovering("QSvgGenerator::outputDevice")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQIODeviceFromPointer(C.QSvgGenerator_OutputDevice(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) Resolution() int {
	defer qt.Recovering("QSvgGenerator::resolution")

	if ptr.Pointer() != nil {
		return int(int32(C.QSvgGenerator_Resolution(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSvgGenerator) SetDescription(description string) {
	defer qt.Recovering("QSvgGenerator::setDescription")

	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QSvgGenerator_SetDescription(ptr.Pointer(), descriptionC)
	}
}

func (ptr *QSvgGenerator) SetFileName(fileName string) {
	defer qt.Recovering("QSvgGenerator::setFileName")

	if ptr.Pointer() != nil {
		var fileNameC = C.CString(fileName)
		defer C.free(unsafe.Pointer(fileNameC))
		C.QSvgGenerator_SetFileName(ptr.Pointer(), fileNameC)
	}
}

func (ptr *QSvgGenerator) SetOutputDevice(outputDevice core.QIODevice_ITF) {
	defer qt.Recovering("QSvgGenerator::setOutputDevice")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetOutputDevice(ptr.Pointer(), core.PointerFromQIODevice(outputDevice))
	}
}

func (ptr *QSvgGenerator) SetResolution(dpi int) {
	defer qt.Recovering("QSvgGenerator::setResolution")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetResolution(ptr.Pointer(), C.int(int32(dpi)))
	}
}

func (ptr *QSvgGenerator) SetSize(size core.QSize_ITF) {
	defer qt.Recovering("QSvgGenerator::setSize")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSvgGenerator) SetTitle(title string) {
	defer qt.Recovering("QSvgGenerator::setTitle")

	if ptr.Pointer() != nil {
		var titleC = C.CString(title)
		defer C.free(unsafe.Pointer(titleC))
		C.QSvgGenerator_SetTitle(ptr.Pointer(), titleC)
	}
}

func (ptr *QSvgGenerator) SetViewBox(viewBox core.QRect_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewBox))
	}
}

func (ptr *QSvgGenerator) SetViewBox2(viewBox core.QRectF_ITF) {
	defer qt.Recovering("QSvgGenerator::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewBox))
	}
}

func (ptr *QSvgGenerator) Size() *core.QSize {
	defer qt.Recovering("QSvgGenerator::size")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSvgGenerator_Size(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) Title() string {
	defer qt.Recovering("QSvgGenerator::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSvgGenerator_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSvgGenerator) ViewBoxF() *core.QRectF {
	defer qt.Recovering("QSvgGenerator::viewBoxF")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSvgGenerator_ViewBoxF(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func NewQSvgGenerator() *QSvgGenerator {
	defer qt.Recovering("QSvgGenerator::QSvgGenerator")

	var tmpValue = NewQSvgGeneratorFromPointer(C.QSvgGenerator_NewQSvgGenerator())
	runtime.SetFinalizer(tmpValue, (*QSvgGenerator).DestroyQSvgGenerator)
	return tmpValue
}

func (ptr *QSvgGenerator) Metric(metric gui.QPaintDevice__PaintDeviceMetric) int {
	defer qt.Recovering("QSvgGenerator::metric")

	if ptr.Pointer() != nil {
		return int(int32(C.QSvgGenerator_Metric(ptr.Pointer(), C.longlong(metric))))
	}
	return 0
}

func (ptr *QSvgGenerator) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QSvgGenerator::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QSvgGenerator_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgGenerator) ViewBox() *core.QRect {
	defer qt.Recovering("QSvgGenerator::viewBox")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QSvgGenerator_ViewBox(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgGenerator) DestroyQSvgGenerator() {
	defer qt.Recovering("QSvgGenerator::~QSvgGenerator")

	if ptr.Pointer() != nil {
		C.QSvgGenerator_DestroyQSvgGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSvgRenderer struct {
	core.QObject
}

type QSvgRenderer_ITF interface {
	core.QObject_ITF
	QSvgRenderer_PTR() *QSvgRenderer
}

func (p *QSvgRenderer) QSvgRenderer_PTR() *QSvgRenderer {
	return p
}

func (p *QSvgRenderer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSvgRenderer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSvgRenderer(ptr QSvgRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSvgRendererFromPointer(ptr unsafe.Pointer) *QSvgRenderer {
	var n = new(QSvgRenderer)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSvgRenderer) FramesPerSecond() int {
	defer qt.Recovering("QSvgRenderer::framesPerSecond")

	if ptr.Pointer() != nil {
		return int(int32(C.QSvgRenderer_FramesPerSecond(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSvgRenderer) SetFramesPerSecond(num int) {
	defer qt.Recovering("QSvgRenderer::setFramesPerSecond")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetFramesPerSecond(ptr.Pointer(), C.int(int32(num)))
	}
}

func (ptr *QSvgRenderer) SetViewBox(viewbox core.QRect_ITF) {
	defer qt.Recovering("QSvgRenderer::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox(ptr.Pointer(), core.PointerFromQRect(viewbox))
	}
}

func (ptr *QSvgRenderer) SetViewBox2(viewbox core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::setViewBox")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_SetViewBox2(ptr.Pointer(), core.PointerFromQRectF(viewbox))
	}
}

func (ptr *QSvgRenderer) ViewBoxF() *core.QRectF {
	defer qt.Recovering("QSvgRenderer::viewBoxF")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFFromPointer(C.QSvgRenderer_ViewBoxF(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func NewQSvgRenderer(parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	var tmpValue = NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer4(contents core.QXmlStreamReader_ITF, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	var tmpValue = NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer4(core.PointerFromQXmlStreamReader(contents), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer3(contents string, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	var contentsC = C.CString(hex.EncodeToString([]byte(contents)))
	defer C.free(unsafe.Pointer(contentsC))
	var tmpValue = NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer3(contentsC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgRenderer2(filename string, parent core.QObject_ITF) *QSvgRenderer {
	defer qt.Recovering("QSvgRenderer::QSvgRenderer")

	var filenameC = C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	var tmpValue = NewQSvgRendererFromPointer(C.QSvgRenderer_NewQSvgRenderer2(filenameC, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSvgRenderer) Animated() bool {
	defer qt.Recovering("QSvgRenderer::animated")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Animated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSvgRenderer) BoundsOnElement(id string) *core.QRectF {
	defer qt.Recovering("QSvgRenderer::boundsOnElement")

	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		var tmpValue = core.NewQRectFFromPointer(C.QSvgRenderer_BoundsOnElement(ptr.Pointer(), idC))
		runtime.SetFinalizer(tmpValue, (*core.QRectF).DestroyQRectF)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) DefaultSize() *core.QSize {
	defer qt.Recovering("QSvgRenderer::defaultSize")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSvgRenderer_DefaultSize(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) ElementExists(id string) bool {
	defer qt.Recovering("QSvgRenderer::elementExists")

	if ptr.Pointer() != nil {
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		return C.QSvgRenderer_ElementExists(ptr.Pointer(), idC) != 0
	}
	return false
}

func (ptr *QSvgRenderer) IsValid() bool {
	defer qt.Recovering("QSvgRenderer::isValid")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_IsValid(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSvgRenderer_Load3
func callbackQSvgRenderer_Load3(ptr unsafe.Pointer, contents unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgRenderer::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::load3"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QXmlStreamReader) bool)(core.NewQXmlStreamReaderFromPointer(contents)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSvgRenderer) ConnectLoad3(f func(contents *core.QXmlStreamReader) bool) {
	defer qt.Recovering("connect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load3", f)
	}
}

func (ptr *QSvgRenderer) DisconnectLoad3(contents core.QXmlStreamReader_ITF) {
	defer qt.Recovering("disconnect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load3")
	}
}

func (ptr *QSvgRenderer) Load3(contents core.QXmlStreamReader_ITF) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Load3(ptr.Pointer(), core.PointerFromQXmlStreamReader(contents)) != 0
	}
	return false
}

//export callbackQSvgRenderer_Load2
func callbackQSvgRenderer_Load2(ptr unsafe.Pointer, contents *C.char) C.char {
	defer qt.Recovering("callback QSvgRenderer::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::load2"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(qt.HexDecodeToString(C.GoString(contents))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSvgRenderer) ConnectLoad2(f func(contents string) bool) {
	defer qt.Recovering("connect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load2", f)
	}
}

func (ptr *QSvgRenderer) DisconnectLoad2(contents string) {
	defer qt.Recovering("disconnect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load2")
	}
}

func (ptr *QSvgRenderer) Load2(contents string) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		var contentsC = C.CString(hex.EncodeToString([]byte(contents)))
		defer C.free(unsafe.Pointer(contentsC))
		return C.QSvgRenderer_Load2(ptr.Pointer(), contentsC) != 0
	}
	return false
}

//export callbackQSvgRenderer_Load
func callbackQSvgRenderer_Load(ptr unsafe.Pointer, filename *C.char) C.char {
	defer qt.Recovering("callback QSvgRenderer::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::load"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(filename)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSvgRenderer) ConnectLoad(f func(filename string) bool) {
	defer qt.Recovering("connect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load", f)
	}
}

func (ptr *QSvgRenderer) DisconnectLoad(filename string) {
	defer qt.Recovering("disconnect QSvgRenderer::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::load")
	}
}

func (ptr *QSvgRenderer) Load(filename string) bool {
	defer qt.Recovering("QSvgRenderer::load")

	if ptr.Pointer() != nil {
		var filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
		return C.QSvgRenderer_Load(ptr.Pointer(), filenameC) != 0
	}
	return false
}

//export callbackQSvgRenderer_Render
func callbackQSvgRenderer_Render(ptr unsafe.Pointer, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::render")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::render"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	}

}

func (ptr *QSvgRenderer) ConnectRender(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRender(painter gui.QPainter_ITF) {
	defer qt.Recovering("disconnect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render")
	}
}

func (ptr *QSvgRenderer) Render(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

//export callbackQSvgRenderer_Render2
func callbackQSvgRenderer_Render2(ptr unsafe.Pointer, painter unsafe.Pointer, bounds unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::render")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::render2"); signal != nil {
		signal.(func(*gui.QPainter, *core.QRectF))(gui.NewQPainterFromPointer(painter), core.NewQRectFFromPointer(bounds))
	}

}

func (ptr *QSvgRenderer) ConnectRender2(f func(painter *gui.QPainter, bounds *core.QRectF)) {
	defer qt.Recovering("connect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render2", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRender2(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {
	defer qt.Recovering("disconnect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render2")
	}
}

func (ptr *QSvgRenderer) Render2(painter gui.QPainter_ITF, bounds core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(bounds))
	}
}

//export callbackQSvgRenderer_Render3
func callbackQSvgRenderer_Render3(ptr unsafe.Pointer, painter unsafe.Pointer, elementId *C.char, bounds unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::render")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::render3"); signal != nil {
		signal.(func(*gui.QPainter, string, *core.QRectF))(gui.NewQPainterFromPointer(painter), C.GoString(elementId), core.NewQRectFFromPointer(bounds))
	}

}

func (ptr *QSvgRenderer) ConnectRender3(f func(painter *gui.QPainter, elementId string, bounds *core.QRectF)) {
	defer qt.Recovering("connect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render3", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRender3(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {
	defer qt.Recovering("disconnect QSvgRenderer::render")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::render3")
	}
}

func (ptr *QSvgRenderer) Render3(painter gui.QPainter_ITF, elementId string, bounds core.QRectF_ITF) {
	defer qt.Recovering("QSvgRenderer::render")

	if ptr.Pointer() != nil {
		var elementIdC = C.CString(elementId)
		defer C.free(unsafe.Pointer(elementIdC))
		C.QSvgRenderer_Render3(ptr.Pointer(), gui.PointerFromQPainter(painter), elementIdC, core.PointerFromQRectF(bounds))
	}
}

//export callbackQSvgRenderer_RepaintNeeded
func callbackQSvgRenderer_RepaintNeeded(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::repaintNeeded")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::repaintNeeded"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSvgRenderer) ConnectRepaintNeeded(f func()) {
	defer qt.Recovering("connect QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectRepaintNeeded(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::repaintNeeded", f)
	}
}

func (ptr *QSvgRenderer) DisconnectRepaintNeeded() {
	defer qt.Recovering("disconnect QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectRepaintNeeded(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::repaintNeeded")
	}
}

func (ptr *QSvgRenderer) RepaintNeeded() {
	defer qt.Recovering("QSvgRenderer::repaintNeeded")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_RepaintNeeded(ptr.Pointer())
	}
}

func (ptr *QSvgRenderer) ViewBox() *core.QRect {
	defer qt.Recovering("QSvgRenderer::viewBox")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQRectFromPointer(C.QSvgRenderer_ViewBox(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgRenderer) DestroyQSvgRenderer() {
	defer qt.Recovering("QSvgRenderer::~QSvgRenderer")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DestroyQSvgRenderer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSvgRenderer_TimerEvent
func callbackQSvgRenderer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::timerEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::timerEvent")
	}
}

func (ptr *QSvgRenderer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSvgRenderer_ChildEvent
func callbackQSvgRenderer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::childEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::childEvent")
	}
}

func (ptr *QSvgRenderer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSvgRenderer_ConnectNotify
func callbackQSvgRenderer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgRendererFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgRenderer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSvgRenderer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::connectNotify", f)
	}
}

func (ptr *QSvgRenderer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSvgRenderer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::connectNotify")
	}
}

func (ptr *QSvgRenderer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgRenderer::connectNotify")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSvgRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgRenderer::connectNotify")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgRenderer_CustomEvent
func callbackQSvgRenderer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgRenderer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::customEvent", f)
	}
}

func (ptr *QSvgRenderer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::customEvent")
	}
}

func (ptr *QSvgRenderer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgRenderer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgRenderer_DeleteLater
func callbackQSvgRenderer_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgRendererFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSvgRenderer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSvgRenderer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::deleteLater", f)
	}
}

func (ptr *QSvgRenderer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSvgRenderer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::deleteLater")
	}
}

func (ptr *QSvgRenderer) DeleteLater() {
	defer qt.Recovering("QSvgRenderer::deleteLater")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSvgRenderer) DeleteLaterDefault() {
	defer qt.Recovering("QSvgRenderer::deleteLater")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSvgRenderer_DisconnectNotify
func callbackQSvgRenderer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSvgRenderer::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgRendererFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgRenderer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSvgRenderer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::disconnectNotify", f)
	}
}

func (ptr *QSvgRenderer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSvgRenderer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::disconnectNotify")
	}
}

func (ptr *QSvgRenderer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgRenderer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSvgRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgRenderer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSvgRenderer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgRenderer_Event
func callbackQSvgRenderer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgRenderer::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSvgRenderer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSvgRenderer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::event", f)
	}
}

func (ptr *QSvgRenderer) DisconnectEvent() {
	defer qt.Recovering("disconnect QSvgRenderer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::event")
	}
}

func (ptr *QSvgRenderer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgRenderer::event")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgRenderer::event")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSvgRenderer_EventFilter
func callbackQSvgRenderer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgRenderer::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgRendererFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSvgRenderer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSvgRenderer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::eventFilter", f)
	}
}

func (ptr *QSvgRenderer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSvgRenderer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::eventFilter")
	}
}

func (ptr *QSvgRenderer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgRenderer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSvgRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgRenderer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSvgRenderer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSvgRenderer_MetaObject
func callbackQSvgRenderer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSvgRenderer::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgRenderer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSvgRendererFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSvgRenderer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSvgRenderer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::metaObject", f)
	}
}

func (ptr *QSvgRenderer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSvgRenderer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgRenderer::metaObject")
	}
}

func (ptr *QSvgRenderer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSvgRenderer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgRenderer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgRenderer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSvgRenderer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgRenderer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSvgWidget struct {
	widgets.QWidget
}

type QSvgWidget_ITF interface {
	widgets.QWidget_ITF
	QSvgWidget_PTR() *QSvgWidget
}

func (p *QSvgWidget) QSvgWidget_PTR() *QSvgWidget {
	return p
}

func (p *QSvgWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QSvgWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
}

func PointerFromQSvgWidget(ptr QSvgWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSvgWidget_PTR().Pointer()
	}
	return nil
}

func NewQSvgWidgetFromPointer(ptr unsafe.Pointer) *QSvgWidget {
	var n = new(QSvgWidget)
	n.SetPointer(ptr)
	return n
}
func NewQSvgWidget(parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	var tmpValue = NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQSvgWidget2(file string, parent widgets.QWidget_ITF) *QSvgWidget {
	defer qt.Recovering("QSvgWidget::QSvgWidget")

	var fileC = C.CString(file)
	defer C.free(unsafe.Pointer(fileC))
	var tmpValue = NewQSvgWidgetFromPointer(C.QSvgWidget_NewQSvgWidget2(fileC, widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSvgWidget_Load2
func callbackQSvgWidget_Load2(ptr unsafe.Pointer, contents *C.char) {
	defer qt.Recovering("callback QSvgWidget::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::load2"); signal != nil {
		signal.(func(string))(qt.HexDecodeToString(C.GoString(contents)))
	}

}

func (ptr *QSvgWidget) ConnectLoad2(f func(contents string)) {
	defer qt.Recovering("connect QSvgWidget::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::load2", f)
	}
}

func (ptr *QSvgWidget) DisconnectLoad2(contents string) {
	defer qt.Recovering("disconnect QSvgWidget::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::load2")
	}
}

func (ptr *QSvgWidget) Load2(contents string) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		var contentsC = C.CString(hex.EncodeToString([]byte(contents)))
		defer C.free(unsafe.Pointer(contentsC))
		C.QSvgWidget_Load2(ptr.Pointer(), contentsC)
	}
}

//export callbackQSvgWidget_Load
func callbackQSvgWidget_Load(ptr unsafe.Pointer, file *C.char) {
	defer qt.Recovering("callback QSvgWidget::load")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::load"); signal != nil {
		signal.(func(string))(C.GoString(file))
	}

}

func (ptr *QSvgWidget) ConnectLoad(f func(file string)) {
	defer qt.Recovering("connect QSvgWidget::load")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::load", f)
	}
}

func (ptr *QSvgWidget) DisconnectLoad(file string) {
	defer qt.Recovering("disconnect QSvgWidget::load")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::load")
	}
}

func (ptr *QSvgWidget) Load(file string) {
	defer qt.Recovering("QSvgWidget::load")

	if ptr.Pointer() != nil {
		var fileC = C.CString(file)
		defer C.free(unsafe.Pointer(fileC))
		C.QSvgWidget_Load(ptr.Pointer(), fileC)
	}
}

func (ptr *QSvgWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSvgWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QSvgWidget) Renderer() *QSvgRenderer {
	defer qt.Recovering("QSvgWidget::renderer")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSvgRendererFromPointer(C.QSvgWidget_Renderer(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QSvgWidget::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSvgWidget_SizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) DestroyQSvgWidget() {
	defer qt.Recovering("QSvgWidget::~QSvgWidget")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DestroyQSvgWidget(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSvgWidget_ActionEvent
func callbackQSvgWidget_ActionEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::actionEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::actionEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::actionEvent")
	}
}

func (ptr *QSvgWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSvgWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSvgWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

//export callbackQSvgWidget_DragEnterEvent
func callbackQSvgWidget_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragEnterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragEnterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragEnterEvent")
	}
}

func (ptr *QSvgWidget) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSvgWidget) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackQSvgWidget_DragLeaveEvent
func callbackQSvgWidget_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragLeaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragLeaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragLeaveEvent")
	}
}

func (ptr *QSvgWidget) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSvgWidget) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackQSvgWidget_DragMoveEvent
func callbackQSvgWidget_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dragMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dragMoveEvent")
	}
}

func (ptr *QSvgWidget) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSvgWidget) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackQSvgWidget_DropEvent
func callbackQSvgWidget_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::dropEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dropEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::dropEvent")
	}
}

func (ptr *QSvgWidget) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSvgWidget) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSvgWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

//export callbackQSvgWidget_EnterEvent
func callbackQSvgWidget_EnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::enterEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::enterEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::enterEvent")
	}
}

func (ptr *QSvgWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_FocusInEvent
func callbackQSvgWidget_FocusInEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusInEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusInEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusInEvent")
	}
}

func (ptr *QSvgWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQSvgWidget_FocusOutEvent
func callbackQSvgWidget_FocusOutEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::focusOutEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusOutEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusOutEvent")
	}
}

func (ptr *QSvgWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSvgWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSvgWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

//export callbackQSvgWidget_HideEvent
func callbackQSvgWidget_HideEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::hideEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hideEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hideEvent")
	}
}

func (ptr *QSvgWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSvgWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSvgWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

//export callbackQSvgWidget_LeaveEvent
func callbackQSvgWidget_LeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::leaveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::leaveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::leaveEvent")
	}
}

func (ptr *QSvgWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_MinimumSizeHint
func callbackQSvgWidget_MinimumSizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSvgWidget::minimumSizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQSvgWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QSvgWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QSvgWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::minimumSizeHint", f)
	}
}

func (ptr *QSvgWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QSvgWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::minimumSizeHint")
	}
}

func (ptr *QSvgWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QSvgWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSvgWidget_MinimumSizeHint(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QSvgWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSvgWidget_MinimumSizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSvgWidget_MoveEvent
func callbackQSvgWidget_MoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::moveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::moveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::moveEvent")
	}
}

func (ptr *QSvgWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSvgWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSvgWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

//export callbackQSvgWidget_SetEnabled
func callbackQSvgWidget_SetEnabled(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QSvgWidget::setEnabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setEnabled"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetEnabledDefault(int8(vbo) != 0)
	}
}

func (ptr *QSvgWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QSvgWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setEnabled", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QSvgWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setEnabled")
	}
}

func (ptr *QSvgWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QSvgWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetEnabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QSvgWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QSvgWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetEnabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQSvgWidget_SetStyleSheet
func callbackQSvgWidget_SetStyleSheet(ptr unsafe.Pointer, styleSheet *C.char) {
	defer qt.Recovering("callback QSvgWidget::setStyleSheet")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQSvgWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QSvgWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QSvgWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setStyleSheet", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QSvgWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setStyleSheet")
	}
}

func (ptr *QSvgWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QSvgWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QSvgWidget_SetStyleSheet(ptr.Pointer(), styleSheetC)
	}
}

func (ptr *QSvgWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QSvgWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		var styleSheetC = C.CString(styleSheet)
		defer C.free(unsafe.Pointer(styleSheetC))
		C.QSvgWidget_SetStyleSheetDefault(ptr.Pointer(), styleSheetC)
	}
}

//export callbackQSvgWidget_SetVisible
func callbackQSvgWidget_SetVisible(ptr unsafe.Pointer, visible C.char) {
	defer qt.Recovering("callback QSvgWidget::setVisible")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setVisible"); signal != nil {
		signal.(func(bool))(int8(visible) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetVisibleDefault(int8(visible) != 0)
	}
}

func (ptr *QSvgWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setVisible", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSvgWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setVisible")
	}
}

func (ptr *QSvgWidget) SetVisible(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisible(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

func (ptr *QSvgWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSvgWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetVisibleDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(visible))))
	}
}

//export callbackQSvgWidget_SetWindowModified
func callbackQSvgWidget_SetWindowModified(ptr unsafe.Pointer, vbo C.char) {
	defer qt.Recovering("callback QSvgWidget::setWindowModified")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setWindowModified"); signal != nil {
		signal.(func(bool))(int8(vbo) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetWindowModifiedDefault(int8(vbo) != 0)
	}
}

func (ptr *QSvgWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QSvgWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setWindowModified", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QSvgWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setWindowModified")
	}
}

func (ptr *QSvgWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QSvgWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetWindowModified(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

func (ptr *QSvgWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QSvgWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetWindowModifiedDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(vbo))))
	}
}

//export callbackQSvgWidget_SetWindowTitle
func callbackQSvgWidget_SetWindowTitle(ptr unsafe.Pointer, vqs *C.char) {
	defer qt.Recovering("callback QSvgWidget::setWindowTitle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQSvgWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QSvgWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QSvgWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setWindowTitle", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QSvgWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setWindowTitle")
	}
}

func (ptr *QSvgWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QSvgWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QSvgWidget_SetWindowTitle(ptr.Pointer(), vqsC)
	}
}

func (ptr *QSvgWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QSvgWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		var vqsC = C.CString(vqs)
		defer C.free(unsafe.Pointer(vqsC))
		C.QSvgWidget_SetWindowTitleDefault(ptr.Pointer(), vqsC)
	}
}

//export callbackQSvgWidget_ShowEvent
func callbackQSvgWidget_ShowEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSvgWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showEvent")
	}
}

func (ptr *QSvgWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSvgWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSvgWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

//export callbackQSvgWidget_ChangeEvent
func callbackQSvgWidget_ChangeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::changeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::changeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::changeEvent")
	}
}

func (ptr *QSvgWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_Close
func callbackQSvgWidget_Close(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgWidget::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::close"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).CloseDefault())))
}

func (ptr *QSvgWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QSvgWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::close", f)
	}
}

func (ptr *QSvgWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QSvgWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::close")
	}
}

func (ptr *QSvgWidget) Close() bool {
	defer qt.Recovering("QSvgWidget::close")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSvgWidget) CloseDefault() bool {
	defer qt.Recovering("QSvgWidget::close")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSvgWidget_CloseEvent
func callbackQSvgWidget_CloseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::closeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::closeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::closeEvent")
	}
}

func (ptr *QSvgWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSvgWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

//export callbackQSvgWidget_ContextMenuEvent
func callbackQSvgWidget_ContextMenuEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::contextMenuEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::contextMenuEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::contextMenuEvent")
	}
}

func (ptr *QSvgWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSvgWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSvgWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

//export callbackQSvgWidget_FocusNextPrevChild
func callbackQSvgWidget_FocusNextPrevChild(ptr unsafe.Pointer, next C.char) C.char {
	defer qt.Recovering("callback QSvgWidget::focusNextPrevChild")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::focusNextPrevChild"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(bool) bool)(int8(next) != 0))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).FocusNextPrevChildDefault(int8(next) != 0))))
}

func (ptr *QSvgWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QSvgWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusNextPrevChild", f)
	}
}

func (ptr *QSvgWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QSvgWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::focusNextPrevChild")
	}
}

func (ptr *QSvgWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QSvgWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_FocusNextPrevChild(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

func (ptr *QSvgWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QSvgWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(next)))) != 0
	}
	return false
}

//export callbackQSvgWidget_HasHeightForWidth
func callbackQSvgWidget_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgWidget::hasHeightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QSvgWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QSvgWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hasHeightForWidth", f)
	}
}

func (ptr *QSvgWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QSvgWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hasHeightForWidth")
	}
}

func (ptr *QSvgWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QSvgWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSvgWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QSvgWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSvgWidget_HeightForWidth
func callbackQSvgWidget_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	defer qt.Recovering("callback QSvgWidget::heightForWidth")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQSvgWidgetFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QSvgWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QSvgWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::heightForWidth", f)
	}
}

func (ptr *QSvgWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QSvgWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::heightForWidth")
	}
}

func (ptr *QSvgWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QSvgWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QSvgWidget_HeightForWidth(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

func (ptr *QSvgWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QSvgWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(int32(C.QSvgWidget_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQSvgWidget_Hide
func callbackQSvgWidget_Hide(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::hide")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::hide"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QSvgWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QSvgWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hide", f)
	}
}

func (ptr *QSvgWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QSvgWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::hide")
	}
}

func (ptr *QSvgWidget) Hide() {
	defer qt.Recovering("QSvgWidget::hide")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) HideDefault() {
	defer qt.Recovering("QSvgWidget::hide")

	if ptr.Pointer() != nil {
		C.QSvgWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_InputMethodEvent
func callbackQSvgWidget_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::inputMethodEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::inputMethodEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::inputMethodEvent")
	}
}

func (ptr *QSvgWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSvgWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSvgWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackQSvgWidget_InputMethodQuery
func callbackQSvgWidget_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	defer qt.Recovering("callback QSvgWidget::inputMethodQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQSvgWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QSvgWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QSvgWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::inputMethodQuery", f)
	}
}

func (ptr *QSvgWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QSvgWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::inputMethodQuery")
	}
}

func (ptr *QSvgWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QSvgWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSvgWidget_InputMethodQuery(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSvgWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QSvgWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSvgWidget_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSvgWidget_KeyPressEvent
func callbackQSvgWidget_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyPressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::keyPressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::keyPressEvent")
	}
}

func (ptr *QSvgWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQSvgWidget_KeyReleaseEvent
func callbackQSvgWidget_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::keyReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::keyReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::keyReleaseEvent")
	}
}

func (ptr *QSvgWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSvgWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSvgWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

//export callbackQSvgWidget_Lower
func callbackQSvgWidget_Lower(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::lower")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::lower"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QSvgWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QSvgWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::lower", f)
	}
}

func (ptr *QSvgWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QSvgWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::lower")
	}
}

func (ptr *QSvgWidget) Lower() {
	defer qt.Recovering("QSvgWidget::lower")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) LowerDefault() {
	defer qt.Recovering("QSvgWidget::lower")

	if ptr.Pointer() != nil {
		C.QSvgWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_MouseDoubleClickEvent
func callbackQSvgWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseDoubleClickEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseDoubleClickEvent")
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MouseMoveEvent
func callbackQSvgWidget_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseMoveEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseMoveEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseMoveEvent")
	}
}

func (ptr *QSvgWidget) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MousePressEvent
func callbackQSvgWidget_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mousePressEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mousePressEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mousePressEvent")
	}
}

func (ptr *QSvgWidget) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_MouseReleaseEvent
func callbackQSvgWidget_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseReleaseEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::mouseReleaseEvent")
	}
}

func (ptr *QSvgWidget) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSvgWidget) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSvgWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

//export callbackQSvgWidget_NativeEvent
func callbackQSvgWidget_NativeEvent(ptr unsafe.Pointer, eventType *C.char, message unsafe.Pointer, result C.long) C.char {
	defer qt.Recovering("callback QSvgWidget::nativeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::nativeEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).NativeEventDefault(qt.HexDecodeToString(C.GoString(eventType)), message, int(int32(result))))))
}

func (ptr *QSvgWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QSvgWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::nativeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::nativeEvent")
	}
}

func (ptr *QSvgWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QSvgWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QSvgWidget_NativeEvent(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

func (ptr *QSvgWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QSvgWidget::nativeEvent")

	if ptr.Pointer() != nil {
		var eventTypeC = C.CString(hex.EncodeToString([]byte(eventType)))
		defer C.free(unsafe.Pointer(eventTypeC))
		return C.QSvgWidget_NativeEventDefault(ptr.Pointer(), eventTypeC, message, C.long(int32(result))) != 0
	}
	return false
}

//export callbackQSvgWidget_Raise
func callbackQSvgWidget_Raise(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::raise")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::raise"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QSvgWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QSvgWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::raise", f)
	}
}

func (ptr *QSvgWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QSvgWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::raise")
	}
}

func (ptr *QSvgWidget) Raise() {
	defer qt.Recovering("QSvgWidget::raise")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) RaiseDefault() {
	defer qt.Recovering("QSvgWidget::raise")

	if ptr.Pointer() != nil {
		C.QSvgWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_Repaint
func callbackQSvgWidget_Repaint(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::repaint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QSvgWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QSvgWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::repaint", f)
	}
}

func (ptr *QSvgWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QSvgWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::repaint")
	}
}

func (ptr *QSvgWidget) Repaint() {
	defer qt.Recovering("QSvgWidget::repaint")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) RepaintDefault() {
	defer qt.Recovering("QSvgWidget::repaint")

	if ptr.Pointer() != nil {
		C.QSvgWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ResizeEvent
func callbackQSvgWidget_ResizeEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::resizeEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::resizeEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::resizeEvent")
	}
}

func (ptr *QSvgWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSvgWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSvgWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

//export callbackQSvgWidget_SetDisabled
func callbackQSvgWidget_SetDisabled(ptr unsafe.Pointer, disable C.char) {
	defer qt.Recovering("callback QSvgWidget::setDisabled")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setDisabled"); signal != nil {
		signal.(func(bool))(int8(disable) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetDisabledDefault(int8(disable) != 0)
	}
}

func (ptr *QSvgWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QSvgWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setDisabled", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QSvgWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setDisabled")
	}
}

func (ptr *QSvgWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QSvgWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetDisabled(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

func (ptr *QSvgWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QSvgWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetDisabledDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(disable))))
	}
}

//export callbackQSvgWidget_SetFocus2
func callbackQSvgWidget_SetFocus2(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::setFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QSvgWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QSvgWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setFocus2", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QSvgWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setFocus2")
	}
}

func (ptr *QSvgWidget) SetFocus2() {
	defer qt.Recovering("QSvgWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) SetFocus2Default() {
	defer qt.Recovering("QSvgWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQSvgWidget_SetHidden
func callbackQSvgWidget_SetHidden(ptr unsafe.Pointer, hidden C.char) {
	defer qt.Recovering("callback QSvgWidget::setHidden")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::setHidden"); signal != nil {
		signal.(func(bool))(int8(hidden) != 0)
	} else {
		NewQSvgWidgetFromPointer(ptr).SetHiddenDefault(int8(hidden) != 0)
	}
}

func (ptr *QSvgWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QSvgWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setHidden", f)
	}
}

func (ptr *QSvgWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QSvgWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::setHidden")
	}
}

func (ptr *QSvgWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QSvgWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetHidden(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

func (ptr *QSvgWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QSvgWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QSvgWidget_SetHiddenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(hidden))))
	}
}

//export callbackQSvgWidget_Show
func callbackQSvgWidget_Show(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::show")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::show"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QSvgWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QSvgWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::show", f)
	}
}

func (ptr *QSvgWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QSvgWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::show")
	}
}

func (ptr *QSvgWidget) Show() {
	defer qt.Recovering("QSvgWidget::show")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Show(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) ShowDefault() {
	defer qt.Recovering("QSvgWidget::show")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowFullScreen
func callbackQSvgWidget_ShowFullScreen(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showFullScreen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QSvgWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QSvgWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showFullScreen", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QSvgWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showFullScreen")
	}
}

func (ptr *QSvgWidget) ShowFullScreen() {
	defer qt.Recovering("QSvgWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QSvgWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowMaximized
func callbackQSvgWidget_ShowMaximized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showMaximized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QSvgWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QSvgWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showMaximized", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QSvgWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showMaximized")
	}
}

func (ptr *QSvgWidget) ShowMaximized() {
	defer qt.Recovering("QSvgWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QSvgWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowMinimized
func callbackQSvgWidget_ShowMinimized(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showMinimized")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QSvgWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QSvgWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showMinimized", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QSvgWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showMinimized")
	}
}

func (ptr *QSvgWidget) ShowMinimized() {
	defer qt.Recovering("QSvgWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QSvgWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_ShowNormal
func callbackQSvgWidget_ShowNormal(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::showNormal")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QSvgWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QSvgWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showNormal", f)
	}
}

func (ptr *QSvgWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QSvgWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::showNormal")
	}
}

func (ptr *QSvgWidget) ShowNormal() {
	defer qt.Recovering("QSvgWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) ShowNormalDefault() {
	defer qt.Recovering("QSvgWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_TabletEvent
func callbackQSvgWidget_TabletEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::tabletEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::tabletEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::tabletEvent")
	}
}

func (ptr *QSvgWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSvgWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSvgWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

//export callbackQSvgWidget_Update
func callbackQSvgWidget_Update(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::update")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::update"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QSvgWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QSvgWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::update", f)
	}
}

func (ptr *QSvgWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QSvgWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::update")
	}
}

func (ptr *QSvgWidget) Update() {
	defer qt.Recovering("QSvgWidget::update")

	if ptr.Pointer() != nil {
		C.QSvgWidget_Update(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) UpdateDefault() {
	defer qt.Recovering("QSvgWidget::update")

	if ptr.Pointer() != nil {
		C.QSvgWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_UpdateMicroFocus
func callbackQSvgWidget_UpdateMicroFocus(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::updateMicroFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QSvgWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QSvgWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::updateMicroFocus", f)
	}
}

func (ptr *QSvgWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QSvgWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::updateMicroFocus")
	}
}

func (ptr *QSvgWidget) UpdateMicroFocus() {
	defer qt.Recovering("QSvgWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QSvgWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QSvgWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QSvgWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QSvgWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQSvgWidget_WheelEvent
func callbackQSvgWidget_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::wheelEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::wheelEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::wheelEvent")
	}
}

func (ptr *QSvgWidget) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QSvgWidget) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSvgWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

//export callbackQSvgWidget_TimerEvent
func callbackQSvgWidget_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::timerEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::timerEvent")
	}
}

func (ptr *QSvgWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSvgWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSvgWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSvgWidget_ChildEvent
func callbackQSvgWidget_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::childEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSvgWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::childEvent")
	}
}

func (ptr *QSvgWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSvgWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSvgWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSvgWidget_ConnectNotify
func callbackQSvgWidget_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSvgWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::connectNotify", f)
	}
}

func (ptr *QSvgWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSvgWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::connectNotify")
	}
}

func (ptr *QSvgWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSvgWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QSvgWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgWidget_CustomEvent
func callbackQSvgWidget_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSvgWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSvgWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::customEvent", f)
	}
}

func (ptr *QSvgWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSvgWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::customEvent")
	}
}

func (ptr *QSvgWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSvgWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSvgWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QSvgWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSvgWidget_DeleteLater
func callbackQSvgWidget_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSvgWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSvgWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSvgWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::deleteLater", f)
	}
}

func (ptr *QSvgWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSvgWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::deleteLater")
	}
}

func (ptr *QSvgWidget) DeleteLater() {
	defer qt.Recovering("QSvgWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSvgWidget) DeleteLaterDefault() {
	defer qt.Recovering("QSvgWidget::deleteLater")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSvgWidget_DisconnectNotify
func callbackQSvgWidget_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSvgWidget::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSvgWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSvgWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSvgWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::disconnectNotify", f)
	}
}

func (ptr *QSvgWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSvgWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::disconnectNotify")
	}
}

func (ptr *QSvgWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSvgWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSvgWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSvgWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSvgWidget_EventFilter
func callbackQSvgWidget_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSvgWidget::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSvgWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSvgWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSvgWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::eventFilter", f)
	}
}

func (ptr *QSvgWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSvgWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::eventFilter")
	}
}

func (ptr *QSvgWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSvgWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSvgWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSvgWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSvgWidget_MetaObject
func callbackQSvgWidget_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSvgWidget::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSvgWidget::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSvgWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSvgWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSvgWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::metaObject", f)
	}
}

func (ptr *QSvgWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSvgWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSvgWidget::metaObject")
	}
}

func (ptr *QSvgWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSvgWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSvgWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSvgWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSvgWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
