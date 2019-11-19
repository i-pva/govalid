package govalid_test

import (
	"testing"
	"time"

	"github.com/twharmon/gosql"
	"github.com/twharmon/govalid"
)

func init() {
	govalid.Register(tm{}, tmMin{}, tmNullMax{})
}

type tm struct {
	T time.Time
}

type tmMin struct {
	T time.Time `govalid:"min:0"`
}

type tmNullMax struct {
	T gosql.NullTime `govalid:"req|max:3600"`
}

func TestTime(t *testing.T) {
	now := time.Now()
	assertNoViolation(t, "no validation rules with empty field", &tm{})
	assertNoViolation(t, "no validation rules with non-empty field", &tm{now})

	assertNoViolation(t, "`min` with empty field", &tmMin{})
	assertNoViolation(t, "`min` with valid field", &tmMin{now})
	assertViolation(t, "`min` with invalid field", &tmMin{now.Add(time.Hour * 10)})

	assertViolation(t, "`req|max` with empty struct field", &tmNullMax{})
	assertViolation(t, "`req|max` with invalid struct field", &tmNullMax{gosql.NullTime{Valid: true, Time: now.AddDate(0, 0, -1)}})
}
