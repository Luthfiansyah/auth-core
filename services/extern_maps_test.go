//
//  extern_maps_test.go
//  Auth Core
//
//  Copyright © 2017 Ice House. All rights reserved.
//

package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	err := InitializeMaps("../map_data.json")
	assert.Nil(t, err)

	bcMap := MustGetMap(M_BANKCODES)
	assert.Equal(t, "014", bcMap["BANK BCA"])

	assert.Panics(t, func() {
		MustGetMap("dummy non existent map")
	}, "getting a non existing map should panic")

	assert.Equal(t, "014", MustGetString(M_BANKCODES, "BANK BCA"))

	assert.Panics(t, func() {
		MustGetString(M_BANKCODES, "Pierre's bank")
	}, "getting a non existing key should panic")

	f := LookupFunc(M_BANKCODES)
	assert.Equal(t, "014", f("BANK BCA"))

	assert.Panics(t, func() {
		f("Pierre's bank")
	}, "bad key should make the function panic")

	assert.Panics(t, func() {
		LookupFunc("my dummy map name")
	}, "bad map name should make the function panic")
}
