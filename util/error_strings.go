//
//  error_strings.go
//  Auth Core
//
//  Copyright © 2017 Ice House. All rights reserved.
//

package util

import (
	"github.com/auth-core/services"

	"strconv"
)

// This file contains error strings suited for user display
// The erro IDs are based on the exisitng dompetku error codes

const (
	DEFAULT_LANGUAGE = iota
	EN               = iota
	ID               = iota
)

var defaultErrorId = -1000

var defaultSuccessMessageId = "default_success_message"

var getErrorMessageId = services.LookupFuncWithDefaultKey(services.M_ERROR_CODES_ID, strconv.Itoa(defaultErrorId))

var getErrorMessageEn = services.LookupFuncWithDefaultKey(services.M_ERROR_CODES_EN, strconv.Itoa(defaultErrorId))

var getSuccessMessageId = services.LookupFuncWithDefaultKey(services.M_SUCCESS_CODES_ID, defaultSuccessMessageId)

var getSuccessMessageEn = services.LookupFuncWithDefaultKey(services.M_SUCCESS_CODES_EN, defaultSuccessMessageId)

func GetErrorString(language int, id int) string {

	if language == ID {
		return getErrorMessageId(strconv.Itoa(id))
	}

	if language == EN {
		return getErrorMessageEn(strconv.Itoa(id))
	}

	return getErrorMessageEn(strconv.Itoa(id))
}

func GetSuccessString(language int, formName string) string {

	if language == ID {
		return getSuccessMessageId(formName)
	}

	if language == EN {
		return getSuccessMessageEn(formName)
	}

	return getSuccessMessageEn(formName)
}
