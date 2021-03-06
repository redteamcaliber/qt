// +build !minimal

package sql

//#include <stdint.h>
//#include <stdlib.h>
//#include "sql.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"runtime"
	"strings"
	"unsafe"
)

type QSqlDatabase struct {
	ptr unsafe.Pointer
}

type QSqlDatabase_ITF interface {
	QSqlDatabase_PTR() *QSqlDatabase
}

func (p *QSqlDatabase) QSqlDatabase_PTR() *QSqlDatabase {
	return p
}

func (p *QSqlDatabase) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlDatabase) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlDatabase(ptr QSqlDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDatabase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDatabaseFromPointer(ptr unsafe.Pointer) *QSqlDatabase {
	var n = new(QSqlDatabase)
	n.SetPointer(ptr)
	return n
}
func QSqlDatabase_AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::addDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase2(PointerFromQSqlDriver(driver), connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) AddDatabase2(driver QSqlDriver_ITF, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::addDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase2(PointerFromQSqlDriver(driver), connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_AddDatabase(ty string, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::addDatabase")

	var tyC = C.CString(ty)
	defer C.free(unsafe.Pointer(tyC))
	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase(tyC, connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) AddDatabase(ty string, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::addDatabase")

	var tyC = C.CString(ty)
	defer C.free(unsafe.Pointer(tyC))
	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_AddDatabase(tyC, connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase() *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase())
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase2(other QSqlDatabase_ITF) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase2(PointerFromQSqlDatabase(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase4(driver QSqlDriver_ITF) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase4(PointerFromQSqlDriver(driver)))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func NewQSqlDatabase3(ty string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::QSqlDatabase")

	var tyC = C.CString(ty)
	defer C.free(unsafe.Pointer(tyC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_NewQSqlDatabase3(tyC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func QSqlDatabase_CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::cloneDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_CloneDatabase(PointerFromQSqlDatabase(other), connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) CloneDatabase(other QSqlDatabase_ITF, connectionName string) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::cloneDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_CloneDatabase(PointerFromQSqlDatabase(other), connectionNameC))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) Close() {
	defer qt.Recovering("QSqlDatabase::close")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_Close(ptr.Pointer())
	}
}

func (ptr *QSqlDatabase) Commit() bool {
	defer qt.Recovering("QSqlDatabase::commit")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Commit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) ConnectOptions() string {
	defer qt.Recovering("QSqlDatabase::connectOptions")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectOptions(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) ConnectionName() string {
	defer qt.Recovering("QSqlDatabase::connectionName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_ConnectionName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_ConnectionNames() []string {
	defer qt.Recovering("QSqlDatabase::connectionNames")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func (ptr *QSqlDatabase) ConnectionNames() []string {
	defer qt.Recovering("QSqlDatabase::connectionNames")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_ConnectionNames()), "|")
}

func QSqlDatabase_Contains(connectionName string) bool {
	defer qt.Recovering("QSqlDatabase::contains")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	return C.QSqlDatabase_QSqlDatabase_Contains(connectionNameC) != 0
}

func (ptr *QSqlDatabase) Contains(connectionName string) bool {
	defer qt.Recovering("QSqlDatabase::contains")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	return C.QSqlDatabase_QSqlDatabase_Contains(connectionNameC) != 0
}

func QSqlDatabase_Database(connectionName string, open bool) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::database")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_Database(connectionNameC, C.char(int8(qt.GoBoolToInt(open)))))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) Database(connectionName string, open bool) *QSqlDatabase {
	defer qt.Recovering("QSqlDatabase::database")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlDatabase_QSqlDatabase_Database(connectionNameC, C.char(int8(qt.GoBoolToInt(open)))))
	runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
	return tmpValue
}

func (ptr *QSqlDatabase) DatabaseName() string {
	defer qt.Recovering("QSqlDatabase::databaseName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DatabaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlDatabase::driver")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlDriverFromPointer(C.QSqlDatabase_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) DriverName() string {
	defer qt.Recovering("QSqlDatabase::driverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_DriverName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_Drivers() []string {
	defer qt.Recovering("QSqlDatabase::drivers")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) Drivers() []string {
	defer qt.Recovering("QSqlDatabase::drivers")

	return strings.Split(C.GoString(C.QSqlDatabase_QSqlDatabase_Drivers()), "|")
}

func (ptr *QSqlDatabase) Exec(query string) *QSqlQuery {
	defer qt.Recovering("QSqlDatabase::exec")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		var tmpValue = NewQSqlQueryFromPointer(C.QSqlDatabase_Exec(ptr.Pointer(), queryC))
		runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) HostName() string {
	defer qt.Recovering("QSqlDatabase::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_HostName(ptr.Pointer()))
	}
	return ""
}

func QSqlDatabase_IsDriverAvailable(name string) bool {
	defer qt.Recovering("QSqlDatabase::isDriverAvailable")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(nameC) != 0
}

func (ptr *QSqlDatabase) IsDriverAvailable(name string) bool {
	defer qt.Recovering("QSqlDatabase::isDriverAvailable")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	return C.QSqlDatabase_QSqlDatabase_IsDriverAvailable(nameC) != 0
}

func (ptr *QSqlDatabase) IsOpen() bool {
	defer qt.Recovering("QSqlDatabase::isOpen")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsOpenError() bool {
	defer qt.Recovering("QSqlDatabase::isOpenError")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) IsValid() bool {
	defer qt.Recovering("QSqlDatabase::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) LastError() *QSqlError {
	defer qt.Recovering("QSqlDatabase::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlErrorFromPointer(C.QSqlDatabase_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) Open() bool {
	defer qt.Recovering("QSqlDatabase::open")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Open2(user string, password string) bool {
	defer qt.Recovering("QSqlDatabase::open")

	if ptr.Pointer() != nil {
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		return C.QSqlDatabase_Open2(ptr.Pointer(), userC, passwordC) != 0
	}
	return false
}

func (ptr *QSqlDatabase) Password() string {
	defer qt.Recovering("QSqlDatabase::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) Port() int {
	defer qt.Recovering("QSqlDatabase::port")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlDatabase_Port(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlDatabase) PrimaryIndex(tablename string) *QSqlIndex {
	defer qt.Recovering("QSqlDatabase::primaryIndex")

	if ptr.Pointer() != nil {
		var tablenameC = C.CString(tablename)
		defer C.free(unsafe.Pointer(tablenameC))
		var tmpValue = NewQSqlIndexFromPointer(C.QSqlDatabase_PrimaryIndex(ptr.Pointer(), tablenameC))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDatabase) Record(tablename string) *QSqlRecord {
	defer qt.Recovering("QSqlDatabase::record")

	if ptr.Pointer() != nil {
		var tablenameC = C.CString(tablename)
		defer C.free(unsafe.Pointer(tablenameC))
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlDatabase_Record(ptr.Pointer(), tablenameC))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func QSqlDatabase_RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	defer qt.Recovering("QSqlDatabase::registerSqlDriver")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(nameC, PointerFromQSqlDriverCreatorBase(creator))
}

func (ptr *QSqlDatabase) RegisterSqlDriver(name string, creator QSqlDriverCreatorBase_ITF) {
	defer qt.Recovering("QSqlDatabase::registerSqlDriver")

	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.QSqlDatabase_QSqlDatabase_RegisterSqlDriver(nameC, PointerFromQSqlDriverCreatorBase(creator))
}

func QSqlDatabase_RemoveDatabase(connectionName string) {
	defer qt.Recovering("QSqlDatabase::removeDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(connectionNameC)
}

func (ptr *QSqlDatabase) RemoveDatabase(connectionName string) {
	defer qt.Recovering("QSqlDatabase::removeDatabase")

	var connectionNameC = C.CString(connectionName)
	defer C.free(unsafe.Pointer(connectionNameC))
	C.QSqlDatabase_QSqlDatabase_RemoveDatabase(connectionNameC)
}

func (ptr *QSqlDatabase) Rollback() bool {
	defer qt.Recovering("QSqlDatabase::rollback")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Rollback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) SetConnectOptions(options string) {
	defer qt.Recovering("QSqlDatabase::setConnectOptions")

	if ptr.Pointer() != nil {
		var optionsC = C.CString(options)
		defer C.free(unsafe.Pointer(optionsC))
		C.QSqlDatabase_SetConnectOptions(ptr.Pointer(), optionsC)
	}
}

func (ptr *QSqlDatabase) SetDatabaseName(name string) {
	defer qt.Recovering("QSqlDatabase::setDatabaseName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlDatabase_SetDatabaseName(ptr.Pointer(), nameC)
	}
}

func (ptr *QSqlDatabase) SetHostName(host string) {
	defer qt.Recovering("QSqlDatabase::setHostName")

	if ptr.Pointer() != nil {
		var hostC = C.CString(host)
		defer C.free(unsafe.Pointer(hostC))
		C.QSqlDatabase_SetHostName(ptr.Pointer(), hostC)
	}
}

func (ptr *QSqlDatabase) SetPassword(password string) {
	defer qt.Recovering("QSqlDatabase::setPassword")

	if ptr.Pointer() != nil {
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		C.QSqlDatabase_SetPassword(ptr.Pointer(), passwordC)
	}
}

func (ptr *QSqlDatabase) SetPort(port int) {
	defer qt.Recovering("QSqlDatabase::setPort")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_SetPort(ptr.Pointer(), C.int(int32(port)))
	}
}

func (ptr *QSqlDatabase) SetUserName(name string) {
	defer qt.Recovering("QSqlDatabase::setUserName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlDatabase_SetUserName(ptr.Pointer(), nameC)
	}
}

func (ptr *QSqlDatabase) Transaction() bool {
	defer qt.Recovering("QSqlDatabase::transaction")

	if ptr.Pointer() != nil {
		return C.QSqlDatabase_Transaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDatabase) UserName() string {
	defer qt.Recovering("QSqlDatabase::userName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDatabase_UserName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlDatabase) DestroyQSqlDatabase() {
	defer qt.Recovering("QSqlDatabase::~QSqlDatabase")

	if ptr.Pointer() != nil {
		C.QSqlDatabase_DestroyQSqlDatabase(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QSqlDatabase_DefaultConnection() string {
	defer qt.Recovering("QSqlDatabase::defaultConnection")

	return C.GoString(C.QSqlDatabase_QSqlDatabase_DefaultConnection())
}

func (ptr *QSqlDatabase) DefaultConnection() string {
	defer qt.Recovering("QSqlDatabase::defaultConnection")

	return C.GoString(C.QSqlDatabase_QSqlDatabase_DefaultConnection())
}

//QSqlDriver::DbmsType
type QSqlDriver__DbmsType int64

const (
	QSqlDriver__UnknownDbms = QSqlDriver__DbmsType(0)
	QSqlDriver__MSSqlServer = QSqlDriver__DbmsType(1)
	QSqlDriver__MySqlServer = QSqlDriver__DbmsType(2)
	QSqlDriver__PostgreSQL  = QSqlDriver__DbmsType(3)
	QSqlDriver__Oracle      = QSqlDriver__DbmsType(4)
	QSqlDriver__Sybase      = QSqlDriver__DbmsType(5)
	QSqlDriver__SQLite      = QSqlDriver__DbmsType(6)
	QSqlDriver__Interbase   = QSqlDriver__DbmsType(7)
	QSqlDriver__DB2         = QSqlDriver__DbmsType(8)
)

//QSqlDriver::DriverFeature
type QSqlDriver__DriverFeature int64

const (
	QSqlDriver__Transactions           = QSqlDriver__DriverFeature(0)
	QSqlDriver__QuerySize              = QSqlDriver__DriverFeature(1)
	QSqlDriver__BLOB                   = QSqlDriver__DriverFeature(2)
	QSqlDriver__Unicode                = QSqlDriver__DriverFeature(3)
	QSqlDriver__PreparedQueries        = QSqlDriver__DriverFeature(4)
	QSqlDriver__NamedPlaceholders      = QSqlDriver__DriverFeature(5)
	QSqlDriver__PositionalPlaceholders = QSqlDriver__DriverFeature(6)
	QSqlDriver__LastInsertId           = QSqlDriver__DriverFeature(7)
	QSqlDriver__BatchOperations        = QSqlDriver__DriverFeature(8)
	QSqlDriver__SimpleLocking          = QSqlDriver__DriverFeature(9)
	QSqlDriver__LowPrecisionNumbers    = QSqlDriver__DriverFeature(10)
	QSqlDriver__EventNotifications     = QSqlDriver__DriverFeature(11)
	QSqlDriver__FinishQuery            = QSqlDriver__DriverFeature(12)
	QSqlDriver__MultipleResultSets     = QSqlDriver__DriverFeature(13)
	QSqlDriver__CancelQuery            = QSqlDriver__DriverFeature(14)
)

//QSqlDriver::IdentifierType
type QSqlDriver__IdentifierType int64

const (
	QSqlDriver__FieldName = QSqlDriver__IdentifierType(0)
	QSqlDriver__TableName = QSqlDriver__IdentifierType(1)
)

//QSqlDriver::NotificationSource
type QSqlDriver__NotificationSource int64

const (
	QSqlDriver__UnknownSource = QSqlDriver__NotificationSource(0)
	QSqlDriver__SelfSource    = QSqlDriver__NotificationSource(1)
	QSqlDriver__OtherSource   = QSqlDriver__NotificationSource(2)
)

//QSqlDriver::StatementType
type QSqlDriver__StatementType int64

const (
	QSqlDriver__WhereStatement  = QSqlDriver__StatementType(0)
	QSqlDriver__SelectStatement = QSqlDriver__StatementType(1)
	QSqlDriver__UpdateStatement = QSqlDriver__StatementType(2)
	QSqlDriver__InsertStatement = QSqlDriver__StatementType(3)
	QSqlDriver__DeleteStatement = QSqlDriver__StatementType(4)
)

type QSqlDriver struct {
	core.QObject
}

type QSqlDriver_ITF interface {
	core.QObject_ITF
	QSqlDriver_PTR() *QSqlDriver
}

func (p *QSqlDriver) QSqlDriver_PTR() *QSqlDriver {
	return p
}

func (p *QSqlDriver) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSqlDriver) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlDriver(ptr QSqlDriver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriver_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverFromPointer(ptr unsafe.Pointer) *QSqlDriver {
	var n = new(QSqlDriver)
	n.SetPointer(ptr)
	return n
}
func NewQSqlDriver(parent core.QObject_ITF) *QSqlDriver {
	defer qt.Recovering("QSqlDriver::QSqlDriver")

	var tmpValue = NewQSqlDriverFromPointer(C.QSqlDriver_NewQSqlDriver(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlDriver_BeginTransaction
func callbackQSqlDriver_BeginTransaction(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::beginTransaction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::beginTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).BeginTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectBeginTransaction(f func() bool) {
	defer qt.Recovering("connect QSqlDriver::beginTransaction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::beginTransaction", f)
	}
}

func (ptr *QSqlDriver) DisconnectBeginTransaction() {
	defer qt.Recovering("disconnect QSqlDriver::beginTransaction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::beginTransaction")
	}
}

func (ptr *QSqlDriver) BeginTransaction() bool {
	defer qt.Recovering("QSqlDriver::beginTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_BeginTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) BeginTransactionDefault() bool {
	defer qt.Recovering("QSqlDriver::beginTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_BeginTransactionDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlDriver_Close
func callbackQSqlDriver_Close(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::close")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSqlDriver) ConnectClose(f func()) {
	defer qt.Recovering("connect QSqlDriver::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::close", f)
	}
}

func (ptr *QSqlDriver) DisconnectClose() {
	defer qt.Recovering("disconnect QSqlDriver::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::close")
	}
}

func (ptr *QSqlDriver) Close() {
	defer qt.Recovering("QSqlDriver::close")

	if ptr.Pointer() != nil {
		C.QSqlDriver_Close(ptr.Pointer())
	}
}

//export callbackQSqlDriver_CommitTransaction
func callbackQSqlDriver_CommitTransaction(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::commitTransaction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::commitTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).CommitTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectCommitTransaction(f func() bool) {
	defer qt.Recovering("connect QSqlDriver::commitTransaction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::commitTransaction", f)
	}
}

func (ptr *QSqlDriver) DisconnectCommitTransaction() {
	defer qt.Recovering("disconnect QSqlDriver::commitTransaction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::commitTransaction")
	}
}

func (ptr *QSqlDriver) CommitTransaction() bool {
	defer qt.Recovering("QSqlDriver::commitTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_CommitTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) CommitTransactionDefault() bool {
	defer qt.Recovering("QSqlDriver::commitTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_CommitTransactionDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlDriver_CreateResult
func callbackQSqlDriver_CreateResult(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriver::createResult")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::createResult"); signal != nil {
		return PointerFromQSqlResult(signal.(func() *QSqlResult)())
	}

	return PointerFromQSqlResult(nil)
}

func (ptr *QSqlDriver) ConnectCreateResult(f func() *QSqlResult) {
	defer qt.Recovering("connect QSqlDriver::createResult")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::createResult", f)
	}
}

func (ptr *QSqlDriver) DisconnectCreateResult() {
	defer qt.Recovering("disconnect QSqlDriver::createResult")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::createResult")
	}
}

func (ptr *QSqlDriver) CreateResult() *QSqlResult {
	defer qt.Recovering("QSqlDriver::createResult")

	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlDriver_CreateResult(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) DbmsType() QSqlDriver__DbmsType {
	defer qt.Recovering("QSqlDriver::dbmsType")

	if ptr.Pointer() != nil {
		return QSqlDriver__DbmsType(C.QSqlDriver_DbmsType(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlDriver_EscapeIdentifier
func callbackQSqlDriver_EscapeIdentifier(ptr unsafe.Pointer, identifier *C.char, ty C.longlong) *C.char {
	defer qt.Recovering("callback QSqlDriver::escapeIdentifier")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::escapeIdentifier"); signal != nil {
		return C.CString(signal.(func(string, QSqlDriver__IdentifierType) string)(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))
	}

	return C.CString(NewQSqlDriverFromPointer(ptr).EscapeIdentifierDefault(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))
}

func (ptr *QSqlDriver) ConnectEscapeIdentifier(f func(identifier string, ty QSqlDriver__IdentifierType) string) {
	defer qt.Recovering("connect QSqlDriver::escapeIdentifier")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::escapeIdentifier", f)
	}
}

func (ptr *QSqlDriver) DisconnectEscapeIdentifier() {
	defer qt.Recovering("disconnect QSqlDriver::escapeIdentifier")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::escapeIdentifier")
	}
}

func (ptr *QSqlDriver) EscapeIdentifier(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::escapeIdentifier")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.GoString(C.QSqlDriver_EscapeIdentifier(ptr.Pointer(), identifierC, C.longlong(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) EscapeIdentifierDefault(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::escapeIdentifier")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.GoString(C.QSqlDriver_EscapeIdentifierDefault(ptr.Pointer(), identifierC, C.longlong(ty)))
	}
	return ""
}

//export callbackQSqlDriver_FormatValue
func callbackQSqlDriver_FormatValue(ptr unsafe.Pointer, field unsafe.Pointer, trimStrings C.char) *C.char {
	defer qt.Recovering("callback QSqlDriver::formatValue")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::formatValue"); signal != nil {
		return C.CString(signal.(func(*QSqlField, bool) string)(NewQSqlFieldFromPointer(field), int8(trimStrings) != 0))
	}

	return C.CString(NewQSqlDriverFromPointer(ptr).FormatValueDefault(NewQSqlFieldFromPointer(field), int8(trimStrings) != 0))
}

func (ptr *QSqlDriver) ConnectFormatValue(f func(field *QSqlField, trimStrings bool) string) {
	defer qt.Recovering("connect QSqlDriver::formatValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::formatValue", f)
	}
}

func (ptr *QSqlDriver) DisconnectFormatValue() {
	defer qt.Recovering("disconnect QSqlDriver::formatValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::formatValue")
	}
}

func (ptr *QSqlDriver) FormatValue(field QSqlField_ITF, trimStrings bool) string {
	defer qt.Recovering("QSqlDriver::formatValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_FormatValue(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(trimStrings)))))
	}
	return ""
}

func (ptr *QSqlDriver) FormatValueDefault(field QSqlField_ITF, trimStrings bool) string {
	defer qt.Recovering("QSqlDriver::formatValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlDriver_FormatValueDefault(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(trimStrings)))))
	}
	return ""
}

//export callbackQSqlDriver_Handle
func callbackQSqlDriver_Handle(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriver::handle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::handle"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlDriverFromPointer(ptr).HandleDefault())
}

func (ptr *QSqlDriver) ConnectHandle(f func() *core.QVariant) {
	defer qt.Recovering("connect QSqlDriver::handle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::handle", f)
	}
}

func (ptr *QSqlDriver) DisconnectHandle() {
	defer qt.Recovering("disconnect QSqlDriver::handle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::handle")
	}
}

func (ptr *QSqlDriver) Handle() *core.QVariant {
	defer qt.Recovering("QSqlDriver::handle")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlDriver_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) HandleDefault() *core.QVariant {
	defer qt.Recovering("QSqlDriver::handle")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlDriver_HandleDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_HasFeature
func callbackQSqlDriver_HasFeature(ptr unsafe.Pointer, feature C.longlong) C.char {
	defer qt.Recovering("callback QSqlDriver::hasFeature")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::hasFeature"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(QSqlDriver__DriverFeature) bool)(QSqlDriver__DriverFeature(feature)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlDriver) ConnectHasFeature(f func(feature QSqlDriver__DriverFeature) bool) {
	defer qt.Recovering("connect QSqlDriver::hasFeature")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::hasFeature", f)
	}
}

func (ptr *QSqlDriver) DisconnectHasFeature(feature QSqlDriver__DriverFeature) {
	defer qt.Recovering("disconnect QSqlDriver::hasFeature")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::hasFeature")
	}
}

func (ptr *QSqlDriver) HasFeature(feature QSqlDriver__DriverFeature) bool {
	defer qt.Recovering("QSqlDriver::hasFeature")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_HasFeature(ptr.Pointer(), C.longlong(feature)) != 0
	}
	return false
}

//export callbackQSqlDriver_IsIdentifierEscaped
func callbackQSqlDriver_IsIdentifierEscaped(ptr unsafe.Pointer, identifier *C.char, ty C.longlong) C.char {
	defer qt.Recovering("callback QSqlDriver::isIdentifierEscaped")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::isIdentifierEscaped"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, QSqlDriver__IdentifierType) bool)(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).IsIdentifierEscapedDefault(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))))
}

func (ptr *QSqlDriver) ConnectIsIdentifierEscaped(f func(identifier string, ty QSqlDriver__IdentifierType) bool) {
	defer qt.Recovering("connect QSqlDriver::isIdentifierEscaped")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::isIdentifierEscaped", f)
	}
}

func (ptr *QSqlDriver) DisconnectIsIdentifierEscaped() {
	defer qt.Recovering("disconnect QSqlDriver::isIdentifierEscaped")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::isIdentifierEscaped")
	}
}

func (ptr *QSqlDriver) IsIdentifierEscaped(identifier string, ty QSqlDriver__IdentifierType) bool {
	defer qt.Recovering("QSqlDriver::isIdentifierEscaped")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.QSqlDriver_IsIdentifierEscaped(ptr.Pointer(), identifierC, C.longlong(ty)) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsIdentifierEscapedDefault(identifier string, ty QSqlDriver__IdentifierType) bool {
	defer qt.Recovering("QSqlDriver::isIdentifierEscaped")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.QSqlDriver_IsIdentifierEscapedDefault(ptr.Pointer(), identifierC, C.longlong(ty)) != 0
	}
	return false
}

//export callbackQSqlDriver_IsOpen
func callbackQSqlDriver_IsOpen(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::isOpen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::isOpen"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).IsOpenDefault())))
}

func (ptr *QSqlDriver) ConnectIsOpen(f func() bool) {
	defer qt.Recovering("connect QSqlDriver::isOpen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::isOpen", f)
	}
}

func (ptr *QSqlDriver) DisconnectIsOpen() {
	defer qt.Recovering("disconnect QSqlDriver::isOpen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::isOpen")
	}
}

func (ptr *QSqlDriver) IsOpen() bool {
	defer qt.Recovering("QSqlDriver::isOpen")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenDefault() bool {
	defer qt.Recovering("QSqlDriver::isOpen")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpenDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) IsOpenError() bool {
	defer qt.Recovering("QSqlDriver::isOpenError")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_IsOpenError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) LastError() *QSqlError {
	defer qt.Recovering("QSqlDriver::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlErrorFromPointer(C.QSqlDriver_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_Notification
func callbackQSqlDriver_Notification(ptr unsafe.Pointer, name *C.char) {
	defer qt.Recovering("callback QSqlDriver::notification")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::notification"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QSqlDriver) ConnectNotification(f func(name string)) {
	defer qt.Recovering("connect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotification(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::notification", f)
	}
}

func (ptr *QSqlDriver) DisconnectNotification() {
	defer qt.Recovering("disconnect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::notification")
	}
}

func (ptr *QSqlDriver) Notification(name string) {
	defer qt.Recovering("QSqlDriver::notification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlDriver_Notification(ptr.Pointer(), nameC)
	}
}

//export callbackQSqlDriver_Notification2
func callbackQSqlDriver_Notification2(ptr unsafe.Pointer, name *C.char, source C.longlong, payload unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::notification")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::notification2"); signal != nil {
		signal.(func(string, QSqlDriver__NotificationSource, *core.QVariant))(C.GoString(name), QSqlDriver__NotificationSource(source), core.NewQVariantFromPointer(payload))
	}

}

func (ptr *QSqlDriver) ConnectNotification2(f func(name string, source QSqlDriver__NotificationSource, payload *core.QVariant)) {
	defer qt.Recovering("connect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotification2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::notification2", f)
	}
}

func (ptr *QSqlDriver) DisconnectNotification2() {
	defer qt.Recovering("disconnect QSqlDriver::notification")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotification2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::notification2")
	}
}

func (ptr *QSqlDriver) Notification2(name string, source QSqlDriver__NotificationSource, payload core.QVariant_ITF) {
	defer qt.Recovering("QSqlDriver::notification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlDriver_Notification2(ptr.Pointer(), nameC, C.longlong(source), core.PointerFromQVariant(payload))
	}
}

//export callbackQSqlDriver_Open
func callbackQSqlDriver_Open(ptr unsafe.Pointer, db *C.char, user *C.char, password *C.char, host *C.char, port C.int, options *C.char) C.char {
	defer qt.Recovering("callback QSqlDriver::open")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string, string, string, string, int, string) bool)(C.GoString(db), C.GoString(user), C.GoString(password), C.GoString(host), int(int32(port)), C.GoString(options)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlDriver) ConnectOpen(f func(db string, user string, password string, host string, port int, options string) bool) {
	defer qt.Recovering("connect QSqlDriver::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::open", f)
	}
}

func (ptr *QSqlDriver) DisconnectOpen(db string, user string, password string, host string, port int, options string) {
	defer qt.Recovering("disconnect QSqlDriver::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::open")
	}
}

func (ptr *QSqlDriver) Open(db string, user string, password string, host string, port int, options string) bool {
	defer qt.Recovering("QSqlDriver::open")

	if ptr.Pointer() != nil {
		var dbC = C.CString(db)
		defer C.free(unsafe.Pointer(dbC))
		var userC = C.CString(user)
		defer C.free(unsafe.Pointer(userC))
		var passwordC = C.CString(password)
		defer C.free(unsafe.Pointer(passwordC))
		var hostC = C.CString(host)
		defer C.free(unsafe.Pointer(hostC))
		var optionsC = C.CString(options)
		defer C.free(unsafe.Pointer(optionsC))
		return C.QSqlDriver_Open(ptr.Pointer(), dbC, userC, passwordC, hostC, C.int(int32(port)), optionsC) != 0
	}
	return false
}

//export callbackQSqlDriver_PrimaryIndex
func callbackQSqlDriver_PrimaryIndex(ptr unsafe.Pointer, tableName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriver::primaryIndex")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::primaryIndex"); signal != nil {
		return PointerFromQSqlIndex(signal.(func(string) *QSqlIndex)(C.GoString(tableName)))
	}

	return PointerFromQSqlIndex(NewQSqlDriverFromPointer(ptr).PrimaryIndexDefault(C.GoString(tableName)))
}

func (ptr *QSqlDriver) ConnectPrimaryIndex(f func(tableName string) *QSqlIndex) {
	defer qt.Recovering("connect QSqlDriver::primaryIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::primaryIndex", f)
	}
}

func (ptr *QSqlDriver) DisconnectPrimaryIndex() {
	defer qt.Recovering("disconnect QSqlDriver::primaryIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::primaryIndex")
	}
}

func (ptr *QSqlDriver) PrimaryIndex(tableName string) *QSqlIndex {
	defer qt.Recovering("QSqlDriver::primaryIndex")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		var tmpValue = NewQSqlIndexFromPointer(C.QSqlDriver_PrimaryIndex(ptr.Pointer(), tableNameC))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) PrimaryIndexDefault(tableName string) *QSqlIndex {
	defer qt.Recovering("QSqlDriver::primaryIndex")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		var tmpValue = NewQSqlIndexFromPointer(C.QSqlDriver_PrimaryIndexDefault(ptr.Pointer(), tableNameC))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_Record
func callbackQSqlDriver_Record(ptr unsafe.Pointer, tableName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriver::record")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::record"); signal != nil {
		return PointerFromQSqlRecord(signal.(func(string) *QSqlRecord)(C.GoString(tableName)))
	}

	return PointerFromQSqlRecord(NewQSqlDriverFromPointer(ptr).RecordDefault(C.GoString(tableName)))
}

func (ptr *QSqlDriver) ConnectRecord(f func(tableName string) *QSqlRecord) {
	defer qt.Recovering("connect QSqlDriver::record")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::record", f)
	}
}

