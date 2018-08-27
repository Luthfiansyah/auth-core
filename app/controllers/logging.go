package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/auth-core/app/models"
	"github.com/auth-core/app/repositories"
	"github.com/auth-core/app/types"
)

func ReadBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func Logging(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	buf, _ := ioutil.ReadAll(c.Request.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

	fmt.Println(ReadBody(rdr1)) // Print request body

	c.Request.Body = rdr2
	c.Next()
}

func Log(data *types.Log) {

	req, err := json.Marshal(data.RequestMessage)
	res, err := json.Marshal(data.ResponseMessage)
	endpoint, err := json.Marshal(data.Endpoint)

	reqTimestamp := int32(data.RequestTime.Unix())
	resTimestamp := int32(data.ResponseTime.Unix())
	elapsedTime := resTimestamp - reqTimestamp

	var log models.Log
	log.Endpoint = string(endpoint)
	log.RequestMessage = string(req)
	log.ResponseMessage = string(res)
	log.RequestTime = data.RequestTime
	log.ResponseTime = data.ResponseTime
	log.ElapsedTime = elapsedTime
	log.CreatedAt = GetCurrentTimestampTimeZone("Asia/Jakarta")

	fmt.Println("LOG API ----------------------------------------------------------------------------------------------")
	fmt.Println("ENDPOINT", data.Endpoint)
	fmt.Println("REQUEST MESSAGE:: ", string(req))
	fmt.Println("RESPONSE MESSAGE :: ", string(res))
	fmt.Println("REQUEST TIME:: ", data.RequestTime)
	fmt.Println("RESPONSE TIME :: ", GetCurrentTimestampTimeZone("Asia/Jakarta"))
	fmt.Println("ELAPSED TIME :: ", elapsedTime, " Second")

	// fmt.Println(log)
	err = repositories.InsertLog(&log)
	if err != nil {
		fmt.Println(err.Error())
	}
}
