// Code generated by ent, DO NOT EDIT.

package placetable

import (
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldID, id))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldNumber, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldName, v))
}

// Capacity applies equality check predicate on the "capacity" field. It's identical to CapacityEQ.
func Capacity(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldCapacity, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldDeletedAt, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsDeleted, v))
}

// QrCode applies equality check predicate on the "qr_code" field. It's identical to QrCodeEQ.
func QrCode(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldQrCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldDescription, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldStatus, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsActive, v))
}

// IsReserved applies equality check predicate on the "is_reserved" field. It's identical to IsReservedEQ.
func IsReserved(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsReserved, v))
}

// IsVip applies equality check predicate on the "is_vip" field. It's identical to IsVipEQ.
func IsVip(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsVip, v))
}

// IsPremium applies equality check predicate on the "is_premium" field. It's identical to IsPremiumEQ.
func IsPremium(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsPremium, v))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldNumber, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldName, v))
}

// CapacityEQ applies the EQ predicate on the "capacity" field.
func CapacityEQ(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldCapacity, v))
}

// CapacityNEQ applies the NEQ predicate on the "capacity" field.
func CapacityNEQ(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldCapacity, v))
}

// CapacityIn applies the In predicate on the "capacity" field.
func CapacityIn(vs ...int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldCapacity, vs...))
}

// CapacityNotIn applies the NotIn predicate on the "capacity" field.
func CapacityNotIn(vs ...int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldCapacity, vs...))
}

// CapacityGT applies the GT predicate on the "capacity" field.
func CapacityGT(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldCapacity, v))
}

// CapacityGTE applies the GTE predicate on the "capacity" field.
func CapacityGTE(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldCapacity, v))
}

// CapacityLT applies the LT predicate on the "capacity" field.
func CapacityLT(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldCapacity, v))
}

// CapacityLTE applies the LTE predicate on the "capacity" field.
func CapacityLTE(v int) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldCapacity, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtContains applies the Contains predicate on the "deleted_at" field.
func DeletedAtContains(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContains(FieldDeletedAt, v))
}

// DeletedAtHasPrefix applies the HasPrefix predicate on the "deleted_at" field.
func DeletedAtHasPrefix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasPrefix(FieldDeletedAt, v))
}

// DeletedAtHasSuffix applies the HasSuffix predicate on the "deleted_at" field.
func DeletedAtHasSuffix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasSuffix(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedAtEqualFold applies the EqualFold predicate on the "deleted_at" field.
func DeletedAtEqualFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldDeletedAt, v))
}

// DeletedAtContainsFold applies the ContainsFold predicate on the "deleted_at" field.
func DeletedAtContainsFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldDeletedAt, v))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldIsDeleted, v))
}

// QrCodeEQ applies the EQ predicate on the "qr_code" field.
func QrCodeEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldQrCode, v))
}

// QrCodeNEQ applies the NEQ predicate on the "qr_code" field.
func QrCodeNEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldQrCode, v))
}

// QrCodeIn applies the In predicate on the "qr_code" field.
func QrCodeIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldQrCode, vs...))
}

// QrCodeNotIn applies the NotIn predicate on the "qr_code" field.
func QrCodeNotIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldQrCode, vs...))
}

// QrCodeGT applies the GT predicate on the "qr_code" field.
func QrCodeGT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldQrCode, v))
}

// QrCodeGTE applies the GTE predicate on the "qr_code" field.
func QrCodeGTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldQrCode, v))
}

// QrCodeLT applies the LT predicate on the "qr_code" field.
func QrCodeLT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldQrCode, v))
}

// QrCodeLTE applies the LTE predicate on the "qr_code" field.
func QrCodeLTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldQrCode, v))
}

// QrCodeContains applies the Contains predicate on the "qr_code" field.
func QrCodeContains(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContains(FieldQrCode, v))
}

// QrCodeHasPrefix applies the HasPrefix predicate on the "qr_code" field.
func QrCodeHasPrefix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasPrefix(FieldQrCode, v))
}

// QrCodeHasSuffix applies the HasSuffix predicate on the "qr_code" field.
func QrCodeHasSuffix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasSuffix(FieldQrCode, v))
}

// QrCodeIsNil applies the IsNil predicate on the "qr_code" field.
func QrCodeIsNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIsNull(FieldQrCode))
}

// QrCodeNotNil applies the NotNil predicate on the "qr_code" field.
func QrCodeNotNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotNull(FieldQrCode))
}

// QrCodeEqualFold applies the EqualFold predicate on the "qr_code" field.
func QrCodeEqualFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldQrCode, v))
}

// QrCodeContainsFold applies the ContainsFold predicate on the "qr_code" field.
func QrCodeContainsFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldQrCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldContainsFold(FieldStatus, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotIn(FieldType, vs...))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNotNull(FieldType))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldIsActive, v))
}

// IsReservedEQ applies the EQ predicate on the "is_reserved" field.
func IsReservedEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsReserved, v))
}

// IsReservedNEQ applies the NEQ predicate on the "is_reserved" field.
func IsReservedNEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldIsReserved, v))
}

// IsVipEQ applies the EQ predicate on the "is_vip" field.
func IsVipEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsVip, v))
}

// IsVipNEQ applies the NEQ predicate on the "is_vip" field.
func IsVipNEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldIsVip, v))
}

// IsPremiumEQ applies the EQ predicate on the "is_premium" field.
func IsPremiumEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldEQ(FieldIsPremium, v))
}

// IsPremiumNEQ applies the NEQ predicate on the "is_premium" field.
func IsPremiumNEQ(v bool) predicate.PlaceTable {
	return predicate.PlaceTable(sql.FieldNEQ(FieldIsPremium, v))
}

// HasPlace applies the HasEdge predicate on the "place" edge.
func HasPlace() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlaceTable, PlaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaceWith applies the HasEdge predicate on the "place" edge with a given conditions (other predicates).
func HasPlaceWith(preds ...predicate.Place) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newPlaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "created_by" edge.
func HasCreatedBy() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "created_by" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.User) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newCreatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpdatedBy applies the HasEdge predicate on the "updated_by" edge.
func HasUpdatedBy() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpdatedByTable, UpdatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpdatedByWith applies the HasEdge predicate on the "updated_by" edge with a given conditions (other predicates).
func HasUpdatedByWith(preds ...predicate.User) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newUpdatedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDeletedBy applies the HasEdge predicate on the "deleted_by" edge.
func HasDeletedBy() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeletedByTable, DeletedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeletedByWith applies the HasEdge predicate on the "deleted_by" edge with a given conditions (other predicates).
func HasDeletedByWith(preds ...predicate.User) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newDeletedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReservedBy applies the HasEdge predicate on the "reserved_by" edge.
func HasReservedBy() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReservedByTable, ReservedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReservedByWith applies the HasEdge predicate on the "reserved_by" edge with a given conditions (other predicates).
func HasReservedByWith(preds ...predicate.User) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newReservedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWaiter applies the HasEdge predicate on the "waiter" edge.
func HasWaiter() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WaiterTable, WaiterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWaiterWith applies the HasEdge predicate on the "waiter" edge with a given conditions (other predicates).
func HasWaiterWith(preds ...predicate.User) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newWaiterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrdersTable, OrdersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.PlaceTable {
	return predicate.PlaceTable(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlaceTable) predicate.PlaceTable {
	return predicate.PlaceTable(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlaceTable) predicate.PlaceTable {
	return predicate.PlaceTable(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlaceTable) predicate.PlaceTable {
	return predicate.PlaceTable(sql.NotPredicates(p))
}