func (ptr *QSqlDriver) DisconnectRecord() {
	defer qt.Recovering("disconnect QSqlDriver::record")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::record")
	}
}

func (ptr *QSqlDriver) Record(tableName string) *QSqlRecord {
	defer qt.Recovering("QSqlDriver::record")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlDriver_Record(ptr.Pointer(), tableNameC))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriver) RecordDefault(tableName string) *QSqlRecord {
	defer qt.Recovering("QSqlDriver::record")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlDriver_RecordDefault(ptr.Pointer(), tableNameC))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriver_RollbackTransaction
func callbackQSqlDriver_RollbackTransaction(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::rollbackTransaction")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::rollbackTransaction"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).RollbackTransactionDefault())))
}

func (ptr *QSqlDriver) ConnectRollbackTransaction(f func() bool) {
	defer qt.Recovering("connect QSqlDriver::rollbackTransaction")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::rollbackTransaction", f)
	}
}

func (ptr *QSqlDriver) DisconnectRollbackTransaction() {
	defer qt.Recovering("disconnect QSqlDriver::rollbackTransaction")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::rollbackTransaction")
	}
}

func (ptr *QSqlDriver) RollbackTransaction() bool {
	defer qt.Recovering("QSqlDriver::rollbackTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_RollbackTransaction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlDriver) RollbackTransactionDefault() bool {
	defer qt.Recovering("QSqlDriver::rollbackTransaction")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_RollbackTransactionDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlDriver_SetLastError
func callbackQSqlDriver_SetLastError(ptr unsafe.Pointer, error unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::setLastError")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::setLastError"); signal != nil {
		signal.(func(*QSqlError))(NewQSqlErrorFromPointer(error))
	} else {
		NewQSqlDriverFromPointer(ptr).SetLastErrorDefault(NewQSqlErrorFromPointer(error))
	}
}

func (ptr *QSqlDriver) ConnectSetLastError(f func(error *QSqlError)) {
	defer qt.Recovering("connect QSqlDriver::setLastError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setLastError", f)
	}
}

func (ptr *QSqlDriver) DisconnectSetLastError() {
	defer qt.Recovering("disconnect QSqlDriver::setLastError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setLastError")
	}
}

func (ptr *QSqlDriver) SetLastError(error QSqlError_ITF) {
	defer qt.Recovering("QSqlDriver::setLastError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlDriver) SetLastErrorDefault(error QSqlError_ITF) {
	defer qt.Recovering("QSqlDriver::setLastError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetLastErrorDefault(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

//export callbackQSqlDriver_SetOpen
func callbackQSqlDriver_SetOpen(ptr unsafe.Pointer, open C.char) {
	defer qt.Recovering("callback QSqlDriver::setOpen")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::setOpen"); signal != nil {
		signal.(func(bool))(int8(open) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenDefault(int8(open) != 0)
	}
}

func (ptr *QSqlDriver) ConnectSetOpen(f func(open bool)) {
	defer qt.Recovering("connect QSqlDriver::setOpen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setOpen", f)
	}
}

func (ptr *QSqlDriver) DisconnectSetOpen() {
	defer qt.Recovering("disconnect QSqlDriver::setOpen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setOpen")
	}
}

func (ptr *QSqlDriver) SetOpen(open bool) {
	defer qt.Recovering("QSqlDriver::setOpen")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpen(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(open))))
	}
}

func (ptr *QSqlDriver) SetOpenDefault(open bool) {
	defer qt.Recovering("QSqlDriver::setOpen")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(open))))
	}
}

//export callbackQSqlDriver_SetOpenError
func callbackQSqlDriver_SetOpenError(ptr unsafe.Pointer, error C.char) {
	defer qt.Recovering("callback QSqlDriver::setOpenError")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::setOpenError"); signal != nil {
		signal.(func(bool))(int8(error) != 0)
	} else {
		NewQSqlDriverFromPointer(ptr).SetOpenErrorDefault(int8(error) != 0)
	}
}

func (ptr *QSqlDriver) ConnectSetOpenError(f func(error bool)) {
	defer qt.Recovering("connect QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setOpenError", f)
	}
}

func (ptr *QSqlDriver) DisconnectSetOpenError() {
	defer qt.Recovering("disconnect QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::setOpenError")
	}
}

func (ptr *QSqlDriver) SetOpenError(error bool) {
	defer qt.Recovering("QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenError(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(error))))
	}
}

func (ptr *QSqlDriver) SetOpenErrorDefault(error bool) {
	defer qt.Recovering("QSqlDriver::setOpenError")

	if ptr.Pointer() != nil {
		C.QSqlDriver_SetOpenErrorDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(error))))
	}
}

//export callbackQSqlDriver_SqlStatement
func callbackQSqlDriver_SqlStatement(ptr unsafe.Pointer, ty C.longlong, tableName *C.char, rec unsafe.Pointer, preparedStatement C.char) *C.char {
	defer qt.Recovering("callback QSqlDriver::sqlStatement")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::sqlStatement"); signal != nil {
		return C.CString(signal.(func(QSqlDriver__StatementType, string, *QSqlRecord, bool) string)(QSqlDriver__StatementType(ty), C.GoString(tableName), NewQSqlRecordFromPointer(rec), int8(preparedStatement) != 0))
	}

	return C.CString(NewQSqlDriverFromPointer(ptr).SqlStatementDefault(QSqlDriver__StatementType(ty), C.GoString(tableName), NewQSqlRecordFromPointer(rec), int8(preparedStatement) != 0))
}

func (ptr *QSqlDriver) ConnectSqlStatement(f func(ty QSqlDriver__StatementType, tableName string, rec *QSqlRecord, preparedStatement bool) string) {
	defer qt.Recovering("connect QSqlDriver::sqlStatement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::sqlStatement", f)
	}
}

func (ptr *QSqlDriver) DisconnectSqlStatement() {
	defer qt.Recovering("disconnect QSqlDriver::sqlStatement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::sqlStatement")
	}
}

func (ptr *QSqlDriver) SqlStatement(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	defer qt.Recovering("QSqlDriver::sqlStatement")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		return C.GoString(C.QSqlDriver_SqlStatement(ptr.Pointer(), C.longlong(ty), tableNameC, PointerFromQSqlRecord(rec), C.char(int8(qt.GoBoolToInt(preparedStatement)))))
	}
	return ""
}

func (ptr *QSqlDriver) SqlStatementDefault(ty QSqlDriver__StatementType, tableName string, rec QSqlRecord_ITF, preparedStatement bool) string {
	defer qt.Recovering("QSqlDriver::sqlStatement")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		return C.GoString(C.QSqlDriver_SqlStatementDefault(ptr.Pointer(), C.longlong(ty), tableNameC, PointerFromQSqlRecord(rec), C.char(int8(qt.GoBoolToInt(preparedStatement)))))
	}
	return ""
}

//export callbackQSqlDriver_StripDelimiters
func callbackQSqlDriver_StripDelimiters(ptr unsafe.Pointer, identifier *C.char, ty C.longlong) *C.char {
	defer qt.Recovering("callback QSqlDriver::stripDelimiters")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::stripDelimiters"); signal != nil {
		return C.CString(signal.(func(string, QSqlDriver__IdentifierType) string)(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))
	}

	return C.CString(NewQSqlDriverFromPointer(ptr).StripDelimitersDefault(C.GoString(identifier), QSqlDriver__IdentifierType(ty)))
}

func (ptr *QSqlDriver) ConnectStripDelimiters(f func(identifier string, ty QSqlDriver__IdentifierType) string) {
	defer qt.Recovering("connect QSqlDriver::stripDelimiters")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::stripDelimiters", f)
	}
}

func (ptr *QSqlDriver) DisconnectStripDelimiters() {
	defer qt.Recovering("disconnect QSqlDriver::stripDelimiters")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::stripDelimiters")
	}
}

func (ptr *QSqlDriver) StripDelimiters(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::stripDelimiters")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.GoString(C.QSqlDriver_StripDelimiters(ptr.Pointer(), identifierC, C.longlong(ty)))
	}
	return ""
}

func (ptr *QSqlDriver) StripDelimitersDefault(identifier string, ty QSqlDriver__IdentifierType) string {
	defer qt.Recovering("QSqlDriver::stripDelimiters")

	if ptr.Pointer() != nil {
		var identifierC = C.CString(identifier)
		defer C.free(unsafe.Pointer(identifierC))
		return C.GoString(C.QSqlDriver_StripDelimitersDefault(ptr.Pointer(), identifierC, C.longlong(ty)))
	}
	return ""
}

//export callbackQSqlDriver_SubscribeToNotification
func callbackQSqlDriver_SubscribeToNotification(ptr unsafe.Pointer, name *C.char) C.char {
	defer qt.Recovering("callback QSqlDriver::subscribeToNotification")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::subscribeToNotification"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).SubscribeToNotificationDefault(C.GoString(name)))))
}

func (ptr *QSqlDriver) ConnectSubscribeToNotification(f func(name string) bool) {
	defer qt.Recovering("connect QSqlDriver::subscribeToNotification")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::subscribeToNotification", f)
	}
}

func (ptr *QSqlDriver) DisconnectSubscribeToNotification() {
	defer qt.Recovering("disconnect QSqlDriver::subscribeToNotification")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::subscribeToNotification")
	}
}

func (ptr *QSqlDriver) SubscribeToNotification(name string) bool {
	defer qt.Recovering("QSqlDriver::subscribeToNotification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlDriver_SubscribeToNotification(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlDriver) SubscribeToNotificationDefault(name string) bool {
	defer qt.Recovering("QSqlDriver::subscribeToNotification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlDriver_SubscribeToNotificationDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

//export callbackQSqlDriver_SubscribedToNotifications
func callbackQSqlDriver_SubscribedToNotifications(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlDriver::subscribedToNotifications")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::subscribedToNotifications"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQSqlDriverFromPointer(ptr).SubscribedToNotificationsDefault(), "|"))
}

func (ptr *QSqlDriver) ConnectSubscribedToNotifications(f func() []string) {
	defer qt.Recovering("connect QSqlDriver::subscribedToNotifications")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::subscribedToNotifications", f)
	}
}

func (ptr *QSqlDriver) DisconnectSubscribedToNotifications() {
	defer qt.Recovering("disconnect QSqlDriver::subscribedToNotifications")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::subscribedToNotifications")
	}
}

func (ptr *QSqlDriver) SubscribedToNotifications() []string {
	defer qt.Recovering("QSqlDriver::subscribedToNotifications")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlDriver_SubscribedToNotifications(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlDriver) SubscribedToNotificationsDefault() []string {
	defer qt.Recovering("QSqlDriver::subscribedToNotifications")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlDriver_SubscribedToNotificationsDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlDriver_UnsubscribeFromNotification
func callbackQSqlDriver_UnsubscribeFromNotification(ptr unsafe.Pointer, name *C.char) C.char {
	defer qt.Recovering("callback QSqlDriver::unsubscribeFromNotification")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::unsubscribeFromNotification"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(name)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).UnsubscribeFromNotificationDefault(C.GoString(name)))))
}

func (ptr *QSqlDriver) ConnectUnsubscribeFromNotification(f func(name string) bool) {
	defer qt.Recovering("connect QSqlDriver::unsubscribeFromNotification")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::unsubscribeFromNotification", f)
	}
}

func (ptr *QSqlDriver) DisconnectUnsubscribeFromNotification() {
	defer qt.Recovering("disconnect QSqlDriver::unsubscribeFromNotification")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::unsubscribeFromNotification")
	}
}

func (ptr *QSqlDriver) UnsubscribeFromNotification(name string) bool {
	defer qt.Recovering("QSqlDriver::unsubscribeFromNotification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlDriver_UnsubscribeFromNotification(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlDriver) UnsubscribeFromNotificationDefault(name string) bool {
	defer qt.Recovering("QSqlDriver::unsubscribeFromNotification")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlDriver_UnsubscribeFromNotificationDefault(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlDriver) DestroyQSqlDriver() {
	defer qt.Recovering("QSqlDriver::~QSqlDriver")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DestroyQSqlDriver(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlDriver_TimerEvent
func callbackQSqlDriver_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::timerEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::timerEvent")
	}
}

func (ptr *QSqlDriver) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriver) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriver::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlDriver_ChildEvent
func callbackQSqlDriver_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlDriver::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::childEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlDriver::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::childEvent")
	}
}

func (ptr *QSqlDriver) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriver::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriver) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriver::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlDriver_ConnectNotify
func callbackQSqlDriver_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriver) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlDriver::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::connectNotify", f)
	}
}

func (ptr *QSqlDriver) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlDriver::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::connectNotify")
	}
}

func (ptr *QSqlDriver) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriver::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlDriver) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriver::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriver_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriver_CustomEvent
func callbackQSqlDriver_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriver) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlDriver::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::customEvent", f)
	}
}

func (ptr *QSqlDriver) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlDriver::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::customEvent")
	}
}

func (ptr *QSqlDriver) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriver::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlDriver) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriver::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriver_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlDriver_DeleteLater
func callbackQSqlDriver_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlDriver) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlDriver::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::deleteLater", f)
	}
}

func (ptr *QSqlDriver) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlDriver::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::deleteLater")
	}
}

func (ptr *QSqlDriver) DeleteLater() {
	defer qt.Recovering("QSqlDriver::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriver) DeleteLaterDefault() {
	defer qt.Recovering("QSqlDriver::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlDriver_DisconnectNotify
func callbackQSqlDriver_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriver::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriver) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlDriver::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::disconnectNotify", f)
	}
}

func (ptr *QSqlDriver) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlDriver::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::disconnectNotify")
	}
}

func (ptr *QSqlDriver) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriver::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlDriver) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriver::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriver_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriver_Event
func callbackQSqlDriver_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlDriver) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlDriver::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::event", f)
	}
}

func (ptr *QSqlDriver) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlDriver::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::event")
	}
}

func (ptr *QSqlDriver) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriver::event")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlDriver) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriver::event")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlDriver_EventFilter
func callbackQSqlDriver_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriver::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlDriver) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlDriver::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::eventFilter", f)
	}
}

func (ptr *QSqlDriver) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSqlDriver::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::eventFilter")
	}
}

func (ptr *QSqlDriver) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriver::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSqlDriver) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriver::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlDriver_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSqlDriver_MetaObject
func callbackQSqlDriver_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriver::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriver::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlDriverFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlDriver) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlDriver::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::metaObject", f)
	}
}

func (ptr *QSqlDriver) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlDriver::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriver::metaObject")
	}
}

func (ptr *QSqlDriver) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlDriver::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriver_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriver) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlDriver::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriver_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSqlDriverCreator struct {
	QSqlDriverCreatorBase
}

type QSqlDriverCreator_ITF interface {
	QSqlDriverCreatorBase_ITF
	QSqlDriverCreator_PTR() *QSqlDriverCreator
}

func (p *QSqlDriverCreator) QSqlDriverCreator_PTR() *QSqlDriverCreator {
	return p
}

func (p *QSqlDriverCreator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func (p *QSqlDriverCreator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSqlDriverCreatorBase_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlDriverCreator(ptr QSqlDriverCreator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreator_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorFromPointer(ptr unsafe.Pointer) *QSqlDriverCreator {
	var n = new(QSqlDriverCreator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlDriverCreator) DestroyQSqlDriverCreator() {
	C.free(ptr.Pointer())
	qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
	ptr.SetPointer(nil)
}

type QSqlDriverCreatorBase struct {
	ptr unsafe.Pointer
}

type QSqlDriverCreatorBase_ITF interface {
	QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase
}

func (p *QSqlDriverCreatorBase) QSqlDriverCreatorBase_PTR() *QSqlDriverCreatorBase {
	return p
}

func (p *QSqlDriverCreatorBase) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlDriverCreatorBase) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlDriverCreatorBase(ptr QSqlDriverCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverCreatorBaseFromPointer(ptr unsafe.Pointer) *QSqlDriverCreatorBase {
	var n = new(QSqlDriverCreatorBase)
	n.SetPointer(ptr)
	return n
}

//export callbackQSqlDriverCreatorBase_CreateObject
func callbackQSqlDriverCreatorBase_CreateObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriverCreatorBase::createObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverCreatorBase::createObject"); signal != nil {
		return PointerFromQSqlDriver(signal.(func() *QSqlDriver)())
	}

	return PointerFromQSqlDriver(nil)
}

func (ptr *QSqlDriverCreatorBase) ConnectCreateObject(f func() *QSqlDriver) {
	defer qt.Recovering("connect QSqlDriverCreatorBase::createObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverCreatorBase::createObject", f)
	}
}

func (ptr *QSqlDriverCreatorBase) DisconnectCreateObject() {
	defer qt.Recovering("disconnect QSqlDriverCreatorBase::createObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverCreatorBase::createObject")
	}
}

