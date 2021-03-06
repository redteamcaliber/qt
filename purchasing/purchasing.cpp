// +build !minimal

#define protected public
#define private public

#include "purchasing.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QInAppProduct>
#include <QInAppStore>
#include <QInAppTransaction>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

class MyQInAppProduct: public QInAppProduct
{
public:
	void purchase() { callbackQInAppProduct_Purchase(this); };
	void timerEvent(QTimerEvent * event) { callbackQInAppProduct_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQInAppProduct_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppProduct_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppProduct_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppProduct_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppProduct_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQInAppProduct_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppProduct_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppProduct_MetaObject(const_cast<MyQInAppProduct*>(this))); };
};

char* QInAppProduct_Description(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppProduct*>(ptr)->description().toUtf8().constData());
}

char* QInAppProduct_Identifier(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppProduct*>(ptr)->identifier().toUtf8().constData());
}

char* QInAppProduct_Price(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppProduct*>(ptr)->price().toUtf8().constData());
}

long long QInAppProduct_ProductType(void* ptr)
{
	return static_cast<QInAppProduct*>(ptr)->productType();
}

char* QInAppProduct_Title(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppProduct*>(ptr)->title().toUtf8().constData());
}

void QInAppProduct_Purchase(void* ptr)
{
	static_cast<QInAppProduct*>(ptr)->purchase();
}

void QInAppProduct_TimerEvent(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppProduct_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppProduct_ChildEvent(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QInAppProduct_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppProduct_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppProduct*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppProduct_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppProduct_CustomEvent(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QInAppProduct_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::customEvent(static_cast<QEvent*>(event));
}

void QInAppProduct_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInAppProduct*>(ptr), "deleteLater");
}

void QInAppProduct_DeleteLaterDefault(void* ptr)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::deleteLater();
}

void QInAppProduct_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppProduct*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppProduct_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppProduct*>(ptr)->QInAppProduct::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QInAppProduct_Event(void* ptr, void* e)
{
	return static_cast<QInAppProduct*>(ptr)->event(static_cast<QEvent*>(e));
}

char QInAppProduct_EventDefault(void* ptr, void* e)
{
	return static_cast<QInAppProduct*>(ptr)->QInAppProduct::event(static_cast<QEvent*>(e));
}

char QInAppProduct_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppProduct*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QInAppProduct_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppProduct*>(ptr)->QInAppProduct::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QInAppProduct_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppProduct*>(ptr)->metaObject());
}

void* QInAppProduct_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppProduct*>(ptr)->QInAppProduct::metaObject());
}

class MyQInAppStore: public QInAppStore
{
public:
	MyQInAppStore(QObject *parent) : QInAppStore(parent) {};
	void Signal_ProductRegistered(QInAppProduct * product) { callbackQInAppStore_ProductRegistered(this, product); };
	void Signal_ProductUnknown(QInAppProduct::ProductType productType, const QString & identifier) { callbackQInAppStore_ProductUnknown(this, productType, const_cast<char*>(identifier.toUtf8().constData())); };
	void Signal_TransactionReady(QInAppTransaction * transaction) { callbackQInAppStore_TransactionReady(this, transaction); };
	void timerEvent(QTimerEvent * event) { callbackQInAppStore_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQInAppStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppStore_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppStore_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQInAppStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppStore_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppStore_MetaObject(const_cast<MyQInAppStore*>(this))); };
};

void* QInAppStore_NewQInAppStore(void* parent)
{
	return new MyQInAppStore(static_cast<QObject*>(parent));
}

void QInAppStore_ConnectProductRegistered(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct *)>(&QInAppStore::productRegistered), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct *)>(&MyQInAppStore::Signal_ProductRegistered));
}

void QInAppStore_DisconnectProductRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct *)>(&QInAppStore::productRegistered), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct *)>(&MyQInAppStore::Signal_ProductRegistered));
}

