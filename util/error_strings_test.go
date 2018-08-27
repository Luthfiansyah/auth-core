//  error_strings_test.go
//  Auth Core
//
//  Copyright © 2017 Ice House. All rights reserved.
//

package util

import (
	"fmt"
	"testing"

	"github.com/auth-core/services"
	"github.com/stretchr/testify/assert"
)

func TestErrorMaps(t *testing.T) {

	err := services.InitializeMaps("../map_data.json")
	if err != nil {
		fmt.Printf("error reading map_data.json: %s", err.Error())
	}

	s := GetErrorString(EN, -1)
	assert.Equal(t, s, getErrorMessageEn("-1"))

	s = GetErrorString(ID, 0)
	assert.Equal(t, s, getErrorMessageId("0"))

	s = GetErrorString(100000, -1)
	assert.Equal(t, s, getErrorMessageEn("-1"))

	s = GetErrorString(EN, 100000)
	assert.Equal(t, s, getErrorMessageEn(string(defaultErrorId)))

	s = GetErrorString(ID, 10000)
	assert.Equal(t, s, getErrorMessageId(string(defaultErrorId)))

	s = GetErrorString(10000, 100000)
	assert.Equal(t, s, getErrorMessageEn(string(defaultErrorId)))

}

func TestSuccessMessage(t *testing.T) {
	err := services.InitializeMaps("../map_data.json")
	assert.Nil(t, err)

	s := getSuccessMessageId("wallet_to_cash_success")
	assert.NotNil(t, s)
	println(s)

	d := getSuccessMessageEn("Random form name should return default success message")
	assert.NotNil(t, d)
	println(d)

}