func (ptr *QSqlDriverCreatorBase) CreateObject() *QSqlDriver {
	defer qt.Recovering("QSqlDriverCreatorBase::createObject")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlDriverFromPointer(C.QSqlDriverCreatorBase_CreateObject(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase
func callbackQSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverCreatorBase::~QSqlDriverCreatorBase"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverCreatorBaseFromPointer(ptr).DestroyQSqlDriverCreatorBaseDefault()
	}
}

func (ptr *QSqlDriverCreatorBase) ConnectDestroyQSqlDriverCreatorBase(f func()) {
	defer qt.Recovering("connect QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverCreatorBase::~QSqlDriverCreatorBase", f)
	}
}

func (ptr *QSqlDriverCreatorBase) DisconnectDestroyQSqlDriverCreatorBase() {
	defer qt.Recovering("disconnect QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverCreatorBase::~QSqlDriverCreatorBase")
	}
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBase() {
	defer qt.Recovering("QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriverCreatorBase) DestroyQSqlDriverCreatorBaseDefault() {
	defer qt.Recovering("QSqlDriverCreatorBase::~QSqlDriverCreatorBase")

	if ptr.Pointer() != nil {
		C.QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBaseDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QSqlDriverPlugin struct {
	core.QObject
}

type QSqlDriverPlugin_ITF interface {
	core.QObject_ITF
	QSqlDriverPlugin_PTR() *QSqlDriverPlugin
}

func (p *QSqlDriverPlugin) QSqlDriverPlugin_PTR() *QSqlDriverPlugin {
	return p
}

func (p *QSqlDriverPlugin) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSqlDriverPlugin) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlDriverPlugin(ptr QSqlDriverPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlDriverPlugin_PTR().Pointer()
	}
	return nil
}

func NewQSqlDriverPluginFromPointer(ptr unsafe.Pointer) *QSqlDriverPlugin {
	var n = new(QSqlDriverPlugin)
	n.SetPointer(ptr)
	return n
}
func NewQSqlDriverPlugin(parent core.QObject_ITF) *QSqlDriverPlugin {
	defer qt.Recovering("QSqlDriverPlugin::QSqlDriverPlugin")

	var tmpValue = NewQSqlDriverPluginFromPointer(C.QSqlDriverPlugin_NewQSqlDriverPlugin(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlDriverPlugin_Create
func callbackQSqlDriverPlugin_Create(ptr unsafe.Pointer, key *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriverPlugin::create")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::create"); signal != nil {
		return PointerFromQSqlDriver(signal.(func(string) *QSqlDriver)(C.GoString(key)))
	}

	return PointerFromQSqlDriver(nil)
}

func (ptr *QSqlDriverPlugin) ConnectCreate(f func(key string) *QSqlDriver) {
	defer qt.Recovering("connect QSqlDriverPlugin::create")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::create", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectCreate(key string) {
	defer qt.Recovering("disconnect QSqlDriverPlugin::create")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::create")
	}
}

func (ptr *QSqlDriverPlugin) Create(key string) *QSqlDriver {
	defer qt.Recovering("QSqlDriverPlugin::create")

	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		var tmpValue = NewQSqlDriverFromPointer(C.QSqlDriverPlugin_Create(ptr.Pointer(), keyC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlDriverPlugin) DestroyQSqlDriverPlugin() {
	defer qt.Recovering("QSqlDriverPlugin::~QSqlDriverPlugin")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DestroyQSqlDriverPlugin(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlDriverPlugin_TimerEvent
func callbackQSqlDriverPlugin_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::timerEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::timerEvent")
	}
}

func (ptr *QSqlDriverPlugin) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlDriverPlugin_ChildEvent
func callbackQSqlDriverPlugin_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::childEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::childEvent")
	}
}

func (ptr *QSqlDriverPlugin) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlDriverPlugin_ConnectNotify
func callbackQSqlDriverPlugin_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriverPlugin) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlDriverPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::connectNotify", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::connectNotify")
	}
}

func (ptr *QSqlDriverPlugin) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlDriverPlugin) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriverPlugin_CustomEvent
func callbackQSqlDriverPlugin_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlDriverPlugin) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::customEvent", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::customEvent")
	}
}

func (ptr *QSqlDriverPlugin) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlDriverPlugin) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlDriverPlugin_DeleteLater
func callbackQSqlDriverPlugin_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlDriverPluginFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlDriverPlugin) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlDriverPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::deleteLater", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::deleteLater")
	}
}

func (ptr *QSqlDriverPlugin) DeleteLater() {
	defer qt.Recovering("QSqlDriverPlugin::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlDriverPlugin) DeleteLaterDefault() {
	defer qt.Recovering("QSqlDriverPlugin::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlDriverPlugin_DisconnectNotify
func callbackQSqlDriverPlugin_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlDriverPlugin::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlDriverPluginFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlDriverPlugin) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlDriverPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::disconnectNotify", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::disconnectNotify")
	}
}

func (ptr *QSqlDriverPlugin) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlDriverPlugin) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlDriverPlugin::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlDriverPlugin_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlDriverPlugin_Event
func callbackQSqlDriverPlugin_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriverPlugin::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverPluginFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlDriverPlugin) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlDriverPlugin::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::event", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::event")
	}
}

func (ptr *QSqlDriverPlugin) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriverPlugin::event")

	if ptr.Pointer() != nil {
		return C.QSqlDriverPlugin_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlDriverPlugin) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriverPlugin::event")

	if ptr.Pointer() != nil {
		return C.QSqlDriverPlugin_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlDriverPlugin_EventFilter
func callbackQSqlDriverPlugin_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlDriverPlugin::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlDriverPluginFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlDriverPlugin) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlDriverPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::eventFilter", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::eventFilter")
	}
}

func (ptr *QSqlDriverPlugin) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriverPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlDriverPlugin_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSqlDriverPlugin) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlDriverPlugin::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlDriverPlugin_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSqlDriverPlugin_MetaObject
func callbackQSqlDriverPlugin_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlDriverPlugin::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlDriverPlugin::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlDriverPluginFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlDriverPlugin) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlDriverPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::metaObject", f)
	}
}

func (ptr *QSqlDriverPlugin) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlDriverPlugin::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlDriverPlugin::metaObject")
	}
}

func (ptr *QSqlDriverPlugin) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlDriverPlugin::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriverPlugin_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlDriverPlugin) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlDriverPlugin::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlDriverPlugin_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSqlError::ErrorType
type QSqlError__ErrorType int64

const (
	QSqlError__NoError          = QSqlError__ErrorType(0)
	QSqlError__ConnectionError  = QSqlError__ErrorType(1)
	QSqlError__StatementError   = QSqlError__ErrorType(2)
	QSqlError__TransactionError = QSqlError__ErrorType(3)
	QSqlError__UnknownError     = QSqlError__ErrorType(4)
)

type QSqlError struct {
	ptr unsafe.Pointer
}

type QSqlError_ITF interface {
	QSqlError_PTR() *QSqlError
}

func (p *QSqlError) QSqlError_PTR() *QSqlError {
	return p
}

func (p *QSqlError) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlError) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlError(ptr QSqlError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlError_PTR().Pointer()
	}
	return nil
}