void QInAppStore_ProductRegistered(void* ptr, void* product)
{
	static_cast<QInAppStore*>(ptr)->productRegistered(static_cast<QInAppProduct*>(product));
}

void QInAppStore_ConnectProductUnknown(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&QInAppStore::productUnknown), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&MyQInAppStore::Signal_ProductUnknown));
}

void QInAppStore_DisconnectProductUnknown(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&QInAppStore::productUnknown), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppProduct::ProductType, const QString &)>(&MyQInAppStore::Signal_ProductUnknown));
}

void QInAppStore_ProductUnknown(void* ptr, long long productType, char* identifier)
{
	static_cast<QInAppStore*>(ptr)->productUnknown(static_cast<QInAppProduct::ProductType>(productType), QString(identifier));
}

void QInAppStore_RegisterProduct(void* ptr, long long productType, char* identifier)
{
	static_cast<QInAppStore*>(ptr)->registerProduct(static_cast<QInAppProduct::ProductType>(productType), QString(identifier));
}

void* QInAppStore_RegisteredProduct(void* ptr, char* identifier)
{
	return static_cast<QInAppStore*>(ptr)->registeredProduct(QString(identifier));
}

void QInAppStore_RestorePurchases(void* ptr)
{
	static_cast<QInAppStore*>(ptr)->restorePurchases();
}

void QInAppStore_SetPlatformProperty(void* ptr, char* propertyName, char* value)
{
	static_cast<QInAppStore*>(ptr)->setPlatformProperty(QString(propertyName), QString(value));
}

void QInAppStore_ConnectTransactionReady(void* ptr)
{
	QObject::connect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppTransaction *)>(&QInAppStore::transactionReady), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppTransaction *)>(&MyQInAppStore::Signal_TransactionReady));
}

void QInAppStore_DisconnectTransactionReady(void* ptr)
{
	QObject::disconnect(static_cast<QInAppStore*>(ptr), static_cast<void (QInAppStore::*)(QInAppTransaction *)>(&QInAppStore::transactionReady), static_cast<MyQInAppStore*>(ptr), static_cast<void (MyQInAppStore::*)(QInAppTransaction *)>(&MyQInAppStore::Signal_TransactionReady));
}

void QInAppStore_TransactionReady(void* ptr, void* transaction)
{
	static_cast<QInAppStore*>(ptr)->transactionReady(static_cast<QInAppTransaction*>(transaction));
}

void QInAppStore_DestroyQInAppStore(void* ptr)
{
	static_cast<QInAppStore*>(ptr)->~QInAppStore();
}

void QInAppStore_TimerEvent(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppStore_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppStore_ChildEvent(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QInAppStore_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppStore_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppStore*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppStore_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppStore_CustomEvent(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QInAppStore_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::customEvent(static_cast<QEvent*>(event));
}

void QInAppStore_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInAppStore*>(ptr), "deleteLater");
}

void QInAppStore_DeleteLaterDefault(void* ptr)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::deleteLater();
}

void QInAppStore_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppStore*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppStore_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppStore*>(ptr)->QInAppStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QInAppStore_Event(void* ptr, void* e)
{
	return static_cast<QInAppStore*>(ptr)->event(static_cast<QEvent*>(e));
}

char QInAppStore_EventDefault(void* ptr, void* e)
{
	return static_cast<QInAppStore*>(ptr)->QInAppStore::event(static_cast<QEvent*>(e));
}

char QInAppStore_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppStore*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QInAppStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppStore*>(ptr)->QInAppStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QInAppStore_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppStore*>(ptr)->metaObject());
}

void* QInAppStore_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppStore*>(ptr)->QInAppStore::metaObject());
}

