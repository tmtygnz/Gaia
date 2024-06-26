//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Defects struct {
	ID                int64 `sql:"primary_key"`
	Placeid           *int64
	Product           *int64
	EncodedBy         *string
	Date              *time.Time
	DefectDescription *string
	Quantity          *int64
	DefPrebTp         *int64
	IsDev             *bool
	Packaging         *int64
	AppVersion        *float64
	DateEncoded       *time.Time
	Batch             *string
}