func NewQSqlErrorFromPointer(ptr unsafe.Pointer) *QSqlError {
	var n = new(QSqlError)
	n.SetPointer(ptr)
	return n
}
func NewQSqlError3(other QSqlError_ITF) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	var tmpValue = NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError3(PointerFromQSqlError(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
	return tmpValue
}

func NewQSqlError(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {
	defer qt.Recovering("QSqlError::QSqlError")

	var driverTextC = C.CString(driverText)
	defer C.free(unsafe.Pointer(driverTextC))
	var databaseTextC = C.CString(databaseText)
	defer C.free(unsafe.Pointer(databaseTextC))
	var codeC = C.CString(code)
	defer C.free(unsafe.Pointer(codeC))
	var tmpValue = NewQSqlErrorFromPointer(C.QSqlError_NewQSqlError(driverTextC, databaseTextC, C.longlong(ty), codeC))
	runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
	return tmpValue
}

func (ptr *QSqlError) DatabaseText() string {
	defer qt.Recovering("QSqlError::databaseText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DatabaseText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) DriverText() string {
	defer qt.Recovering("QSqlError::driverText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DriverText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) IsValid() bool {
	defer qt.Recovering("QSqlError::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlError_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlError) NativeErrorCode() string {
	defer qt.Recovering("QSqlError::nativeErrorCode")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_NativeErrorCode(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Text() string {
	defer qt.Recovering("QSqlError::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {
	defer qt.Recovering("QSqlError::type")

	if ptr.Pointer() != nil {
		return QSqlError__ErrorType(C.QSqlError_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlError) DestroyQSqlError() {
	defer qt.Recovering("QSqlError::~QSqlError")

	if ptr.Pointer() != nil {
		C.QSqlError_DestroyQSqlError(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QSqlField::RequiredStatus
type QSqlField__RequiredStatus int64

const (
	QSqlField__Unknown  = QSqlField__RequiredStatus(-1)
	QSqlField__Optional = QSqlField__RequiredStatus(0)
	QSqlField__Required = QSqlField__RequiredStatus(1)
)

type QSqlField struct {
	ptr unsafe.Pointer
}

type QSqlField_ITF interface {
	QSqlField_PTR() *QSqlField
}

func (p *QSqlField) QSqlField_PTR() *QSqlField {
	return p
}

func (p *QSqlField) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlField) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlField(ptr QSqlField_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlField_PTR().Pointer()
	}
	return nil
}

func NewQSqlFieldFromPointer(ptr unsafe.Pointer) *QSqlField {
	var n = new(QSqlField)
	n.SetPointer(ptr)
	return n
}
func NewQSqlField2(other QSqlField_ITF) *QSqlField {
	defer qt.Recovering("QSqlField::QSqlField")

	var tmpValue = NewQSqlFieldFromPointer(C.QSqlField_NewQSqlField2(PointerFromQSqlField(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
	return tmpValue
}

func (ptr *QSqlField) Clear() {
	defer qt.Recovering("QSqlField::clear")

	if ptr.Pointer() != nil {
		C.QSqlField_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlField) DefaultValue() *core.QVariant {
	defer qt.Recovering("QSqlField::defaultValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlField_DefaultValue(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlField) IsAutoValue() bool {
	defer qt.Recovering("QSqlField::isAutoValue")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsAutoValue(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsGenerated() bool {
	defer qt.Recovering("QSqlField::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsGenerated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsNull() bool {
	defer qt.Recovering("QSqlField::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsReadOnly() bool {
	defer qt.Recovering("QSqlField::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) IsValid() bool {
	defer qt.Recovering("QSqlField::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlField_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlField) Length() int {
	defer qt.Recovering("QSqlField::length")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlField_Length(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) Name() string {
	defer qt.Recovering("QSqlField::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlField_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlField) Precision() int {
	defer qt.Recovering("QSqlField::precision")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlField_Precision(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlField) RequiredStatus() QSqlField__RequiredStatus {
	defer qt.Recovering("QSqlField::requiredStatus")

	if ptr.Pointer() != nil {
		return QSqlField__RequiredStatus(C.QSqlField_RequiredStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlField) SetAutoValue(autoVal bool) {
	defer qt.Recovering("QSqlField::setAutoValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetAutoValue(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(autoVal))))
	}
}

func (ptr *QSqlField) SetDefaultValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setDefaultValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetDefaultValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) SetGenerated(gen bool) {
	defer qt.Recovering("QSqlField::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlField_SetGenerated(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(gen))))
	}
}

func (ptr *QSqlField) SetLength(fieldLength int) {
	defer qt.Recovering("QSqlField::setLength")

	if ptr.Pointer() != nil {
		C.QSqlField_SetLength(ptr.Pointer(), C.int(int32(fieldLength)))
	}
}

func (ptr *QSqlField) SetName(name string) {
	defer qt.Recovering("QSqlField::setName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlField_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QSqlField) SetPrecision(precision int) {
	defer qt.Recovering("QSqlField::setPrecision")

	if ptr.Pointer() != nil {
		C.QSqlField_SetPrecision(ptr.Pointer(), C.int(int32(precision)))
	}
}

func (ptr *QSqlField) SetReadOnly(readOnly bool) {
	defer qt.Recovering("QSqlField::setReadOnly")

	if ptr.Pointer() != nil {
		C.QSqlField_SetReadOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(readOnly))))
	}
}

func (ptr *QSqlField) SetRequired(required bool) {
	defer qt.Recovering("QSqlField::setRequired")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequired(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(required))))
	}
}

func (ptr *QSqlField) SetRequiredStatus(required QSqlField__RequiredStatus) {
	defer qt.Recovering("QSqlField::setRequiredStatus")

	if ptr.Pointer() != nil {
		C.QSqlField_SetRequiredStatus(ptr.Pointer(), C.longlong(required))
	}
}

func (ptr *QSqlField) SetValue(value core.QVariant_ITF) {
	defer qt.Recovering("QSqlField::setValue")

	if ptr.Pointer() != nil {
		C.QSqlField_SetValue(ptr.Pointer(), core.PointerFromQVariant(value))
	}
}

func (ptr *QSqlField) Value() *core.QVariant {
	defer qt.Recovering("QSqlField::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlField_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlField) DestroyQSqlField() {
	defer qt.Recovering("QSqlField::~QSqlField")

	if ptr.Pointer() != nil {
		C.QSqlField_DestroyQSqlField(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndex_ITF interface {
	QSqlRecord_ITF
	QSqlIndex_PTR() *QSqlIndex
}

func (p *QSqlIndex) QSqlIndex_PTR() *QSqlIndex {
	return p
}

func (p *QSqlIndex) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func (p *QSqlIndex) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSqlRecord_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlIndex(ptr QSqlIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndex_PTR().Pointer()
	}
	return nil
}

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = new(QSqlIndex)
	n.SetPointer(ptr)
	return n
}
func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	var tmpValue = NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex2(PointerFromQSqlIndex(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
	return tmpValue
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	defer qt.Recovering("QSqlIndex::QSqlIndex")

	var cursornameC = C.CString(cursorname)
	defer C.free(unsafe.Pointer(cursornameC))
	var nameC = C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	var tmpValue = NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex(cursornameC, nameC))
	runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
	return tmpValue
}

func (ptr *QSqlIndex) Append(field QSqlField_ITF) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {
	defer qt.Recovering("QSqlIndex::append")

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(ptr.Pointer(), PointerFromQSqlField(field), C.char(int8(qt.GoBoolToInt(desc))))
	}
}

func (ptr *QSqlIndex) CursorName() string {
	defer qt.Recovering("QSqlIndex::cursorName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_CursorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	defer qt.Recovering("QSqlIndex::isDescending")

	if ptr.Pointer() != nil {
		return C.QSqlIndex_IsDescending(ptr.Pointer(), C.int(int32(i))) != 0
	}
	return false
}

func (ptr *QSqlIndex) Name() string {
	defer qt.Recovering("QSqlIndex::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	defer qt.Recovering("QSqlIndex::setCursorName")

	if ptr.Pointer() != nil {
		var cursorNameC = C.CString(cursorName)
		defer C.free(unsafe.Pointer(cursorNameC))
		C.QSqlIndex_SetCursorName(ptr.Pointer(), cursorNameC)
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	defer qt.Recovering("QSqlIndex::setDescending")

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(ptr.Pointer(), C.int(int32(i)), C.char(int8(qt.GoBoolToInt(desc))))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	defer qt.Recovering("QSqlIndex::setName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlIndex_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	defer qt.Recovering("QSqlIndex::~QSqlIndex")

	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QSqlQuery::BatchExecutionMode
type QSqlQuery__BatchExecutionMode int64

const (
	QSqlQuery__ValuesAsRows    = QSqlQuery__BatchExecutionMode(0)
	QSqlQuery__ValuesAsColumns = QSqlQuery__BatchExecutionMode(1)
)

type QSqlQuery struct {
	ptr unsafe.Pointer
}

type QSqlQuery_ITF interface {
	QSqlQuery_PTR() *QSqlQuery
}

func (p *QSqlQuery) QSqlQuery_PTR() *QSqlQuery {
	return p
}

func (p *QSqlQuery) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlQuery) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlQuery(ptr QSqlQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQuery_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryFromPointer(ptr unsafe.Pointer) *QSqlQuery {
	var n = new(QSqlQuery)
	n.SetPointer(ptr)
	return n
}
func NewQSqlQuery3(db QSqlDatabase_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	var tmpValue = NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery3(PointerFromQSqlDatabase(db)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery(result QSqlResult_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	var tmpValue = NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery(PointerFromQSqlResult(result)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery4(other QSqlQuery_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	var tmpValue = NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery4(PointerFromQSqlQuery(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func NewQSqlQuery2(query string, db QSqlDatabase_ITF) *QSqlQuery {
	defer qt.Recovering("QSqlQuery::QSqlQuery")

	var queryC = C.CString(query)
	defer C.free(unsafe.Pointer(queryC))
	var tmpValue = NewQSqlQueryFromPointer(C.QSqlQuery_NewQSqlQuery2(queryC, PointerFromQSqlDatabase(db)))
	runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
	return tmpValue
}

func (ptr *QSqlQuery) At() int {
	defer qt.Recovering("QSqlQuery::at")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_At(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) BoundValue(placeholder string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		var placeholderC = C.CString(placeholder)
		defer C.free(unsafe.Pointer(placeholderC))
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQuery_BoundValue(ptr.Pointer(), placeholderC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) BoundValue2(pos int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::boundValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQuery_BoundValue2(ptr.Pointer(), C.int(int32(pos))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Clear() {
	defer qt.Recovering("QSqlQuery::clear")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlQuery::driver")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlDriverFromPointer(C.QSqlQuery_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Exec2() bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Exec2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Exec(query string) bool {
	defer qt.Recovering("QSqlQuery::exec")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlQuery_Exec(ptr.Pointer(), queryC) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecBatch(mode QSqlQuery__BatchExecutionMode) bool {
	defer qt.Recovering("QSqlQuery::execBatch")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_ExecBatch(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

func (ptr *QSqlQuery) ExecutedQuery() string {
	defer qt.Recovering("QSqlQuery::executedQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Finish() {
	defer qt.Recovering("QSqlQuery::finish")

	if ptr.Pointer() != nil {
		C.QSqlQuery_Finish(ptr.Pointer())
	}
}

func (ptr *QSqlQuery) First() bool {
	defer qt.Recovering("QSqlQuery::first")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_First(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsActive() bool {
	defer qt.Recovering("QSqlQuery::isActive")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsForwardOnly() bool {
	defer qt.Recovering("QSqlQuery::isForwardOnly")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsForwardOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull2(name string) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlQuery_IsNull2(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsNull(field int) bool {
	defer qt.Recovering("QSqlQuery::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsNull(ptr.Pointer(), C.int(int32(field))) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsSelect() bool {
	defer qt.Recovering("QSqlQuery::isSelect")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) IsValid() bool {
	defer qt.Recovering("QSqlQuery::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Last() bool {
	defer qt.Recovering("QSqlQuery::last")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Last(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) LastError() *QSqlError {
	defer qt.Recovering("QSqlQuery::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlErrorFromPointer(C.QSqlQuery_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlQuery::lastInsertId")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQuery_LastInsertId(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) LastQuery() string {
	defer qt.Recovering("QSqlQuery::lastQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQuery_LastQuery(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlQuery) Next() bool {
	defer qt.Recovering("QSqlQuery::next")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Next(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NextResult() bool {
	defer qt.Recovering("QSqlQuery::nextResult")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_NextResult(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) NumRowsAffected() int {
	defer qt.Recovering("QSqlQuery::numRowsAffected")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_NumRowsAffected(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) Prepare(query string) bool {
	defer qt.Recovering("QSqlQuery::prepare")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlQuery_Prepare(ptr.Pointer(), queryC) != 0
	}
	return false
}

func (ptr *QSqlQuery) Previous() bool {
	defer qt.Recovering("QSqlQuery::previous")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Previous(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQuery) Record() *QSqlRecord {
	defer qt.Recovering("QSqlQuery::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlQuery_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Result() *QSqlResult {
	defer qt.Recovering("QSqlQuery::result")

	if ptr.Pointer() != nil {
		return NewQSqlResultFromPointer(C.QSqlQuery_Result(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQuery) Seek(index int, relative bool) bool {
	defer qt.Recovering("QSqlQuery::seek")

	if ptr.Pointer() != nil {
		return C.QSqlQuery_Seek(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(relative)))) != 0
	}
	return false
}

func (ptr *QSqlQuery) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlQuery::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlQuery_SetForwardOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

func (ptr *QSqlQuery) Size() int {
	defer qt.Recovering("QSqlQuery::size")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQuery_Size(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlQuery) Value2(name string) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQuery_Value2(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) Value(index int) *core.QVariant {
	defer qt.Recovering("QSqlQuery::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQuery_Value(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQuery) DestroyQSqlQuery() {
	defer qt.Recovering("QSqlQuery::~QSqlQuery")

	if ptr.Pointer() != nil {
		C.QSqlQuery_DestroyQSqlQuery(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModel_ITF interface {
	core.QAbstractTableModel_ITF
	QSqlQueryModel_PTR() *QSqlQueryModel
}

func (p *QSqlQueryModel) QSqlQueryModel_PTR() *QSqlQueryModel {
	return p
}

func (p *QSqlQueryModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QAbstractTableModel_PTR().Pointer()
	}
	return nil
}

func (p *QSqlQueryModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QAbstractTableModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	return n
}
func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::rowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func NewQSqlQueryModel(parent core.QObject_ITF) *QSqlQueryModel {
	defer qt.Recovering("QSqlQueryModel::QSqlQueryModel")

	var tmpValue = NewQSqlQueryModelFromPointer(C.QSqlQueryModel_NewQSqlQueryModel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQueryModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::canFetchMore")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Clear
func callbackQSqlQueryModel_Clear(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::clear")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::clear", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlQueryModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::clear")
	}
}

func (ptr *QSqlQueryModel) Clear() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ClearDefault() {
	defer qt.Recovering("QSqlQueryModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlQueryModel::columnCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlQueryModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QSqlQueryModel) FetchMore(parent core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlQueryModel::fetchMore")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlQueryModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_IndexInQuery
func callbackQSqlQueryModel_IndexInQuery(ptr unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::indexInQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::indexInQuery"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(item)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).IndexInQueryDefault(core.NewQModelIndexFromPointer(item)))
}

func (ptr *QSqlQueryModel) ConnectIndexInQuery(f func(item *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::indexInQuery", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectIndexInQuery() {
	defer qt.Recovering("disconnect QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::indexInQuery")
	}
}

func (ptr *QSqlQueryModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) IndexInQueryDefault(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexInQueryDefault(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::insertColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) LastError() *QSqlError {
	defer qt.Recovering("QSqlQueryModel::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlErrorFromPointer(C.QSqlQueryModel_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) Query() *QSqlQuery {
	defer qt.Recovering("QSqlQueryModel::query")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlQueryFromPointer(C.QSqlQueryModel_Query(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlQuery).DestroyQSqlQuery)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_QueryChange
func callbackQSqlQueryModel_QueryChange(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::queryChange")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::queryChange", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::queryChange")
	}
}

func (ptr *QSqlQueryModel) QueryChange() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlQueryModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_QueryChangeDefault(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) Record2() *QSqlRecord {
	defer qt.Recovering("QSqlQueryModel::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlQueryModel_Record2(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) Record(row int) *QSqlRecord {
	defer qt.Recovering("QSqlQueryModel::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlQueryModel_Record(ptr.Pointer(), C.int(int32(row))))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlQueryModel::setHeaderData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetHeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetLastError(error QSqlError_ITF) {
	defer qt.Recovering("QSqlQueryModel::setLastError")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {
	defer qt.Recovering("QSqlQueryModel::setQuery")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		C.QSqlQueryModel_SetQuery2(ptr.Pointer(), queryC, PointerFromQSqlDatabase(db))
	}
}

//export callbackQSqlQueryModel_DestroyQSqlQueryModel
func callbackQSqlQueryModel_DestroyQSqlQueryModel(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::~QSqlQueryModel")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::~QSqlQueryModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).DestroyQSqlQueryModelDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectDestroyQSqlQueryModel(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::~QSqlQueryModel", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectDestroyQSqlQueryModel() {
	defer qt.Recovering("disconnect QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::~QSqlQueryModel")
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	defer qt.Recovering("QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModelDefault() {
	defer qt.Recovering("QSqlQueryModel::~QSqlQueryModel")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlQueryModel_Index
func callbackQSqlQueryModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::index")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QSqlQueryModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlQueryModel::index")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::index", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectIndex() {
	defer qt.Recovering("disconnect QSqlQueryModel::index")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::index")
	}
}

func (ptr *QSqlQueryModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_DropMimeData
func callbackQSqlQueryModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::dropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlQueryModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::dropMimeData", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectDropMimeData() {
	defer qt.Recovering("disconnect QSqlQueryModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::dropMimeData")
	}
}

func (ptr *QSqlQueryModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Flags
func callbackQSqlQueryModel_Flags(ptr unsafe.Pointer, index unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlQueryModel::flags")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::flags"); signal != nil {
		return C.longlong(signal.(func(*core.QModelIndex) core.Qt__ItemFlag)(core.NewQModelIndexFromPointer(index)))
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).FlagsDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) ConnectFlags(f func(index *core.QModelIndex) core.Qt__ItemFlag) {
	defer qt.Recovering("connect QSqlQueryModel::flags")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::flags", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectFlags() {
	defer qt.Recovering("disconnect QSqlQueryModel::flags")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::flags")
	}
}

func (ptr *QSqlQueryModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QSqlQueryModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlQueryModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlQueryModel) FlagsDefault(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QSqlQueryModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlQueryModel_FlagsDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

//export callbackQSqlQueryModel_Sibling
func callbackQSqlQueryModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::sibling")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QSqlQueryModel) ConnectSibling(f func(row int, column int, idx *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlQueryModel::sibling")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::sibling", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSibling() {
	defer qt.Recovering("disconnect QSqlQueryModel::sibling")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::sibling")
	}
}

func (ptr *QSqlQueryModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Buddy
func callbackQSqlQueryModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::buddy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlQueryModel::buddy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::buddy", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectBuddy() {
	defer qt.Recovering("disconnect QSqlQueryModel::buddy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::buddy")
	}
}

func (ptr *QSqlQueryModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_CanDropMimeData
func callbackQSqlQueryModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::canDropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlQueryModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::canDropMimeData", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectCanDropMimeData() {
	defer qt.Recovering("disconnect QSqlQueryModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::canDropMimeData")
	}
}

func (ptr *QSqlQueryModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_HasChildren
func callbackQSqlQueryModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::hasChildren")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlQueryModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::hasChildren", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectHasChildren() {
	defer qt.Recovering("disconnect QSqlQueryModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::hasChildren")
	}
}

func (ptr *QSqlQueryModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_InsertRows
func callbackQSqlQueryModel_InsertRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::insertRows")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::insertRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).InsertRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) ConnectInsertRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlQueryModel::insertRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::insertRows", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectInsertRows() {
	defer qt.Recovering("disconnect QSqlQueryModel::insertRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::insertRows")
	}
}

func (ptr *QSqlQueryModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) InsertRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MimeTypes
func callbackQSqlQueryModel_MimeTypes(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlQueryModel::mimeTypes")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQSqlQueryModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QSqlQueryModel) ConnectMimeTypes(f func() []string) {
	defer qt.Recovering("connect QSqlQueryModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::mimeTypes", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectMimeTypes() {
	defer qt.Recovering("disconnect QSqlQueryModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::mimeTypes")
	}
}

func (ptr *QSqlQueryModel) MimeTypes() []string {
	defer qt.Recovering("QSqlQueryModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlQueryModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlQueryModel) MimeTypesDefault() []string {
	defer qt.Recovering("QSqlQueryModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlQueryModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlQueryModel_MoveColumns
func callbackQSqlQueryModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlQueryModel::moveColumns")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlQueryModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlQueryModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::moveColumns", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectMoveColumns() {
	defer qt.Recovering("disconnect QSqlQueryModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::moveColumns")
	}
}

func (ptr *QSqlQueryModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlQueryModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlQueryModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MoveRows
func callbackQSqlQueryModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlQueryModel::moveRows")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlQueryModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlQueryModel::moveRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::moveRows", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectMoveRows() {
	defer qt.Recovering("disconnect QSqlQueryModel::moveRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::moveRows")
	}
}

func (ptr *QSqlQueryModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlQueryModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlQueryModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Parent
func callbackQSqlQueryModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::parent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlQueryModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlQueryModel::parent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::parent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectParent() {
	defer qt.Recovering("disconnect QSqlQueryModel::parent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::parent")
	}
}

func (ptr *QSqlQueryModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlQueryModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlQueryModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_RemoveRows
func callbackQSqlQueryModel_RemoveRows(ptr unsafe.Pointer, row C.int, count C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::removeRows")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::removeRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, int, *core.QModelIndex) bool)(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).RemoveRowsDefault(int(int32(row)), int(int32(count)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlQueryModel) ConnectRemoveRows(f func(row int, count int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlQueryModel::removeRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::removeRows", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectRemoveRows() {
	defer qt.Recovering("disconnect QSqlQueryModel::removeRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::removeRows")
	}
}

func (ptr *QSqlQueryModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) RemoveRowsDefault(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveRowsDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_ResetInternalData
func callbackQSqlQueryModel_ResetInternalData(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::resetInternalData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectResetInternalData(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::resetInternalData", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectResetInternalData() {
	defer qt.Recovering("disconnect QSqlQueryModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::resetInternalData")
	}
}

func (ptr *QSqlQueryModel) ResetInternalData() {
	defer qt.Recovering("QSqlQueryModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ResetInternalDataDefault() {
	defer qt.Recovering("QSqlQueryModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQSqlQueryModel_Revert
func callbackQSqlQueryModel_Revert(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::revert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::revert"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::revert", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlQueryModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::revert")
	}
}

func (ptr *QSqlQueryModel) Revert() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) RevertDefault() {
	defer qt.Recovering("QSqlQueryModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQSqlQueryModel_SetData
func callbackQSqlQueryModel_SetData(ptr unsafe.Pointer, index unsafe.Pointer, value unsafe.Pointer, role C.int) C.char {
	defer qt.Recovering("callback QSqlQueryModel::setData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::setData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, *core.QVariant, int) bool)(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SetDataDefault(core.NewQModelIndexFromPointer(index), core.NewQVariantFromPointer(value), int(int32(role))))))
}

func (ptr *QSqlQueryModel) ConnectSetData(f func(index *core.QModelIndex, value *core.QVariant, role int) bool) {
	defer qt.Recovering("connect QSqlQueryModel::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::setData", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSetData() {
	defer qt.Recovering("disconnect QSqlQueryModel::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::setData")
	}
}

func (ptr *QSqlQueryModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlQueryModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetDataDefault(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlQueryModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetDataDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQSqlQueryModel_Sort
func callbackQSqlQueryModel_Sort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	defer qt.Recovering("callback QSqlQueryModel::sort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::sort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQSqlQueryModelFromPointer(ptr).SortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlQueryModel) ConnectSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::sort", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSort() {
	defer qt.Recovering("disconnect QSqlQueryModel::sort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::sort")
	}
}

func (ptr *QSqlQueryModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QSqlQueryModel) SortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlQueryModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlQueryModel_Span
func callbackQSqlQueryModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::span")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlQueryModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlQueryModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QSqlQueryModel::span")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::span", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSpan() {
	defer qt.Recovering("disconnect QSqlQueryModel::span")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::span")
	}
}

func (ptr *QSqlQueryModel) Span(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlQueryModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlQueryModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlQueryModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlQueryModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlQueryModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlQueryModel_Submit
func callbackQSqlQueryModel_Submit(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::submit")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QSqlQueryModel) ConnectSubmit(f func() bool) {
	defer qt.Recovering("connect QSqlQueryModel::submit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::submit", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSubmit() {
	defer qt.Recovering("disconnect QSqlQueryModel::submit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::submit")
	}
}

func (ptr *QSqlQueryModel) Submit() bool {
	defer qt.Recovering("QSqlQueryModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SubmitDefault() bool {
	defer qt.Recovering("QSqlQueryModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlQueryModel_SupportedDragActions
func callbackQSqlQueryModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlQueryModel::supportedDragActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QSqlQueryModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlQueryModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::supportedDragActions", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSupportedDragActions() {
	defer qt.Recovering("disconnect QSqlQueryModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::supportedDragActions")
	}
}

func (ptr *QSqlQueryModel) SupportedDragActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlQueryModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQueryModel) SupportedDragActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlQueryModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlQueryModel_SupportedDropActions
func callbackQSqlQueryModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlQueryModel::supportedDropActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlQueryModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QSqlQueryModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlQueryModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::supportedDropActions", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectSupportedDropActions() {
	defer qt.Recovering("disconnect QSqlQueryModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::supportedDropActions")
	}
}

func (ptr *QSqlQueryModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlQueryModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlQueryModel) SupportedDropActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlQueryModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlQueryModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlQueryModel_TimerEvent
func callbackQSqlQueryModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::timerEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::timerEvent")
	}
}

func (ptr *QSqlQueryModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlQueryModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlQueryModel_ChildEvent
func callbackQSqlQueryModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::childEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::childEvent")
	}
}

func (ptr *QSqlQueryModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlQueryModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlQueryModel_ConnectNotify
func callbackQSqlQueryModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlQueryModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlQueryModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlQueryModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::connectNotify", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlQueryModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::connectNotify")
	}
}

func (ptr *QSqlQueryModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlQueryModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlQueryModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlQueryModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlQueryModel_CustomEvent
func callbackQSqlQueryModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlQueryModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlQueryModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::customEvent", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::customEvent")
	}
}

func (ptr *QSqlQueryModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlQueryModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlQueryModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlQueryModel_DeleteLater
func callbackQSqlQueryModel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlQueryModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlQueryModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlQueryModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::deleteLater", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlQueryModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::deleteLater")
	}
}

func (ptr *QSqlQueryModel) DeleteLater() {
	defer qt.Recovering("QSqlQueryModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlQueryModel) DeleteLaterDefault() {
	defer qt.Recovering("QSqlQueryModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlQueryModel_DisconnectNotify
func callbackQSqlQueryModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlQueryModel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlQueryModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlQueryModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlQueryModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::disconnectNotify", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlQueryModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::disconnectNotify")
	}
}

func (ptr *QSqlQueryModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlQueryModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlQueryModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlQueryModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlQueryModel_Event
func callbackQSqlQueryModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlQueryModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlQueryModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::event", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlQueryModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::event")
	}
}

func (ptr *QSqlQueryModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_EventFilter
func callbackQSqlQueryModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlQueryModel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlQueryModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlQueryModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlQueryModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::eventFilter", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSqlQueryModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::eventFilter")
	}
}

func (ptr *QSqlQueryModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlQueryModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSqlQueryModel_MetaObject
func callbackQSqlQueryModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlQueryModel::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlQueryModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlQueryModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlQueryModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlQueryModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::metaObject", f)
	}
}

func (ptr *QSqlQueryModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlQueryModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlQueryModel::metaObject")
	}
}

func (ptr *QSqlQueryModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlQueryModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlQueryModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlQueryModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlQueryModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlQueryModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSqlRecord struct {
	ptr unsafe.Pointer
}

type QSqlRecord_ITF interface {
	QSqlRecord_PTR() *QSqlRecord
}

func (p *QSqlRecord) QSqlRecord_PTR() *QSqlRecord {
	return p
}

func (p *QSqlRecord) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlRecord) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlRecord(ptr QSqlRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRecord_PTR().Pointer()
	}
	return nil
}

func NewQSqlRecordFromPointer(ptr unsafe.Pointer) *QSqlRecord {
	var n = new(QSqlRecord)
	n.SetPointer(ptr)
	return n
}
func NewQSqlRecord() *QSqlRecord {
	defer qt.Recovering("QSqlRecord::QSqlRecord")

	var tmpValue = NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord())
	runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
	return tmpValue
}

func NewQSqlRecord2(other QSqlRecord_ITF) *QSqlRecord {
	defer qt.Recovering("QSqlRecord::QSqlRecord")

	var tmpValue = NewQSqlRecordFromPointer(C.QSqlRecord_NewQSqlRecord2(PointerFromQSqlRecord(other)))
	runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
	return tmpValue
}

func (ptr *QSqlRecord) Append(field QSqlField_ITF) {
	defer qt.Recovering("QSqlRecord::append")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) Clear() {
	defer qt.Recovering("QSqlRecord::clear")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) ClearValues() {
	defer qt.Recovering("QSqlRecord::clearValues")

	if ptr.Pointer() != nil {
		C.QSqlRecord_ClearValues(ptr.Pointer())
	}
}

func (ptr *QSqlRecord) Contains(name string) bool {
	defer qt.Recovering("QSqlRecord::contains")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlRecord_Contains(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlRecord) Count() int {
	defer qt.Recovering("QSqlRecord::count")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlRecord_Count(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlRecord) Field2(name string) *QSqlField {
	defer qt.Recovering("QSqlRecord::field")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = NewQSqlFieldFromPointer(C.QSqlRecord_Field2(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Field(index int) *QSqlField {
	defer qt.Recovering("QSqlRecord::field")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlFieldFromPointer(C.QSqlRecord_Field(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*QSqlField).DestroyQSqlField)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) FieldName(index int) string {
	defer qt.Recovering("QSqlRecord::fieldName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRecord_FieldName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QSqlRecord) IndexOf(name string) int {
	defer qt.Recovering("QSqlRecord::indexOf")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return int(int32(C.QSqlRecord_IndexOf(ptr.Pointer(), nameC)))
	}
	return 0
}

func (ptr *QSqlRecord) Insert(pos int, field QSqlField_ITF) {
	defer qt.Recovering("QSqlRecord::insert")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Insert(ptr.Pointer(), C.int(int32(pos)), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) IsEmpty() bool {
	defer qt.Recovering("QSqlRecord::isEmpty")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated(name string) bool {
	defer qt.Recovering("QSqlRecord::isGenerated")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlRecord_IsGenerated(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsGenerated2(index int) bool {
	defer qt.Recovering("QSqlRecord::isGenerated")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsGenerated2(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull(name string) bool {
	defer qt.Recovering("QSqlRecord::isNull")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		return C.QSqlRecord_IsNull(ptr.Pointer(), nameC) != 0
	}
	return false
}

func (ptr *QSqlRecord) IsNull2(index int) bool {
	defer qt.Recovering("QSqlRecord::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlRecord_IsNull2(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QSqlRecord) KeyValues(keyFields QSqlRecord_ITF) *QSqlRecord {
	defer qt.Recovering("QSqlRecord::keyValues")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlRecord_KeyValues(ptr.Pointer(), PointerFromQSqlRecord(keyFields)))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Remove(pos int) {
	defer qt.Recovering("QSqlRecord::remove")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Remove(ptr.Pointer(), C.int(int32(pos)))
	}
}

func (ptr *QSqlRecord) Replace(pos int, field QSqlField_ITF) {
	defer qt.Recovering("QSqlRecord::replace")

	if ptr.Pointer() != nil {
		C.QSqlRecord_Replace(ptr.Pointer(), C.int(int32(pos)), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlRecord) SetGenerated(name string, generated bool) {
	defer qt.Recovering("QSqlRecord::setGenerated")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlRecord_SetGenerated(ptr.Pointer(), nameC, C.char(int8(qt.GoBoolToInt(generated))))
	}
}

func (ptr *QSqlRecord) SetGenerated2(index int, generated bool) {
	defer qt.Recovering("QSqlRecord::setGenerated")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetGenerated2(ptr.Pointer(), C.int(int32(index)), C.char(int8(qt.GoBoolToInt(generated))))
	}
}

func (ptr *QSqlRecord) SetNull2(name string) {
	defer qt.Recovering("QSqlRecord::setNull")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlRecord_SetNull2(ptr.Pointer(), nameC)
	}
}

func (ptr *QSqlRecord) SetNull(index int) {
	defer qt.Recovering("QSqlRecord::setNull")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetNull(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSqlRecord) SetValue2(name string, val core.QVariant_ITF) {
	defer qt.Recovering("QSqlRecord::setValue")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QSqlRecord_SetValue2(ptr.Pointer(), nameC, core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) SetValue(index int, val core.QVariant_ITF) {
	defer qt.Recovering("QSqlRecord::setValue")

	if ptr.Pointer() != nil {
		C.QSqlRecord_SetValue(ptr.Pointer(), C.int(int32(index)), core.PointerFromQVariant(val))
	}
}

func (ptr *QSqlRecord) Value2(name string) *core.QVariant {
	defer qt.Recovering("QSqlRecord::value")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		var tmpValue = core.NewQVariantFromPointer(C.QSqlRecord_Value2(ptr.Pointer(), nameC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) Value(index int) *core.QVariant {
	defer qt.Recovering("QSqlRecord::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlRecord_Value(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRecord) DestroyQSqlRecord() {
	defer qt.Recovering("QSqlRecord::~QSqlRecord")

	if ptr.Pointer() != nil {
		C.QSqlRecord_DestroyQSqlRecord(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QSqlRelation struct {
	ptr unsafe.Pointer
}

type QSqlRelation_ITF interface {
	QSqlRelation_PTR() *QSqlRelation
}

func (p *QSqlRelation) QSqlRelation_PTR() *QSqlRelation {
	return p
}

func (p *QSqlRelation) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlRelation) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlRelation(ptr QSqlRelation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelation_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationFromPointer(ptr unsafe.Pointer) *QSqlRelation {
	var n = new(QSqlRelation)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlRelation) DestroyQSqlRelation() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQSqlRelation() *QSqlRelation {
	defer qt.Recovering("QSqlRelation::QSqlRelation")

	var tmpValue = NewQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation())
	runtime.SetFinalizer(tmpValue, (*QSqlRelation).DestroyQSqlRelation)
	return tmpValue
}

func NewQSqlRelation2(tableName string, indexColumn string, displayColumn string) *QSqlRelation {
	defer qt.Recovering("QSqlRelation::QSqlRelation")

	var tableNameC = C.CString(tableName)
	defer C.free(unsafe.Pointer(tableNameC))
	var indexColumnC = C.CString(indexColumn)
	defer C.free(unsafe.Pointer(indexColumnC))
	var displayColumnC = C.CString(displayColumn)
	defer C.free(unsafe.Pointer(displayColumnC))
	var tmpValue = NewQSqlRelationFromPointer(C.QSqlRelation_NewQSqlRelation2(tableNameC, indexColumnC, displayColumnC))
	runtime.SetFinalizer(tmpValue, (*QSqlRelation).DestroyQSqlRelation)
	return tmpValue
}

func (ptr *QSqlRelation) DisplayColumn() string {
	defer qt.Recovering("QSqlRelation::displayColumn")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_DisplayColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IndexColumn() string {
	defer qt.Recovering("QSqlRelation::indexColumn")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_IndexColumn(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelation) IsValid() bool {
	defer qt.Recovering("QSqlRelation::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlRelation_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelation) TableName() string {
	defer qt.Recovering("QSqlRelation::tableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelation_TableName(ptr.Pointer()))
	}
	return ""
}

type QSqlRelationalDelegate struct {
	widgets.QItemDelegate
}

type QSqlRelationalDelegate_ITF interface {
	widgets.QItemDelegate_ITF
	QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate
}

func (p *QSqlRelationalDelegate) QSqlRelationalDelegate_PTR() *QSqlRelationalDelegate {
	return p
}

func (p *QSqlRelationalDelegate) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QItemDelegate_PTR().Pointer()
	}
	return nil
}

func (p *QSqlRelationalDelegate) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QItemDelegate_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlRelationalDelegate(ptr QSqlRelationalDelegate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalDelegate_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalDelegateFromPointer(ptr unsafe.Pointer) *QSqlRelationalDelegate {
	var n = new(QSqlRelationalDelegate)
	n.SetPointer(ptr)
	return n
}
func NewQSqlRelationalDelegate(parent core.QObject_ITF) *QSqlRelationalDelegate {
	defer qt.Recovering("QSqlRelationalDelegate::QSqlRelationalDelegate")

	var tmpValue = NewQSqlRelationalDelegateFromPointer(C.QSqlRelationalDelegate_NewQSqlRelationalDelegate(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QSqlRelationalDelegate) CreateEditor(parent widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *widgets.QWidget {
	defer qt.Recovering("QSqlRelationalDelegate::createEditor")

	if ptr.Pointer() != nil {
		var tmpValue = widgets.NewQWidgetFromPointer(C.QSqlRelationalDelegate_CreateEditor(ptr.Pointer(), widgets.PointerFromQWidget(parent), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) SetModelData(editor widgets.QWidget_ITF, model core.QAbstractItemModel_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::setModelData")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_SetModelData(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQAbstractItemModel(model), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) DestroyQSqlRelationalDelegate() {
	defer qt.Recovering("QSqlRelationalDelegate::~QSqlRelationalDelegate")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyQSqlRelationalDelegate(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlRelationalDelegate_DrawCheck
func callbackQSqlRelationalDelegate_DrawCheck(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, state C.longlong) {
	defer qt.Recovering("callback QSqlRelationalDelegate::drawCheck")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::drawCheck"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, core.Qt__CheckState))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), core.Qt__CheckState(state))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawCheckDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), core.Qt__CheckState(state))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDrawCheck(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, rect *core.QRect, state core.Qt__CheckState)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::drawCheck")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawCheck", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDrawCheck() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::drawCheck")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawCheck")
	}
}

func (ptr *QSqlRelationalDelegate) DrawCheck(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, state core.Qt__CheckState) {
	defer qt.Recovering("QSqlRelationalDelegate::drawCheck")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawCheck(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), C.longlong(state))
	}
}

func (ptr *QSqlRelationalDelegate) DrawCheckDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, state core.Qt__CheckState) {
	defer qt.Recovering("QSqlRelationalDelegate::drawCheck")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawCheckDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), C.longlong(state))
	}
}

//export callbackQSqlRelationalDelegate_DrawDecoration
func callbackQSqlRelationalDelegate_DrawDecoration(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, pixmap unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::drawDecoration")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::drawDecoration"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, *gui.QPixmap))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), gui.NewQPixmapFromPointer(pixmap))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawDecorationDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), gui.NewQPixmapFromPointer(pixmap))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDrawDecoration(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, rect *core.QRect, pixmap *gui.QPixmap)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::drawDecoration")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawDecoration", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDrawDecoration() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::drawDecoration")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawDecoration")
	}
}

func (ptr *QSqlRelationalDelegate) DrawDecoration(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::drawDecoration")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawDecoration(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QSqlRelationalDelegate) DrawDecorationDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, pixmap gui.QPixmap_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::drawDecoration")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawDecorationDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), gui.PointerFromQPixmap(pixmap))
	}
}

//export callbackQSqlRelationalDelegate_DrawDisplay
func callbackQSqlRelationalDelegate_DrawDisplay(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer, text *C.char) {
	defer qt.Recovering("callback QSqlRelationalDelegate::drawDisplay")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::drawDisplay"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect, string))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), C.GoString(text))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawDisplayDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect), C.GoString(text))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDrawDisplay(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, rect *core.QRect, text string)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::drawDisplay")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawDisplay", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDrawDisplay() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::drawDisplay")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawDisplay")
	}
}

func (ptr *QSqlRelationalDelegate) DrawDisplay(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, text string) {
	defer qt.Recovering("QSqlRelationalDelegate::drawDisplay")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QSqlRelationalDelegate_DrawDisplay(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), textC)
	}
}

func (ptr *QSqlRelationalDelegate) DrawDisplayDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF, text string) {
	defer qt.Recovering("QSqlRelationalDelegate::drawDisplay")

	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QSqlRelationalDelegate_DrawDisplayDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect), textC)
	}
}

//export callbackQSqlRelationalDelegate_DrawFocus
func callbackQSqlRelationalDelegate_DrawFocus(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, rect unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::drawFocus")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::drawFocus"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QRect))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DrawFocusDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQRectFromPointer(rect))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDrawFocus(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, rect *core.QRect)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::drawFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawFocus", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDrawFocus() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::drawFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::drawFocus")
	}
}

func (ptr *QSqlRelationalDelegate) DrawFocus(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::drawFocus")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawFocus(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect))
	}
}

func (ptr *QSqlRelationalDelegate) DrawFocusDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, rect core.QRect_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::drawFocus")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DrawFocusDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQRect(rect))
	}
}

//export callbackQSqlRelationalDelegate_EditorEvent
func callbackQSqlRelationalDelegate_EditorEvent(ptr unsafe.Pointer, event unsafe.Pointer, model unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalDelegate::editorEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::editorEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent, *core.QAbstractItemModel, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).EditorEventDefault(core.NewQEventFromPointer(event), core.NewQAbstractItemModelFromPointer(model), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QSqlRelationalDelegate) ConnectEditorEvent(f func(event *core.QEvent, model *core.QAbstractItemModel, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlRelationalDelegate::editorEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::editorEvent", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectEditorEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::editorEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::editorEvent")
	}
}

func (ptr *QSqlRelationalDelegate) EditorEvent(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_EditorEvent(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QSqlRelationalDelegate) EditorEventDefault(event core.QEvent_ITF, model core.QAbstractItemModel_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::editorEvent")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_EditorEventDefault(ptr.Pointer(), core.PointerFromQEvent(event), core.PointerFromQAbstractItemModel(model), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_Paint
func callbackQSqlRelationalDelegate_Paint(ptr unsafe.Pointer, painter unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::paint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::paint"); signal != nil {
		signal.(func(*gui.QPainter, *widgets.QStyleOptionViewItem, *core.QModelIndex))(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).PaintDefault(gui.NewQPainterFromPointer(painter), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectPaint(f func(painter *gui.QPainter, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::paint", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectPaint() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::paint")
	}
}

func (ptr *QSqlRelationalDelegate) Paint(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::paint")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) PaintDefault(painter gui.QPainter_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::paint")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_SetEditorData
func callbackQSqlRelationalDelegate_SetEditorData(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::setEditorData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::setEditorData"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).SetEditorDataDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectSetEditorData(f func(editor *widgets.QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::setEditorData", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectSetEditorData() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::setEditorData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::setEditorData")
	}
}

func (ptr *QSqlRelationalDelegate) SetEditorData(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_SetEditorData(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) SetEditorDataDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::setEditorData")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_SetEditorDataDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_SizeHint
func callbackQSqlRelationalDelegate_SizeHint(ptr unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalDelegate::sizeHint")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func(*widgets.QStyleOptionViewItem, *core.QModelIndex) *core.QSize)(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlRelationalDelegateFromPointer(ptr).SizeHintDefault(widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalDelegate) ConnectSizeHint(f func(option *widgets.QStyleOptionViewItem, index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QSqlRelationalDelegate::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::sizeHint", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::sizeHint")
	}
}

func (ptr *QSqlRelationalDelegate) SizeHint(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlRelationalDelegate::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlRelationalDelegate_SizeHint(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) SizeHintDefault(option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlRelationalDelegate::sizeHint")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlRelationalDelegate_SizeHintDefault(ptr.Pointer(), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalDelegate_UpdateEditorGeometry
func callbackQSqlRelationalDelegate_UpdateEditorGeometry(ptr unsafe.Pointer, editor unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::updateEditorGeometry")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::updateEditorGeometry"); signal != nil {
		signal.(func(*widgets.QWidget, *widgets.QStyleOptionViewItem, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).UpdateEditorGeometryDefault(widgets.NewQWidgetFromPointer(editor), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectUpdateEditorGeometry(f func(editor *widgets.QWidget, option *widgets.QStyleOptionViewItem, index *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::updateEditorGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::updateEditorGeometry", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectUpdateEditorGeometry() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::updateEditorGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::updateEditorGeometry")
	}
}

func (ptr *QSqlRelationalDelegate) UpdateEditorGeometry(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::updateEditorGeometry")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_UpdateEditorGeometry(ptr.Pointer(), widgets.PointerFromQWidget(editor), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) UpdateEditorGeometryDefault(editor widgets.QWidget_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::updateEditorGeometry")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_UpdateEditorGeometryDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_DestroyEditor
func callbackQSqlRelationalDelegate_DestroyEditor(ptr unsafe.Pointer, editor unsafe.Pointer, index unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::destroyEditor")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::destroyEditor"); signal != nil {
		signal.(func(*widgets.QWidget, *core.QModelIndex))(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DestroyEditorDefault(widgets.NewQWidgetFromPointer(editor), core.NewQModelIndexFromPointer(index))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDestroyEditor(f func(editor *widgets.QWidget, index *core.QModelIndex)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::destroyEditor", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDestroyEditor() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::destroyEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::destroyEditor")
	}
}

func (ptr *QSqlRelationalDelegate) DestroyEditor(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyEditor(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QSqlRelationalDelegate) DestroyEditorDefault(editor widgets.QWidget_ITF, index core.QModelIndex_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::destroyEditor")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DestroyEditorDefault(ptr.Pointer(), widgets.PointerFromQWidget(editor), core.PointerFromQModelIndex(index))
	}
}

//export callbackQSqlRelationalDelegate_HelpEvent
func callbackQSqlRelationalDelegate_HelpEvent(ptr unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer, option unsafe.Pointer, index unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalDelegate::helpEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::helpEvent"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*gui.QHelpEvent, *widgets.QAbstractItemView, *widgets.QStyleOptionViewItem, *core.QModelIndex) bool)(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).HelpEventDefault(gui.NewQHelpEventFromPointer(event), widgets.NewQAbstractItemViewFromPointer(view), widgets.NewQStyleOptionViewItemFromPointer(option), core.NewQModelIndexFromPointer(index)))))
}

func (ptr *QSqlRelationalDelegate) ConnectHelpEvent(f func(event *gui.QHelpEvent, view *widgets.QAbstractItemView, option *widgets.QStyleOptionViewItem, index *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlRelationalDelegate::helpEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::helpEvent", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectHelpEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::helpEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::helpEvent")
	}
}

func (ptr *QSqlRelationalDelegate) HelpEvent(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::helpEvent")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_HelpEvent(ptr.Pointer(), gui.PointerFromQHelpEvent(event), widgets.PointerFromQAbstractItemView(view), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QSqlRelationalDelegate) HelpEventDefault(event gui.QHelpEvent_ITF, view widgets.QAbstractItemView_ITF, option widgets.QStyleOptionViewItem_ITF, index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::helpEvent")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_HelpEventDefault(ptr.Pointer(), gui.PointerFromQHelpEvent(event), widgets.PointerFromQAbstractItemView(view), widgets.PointerFromQStyleOptionViewItem(option), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_TimerEvent
func callbackQSqlRelationalDelegate_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::timerEvent", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::timerEvent")
	}
}

func (ptr *QSqlRelationalDelegate) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalDelegate) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlRelationalDelegate_ChildEvent
func callbackQSqlRelationalDelegate_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::childEvent", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::childEvent")
	}
}

func (ptr *QSqlRelationalDelegate) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalDelegate) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlRelationalDelegate_ConnectNotify
func callbackQSqlRelationalDelegate_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::connectNotify", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::connectNotify")
	}
}

func (ptr *QSqlRelationalDelegate) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalDelegate_CustomEvent
func callbackQSqlRelationalDelegate_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::customEvent", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::customEvent")
	}
}

func (ptr *QSqlRelationalDelegate) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlRelationalDelegate) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlRelationalDelegate_DeleteLater
func callbackQSqlRelationalDelegate_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlRelationalDelegate::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::deleteLater", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::deleteLater")
	}
}

func (ptr *QSqlRelationalDelegate) DeleteLater() {
	defer qt.Recovering("QSqlRelationalDelegate::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlRelationalDelegate) DeleteLaterDefault() {
	defer qt.Recovering("QSqlRelationalDelegate::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlRelationalDelegate_DisconnectNotify
func callbackQSqlRelationalDelegate_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalDelegate::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalDelegateFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalDelegate) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlRelationalDelegate::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::disconnectNotify", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::disconnectNotify")
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalDelegate::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalDelegate_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalDelegate_Event
func callbackQSqlRelationalDelegate_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalDelegate::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalDelegateFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlRelationalDelegate) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlRelationalDelegate::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::event", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::event")
	}
}

func (ptr *QSqlRelationalDelegate) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::event")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlRelationalDelegate) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalDelegate::event")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalDelegate_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlRelationalDelegate_MetaObject
func callbackQSqlRelationalDelegate_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalDelegate::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalDelegate::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlRelationalDelegateFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlRelationalDelegate) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlRelationalDelegate::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::metaObject", f)
	}
}

func (ptr *QSqlRelationalDelegate) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlRelationalDelegate::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalDelegate::metaObject")
	}
}

func (ptr *QSqlRelationalDelegate) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlRelationalDelegate::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlRelationalDelegate_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlRelationalDelegate) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlRelationalDelegate::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlRelationalDelegate_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int64

const (
	QSqlRelationalTableModel__InnerJoin = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  = QSqlRelationalTableModel__JoinMode(1)
)

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModel_ITF interface {
	QSqlTableModel_ITF
	QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel
}

func (p *QSqlRelationalTableModel) QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel {
	return p
}

func (p *QSqlRelationalTableModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func (p *QSqlRelationalTableModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSqlTableModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
	return n
}
func NewQSqlRelationalTableModel(parent core.QObject_ITF, db QSqlDatabase_ITF) *QSqlRelationalTableModel {
	defer qt.Recovering("QSqlRelationalTableModel::QSqlRelationalTableModel")

	var tmpValue = NewQSqlRelationalTableModelFromPointer(C.QSqlRelationalTableModel_NewQSqlRelationalTableModel(core.PointerFromQObject(parent), PointerFromQSqlDatabase(db)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlRelationalTableModel_Clear
func callbackQSqlRelationalTableModel_Clear(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::clear")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::clear", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::clear")
	}
}

func (ptr *QSqlRelationalTableModel) Clear() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ClearDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlRelationalTableModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlRelationalTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_InsertRowIntoTable
func callbackQSqlRelationalTableModel_InsertRowIntoTable(ptr unsafe.Pointer, values unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::insertRowIntoTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::insertRowIntoTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QSqlRecord) bool)(NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).InsertRowIntoTableDefault(NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlRelationalTableModel) ConnectInsertRowIntoTable(f func(values *QSqlRecord) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::insertRowIntoTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectInsertRowIntoTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::insertRowIntoTable")
	}
}

func (ptr *QSqlRelationalTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) InsertRowIntoTableDefault(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_InsertRowIntoTableDefault(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_OrderByClause
func callbackQSqlRelationalTableModel_OrderByClause(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::orderByClause")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::orderByClause"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSqlRelationalTableModelFromPointer(ptr).OrderByClauseDefault())
}

func (ptr *QSqlRelationalTableModel) ConnectOrderByClause(f func() string) {
	defer qt.Recovering("connect QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::orderByClause", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectOrderByClause() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::orderByClause")
	}
}

func (ptr *QSqlRelationalTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) OrderByClauseDefault() string {
	defer qt.Recovering("QSqlRelationalTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_OrderByClauseDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlRelationalTableModel_RelationModel
func callbackQSqlRelationalTableModel_RelationModel(ptr unsafe.Pointer, column C.int) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::relationModel")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::relationModel"); signal != nil {
		return PointerFromQSqlTableModel(signal.(func(int) *QSqlTableModel)(int(int32(column))))
	}

	return PointerFromQSqlTableModel(NewQSqlRelationalTableModelFromPointer(ptr).RelationModelDefault(int(int32(column))))
}

func (ptr *QSqlRelationalTableModel) ConnectRelationModel(f func(column int) *QSqlTableModel) {
	defer qt.Recovering("connect QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::relationModel", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRelationModel() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::relationModel")
	}
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	defer qt.Recovering("QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModel(ptr.Pointer(), C.int(int32(column))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RelationModelDefault(column int) *QSqlTableModel {
	defer qt.Recovering("QSqlRelationalTableModel::relationModel")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModelDefault(ptr.Pointer(), C.int(int32(column))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_RevertRow
func callbackQSqlRelationalTableModel_RevertRow(ptr unsafe.Pointer, row C.int) {
	defer qt.Recovering("callback QSqlRelationalTableModel::revertRow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::revertRow"); signal != nil {
		signal.(func(int))(int(int32(row)))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).RevertRowDefault(int(int32(row)))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revertRow", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revertRow")
	}
}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlRelationalTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlRelationalTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRowDefault(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackQSqlRelationalTableModel_Select
func callbackQSqlRelationalTableModel_Select(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::select")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::select"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).SelectDefault())))
}

func (ptr *QSqlRelationalTableModel) ConnectSelect(f func() bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::select", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSelect() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::select")
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	defer qt.Recovering("QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SelectDefault() bool {
	defer qt.Recovering("QSqlRelationalTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SelectDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_SelectStatement
func callbackQSqlRelationalTableModel_SelectStatement(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::selectStatement")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::selectStatement"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSqlRelationalTableModelFromPointer(ptr).SelectStatementDefault())
}

func (ptr *QSqlRelationalTableModel) ConnectSelectStatement(f func() string) {
	defer qt.Recovering("connect QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::selectStatement", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSelectStatement() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::selectStatement")
	}
}

func (ptr *QSqlRelationalTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) SelectStatementDefault() string {
	defer qt.Recovering("QSqlRelationalTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_SelectStatementDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	defer qt.Recovering("QSqlRelationalTableModel::setJoinMode")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(ptr.Pointer(), C.longlong(joinMode))
	}
}

//export callbackQSqlRelationalTableModel_SetRelation
func callbackQSqlRelationalTableModel_SetRelation(ptr unsafe.Pointer, column C.int, relation unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setRelation")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::setRelation"); signal != nil {
		signal.(func(int, *QSqlRelation))(int(int32(column)), NewQSqlRelationFromPointer(relation))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetRelationDefault(int(int32(column)), NewQSqlRelationFromPointer(relation))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetRelation(f func(column int, relation *QSqlRelation)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setRelation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setRelation", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetRelation() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setRelation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setRelation")
	}
}

func (ptr *QSqlRelationalTableModel) SetRelation(column int, relation QSqlRelation_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::setRelation")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelation(ptr.Pointer(), C.int(int32(column)), PointerFromQSqlRelation(relation))
	}
}

func (ptr *QSqlRelationalTableModel) SetRelationDefault(column int, relation QSqlRelation_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::setRelation")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelationDefault(ptr.Pointer(), C.int(int32(column)), PointerFromQSqlRelation(relation))
	}
}

//export callbackQSqlRelationalTableModel_SetTable
func callbackQSqlRelationalTableModel_SetTable(ptr unsafe.Pointer, table *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::setTable"); signal != nil {
		signal.(func(string))(C.GoString(table))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetTableDefault(C.GoString(table))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetTable(f func(table string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setTable")
	}
}

func (ptr *QSqlRelationalTableModel) SetTable(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		var tableC = C.CString(table)
		defer C.free(unsafe.Pointer(tableC))
		C.QSqlRelationalTableModel_SetTable(ptr.Pointer(), tableC)
	}
}

func (ptr *QSqlRelationalTableModel) SetTableDefault(table string) {
	defer qt.Recovering("QSqlRelationalTableModel::setTable")

	if ptr.Pointer() != nil {
		var tableC = C.CString(table)
		defer C.free(unsafe.Pointer(tableC))
		C.QSqlRelationalTableModel_SetTableDefault(ptr.Pointer(), tableC)
	}
}

//export callbackQSqlRelationalTableModel_UpdateRowInTable
func callbackQSqlRelationalTableModel_UpdateRowInTable(ptr unsafe.Pointer, row C.int, values unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::updateRowInTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::updateRowInTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *QSqlRecord) bool)(int(int32(row)), NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).UpdateRowInTableDefault(int(int32(row)), NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlRelationalTableModel) ConnectUpdateRowInTable(f func(row int, values *QSqlRecord) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::updateRowInTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectUpdateRowInTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::updateRowInTable")
	}
}

func (ptr *QSqlRelationalTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_UpdateRowInTable(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) UpdateRowInTableDefault(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_UpdateRowInTableDefault(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel
func callbackQSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::~QSqlRelationalTableModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).DestroyQSqlRelationalTableModelDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectDestroyQSqlRelationalTableModel(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::~QSqlRelationalTableModel", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDestroyQSqlRelationalTableModel() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::~QSqlRelationalTableModel")
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	defer qt.Recovering("QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModelDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::~QSqlRelationalTableModel")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlRelationalTableModel_DeleteRowFromTable
func callbackQSqlRelationalTableModel_DeleteRowFromTable(ptr unsafe.Pointer, row C.int) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::deleteRowFromTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::deleteRowFromTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).DeleteRowFromTableDefault(int(int32(row))))))
}

func (ptr *QSqlRelationalTableModel) ConnectDeleteRowFromTable(f func(row int) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::deleteRowFromTable", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDeleteRowFromTable() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::deleteRowFromTable")
	}
}

func (ptr *QSqlRelationalTableModel) DeleteRowFromTable(row int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_DeleteRowFromTable(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) DeleteRowFromTableDefault(row int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_DeleteRowFromTableDefault(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_IndexInQuery
func callbackQSqlRelationalTableModel_IndexInQuery(ptr unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::indexInQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::indexInQuery"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(item)))
	}

	return core.PointerFromQModelIndex(NewQSqlRelationalTableModelFromPointer(ptr).IndexInQueryDefault(core.NewQModelIndexFromPointer(item)))
}

func (ptr *QSqlRelationalTableModel) ConnectIndexInQuery(f func(item *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlRelationalTableModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::indexInQuery", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectIndexInQuery() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::indexInQuery")
	}
}

func (ptr *QSqlRelationalTableModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) IndexInQueryDefault(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_IndexInQueryDefault(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_Revert
func callbackQSqlRelationalTableModel_Revert(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::revert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::revert"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).RevertDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revert", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revert")
	}
}

func (ptr *QSqlRelationalTableModel) Revert() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) RevertDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertDefault(ptr.Pointer())
	}
}

//export callbackQSqlRelationalTableModel_RevertAll
func callbackQSqlRelationalTableModel_RevertAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::revertAll")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::revertAll"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).RevertAllDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectRevertAll(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::revertAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revertAll", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectRevertAll() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::revertAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::revertAll")
	}
}

func (ptr *QSqlRelationalTableModel) RevertAll() {
	defer qt.Recovering("QSqlRelationalTableModel::revertAll")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertAll(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) RevertAllDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::revertAll")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertAllDefault(ptr.Pointer())
	}
}

//export callbackQSqlRelationalTableModel_SelectRow
func callbackQSqlRelationalTableModel_SelectRow(ptr unsafe.Pointer, row C.int) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::selectRow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::selectRow"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).SelectRowDefault(int(int32(row))))))
}

func (ptr *QSqlRelationalTableModel) ConnectSelectRow(f func(row int) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::selectRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::selectRow", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSelectRow() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::selectRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::selectRow")
	}
}

func (ptr *QSqlRelationalTableModel) SelectRow(row int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SelectRow(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SelectRowDefault(row int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SelectRowDefault(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_SetEditStrategy
func callbackQSqlRelationalTableModel_SetEditStrategy(ptr unsafe.Pointer, strategy C.longlong) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setEditStrategy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setEditStrategy", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setEditStrategy")
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategy(ptr.Pointer(), C.longlong(strategy))
	}
}

func (ptr *QSqlRelationalTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlRelationalTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetEditStrategyDefault(ptr.Pointer(), C.longlong(strategy))
	}
}

//export callbackQSqlRelationalTableModel_SetFilter
func callbackQSqlRelationalTableModel_SetFilter(ptr unsafe.Pointer, filter *C.char) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setFilter", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setFilter")
	}
}

