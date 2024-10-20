package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/acd19ml/EventCOM_MySQL/mcube/http/label"
	"github.com/acd19ml/EventCOM_MySQL/mcube/pb/http"
)

func TestEntry(t *testing.T) {
	should := assert.New(t)

	e := http.NewEntry("/mcube/v1/", "GET", "Monkey")
	e.EnableAuth()
	e.EnablePermission()
	e.AddLabel(label.Get)

	should.Equal("Monkey", e.Resource)

	set := http.NewEntrySet()
	set.AddEntry(*e, *e)
	should.Equal(2, len(set.Items))
}
