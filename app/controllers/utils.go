package controllers

import (
	"errors"
	"fmt"
	"os/exec"
	"time"

	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/auth-core/config"
	"github.com/auth-core/database"
	"github.com/auth-core/util"
	"golang.org/x/crypto/bcrypt"
)

type GeneralResponseType struct {
	ResponseStatus    bool      `json:"response_status"`
	ResponseCode      string    `json:"response_code"`
	ResponseMessage   string    `json:"response_message"`
	ResponseTimestamp time.Time `json:"response_timestamp"`
}

func GeneralResponseErrorBuild(ResponseTime time.Time, ResponseStatus bool, ResponseCode string, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = ResponseTime
	generalResponseType.ResponseStatus = ResponseStatus
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

func GeneralResponseSuccessBuild(ResponseTime time.Time, ResponseStatus bool, ResponseCode string, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = ResponseTime
	generalResponseType.ResponseStatus = ResponseStatus
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

type GenericResponse struct {
	Success   bool        `json:"success"`
	Errorid   int         `json:"errorid"`
	MessageEN string      `json:"error_en"`
	MessageIN string      `json:"error_id"`
	Data      interface{} `json:"data"`
	//Transactionid string `json:"transactionid"`
}

var InvalidMsIsdnError = errors.New("invalid msisdn number")

var BadRequestResponse *GenericResponse
var SystemErrorResponse *GenericResponse
var InvalidNumberResponse *GenericResponse
var InvalidLoginCodeResponse *GenericResponse
var TransactionTimeoutResponse *GenericResponse

const (
	P_BASE_URL    = "baseUrl"
	P_DB_DRIVER   = "db_driver"
	P_DB_HOST     = "db_host"
	P_DB_PORT     = "db_port"
	P_DB_NAME     = "db_name"
	P_DB_USERNAME = "db_username"
	P_DB_PASSWORD = "db_password"
	R_KEY_CLIENT  = "client_"
)

func NewGenericResponse(errorId int, formName string) *GenericResponse {

	messageEn := util.GetErrorString(util.EN, errorId)
	messageId := util.GetErrorString(util.ID, errorId)

	if formName != "" && errorId == 0 {
		messageEn = util.GetSuccessString(util.EN, formName)
		messageId = util.GetSuccessString(util.ID, formName)
	}

	return &GenericResponse{
		Success:   errorId == 0,
		Errorid:   errorId,
		MessageEN: messageEn,
		MessageIN: messageId,
	}
}

func GetCurrentTimeTimeZone(timeZone string) string {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	time := theTime.Format("[02 January 2006] 15:04:05 MST")
	return time
}

func GetCurrentTimeTimeZoneUnix(timeZone string) int64 {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	time := theTime.Unix()
	return time
}

func GetCurrentTimeUTC() string {
	t := time.Now().UTC()
	timeFormat := t.Format("2006-01-02 15:04:05")
	return timeFormat
}

func GetCurrentTimestamp() time.Time {
	t := time.Now().UTC()
	// timeFormat := t.Format("2006-01-02 15:04:05")
	return t
}

func GetCurrentTimestampTimeZone(timeZone string) time.Time {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	// theTime := theTime.Format("2006-01-02 15:04:05") // return must be string
	return theTime
}

func GetRunMode() string {
	serverMode := config.MustGetString("server.mode")
	return serverMode
}

func msisdnToLocalID(msisdn string) (string, error) {
	if len(msisdn) < 4 || msisdn[0:3] != "+62" {
		return "", InvalidMsIsdnError
	}

	local := msisdn[3:]
	// did the user enter a leading by mistake?
	if local[0] == '0' {
		return "", InvalidMsIsdnError
	}
	// add leading 0
	local = "0" + local

	// mobile number?
	if local[1] != '8' && local[1:4] != "999" {
		return "", InvalidMsIsdnError
	}

	//must be in ascii plane of UTF-8, so can do straight byte comparisons
	for _, c := range local {
		if c < '0' || c > '9' {
			return "", InvalidMsIsdnError
		}
	}
	return local, nil
}

func msisdnToLocalIDReformat(msisdn string) (string, error) {
	if len(msisdn) < 4 || msisdn[0:3] != "+62" {
		return "", InvalidMsIsdnError
	}

	local := msisdn[3:]
	// did the user enter a leading by mistake?
	if local[0] == '0' {
		return "", InvalidMsIsdnError
	}

	if len(local) == 9 {
		local = "000" + local
	} else if len(local) == 10 {
		local = "00" + local
	} else if len(local) == 11 {
		local = "0" + local
	} else if len(local) == 12 {
		local = local
	} else {
		// add leading 0
		local = "0" + local
	}

	fmt.Println("$$$$$$", local)
	return local, nil
}

func encryptData() {

}

func convertToInt(value string) int {
	i, err2 := strconv.ParseInt(value, 10, 32)
	if err2 != nil {
		panic(err2)
	}
	return int(i)
}

func convertToInt64(value string) int64 {
	i, err2 := strconv.ParseInt(value, 10, 32)
	if err2 != nil {
		panic(err2)
	}
	return int64(i)
}

func showError(err error) {
	if err != nil {
		log.Info("ERROR ", err.Error())
	}
}

func showFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func showResponseError(c *gin.Context, errorCode string, errorMessage string) {
	gr := GeneralResponseErrorBuild(time.Now().UTC(), false, errorCode, errorMessage)
	c.JSON(http.StatusOK, gin.H{
		"general_response": gr,
	})
}

func GetSignature(data string) (string, error) {
	serverMode := config.MustGetString("server.mode")
	password := config.MustGetString(serverMode + ".signature_password")
	alias := config.MustGetString(serverMode + ".signature_alias")

	cmd := exec.Command("sh", "signature.sh", password, alias, data)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
		fmt.Println("cmd.Run() failed with %s\n", err)
	}
	// fmt.Printf("combined out:\n%s\n", string(out))
	return string(out), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RedisSet(key string, value string) error {

	client := database.RedisOpen()

	// SET KEY
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func RedisGet(key string) (string, error) {

	client := database.RedisOpen()

	// GET KEY
	value, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return value, nil
}