func (ptr *QSqlRelationalTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		C.QSqlRelationalTableModel_SetFilter(ptr.Pointer(), filterC)
	}
}

func (ptr *QSqlRelationalTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlRelationalTableModel::setFilter")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		C.QSqlRelationalTableModel_SetFilterDefault(ptr.Pointer(), filterC)
	}
}

//export callbackQSqlRelationalTableModel_SetSort
func callbackQSqlRelationalTableModel_SetSort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	defer qt.Recovering("callback QSqlRelationalTableModel::setSort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).SetSortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setSort", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::setSort")
	}
}

func (ptr *QSqlRelationalTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QSqlRelationalTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlRelationalTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetSortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlRelationalTableModel_Submit
func callbackQSqlRelationalTableModel_Submit(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::submit")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).SubmitDefault())))
}

func (ptr *QSqlRelationalTableModel) ConnectSubmit(f func() bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::submit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::submit", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSubmit() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::submit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::submit")
	}
}

func (ptr *QSqlRelationalTableModel) Submit() bool {
	defer qt.Recovering("QSqlRelationalTableModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SubmitDefault() bool {
	defer qt.Recovering("QSqlRelationalTableModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SubmitDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_SubmitAll
func callbackQSqlRelationalTableModel_SubmitAll(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::submitAll")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::submitAll"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).SubmitAllDefault())))
}

