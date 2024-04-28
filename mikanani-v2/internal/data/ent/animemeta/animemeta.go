// Code generated by ent, DO NOT EDIT.

package animemeta

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the animemeta type in the database.
	Label = "anime_meta"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDownloadBitmap holds the string denoting the downloadbitmap field in the database.
	FieldDownloadBitmap = "download_bitmap"
	// FieldIsActive holds the string denoting the isactive field in the database.
	FieldIsActive = "is_active"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the animemeta in the database.
	Table = "anime_meta"
)

// Columns holds all SQL columns for animemeta fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldName,
	FieldDownloadBitmap,
	FieldIsActive,
	FieldTags,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDownloadBitmap holds the default value on creation for the "downloadBitmap" field.
	DefaultDownloadBitmap int64
	// DefaultIsActive holds the default value on creation for the "isActive" field.
	DefaultIsActive bool
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the AnimeMeta queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUID orders the results by the uid field.
func ByUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDownloadBitmap orders the results by the downloadBitmap field.
func ByDownloadBitmap(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadBitmap, opts...).ToFunc()
}

// ByIsActive orders the results by the isActive field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByCreateTime orders the results by the createTime field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the updateTime field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