class MyQInAppTransaction: public QInAppTransaction
{
public:
	QString errorString() const { return QString(callbackQInAppTransaction_ErrorString(const_cast<MyQInAppTransaction*>(this))); };
	FailureReason failureReason() const { return static_cast<QInAppTransaction::FailureReason>(callbackQInAppTransaction_FailureReason(const_cast<MyQInAppTransaction*>(this))); };
	QString orderId() const { return QString(callbackQInAppTransaction_OrderId(const_cast<MyQInAppTransaction*>(this))); };
	QDateTime timestamp() const { return *static_cast<QDateTime*>(callbackQInAppTransaction_Timestamp(const_cast<MyQInAppTransaction*>(this))); };
	void finalize() { callbackQInAppTransaction_Finalize(this); };
	QString platformProperty(const QString & propertyName) const { return QString(callbackQInAppTransaction_PlatformProperty(const_cast<MyQInAppTransaction*>(this), const_cast<char*>(propertyName.toUtf8().constData()))); };
	void timerEvent(QTimerEvent * event) { callbackQInAppTransaction_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQInAppTransaction_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInAppTransaction_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInAppTransaction_CustomEvent(this, event); };
	void deleteLater() { callbackQInAppTransaction_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInAppTransaction_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQInAppTransaction_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInAppTransaction_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInAppTransaction_MetaObject(const_cast<MyQInAppTransaction*>(this))); };
};

char* QInAppTransaction_ErrorString(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->errorString().toUtf8().constData());
}

char* QInAppTransaction_ErrorStringDefault(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::errorString().toUtf8().constData());
}

long long QInAppTransaction_FailureReason(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->failureReason();
}

long long QInAppTransaction_FailureReasonDefault(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::failureReason();
}

char* QInAppTransaction_OrderId(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->orderId().toUtf8().constData());
}

char* QInAppTransaction_OrderIdDefault(void* ptr)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::orderId().toUtf8().constData());
}

void* QInAppTransaction_Product(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->product();
}

long long QInAppTransaction_Status(void* ptr)
{
	return static_cast<QInAppTransaction*>(ptr)->status();
}

void* QInAppTransaction_Timestamp(void* ptr)
{
	return new QDateTime(static_cast<QInAppTransaction*>(ptr)->timestamp());
}

void* QInAppTransaction_TimestampDefault(void* ptr)
{
	return new QDateTime(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::timestamp());
}

void QInAppTransaction_Finalize(void* ptr)
{
	static_cast<QInAppTransaction*>(ptr)->finalize();
}

char* QInAppTransaction_PlatformProperty(void* ptr, char* propertyName)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->platformProperty(QString(propertyName)).toUtf8().constData());
}

char* QInAppTransaction_PlatformPropertyDefault(void* ptr, char* propertyName)
{
	return const_cast<char*>(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::platformProperty(QString(propertyName)).toUtf8().constData());
}

void QInAppTransaction_TimerEvent(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppTransaction_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::timerEvent(static_cast<QTimerEvent*>(event));
}

void QInAppTransaction_ChildEvent(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QInAppTransaction_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::childEvent(static_cast<QChildEvent*>(event));
}

void QInAppTransaction_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppTransaction*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppTransaction_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppTransaction_CustomEvent(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QInAppTransaction_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::customEvent(static_cast<QEvent*>(event));
}

void QInAppTransaction_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInAppTransaction*>(ptr), "deleteLater");
}

void QInAppTransaction_DeleteLaterDefault(void* ptr)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::deleteLater();
}

void QInAppTransaction_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QInAppTransaction*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInAppTransaction_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QInAppTransaction_Event(void* ptr, void* e)
{
	return static_cast<QInAppTransaction*>(ptr)->event(static_cast<QEvent*>(e));
}

char QInAppTransaction_EventDefault(void* ptr, void* e)
{
	return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::event(static_cast<QEvent*>(e));
}

char QInAppTransaction_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppTransaction*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QInAppTransaction_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QInAppTransaction_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppTransaction*>(ptr)->metaObject());
}

void* QInAppTransaction_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInAppTransaction*>(ptr)->QInAppTransaction::metaObject());
}