func (ptr *QSqlRelationalTableModel) ConnectSubmitAll(f func() bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::submitAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::submitAll", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSubmitAll() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::submitAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::submitAll")
	}
}

func (ptr *QSqlRelationalTableModel) SubmitAll() bool {
	defer qt.Recovering("QSqlRelationalTableModel::submitAll")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SubmitAll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SubmitAllDefault() bool {
	defer qt.Recovering("QSqlRelationalTableModel::submitAll")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SubmitAllDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_QueryChange
func callbackQSqlRelationalTableModel_QueryChange(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::queryChange")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::queryChange", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::queryChange")
	}
}

func (ptr *QSqlRelationalTableModel) QueryChange() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

//export callbackQSqlRelationalTableModel_Index
func callbackQSqlRelationalTableModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::index")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQSqlRelationalTableModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QSqlRelationalTableModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlRelationalTableModel::index")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::index", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectIndex() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::index")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::index")
	}
}

func (ptr *QSqlRelationalTableModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_DropMimeData
func callbackQSqlRelationalTableModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::dropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlRelationalTableModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::dropMimeData", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDropMimeData() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::dropMimeData")
	}
}

func (ptr *QSqlRelationalTableModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_Sibling
func callbackQSqlRelationalTableModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::sibling")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQSqlRelationalTableModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QSqlRelationalTableModel) ConnectSibling(f func(row int, column int, idx *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlRelationalTableModel::sibling")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::sibling", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSibling() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::sibling")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::sibling")
	}
}

func (ptr *QSqlRelationalTableModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_Buddy
func callbackQSqlRelationalTableModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::buddy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlRelationalTableModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalTableModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlRelationalTableModel::buddy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::buddy", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectBuddy() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::buddy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::buddy")
	}
}

func (ptr *QSqlRelationalTableModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_CanDropMimeData
func callbackQSqlRelationalTableModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::canDropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlRelationalTableModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::canDropMimeData", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectCanDropMimeData() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::canDropMimeData")
	}
}

func (ptr *QSqlRelationalTableModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_HasChildren
func callbackQSqlRelationalTableModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::hasChildren")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlRelationalTableModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::hasChildren", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectHasChildren() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::hasChildren")
	}
}

func (ptr *QSqlRelationalTableModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_MimeTypes
func callbackQSqlRelationalTableModel_MimeTypes(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::mimeTypes")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQSqlRelationalTableModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QSqlRelationalTableModel) ConnectMimeTypes(f func() []string) {
	defer qt.Recovering("connect QSqlRelationalTableModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::mimeTypes", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectMimeTypes() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::mimeTypes")
	}
}

func (ptr *QSqlRelationalTableModel) MimeTypes() []string {
	defer qt.Recovering("QSqlRelationalTableModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlRelationalTableModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlRelationalTableModel) MimeTypesDefault() []string {
	defer qt.Recovering("QSqlRelationalTableModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlRelationalTableModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlRelationalTableModel_MoveColumns
func callbackQSqlRelationalTableModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::moveColumns")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlRelationalTableModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::moveColumns", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectMoveColumns() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::moveColumns")
	}
}

func (ptr *QSqlRelationalTableModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_MoveRows
func callbackQSqlRelationalTableModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::moveRows")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlRelationalTableModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::moveRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::moveRows", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectMoveRows() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::moveRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::moveRows")
	}
}

func (ptr *QSqlRelationalTableModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlRelationalTableModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_Parent
func callbackQSqlRelationalTableModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::parent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlRelationalTableModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalTableModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlRelationalTableModel::parent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::parent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectParent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::parent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::parent")
	}
}

func (ptr *QSqlRelationalTableModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlRelationalTableModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlRelationalTableModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_ResetInternalData
func callbackQSqlRelationalTableModel_ResetInternalData(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::resetInternalData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectResetInternalData(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::resetInternalData", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectResetInternalData() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::resetInternalData")
	}
}

func (ptr *QSqlRelationalTableModel) ResetInternalData() {
	defer qt.Recovering("QSqlRelationalTableModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) ResetInternalDataDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQSqlRelationalTableModel_Span
func callbackQSqlRelationalTableModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::span")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlRelationalTableModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlRelationalTableModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QSqlRelationalTableModel::span")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::span", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSpan() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::span")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::span")
	}
}

func (ptr *QSqlRelationalTableModel) Span(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlRelationalTableModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlRelationalTableModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlRelationalTableModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlRelationalTableModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlRelationalTableModel_SupportedDragActions
func callbackQSqlRelationalTableModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlRelationalTableModel::supportedDragActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlRelationalTableModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QSqlRelationalTableModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlRelationalTableModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::supportedDragActions", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSupportedDragActions() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::supportedDragActions")
	}
}

func (ptr *QSqlRelationalTableModel) SupportedDragActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlRelationalTableModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlRelationalTableModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlRelationalTableModel) SupportedDragActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlRelationalTableModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlRelationalTableModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlRelationalTableModel_SupportedDropActions
func callbackQSqlRelationalTableModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlRelationalTableModel::supportedDropActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlRelationalTableModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QSqlRelationalTableModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlRelationalTableModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::supportedDropActions", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectSupportedDropActions() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::supportedDropActions")
	}
}

func (ptr *QSqlRelationalTableModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlRelationalTableModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlRelationalTableModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlRelationalTableModel) SupportedDropActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlRelationalTableModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlRelationalTableModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlRelationalTableModel_TimerEvent
func callbackQSqlRelationalTableModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::timerEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::timerEvent")
	}
}

func (ptr *QSqlRelationalTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlRelationalTableModel_ChildEvent
func callbackQSqlRelationalTableModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::childEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::childEvent")
	}
}

func (ptr *QSqlRelationalTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlRelationalTableModel_ConnectNotify
func callbackQSqlRelationalTableModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::connectNotify", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::connectNotify")
	}
}

func (ptr *QSqlRelationalTableModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalTableModel_CustomEvent
func callbackQSqlRelationalTableModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::customEvent", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::customEvent")
	}
}

func (ptr *QSqlRelationalTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlRelationalTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlRelationalTableModel_DeleteLater
func callbackQSqlRelationalTableModel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlRelationalTableModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlRelationalTableModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::deleteLater", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::deleteLater")
	}
}

func (ptr *QSqlRelationalTableModel) DeleteLater() {
	defer qt.Recovering("QSqlRelationalTableModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlRelationalTableModel) DeleteLaterDefault() {
	defer qt.Recovering("QSqlRelationalTableModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlRelationalTableModel_DisconnectNotify
func callbackQSqlRelationalTableModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlRelationalTableModel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlRelationalTableModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlRelationalTableModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlRelationalTableModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::disconnectNotify", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::disconnectNotify")
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlRelationalTableModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlRelationalTableModel_Event
func callbackQSqlRelationalTableModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlRelationalTableModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::event", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::event")
	}
}

func (ptr *QSqlRelationalTableModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_EventFilter
func callbackQSqlRelationalTableModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlRelationalTableModel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlRelationalTableModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlRelationalTableModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlRelationalTableModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::eventFilter", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::eventFilter")
	}
}

func (ptr *QSqlRelationalTableModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlRelationalTableModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSqlRelationalTableModel_MetaObject
func callbackQSqlRelationalTableModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlRelationalTableModel::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlRelationalTableModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlRelationalTableModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlRelationalTableModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlRelationalTableModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::metaObject", f)
	}
}

func (ptr *QSqlRelationalTableModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlRelationalTableModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlRelationalTableModel::metaObject")
	}
}

func (ptr *QSqlRelationalTableModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlRelationalTableModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlRelationalTableModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlRelationalTableModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlRelationalTableModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSqlResult::BindingSyntax
type QSqlResult__BindingSyntax int64

const (
	QSqlResult__PositionalBinding = QSqlResult__BindingSyntax(0)
	QSqlResult__NamedBinding      = QSqlResult__BindingSyntax(1)
)

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (p *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return p
}

func (p *QSqlResult) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSqlResult) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = new(QSqlResult)
	n.SetPointer(ptr)
	return n
}
func NewQSqlResult(db QSqlDriver_ITF) *QSqlResult {
	defer qt.Recovering("QSqlResult::QSqlResult")

	return NewQSqlResultFromPointer(C.QSqlResult_NewQSqlResult(PointerFromQSqlDriver(db)))
}

func (ptr *QSqlResult) At() int {
	defer qt.Recovering("QSqlResult::at")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_At(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlResult) BindingSyntax() QSqlResult__BindingSyntax {
	defer qt.Recovering("QSqlResult::bindingSyntax")

	if ptr.Pointer() != nil {
		return QSqlResult__BindingSyntax(C.QSqlResult_BindingSyntax(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlResult) BoundValue2(placeholder string) *core.QVariant {
	defer qt.Recovering("QSqlResult::boundValue")

	if ptr.Pointer() != nil {
		var placeholderC = C.CString(placeholder)
		defer C.free(unsafe.Pointer(placeholderC))
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_BoundValue2(ptr.Pointer(), placeholderC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) BoundValue(index int) *core.QVariant {
	defer qt.Recovering("QSqlResult::boundValue")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_BoundValue(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) BoundValueCount() int {
	defer qt.Recovering("QSqlResult::boundValueCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_BoundValueCount(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlResult) BoundValueName(index int) string {
	defer qt.Recovering("QSqlResult::boundValueName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_BoundValueName(ptr.Pointer(), C.int(int32(index))))
	}
	return ""
}

func (ptr *QSqlResult) Clear() {
	defer qt.Recovering("QSqlResult::clear")

	if ptr.Pointer() != nil {
		C.QSqlResult_Clear(ptr.Pointer())
	}
}

//export callbackQSqlResult_Data
func callbackQSqlResult_Data(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	defer qt.Recovering("callback QSqlResult::data")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::data"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(index))))
	}

	return core.PointerFromQVariant(nil)
}

func (ptr *QSqlResult) ConnectData(f func(index int) *core.QVariant) {
	defer qt.Recovering("connect QSqlResult::data")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::data", f)
	}
}

func (ptr *QSqlResult) DisconnectData(index int) {
	defer qt.Recovering("disconnect QSqlResult::data")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::data")
	}
}

func (ptr *QSqlResult) Data(index int) *core.QVariant {
	defer qt.Recovering("QSqlResult::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_Data(ptr.Pointer(), C.int(int32(index))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) Driver() *QSqlDriver {
	defer qt.Recovering("QSqlResult::driver")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlDriverFromPointer(C.QSqlResult_Driver(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_Exec
func callbackQSqlResult_Exec(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlResult::exec")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::exec"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).ExecDefault())))
}

func (ptr *QSqlResult) ConnectExec(f func() bool) {
	defer qt.Recovering("connect QSqlResult::exec")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::exec", f)
	}
}

func (ptr *QSqlResult) DisconnectExec() {
	defer qt.Recovering("disconnect QSqlResult::exec")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::exec")
	}
}

func (ptr *QSqlResult) Exec() bool {
	defer qt.Recovering("QSqlResult::exec")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Exec(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) ExecDefault() bool {
	defer qt.Recovering("QSqlResult::exec")

	if ptr.Pointer() != nil {
		return C.QSqlResult_ExecDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) ExecutedQuery() string {
	defer qt.Recovering("QSqlResult::executedQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_ExecutedQuery(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlResult_Fetch
func callbackQSqlResult_Fetch(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QSqlResult::fetch")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::fetch"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetch(f func(index int) bool) {
	defer qt.Recovering("connect QSqlResult::fetch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetch", f)
	}
}

func (ptr *QSqlResult) DisconnectFetch(index int) {
	defer qt.Recovering("disconnect QSqlResult::fetch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetch")
	}
}

func (ptr *QSqlResult) Fetch(index int) bool {
	defer qt.Recovering("QSqlResult::fetch")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Fetch(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

//export callbackQSqlResult_FetchFirst
func callbackQSqlResult_FetchFirst(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlResult::fetchFirst")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::fetchFirst"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetchFirst(f func() bool) {
	defer qt.Recovering("connect QSqlResult::fetchFirst")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchFirst", f)
	}
}

func (ptr *QSqlResult) DisconnectFetchFirst() {
	defer qt.Recovering("disconnect QSqlResult::fetchFirst")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchFirst")
	}
}

func (ptr *QSqlResult) FetchFirst() bool {
	defer qt.Recovering("QSqlResult::fetchFirst")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchFirst(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlResult_FetchLast
func callbackQSqlResult_FetchLast(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlResult::fetchLast")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::fetchLast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectFetchLast(f func() bool) {
	defer qt.Recovering("connect QSqlResult::fetchLast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchLast", f)
	}
}

func (ptr *QSqlResult) DisconnectFetchLast() {
	defer qt.Recovering("disconnect QSqlResult::fetchLast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchLast")
	}
}

func (ptr *QSqlResult) FetchLast() bool {
	defer qt.Recovering("QSqlResult::fetchLast")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchLast(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlResult_FetchNext
func callbackQSqlResult_FetchNext(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlResult::fetchNext")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::fetchNext"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).FetchNextDefault())))
}

func (ptr *QSqlResult) ConnectFetchNext(f func() bool) {
	defer qt.Recovering("connect QSqlResult::fetchNext")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchNext", f)
	}
}

func (ptr *QSqlResult) DisconnectFetchNext() {
	defer qt.Recovering("disconnect QSqlResult::fetchNext")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchNext")
	}
}

func (ptr *QSqlResult) FetchNext() bool {
	defer qt.Recovering("QSqlResult::fetchNext")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchNext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchNextDefault() bool {
	defer qt.Recovering("QSqlResult::fetchNext")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchNextDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlResult_FetchPrevious
func callbackQSqlResult_FetchPrevious(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlResult::fetchPrevious")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::fetchPrevious"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).FetchPreviousDefault())))
}

func (ptr *QSqlResult) ConnectFetchPrevious(f func() bool) {
	defer qt.Recovering("connect QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchPrevious", f)
	}
}

func (ptr *QSqlResult) DisconnectFetchPrevious() {
	defer qt.Recovering("disconnect QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::fetchPrevious")
	}
}

func (ptr *QSqlResult) FetchPrevious() bool {
	defer qt.Recovering("QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchPrevious(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchPreviousDefault() bool {
	defer qt.Recovering("QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchPreviousDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlResult_Handle
func callbackQSqlResult_Handle(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlResult::handle")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::handle"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlResultFromPointer(ptr).HandleDefault())
}

func (ptr *QSqlResult) ConnectHandle(f func() *core.QVariant) {
	defer qt.Recovering("connect QSqlResult::handle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::handle", f)
	}
}

func (ptr *QSqlResult) DisconnectHandle() {
	defer qt.Recovering("disconnect QSqlResult::handle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::handle")
	}
}

func (ptr *QSqlResult) Handle() *core.QVariant {
	defer qt.Recovering("QSqlResult::handle")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) HandleDefault() *core.QVariant {
	defer qt.Recovering("QSqlResult::handle")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_HandleDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) HasOutValues() bool {
	defer qt.Recovering("QSqlResult::hasOutValues")

	if ptr.Pointer() != nil {
		return C.QSqlResult_HasOutValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) IsActive() bool {
	defer qt.Recovering("QSqlResult::isActive")

	if ptr.Pointer() != nil {
		return C.QSqlResult_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) IsForwardOnly() bool {
	defer qt.Recovering("QSqlResult::isForwardOnly")

	if ptr.Pointer() != nil {
		return C.QSqlResult_IsForwardOnly(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlResult_IsNull
func callbackQSqlResult_IsNull(ptr unsafe.Pointer, index C.int) C.char {
	defer qt.Recovering("callback QSqlResult::isNull")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::isNull"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(index))))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectIsNull(f func(index int) bool) {
	defer qt.Recovering("connect QSqlResult::isNull")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::isNull", f)
	}
}

func (ptr *QSqlResult) DisconnectIsNull(index int) {
	defer qt.Recovering("disconnect QSqlResult::isNull")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::isNull")
	}
}

func (ptr *QSqlResult) IsNull(index int) bool {
	defer qt.Recovering("QSqlResult::isNull")

	if ptr.Pointer() != nil {
		return C.QSqlResult_IsNull(ptr.Pointer(), C.int(int32(index))) != 0
	}
	return false
}

func (ptr *QSqlResult) IsSelect() bool {
	defer qt.Recovering("QSqlResult::isSelect")

	if ptr.Pointer() != nil {
		return C.QSqlResult_IsSelect(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) IsValid() bool {
	defer qt.Recovering("QSqlResult::isValid")

	if ptr.Pointer() != nil {
		return C.QSqlResult_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) LastError() *QSqlError {
	defer qt.Recovering("QSqlResult::lastError")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlErrorFromPointer(C.QSqlResult_LastError(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlError).DestroyQSqlError)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_LastInsertId
func callbackQSqlResult_LastInsertId(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlResult::lastInsertId")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::lastInsertId"); signal != nil {
		return core.PointerFromQVariant(signal.(func() *core.QVariant)())
	}

	return core.PointerFromQVariant(NewQSqlResultFromPointer(ptr).LastInsertIdDefault())
}

func (ptr *QSqlResult) ConnectLastInsertId(f func() *core.QVariant) {
	defer qt.Recovering("connect QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::lastInsertId", f)
	}
}

func (ptr *QSqlResult) DisconnectLastInsertId() {
	defer qt.Recovering("disconnect QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::lastInsertId")
	}
}

func (ptr *QSqlResult) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_LastInsertId(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) LastInsertIdDefault() *core.QVariant {
	defer qt.Recovering("QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlResult_LastInsertIdDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) LastQuery() string {
	defer qt.Recovering("QSqlResult::lastQuery")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_LastQuery(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlResult_NumRowsAffected
func callbackQSqlResult_NumRowsAffected(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSqlResult::numRowsAffected")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::numRowsAffected"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSqlResult) ConnectNumRowsAffected(f func() int) {
	defer qt.Recovering("connect QSqlResult::numRowsAffected")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::numRowsAffected", f)
	}
}

func (ptr *QSqlResult) DisconnectNumRowsAffected() {
	defer qt.Recovering("disconnect QSqlResult::numRowsAffected")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::numRowsAffected")
	}
}

func (ptr *QSqlResult) NumRowsAffected() int {
	defer qt.Recovering("QSqlResult::numRowsAffected")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_NumRowsAffected(ptr.Pointer())))
	}
	return 0
}

//export callbackQSqlResult_Prepare
func callbackQSqlResult_Prepare(ptr unsafe.Pointer, query *C.char) C.char {
	defer qt.Recovering("callback QSqlResult::prepare")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::prepare"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).PrepareDefault(C.GoString(query)))))
}

func (ptr *QSqlResult) ConnectPrepare(f func(query string) bool) {
	defer qt.Recovering("connect QSqlResult::prepare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::prepare", f)
	}
}

func (ptr *QSqlResult) DisconnectPrepare() {
	defer qt.Recovering("disconnect QSqlResult::prepare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::prepare")
	}
}

func (ptr *QSqlResult) Prepare(query string) bool {
	defer qt.Recovering("QSqlResult::prepare")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlResult_Prepare(ptr.Pointer(), queryC) != 0
	}
	return false
}

func (ptr *QSqlResult) PrepareDefault(query string) bool {
	defer qt.Recovering("QSqlResult::prepare")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlResult_PrepareDefault(ptr.Pointer(), queryC) != 0
	}
	return false
}

//export callbackQSqlResult_Record
func callbackQSqlResult_Record(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlResult::record")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::record"); signal != nil {
		return PointerFromQSqlRecord(signal.(func() *QSqlRecord)())
	}

	return PointerFromQSqlRecord(NewQSqlResultFromPointer(ptr).RecordDefault())
}

func (ptr *QSqlResult) ConnectRecord(f func() *QSqlRecord) {
	defer qt.Recovering("connect QSqlResult::record")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::record", f)
	}
}

func (ptr *QSqlResult) DisconnectRecord() {
	defer qt.Recovering("disconnect QSqlResult::record")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::record")
	}
}

func (ptr *QSqlResult) Record() *QSqlRecord {
	defer qt.Recovering("QSqlResult::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlResult_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlResult) RecordDefault() *QSqlRecord {
	defer qt.Recovering("QSqlResult::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlResult_RecordDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

//export callbackQSqlResult_Reset
func callbackQSqlResult_Reset(ptr unsafe.Pointer, query *C.char) C.char {
	defer qt.Recovering("callback QSqlResult::reset")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlResult) ConnectReset(f func(query string) bool) {
	defer qt.Recovering("connect QSqlResult::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::reset", f)
	}
}

func (ptr *QSqlResult) DisconnectReset(query string) {
	defer qt.Recovering("disconnect QSqlResult::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::reset")
	}
}

func (ptr *QSqlResult) Reset(query string) bool {
	defer qt.Recovering("QSqlResult::reset")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlResult_Reset(ptr.Pointer(), queryC) != 0
	}
	return false
}

func (ptr *QSqlResult) ResetBindCount() {
	defer qt.Recovering("QSqlResult::resetBindCount")

	if ptr.Pointer() != nil {
		C.QSqlResult_ResetBindCount(ptr.Pointer())
	}
}

