package godb

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_MySQL_FormatDSN(t *testing.T) {
	c := MySQLConfig{
		Net:  "tcp",
		Addr: "db.com",
	}

	assert.Equal(t, "tcp(db.com)/", c.FormatDSN())
}