//export callbackQSqlResult_SavePrepare
func callbackQSqlResult_SavePrepare(ptr unsafe.Pointer, query *C.char) C.char {
	defer qt.Recovering("callback QSqlResult::savePrepare")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::savePrepare"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(string) bool)(C.GoString(query)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlResultFromPointer(ptr).SavePrepareDefault(C.GoString(query)))))
}

func (ptr *QSqlResult) ConnectSavePrepare(f func(query string) bool) {
	defer qt.Recovering("connect QSqlResult::savePrepare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::savePrepare", f)
	}
}

func (ptr *QSqlResult) DisconnectSavePrepare() {
	defer qt.Recovering("disconnect QSqlResult::savePrepare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::savePrepare")
	}
}

func (ptr *QSqlResult) SavePrepare(query string) bool {
	defer qt.Recovering("QSqlResult::savePrepare")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlResult_SavePrepare(ptr.Pointer(), queryC) != 0
	}
	return false
}

func (ptr *QSqlResult) SavePrepareDefault(query string) bool {
	defer qt.Recovering("QSqlResult::savePrepare")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		return C.QSqlResult_SavePrepareDefault(ptr.Pointer(), queryC) != 0
	}
	return false
}

//export callbackQSqlResult_SetActive
func callbackQSqlResult_SetActive(ptr unsafe.Pointer, active C.char) {
	defer qt.Recovering("callback QSqlResult::setActive")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setActive"); signal != nil {
		signal.(func(bool))(int8(active) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetActiveDefault(int8(active) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetActive(f func(active bool)) {
	defer qt.Recovering("connect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setActive", f)
	}
}

func (ptr *QSqlResult) DisconnectSetActive() {
	defer qt.Recovering("disconnect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setActive")
	}
}

func (ptr *QSqlResult) SetActive(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActive(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

func (ptr *QSqlResult) SetActiveDefault(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActiveDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(active))))
	}
}

//export callbackQSqlResult_SetAt
func callbackQSqlResult_SetAt(ptr unsafe.Pointer, index C.int) {
	defer qt.Recovering("callback QSqlResult::setAt")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setAt"); signal != nil {
		signal.(func(int))(int(int32(index)))
	} else {
		NewQSqlResultFromPointer(ptr).SetAtDefault(int(int32(index)))
	}
}

func (ptr *QSqlResult) ConnectSetAt(f func(index int)) {
	defer qt.Recovering("connect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setAt", f)
	}
}

func (ptr *QSqlResult) DisconnectSetAt() {
	defer qt.Recovering("disconnect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setAt")
	}
}

func (ptr *QSqlResult) SetAt(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAt(ptr.Pointer(), C.int(int32(index)))
	}
}

func (ptr *QSqlResult) SetAtDefault(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAtDefault(ptr.Pointer(), C.int(int32(index)))
	}
}

//export callbackQSqlResult_SetForwardOnly
func callbackQSqlResult_SetForwardOnly(ptr unsafe.Pointer, forward C.char) {
	defer qt.Recovering("callback QSqlResult::setForwardOnly")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setForwardOnly"); signal != nil {
		signal.(func(bool))(int8(forward) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetForwardOnlyDefault(int8(forward) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetForwardOnly(f func(forward bool)) {
	defer qt.Recovering("connect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setForwardOnly", f)
	}
}

func (ptr *QSqlResult) DisconnectSetForwardOnly() {
	defer qt.Recovering("disconnect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setForwardOnly")
	}
}

func (ptr *QSqlResult) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnly(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

func (ptr *QSqlResult) SetForwardOnlyDefault(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnlyDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(forward))))
	}
}

//export callbackQSqlResult_SetLastError
func callbackQSqlResult_SetLastError(ptr unsafe.Pointer, error unsafe.Pointer) {
	defer qt.Recovering("callback QSqlResult::setLastError")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setLastError"); signal != nil {
		signal.(func(*QSqlError))(NewQSqlErrorFromPointer(error))
	} else {
		NewQSqlResultFromPointer(ptr).SetLastErrorDefault(NewQSqlErrorFromPointer(error))
	}
}

func (ptr *QSqlResult) ConnectSetLastError(f func(error *QSqlError)) {
	defer qt.Recovering("connect QSqlResult::setLastError")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setLastError", f)
	}
}

func (ptr *QSqlResult) DisconnectSetLastError() {
	defer qt.Recovering("disconnect QSqlResult::setLastError")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setLastError")
	}
}

func (ptr *QSqlResult) SetLastError(error QSqlError_ITF) {
	defer qt.Recovering("QSqlResult::setLastError")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetLastError(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

func (ptr *QSqlResult) SetLastErrorDefault(error QSqlError_ITF) {
	defer qt.Recovering("QSqlResult::setLastError")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetLastErrorDefault(ptr.Pointer(), PointerFromQSqlError(error))
	}
}

//export callbackQSqlResult_SetQuery
func callbackQSqlResult_SetQuery(ptr unsafe.Pointer, query *C.char) {
	defer qt.Recovering("callback QSqlResult::setQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setQuery"); signal != nil {
		signal.(func(string))(C.GoString(query))
	} else {
		NewQSqlResultFromPointer(ptr).SetQueryDefault(C.GoString(query))
	}
}

func (ptr *QSqlResult) ConnectSetQuery(f func(query string)) {
	defer qt.Recovering("connect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setQuery", f)
	}
}

func (ptr *QSqlResult) DisconnectSetQuery() {
	defer qt.Recovering("disconnect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setQuery")
	}
}

func (ptr *QSqlResult) SetQuery(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		C.QSqlResult_SetQuery(ptr.Pointer(), queryC)
	}
}

func (ptr *QSqlResult) SetQueryDefault(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		var queryC = C.CString(query)
		defer C.free(unsafe.Pointer(queryC))
		C.QSqlResult_SetQueryDefault(ptr.Pointer(), queryC)
	}
}

//export callbackQSqlResult_SetSelect
func callbackQSqlResult_SetSelect(ptr unsafe.Pointer, sele C.char) {
	defer qt.Recovering("callback QSqlResult::setSelect")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::setSelect"); signal != nil {
		signal.(func(bool))(int8(sele) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetSelectDefault(int8(sele) != 0)
	}
}

func (ptr *QSqlResult) ConnectSetSelect(f func(sele bool)) {
	defer qt.Recovering("connect QSqlResult::setSelect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setSelect", f)
	}
}

func (ptr *QSqlResult) DisconnectSetSelect() {
	defer qt.Recovering("disconnect QSqlResult::setSelect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::setSelect")
	}
}

func (ptr *QSqlResult) SetSelect(sele bool) {
	defer qt.Recovering("QSqlResult::setSelect")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetSelect(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

func (ptr *QSqlResult) SetSelectDefault(sele bool) {
	defer qt.Recovering("QSqlResult::setSelect")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetSelectDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(sele))))
	}
}

//export callbackQSqlResult_Size
func callbackQSqlResult_Size(ptr unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSqlResult::size")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::size"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(0))
}

func (ptr *QSqlResult) ConnectSize(f func() int) {
	defer qt.Recovering("connect QSqlResult::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::size", f)
	}
}

func (ptr *QSqlResult) DisconnectSize() {
	defer qt.Recovering("disconnect QSqlResult::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::size")
	}
}

func (ptr *QSqlResult) Size() int {
	defer qt.Recovering("QSqlResult::size")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlResult_Size(ptr.Pointer())))
	}
	return 0
}

//export callbackQSqlResult_DestroyQSqlResult
func callbackQSqlResult_DestroyQSqlResult(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlResult::~QSqlResult")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlResult::~QSqlResult"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlResultFromPointer(ptr).DestroyQSqlResultDefault()
	}
}

func (ptr *QSqlResult) ConnectDestroyQSqlResult(f func()) {
	defer qt.Recovering("connect QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::~QSqlResult", f)
	}
}

func (ptr *QSqlResult) DisconnectDestroyQSqlResult() {
	defer qt.Recovering("disconnect QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlResult::~QSqlResult")
	}
}

func (ptr *QSqlResult) DestroyQSqlResult() {
	defer qt.Recovering("QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlResult) DestroyQSqlResultDefault() {
	defer qt.Recovering("QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResultDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int64

const (
	QSqlTableModel__OnFieldChange  = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit = QSqlTableModel__EditStrategy(2)
)

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModel_ITF interface {
	QSqlQueryModel_ITF
	QSqlTableModel_PTR() *QSqlTableModel
}

func (p *QSqlTableModel) QSqlTableModel_PTR() *QSqlTableModel {
	return p
}

func (p *QSqlTableModel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func (p *QSqlTableModel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSqlQueryModel_PTR().SetPointer(ptr)
	}
}

func PointerFromQSqlTableModel(ptr QSqlTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = new(QSqlTableModel)
	n.SetPointer(ptr)
	return n
}
func NewQSqlTableModel(parent core.QObject_ITF, db QSqlDatabase_ITF) *QSqlTableModel {
	defer qt.Recovering("QSqlTableModel::QSqlTableModel")

	var tmpValue = NewQSqlTableModelFromPointer(C.QSqlTableModel_NewQSqlTableModel(core.PointerFromQObject(parent), PointerFromQSqlDatabase(db)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQSqlTableModel_BeforeDelete
func callbackQSqlTableModel_BeforeDelete(ptr unsafe.Pointer, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::beforeDelete")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::beforeDelete"); signal != nil {
		signal.(func(int))(int(int32(row)))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeDelete(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeDelete", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	defer qt.Recovering("disconnect QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeDelete")
	}
}

func (ptr *QSqlTableModel) BeforeDelete(row int) {
	defer qt.Recovering("QSqlTableModel::beforeDelete")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeDelete(ptr.Pointer(), C.int(int32(row)))
	}
}

//export callbackQSqlTableModel_BeforeInsert
func callbackQSqlTableModel_BeforeInsert(ptr unsafe.Pointer, record unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::beforeInsert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::beforeInsert"); signal != nil {
		signal.(func(*QSqlRecord))(NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeInsert(f func(record *QSqlRecord)) {
	defer qt.Recovering("connect QSqlTableModel::beforeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeInsert(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeInsert", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeInsert() {
	defer qt.Recovering("disconnect QSqlTableModel::beforeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeInsert(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeInsert")
	}
}

func (ptr *QSqlTableModel) BeforeInsert(record QSqlRecord_ITF) {
	defer qt.Recovering("QSqlTableModel::beforeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeInsert(ptr.Pointer(), PointerFromQSqlRecord(record))
	}
}

//export callbackQSqlTableModel_BeforeUpdate
func callbackQSqlTableModel_BeforeUpdate(ptr unsafe.Pointer, row C.int, record unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::beforeUpdate")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::beforeUpdate"); signal != nil {
		signal.(func(int, *QSqlRecord))(int(int32(row)), NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectBeforeUpdate(f func(row int, record *QSqlRecord)) {
	defer qt.Recovering("connect QSqlTableModel::beforeUpdate")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeUpdate(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeUpdate", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeUpdate() {
	defer qt.Recovering("disconnect QSqlTableModel::beforeUpdate")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeUpdate(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::beforeUpdate")
	}
}

func (ptr *QSqlTableModel) BeforeUpdate(row int, record QSqlRecord_ITF) {
	defer qt.Recovering("QSqlTableModel::beforeUpdate")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_BeforeUpdate(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record))
	}
}

//export callbackQSqlTableModel_Clear
func callbackQSqlTableModel_Clear(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::clear")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::clear"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QSqlTableModel) ConnectClear(f func()) {
	defer qt.Recovering("connect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::clear", f)
	}
}

func (ptr *QSqlTableModel) DisconnectClear() {
	defer qt.Recovering("disconnect QSqlTableModel::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::clear")
	}
}

func (ptr *QSqlTableModel) Clear() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ClearDefault() {
	defer qt.Recovering("QSqlTableModel::clear")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::data")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) Database() *QSqlDatabase {
	defer qt.Recovering("QSqlTableModel::database")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlDatabaseFromPointer(C.QSqlTableModel_Database(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlDatabase).DestroyQSqlDatabase)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_DeleteRowFromTable
func callbackQSqlTableModel_DeleteRowFromTable(ptr unsafe.Pointer, row C.int) C.char {
	defer qt.Recovering("callback QSqlTableModel::deleteRowFromTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::deleteRowFromTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).DeleteRowFromTableDefault(int(int32(row))))))
}

func (ptr *QSqlTableModel) ConnectDeleteRowFromTable(f func(row int) bool) {
	defer qt.Recovering("connect QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::deleteRowFromTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectDeleteRowFromTable() {
	defer qt.Recovering("disconnect QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::deleteRowFromTable")
	}
}

func (ptr *QSqlTableModel) DeleteRowFromTable(row int) bool {
	defer qt.Recovering("QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DeleteRowFromTable(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) DeleteRowFromTableDefault(row int) bool {
	defer qt.Recovering("QSqlTableModel::deleteRowFromTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DeleteRowFromTableDefault(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	defer qt.Recovering("QSqlTableModel::editStrategy")

	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	defer qt.Recovering("QSqlTableModel::fieldIndex")

	if ptr.Pointer() != nil {
		var fieldNameC = C.CString(fieldName)
		defer C.free(unsafe.Pointer(fieldNameC))
		return int(int32(C.QSqlTableModel_FieldIndex(ptr.Pointer(), fieldNameC)))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	defer qt.Recovering("QSqlTableModel::filter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Filter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer qt.Recovering("QSqlTableModel::flags")

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlTableModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlTableModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer qt.Recovering("QSqlTableModel::headerData")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QSqlTableModel_HeaderData(ptr.Pointer(), C.int(int32(section)), C.longlong(orientation), C.int(int32(role))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_IndexInQuery
func callbackQSqlTableModel_IndexInQuery(ptr unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::indexInQuery")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::indexInQuery"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(item)))
	}

	return core.PointerFromQModelIndex(NewQSqlTableModelFromPointer(ptr).IndexInQueryDefault(core.NewQModelIndexFromPointer(item)))
}

func (ptr *QSqlTableModel) ConnectIndexInQuery(f func(item *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::indexInQuery", f)
	}
}

func (ptr *QSqlTableModel) DisconnectIndexInQuery() {
	defer qt.Recovering("disconnect QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::indexInQuery")
	}
}

func (ptr *QSqlTableModel) IndexInQuery(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_IndexInQuery(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) IndexInQueryDefault(item core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::indexInQuery")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_IndexInQueryDefault(ptr.Pointer(), core.PointerFromQModelIndex(item)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRecord(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record)) != 0
	}
	return false
}

//export callbackQSqlTableModel_InsertRowIntoTable
func callbackQSqlTableModel_InsertRowIntoTable(ptr unsafe.Pointer, values unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::insertRowIntoTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::insertRowIntoTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QSqlRecord) bool)(NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).InsertRowIntoTableDefault(NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlTableModel) ConnectInsertRowIntoTable(f func(values *QSqlRecord) bool) {
	defer qt.Recovering("connect QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::insertRowIntoTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectInsertRowIntoTable() {
	defer qt.Recovering("disconnect QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::insertRowIntoTable")
	}
}

func (ptr *QSqlTableModel) InsertRowIntoTable(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRowIntoTable(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRowIntoTableDefault(values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRowIntoTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRowIntoTableDefault(ptr.Pointer(), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::insertRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::isDirty")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

//export callbackQSqlTableModel_OrderByClause
func callbackQSqlTableModel_OrderByClause(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlTableModel::orderByClause")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::orderByClause"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSqlTableModelFromPointer(ptr).OrderByClauseDefault())
}

func (ptr *QSqlTableModel) ConnectOrderByClause(f func() string) {
	defer qt.Recovering("connect QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::orderByClause", f)
	}
}

func (ptr *QSqlTableModel) DisconnectOrderByClause() {
	defer qt.Recovering("disconnect QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::orderByClause")
	}
}

func (ptr *QSqlTableModel) OrderByClause() string {
	defer qt.Recovering("QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_OrderByClause(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) OrderByClauseDefault() string {
	defer qt.Recovering("QSqlTableModel::orderByClause")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_OrderByClauseDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) PrimaryKey() *QSqlIndex {
	defer qt.Recovering("QSqlTableModel::primaryKey")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlIndexFromPointer(C.QSqlTableModel_PrimaryKey(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlIndex).DestroyQSqlIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) PrimaryValues(row int) *QSqlRecord {
	defer qt.Recovering("QSqlTableModel::primaryValues")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlTableModel_PrimaryValues(ptr.Pointer(), C.int(int32(row))))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_PrimeInsert
func callbackQSqlTableModel_PrimeInsert(ptr unsafe.Pointer, row C.int, record unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::primeInsert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::primeInsert"); signal != nil {
		signal.(func(int, *QSqlRecord))(int(int32(row)), NewQSqlRecordFromPointer(record))
	}

}

func (ptr *QSqlTableModel) ConnectPrimeInsert(f func(row int, record *QSqlRecord)) {
	defer qt.Recovering("connect QSqlTableModel::primeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectPrimeInsert(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::primeInsert", f)
	}
}

func (ptr *QSqlTableModel) DisconnectPrimeInsert() {
	defer qt.Recovering("disconnect QSqlTableModel::primeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectPrimeInsert(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::primeInsert")
	}
}

func (ptr *QSqlTableModel) PrimeInsert(row int, record QSqlRecord_ITF) {
	defer qt.Recovering("QSqlTableModel::primeInsert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_PrimeInsert(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(record))
	}
}

func (ptr *QSqlTableModel) Record() *QSqlRecord {
	defer qt.Recovering("QSqlTableModel::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlTableModel_Record(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) Record2(row int) *QSqlRecord {
	defer qt.Recovering("QSqlTableModel::record")

	if ptr.Pointer() != nil {
		var tmpValue = NewQSqlRecordFromPointer(C.QSqlTableModel_Record2(ptr.Pointer(), C.int(int32(row))))
		runtime.SetFinalizer(tmpValue, (*QSqlRecord).DestroyQSqlRecord)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeColumns")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveColumns(ptr.Pointer(), C.int(int32(column)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::removeRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveRows(ptr.Pointer(), C.int(int32(row)), C.int(int32(count)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlTableModel_Revert
func callbackQSqlTableModel_Revert(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::revert")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::revert"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSqlTableModel) ConnectRevert(f func()) {
	defer qt.Recovering("connect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revert", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevert() {
	defer qt.Recovering("disconnect QSqlTableModel::revert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revert")
	}
}

func (ptr *QSqlTableModel) Revert() {
	defer qt.Recovering("QSqlTableModel::revert")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Revert(ptr.Pointer())
	}
}

//export callbackQSqlTableModel_RevertAll
func callbackQSqlTableModel_RevertAll(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::revertAll")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::revertAll"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSqlTableModel) ConnectRevertAll(f func()) {
	defer qt.Recovering("connect QSqlTableModel::revertAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revertAll", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevertAll() {
	defer qt.Recovering("disconnect QSqlTableModel::revertAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revertAll")
	}
}

func (ptr *QSqlTableModel) RevertAll() {
	defer qt.Recovering("QSqlTableModel::revertAll")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(ptr.Pointer())
	}
}

//export callbackQSqlTableModel_RevertRow
func callbackQSqlTableModel_RevertRow(ptr unsafe.Pointer, row C.int) {
	defer qt.Recovering("callback QSqlTableModel::revertRow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::revertRow"); signal != nil {
		signal.(func(int))(int(int32(row)))
	} else {
		NewQSqlTableModelFromPointer(ptr).RevertRowDefault(int(int32(row)))
	}
}

func (ptr *QSqlTableModel) ConnectRevertRow(f func(row int)) {
	defer qt.Recovering("connect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revertRow", f)
	}
}

func (ptr *QSqlTableModel) DisconnectRevertRow() {
	defer qt.Recovering("disconnect QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::revertRow")
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlTableModel) RevertRowDefault(row int) {
	defer qt.Recovering("QSqlTableModel::revertRow")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRowDefault(ptr.Pointer(), C.int(int32(row)))
	}
}

func (ptr *QSqlTableModel) RowCount(parent core.QModelIndex_ITF) int {
	defer qt.Recovering("QSqlTableModel::rowCount")

	if ptr.Pointer() != nil {
		return int(int32(C.QSqlTableModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent))))
	}
	return 0
}

//export callbackQSqlTableModel_Select
func callbackQSqlTableModel_Select(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::select")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::select"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).SelectDefault())))
}

func (ptr *QSqlTableModel) ConnectSelect(f func() bool) {
	defer qt.Recovering("connect QSqlTableModel::select")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::select", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSelect() {
	defer qt.Recovering("disconnect QSqlTableModel::select")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::select")
	}
}

func (ptr *QSqlTableModel) Select() bool {
	defer qt.Recovering("QSqlTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectDefault() bool {
	defer qt.Recovering("QSqlTableModel::select")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlTableModel_SelectRow
func callbackQSqlTableModel_SelectRow(ptr unsafe.Pointer, row C.int) C.char {
	defer qt.Recovering("callback QSqlTableModel::selectRow")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::selectRow"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(row))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).SelectRowDefault(int(int32(row))))))
}

func (ptr *QSqlTableModel) ConnectSelectRow(f func(row int) bool) {
	defer qt.Recovering("connect QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::selectRow", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSelectRow() {
	defer qt.Recovering("disconnect QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::selectRow")
	}
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	defer qt.Recovering("QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRow(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRowDefault(row int) bool {
	defer qt.Recovering("QSqlTableModel::selectRow")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRowDefault(ptr.Pointer(), C.int(int32(row))) != 0
	}
	return false
}

//export callbackQSqlTableModel_SelectStatement
func callbackQSqlTableModel_SelectStatement(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlTableModel::selectStatement")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::selectStatement"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSqlTableModelFromPointer(ptr).SelectStatementDefault())
}

func (ptr *QSqlTableModel) ConnectSelectStatement(f func() string) {
	defer qt.Recovering("connect QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::selectStatement", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSelectStatement() {
	defer qt.Recovering("disconnect QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::selectStatement")
	}
}

func (ptr *QSqlTableModel) SelectStatement() string {
	defer qt.Recovering("QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_SelectStatement(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) SelectStatementDefault() string {
	defer qt.Recovering("QSqlTableModel::selectStatement")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_SelectStatementDefault(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer qt.Recovering("QSqlTableModel::setData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(int32(role))) != 0
	}
	return false
}

//export callbackQSqlTableModel_SetEditStrategy
func callbackQSqlTableModel_SetEditStrategy(ptr unsafe.Pointer, strategy C.longlong) {
	defer qt.Recovering("callback QSqlTableModel::setEditStrategy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::setEditStrategy"); signal != nil {
		signal.(func(QSqlTableModel__EditStrategy))(QSqlTableModel__EditStrategy(strategy))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetEditStrategyDefault(QSqlTableModel__EditStrategy(strategy))
	}
}

func (ptr *QSqlTableModel) ConnectSetEditStrategy(f func(strategy QSqlTableModel__EditStrategy)) {
	defer qt.Recovering("connect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setEditStrategy", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetEditStrategy() {
	defer qt.Recovering("disconnect QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setEditStrategy")
	}
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(ptr.Pointer(), C.longlong(strategy))
	}
}

func (ptr *QSqlTableModel) SetEditStrategyDefault(strategy QSqlTableModel__EditStrategy) {
	defer qt.Recovering("QSqlTableModel::setEditStrategy")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategyDefault(ptr.Pointer(), C.longlong(strategy))
	}
}

//export callbackQSqlTableModel_SetFilter
func callbackQSqlTableModel_SetFilter(ptr unsafe.Pointer, filter *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::setFilter"); signal != nil {
		signal.(func(string))(C.GoString(filter))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetFilterDefault(C.GoString(filter))
	}
}

func (ptr *QSqlTableModel) ConnectSetFilter(f func(filter string)) {
	defer qt.Recovering("connect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setFilter", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetFilter() {
	defer qt.Recovering("disconnect QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setFilter")
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		C.QSqlTableModel_SetFilter(ptr.Pointer(), filterC)
	}
}

func (ptr *QSqlTableModel) SetFilterDefault(filter string) {
	defer qt.Recovering("QSqlTableModel::setFilter")

	if ptr.Pointer() != nil {
		var filterC = C.CString(filter)
		defer C.free(unsafe.Pointer(filterC))
		C.QSqlTableModel_SetFilterDefault(ptr.Pointer(), filterC)
	}
}

func (ptr *QSqlTableModel) SetPrimaryKey(key QSqlIndex_ITF) {
	defer qt.Recovering("QSqlTableModel::setPrimaryKey")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetPrimaryKey(ptr.Pointer(), PointerFromQSqlIndex(key))
	}
}

func (ptr *QSqlTableModel) SetQuery(query QSqlQuery_ITF) {
	defer qt.Recovering("QSqlTableModel::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::setRecord")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetRecord(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

//export callbackQSqlTableModel_SetSort
func callbackQSqlTableModel_SetSort(ptr unsafe.Pointer, column C.int, order C.longlong) {
	defer qt.Recovering("callback QSqlTableModel::setSort")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::setSort"); signal != nil {
		signal.(func(int, core.Qt__SortOrder))(int(int32(column)), core.Qt__SortOrder(order))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetSortDefault(int(int32(column)), core.Qt__SortOrder(order))
	}
}

func (ptr *QSqlTableModel) ConnectSetSort(f func(column int, order core.Qt__SortOrder)) {
	defer qt.Recovering("connect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setSort", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetSort() {
	defer qt.Recovering("disconnect QSqlTableModel::setSort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setSort")
	}
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

func (ptr *QSqlTableModel) SetSortDefault(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::setSort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSortDefault(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlTableModel_SetTable
func callbackQSqlTableModel_SetTable(ptr unsafe.Pointer, tableName *C.char) {
	defer qt.Recovering("callback QSqlTableModel::setTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::setTable"); signal != nil {
		signal.(func(string))(C.GoString(tableName))
	} else {
		NewQSqlTableModelFromPointer(ptr).SetTableDefault(C.GoString(tableName))
	}
}

func (ptr *QSqlTableModel) ConnectSetTable(f func(tableName string)) {
	defer qt.Recovering("connect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSetTable() {
	defer qt.Recovering("disconnect QSqlTableModel::setTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::setTable")
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		C.QSqlTableModel_SetTable(ptr.Pointer(), tableNameC)
	}
}

func (ptr *QSqlTableModel) SetTableDefault(tableName string) {
	defer qt.Recovering("QSqlTableModel::setTable")

	if ptr.Pointer() != nil {
		var tableNameC = C.CString(tableName)
		defer C.free(unsafe.Pointer(tableNameC))
		C.QSqlTableModel_SetTableDefault(ptr.Pointer(), tableNameC)
	}
}

func (ptr *QSqlTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer qt.Recovering("QSqlTableModel::sort")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Sort(ptr.Pointer(), C.int(int32(column)), C.longlong(order))
	}
}

//export callbackQSqlTableModel_Submit
func callbackQSqlTableModel_Submit(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::submit")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::submit"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlTableModel) ConnectSubmit(f func() bool) {
	defer qt.Recovering("connect QSqlTableModel::submit")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::submit", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSubmit() {
	defer qt.Recovering("disconnect QSqlTableModel::submit")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::submit")
	}
}

func (ptr *QSqlTableModel) Submit() bool {
	defer qt.Recovering("QSqlTableModel::submit")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSqlTableModel_SubmitAll
func callbackQSqlTableModel_SubmitAll(ptr unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::submitAll")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::submitAll"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QSqlTableModel) ConnectSubmitAll(f func() bool) {
	defer qt.Recovering("connect QSqlTableModel::submitAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::submitAll", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSubmitAll() {
	defer qt.Recovering("disconnect QSqlTableModel::submitAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::submitAll")
	}
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	defer qt.Recovering("QSqlTableModel::submitAll")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SubmitAll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) TableName() string {
	defer qt.Recovering("QSqlTableModel::tableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_TableName(ptr.Pointer()))
	}
	return ""
}

//export callbackQSqlTableModel_UpdateRowInTable
func callbackQSqlTableModel_UpdateRowInTable(ptr unsafe.Pointer, row C.int, values unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::updateRowInTable")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::updateRowInTable"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *QSqlRecord) bool)(int(int32(row)), NewQSqlRecordFromPointer(values)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).UpdateRowInTableDefault(int(int32(row)), NewQSqlRecordFromPointer(values)))))
}

func (ptr *QSqlTableModel) ConnectUpdateRowInTable(f func(row int, values *QSqlRecord) bool) {
	defer qt.Recovering("connect QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::updateRowInTable", f)
	}
}

func (ptr *QSqlTableModel) DisconnectUpdateRowInTable() {
	defer qt.Recovering("disconnect QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::updateRowInTable")
	}
}

func (ptr *QSqlTableModel) UpdateRowInTable(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_UpdateRowInTable(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) UpdateRowInTableDefault(row int, values QSqlRecord_ITF) bool {
	defer qt.Recovering("QSqlTableModel::updateRowInTable")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_UpdateRowInTableDefault(ptr.Pointer(), C.int(int32(row)), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

//export callbackQSqlTableModel_DestroyQSqlTableModel
func callbackQSqlTableModel_DestroyQSqlTableModel(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::~QSqlTableModel")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::~QSqlTableModel"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).DestroyQSqlTableModelDefault()
	}
}

func (ptr *QSqlTableModel) ConnectDestroyQSqlTableModel(f func()) {
	defer qt.Recovering("connect QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::~QSqlTableModel", f)
	}
}

func (ptr *QSqlTableModel) DisconnectDestroyQSqlTableModel() {
	defer qt.Recovering("disconnect QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::~QSqlTableModel")
	}
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	defer qt.Recovering("QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlTableModel) DestroyQSqlTableModelDefault() {
	defer qt.Recovering("QSqlTableModel::~QSqlTableModel")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModelDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlTableModel_QueryChange
func callbackQSqlTableModel_QueryChange(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::queryChange")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::queryChange"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).QueryChangeDefault()
	}
}

func (ptr *QSqlTableModel) ConnectQueryChange(f func()) {
	defer qt.Recovering("connect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::queryChange", f)
	}
}

func (ptr *QSqlTableModel) DisconnectQueryChange() {
	defer qt.Recovering("disconnect QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::queryChange")
	}
}

func (ptr *QSqlTableModel) QueryChange() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChange(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) QueryChangeDefault() {
	defer qt.Recovering("QSqlTableModel::queryChange")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_QueryChangeDefault(ptr.Pointer())
	}
}

//export callbackQSqlTableModel_Index
func callbackQSqlTableModel_Index(ptr unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::index")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::index"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
	}

	return core.PointerFromQModelIndex(NewQSqlTableModelFromPointer(ptr).IndexDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))
}

func (ptr *QSqlTableModel) ConnectIndex(f func(row int, column int, parent *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlTableModel::index")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::index", f)
	}
}

func (ptr *QSqlTableModel) DisconnectIndex() {
	defer qt.Recovering("disconnect QSqlTableModel::index")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::index")
	}
}

func (ptr *QSqlTableModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_Index(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) IndexDefault(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::index")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_IndexDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_DropMimeData
func callbackQSqlTableModel_DropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::dropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::dropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).DropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlTableModel) ConnectDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlTableModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::dropMimeData", f)
	}
}

func (ptr *QSqlTableModel) DisconnectDropMimeData() {
	defer qt.Recovering("disconnect QSqlTableModel::dropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::dropMimeData")
	}
}

func (ptr *QSqlTableModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) DropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::dropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_DropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlTableModel_Sibling
func callbackQSqlTableModel_Sibling(ptr unsafe.Pointer, row C.int, column C.int, idx unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::sibling")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::sibling"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(int, int, *core.QModelIndex) *core.QModelIndex)(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
	}

	return core.PointerFromQModelIndex(NewQSqlTableModelFromPointer(ptr).SiblingDefault(int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(idx)))
}

func (ptr *QSqlTableModel) ConnectSibling(f func(row int, column int, idx *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlTableModel::sibling")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::sibling", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSibling() {
	defer qt.Recovering("disconnect QSqlTableModel::sibling")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::sibling")
	}
}

func (ptr *QSqlTableModel) Sibling(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_Sibling(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) SiblingDefault(row int, column int, idx core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::sibling")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_SiblingDefault(ptr.Pointer(), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(idx)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_Buddy
func callbackQSqlTableModel_Buddy(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::buddy")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::buddy"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlTableModelFromPointer(ptr).BuddyDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlTableModel) ConnectBuddy(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlTableModel::buddy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::buddy", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBuddy() {
	defer qt.Recovering("disconnect QSqlTableModel::buddy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::buddy")
	}
}

func (ptr *QSqlTableModel) Buddy(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_Buddy(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) BuddyDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::buddy")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_BuddyDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_CanDropMimeData
func callbackQSqlTableModel_CanDropMimeData(ptr unsafe.Pointer, data unsafe.Pointer, action C.longlong, row C.int, column C.int, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::canDropMimeData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::canDropMimeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QMimeData, core.Qt__DropAction, int, int, *core.QModelIndex) bool)(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).CanDropMimeDataDefault(core.NewQMimeDataFromPointer(data), core.Qt__DropAction(action), int(int32(row)), int(int32(column)), core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlTableModel) ConnectCanDropMimeData(f func(data *core.QMimeData, action core.Qt__DropAction, row int, column int, parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlTableModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::canDropMimeData", f)
	}
}

func (ptr *QSqlTableModel) DisconnectCanDropMimeData() {
	defer qt.Recovering("disconnect QSqlTableModel::canDropMimeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::canDropMimeData")
	}
}

func (ptr *QSqlTableModel) CanDropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_CanDropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) CanDropMimeDataDefault(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::canDropMimeData")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_CanDropMimeDataDefault(ptr.Pointer(), core.PointerFromQMimeData(data), C.longlong(action), C.int(int32(row)), C.int(int32(column)), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlTableModel_HasChildren
func callbackQSqlTableModel_HasChildren(ptr unsafe.Pointer, parent unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::hasChildren")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::hasChildren"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex) bool)(core.NewQModelIndexFromPointer(parent)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).HasChildrenDefault(core.NewQModelIndexFromPointer(parent)))))
}

func (ptr *QSqlTableModel) ConnectHasChildren(f func(parent *core.QModelIndex) bool) {
	defer qt.Recovering("connect QSqlTableModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::hasChildren", f)
	}
}

func (ptr *QSqlTableModel) DisconnectHasChildren() {
	defer qt.Recovering("disconnect QSqlTableModel::hasChildren")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::hasChildren")
	}
}

func (ptr *QSqlTableModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) HasChildrenDefault(parent core.QModelIndex_ITF) bool {
	defer qt.Recovering("QSqlTableModel::hasChildren")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_HasChildrenDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

//export callbackQSqlTableModel_MimeTypes
func callbackQSqlTableModel_MimeTypes(ptr unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QSqlTableModel::mimeTypes")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::mimeTypes"); signal != nil {
		return C.CString(strings.Join(signal.(func() []string)(), "|"))
	}

	return C.CString(strings.Join(NewQSqlTableModelFromPointer(ptr).MimeTypesDefault(), "|"))
}

func (ptr *QSqlTableModel) ConnectMimeTypes(f func() []string) {
	defer qt.Recovering("connect QSqlTableModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::mimeTypes", f)
	}
}

func (ptr *QSqlTableModel) DisconnectMimeTypes() {
	defer qt.Recovering("disconnect QSqlTableModel::mimeTypes")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::mimeTypes")
	}
}

func (ptr *QSqlTableModel) MimeTypes() []string {
	defer qt.Recovering("QSqlTableModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlTableModel_MimeTypes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QSqlTableModel) MimeTypesDefault() []string {
	defer qt.Recovering("QSqlTableModel::mimeTypes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSqlTableModel_MimeTypesDefault(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

//export callbackQSqlTableModel_MoveColumns
func callbackQSqlTableModel_MoveColumns(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceColumn C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlTableModel::moveColumns")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::moveColumns"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).MoveColumnsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceColumn)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlTableModel) ConnectMoveColumns(f func(sourceParent *core.QModelIndex, sourceColumn int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlTableModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::moveColumns", f)
	}
}

func (ptr *QSqlTableModel) DisconnectMoveColumns() {
	defer qt.Recovering("disconnect QSqlTableModel::moveColumns")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::moveColumns")
	}
}

func (ptr *QSqlTableModel) MoveColumns(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlTableModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_MoveColumns(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) MoveColumnsDefault(sourceParent core.QModelIndex_ITF, sourceColumn int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlTableModel::moveColumns")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_MoveColumnsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceColumn)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlTableModel_MoveRows
func callbackQSqlTableModel_MoveRows(ptr unsafe.Pointer, sourceParent unsafe.Pointer, sourceRow C.int, count C.int, destinationParent unsafe.Pointer, destinationChild C.int) C.char {
	defer qt.Recovering("callback QSqlTableModel::moveRows")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::moveRows"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QModelIndex, int, int, *core.QModelIndex, int) bool)(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).MoveRowsDefault(core.NewQModelIndexFromPointer(sourceParent), int(int32(sourceRow)), int(int32(count)), core.NewQModelIndexFromPointer(destinationParent), int(int32(destinationChild))))))
}

func (ptr *QSqlTableModel) ConnectMoveRows(f func(sourceParent *core.QModelIndex, sourceRow int, count int, destinationParent *core.QModelIndex, destinationChild int) bool) {
	defer qt.Recovering("connect QSqlTableModel::moveRows")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::moveRows", f)
	}
}

func (ptr *QSqlTableModel) DisconnectMoveRows() {
	defer qt.Recovering("disconnect QSqlTableModel::moveRows")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::moveRows")
	}
}

func (ptr *QSqlTableModel) MoveRows(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlTableModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_MoveRows(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) MoveRowsDefault(sourceParent core.QModelIndex_ITF, sourceRow int, count int, destinationParent core.QModelIndex_ITF, destinationChild int) bool {
	defer qt.Recovering("QSqlTableModel::moveRows")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_MoveRowsDefault(ptr.Pointer(), core.PointerFromQModelIndex(sourceParent), C.int(int32(sourceRow)), C.int(int32(count)), core.PointerFromQModelIndex(destinationParent), C.int(int32(destinationChild))) != 0
	}
	return false
}

//export callbackQSqlTableModel_Parent
func callbackQSqlTableModel_Parent(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::parent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::parent"); signal != nil {
		return core.PointerFromQModelIndex(signal.(func(*core.QModelIndex) *core.QModelIndex)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQModelIndex(NewQSqlTableModelFromPointer(ptr).ParentDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlTableModel) ConnectParent(f func(index *core.QModelIndex) *core.QModelIndex) {
	defer qt.Recovering("connect QSqlTableModel::parent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::parent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectParent() {
	defer qt.Recovering("disconnect QSqlTableModel::parent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::parent")
	}
}

func (ptr *QSqlTableModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) ParentDefault(index core.QModelIndex_ITF) *core.QModelIndex {
	defer qt.Recovering("QSqlTableModel::parent")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQModelIndexFromPointer(C.QSqlTableModel_ParentDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QModelIndex).DestroyQModelIndex)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_ResetInternalData
func callbackQSqlTableModel_ResetInternalData(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::resetInternalData")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::resetInternalData"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).ResetInternalDataDefault()
	}
}

func (ptr *QSqlTableModel) ConnectResetInternalData(f func()) {
	defer qt.Recovering("connect QSqlTableModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::resetInternalData", f)
	}
}

func (ptr *QSqlTableModel) DisconnectResetInternalData() {
	defer qt.Recovering("disconnect QSqlTableModel::resetInternalData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::resetInternalData")
	}
}

func (ptr *QSqlTableModel) ResetInternalData() {
	defer qt.Recovering("QSqlTableModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ResetInternalData(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) ResetInternalDataDefault() {
	defer qt.Recovering("QSqlTableModel::resetInternalData")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ResetInternalDataDefault(ptr.Pointer())
	}
}

//export callbackQSqlTableModel_Span
func callbackQSqlTableModel_Span(ptr unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::span")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::span"); signal != nil {
		return core.PointerFromQSize(signal.(func(*core.QModelIndex) *core.QSize)(core.NewQModelIndexFromPointer(index)))
	}

	return core.PointerFromQSize(NewQSqlTableModelFromPointer(ptr).SpanDefault(core.NewQModelIndexFromPointer(index)))
}

func (ptr *QSqlTableModel) ConnectSpan(f func(index *core.QModelIndex) *core.QSize) {
	defer qt.Recovering("connect QSqlTableModel::span")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::span", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSpan() {
	defer qt.Recovering("disconnect QSqlTableModel::span")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::span")
	}
}

func (ptr *QSqlTableModel) Span(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlTableModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlTableModel_Span(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

func (ptr *QSqlTableModel) SpanDefault(index core.QModelIndex_ITF) *core.QSize {
	defer qt.Recovering("QSqlTableModel::span")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQSizeFromPointer(C.QSqlTableModel_SpanDefault(ptr.Pointer(), core.PointerFromQModelIndex(index)))
		runtime.SetFinalizer(tmpValue, (*core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQSqlTableModel_SupportedDragActions
func callbackQSqlTableModel_SupportedDragActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlTableModel::supportedDragActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::supportedDragActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlTableModelFromPointer(ptr).SupportedDragActionsDefault())
}

func (ptr *QSqlTableModel) ConnectSupportedDragActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlTableModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::supportedDragActions", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSupportedDragActions() {
	defer qt.Recovering("disconnect QSqlTableModel::supportedDragActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::supportedDragActions")
	}
}

func (ptr *QSqlTableModel) SupportedDragActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlTableModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlTableModel_SupportedDragActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) SupportedDragActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlTableModel::supportedDragActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlTableModel_SupportedDragActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlTableModel_SupportedDropActions
func callbackQSqlTableModel_SupportedDropActions(ptr unsafe.Pointer) C.longlong {
	defer qt.Recovering("callback QSqlTableModel::supportedDropActions")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::supportedDropActions"); signal != nil {
		return C.longlong(signal.(func() core.Qt__DropAction)())
	}

	return C.longlong(NewQSqlTableModelFromPointer(ptr).SupportedDropActionsDefault())
}

func (ptr *QSqlTableModel) ConnectSupportedDropActions(f func() core.Qt__DropAction) {
	defer qt.Recovering("connect QSqlTableModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::supportedDropActions", f)
	}
}

func (ptr *QSqlTableModel) DisconnectSupportedDropActions() {
	defer qt.Recovering("disconnect QSqlTableModel::supportedDropActions")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::supportedDropActions")
	}
}

func (ptr *QSqlTableModel) SupportedDropActions() core.Qt__DropAction {
	defer qt.Recovering("QSqlTableModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlTableModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) SupportedDropActionsDefault() core.Qt__DropAction {
	defer qt.Recovering("QSqlTableModel::supportedDropActions")

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QSqlTableModel_SupportedDropActionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQSqlTableModel_TimerEvent
func callbackQSqlTableModel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::timerEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::timerEvent")
	}
}

func (ptr *QSqlTableModel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSqlTableModel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::timerEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQSqlTableModel_ChildEvent
func callbackQSqlTableModel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::childEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::childEvent")
	}
}

func (ptr *QSqlTableModel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSqlTableModel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::childEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQSqlTableModel_ConnectNotify
func callbackQSqlTableModel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlTableModelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlTableModel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlTableModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::connectNotify", f)
	}
}

func (ptr *QSqlTableModel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSqlTableModel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::connectNotify")
	}
}

func (ptr *QSqlTableModel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlTableModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlTableModel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlTableModel::connectNotify")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlTableModel_CustomEvent
func callbackQSqlTableModel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSqlTableModelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSqlTableModel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::customEvent", f)
	}
}

func (ptr *QSqlTableModel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::customEvent")
	}
}

func (ptr *QSqlTableModel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSqlTableModel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSqlTableModel::customEvent")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQSqlTableModel_DeleteLater
func callbackQSqlTableModel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSqlTableModelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSqlTableModel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSqlTableModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::deleteLater", f)
	}
}

func (ptr *QSqlTableModel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSqlTableModel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::deleteLater")
	}
}

func (ptr *QSqlTableModel) DeleteLater() {
	defer qt.Recovering("QSqlTableModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QSqlTableModel) DeleteLaterDefault() {
	defer qt.Recovering("QSqlTableModel::deleteLater")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQSqlTableModel_DisconnectNotify
func callbackQSqlTableModel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSqlTableModel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSqlTableModelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSqlTableModel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSqlTableModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::disconnectNotify", f)
	}
}

func (ptr *QSqlTableModel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSqlTableModel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::disconnectNotify")
	}
}

func (ptr *QSqlTableModel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlTableModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSqlTableModel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSqlTableModel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSqlTableModel_Event
func callbackQSqlTableModel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QSqlTableModel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlTableModel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::event", f)
	}
}

func (ptr *QSqlTableModel) DisconnectEvent() {
	defer qt.Recovering("disconnect QSqlTableModel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::event")
	}
}

func (ptr *QSqlTableModel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlTableModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlTableModel::event")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSqlTableModel_EventFilter
func callbackQSqlTableModel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QSqlTableModel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQSqlTableModelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QSqlTableModel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSqlTableModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::eventFilter", f)
	}
}

func (ptr *QSqlTableModel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSqlTableModel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::eventFilter")
	}
}

func (ptr *QSqlTableModel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlTableModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSqlTableModel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSqlTableModel_MetaObject
func callbackQSqlTableModel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QSqlTableModel::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QSqlTableModel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSqlTableModelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSqlTableModel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSqlTableModel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::metaObject", f)
	}
}

func (ptr *QSqlTableModel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSqlTableModel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QSqlTableModel::metaObject")
	}
}

func (ptr *QSqlTableModel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSqlTableModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlTableModel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlTableModel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSqlTableModel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSqlTableModel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
