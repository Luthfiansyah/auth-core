package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/auth-core/app/repositories"
	"github.com/auth-core/app/types"
)

func GetCity(c *gin.Context) {

	reqTime := GetCurrentTimestampTimeZone("Asia/Jakarta")
	fmt.Println(reqTime)

	// var requestCity types.RequestCity
	// if err := c.BindJSON(&requestCity); err != nil {
	// 	log.Errorf("", "response parsing request: %s", err.Error())
	// 	c.JSON(http.StatusBadRequest, BadRequestResponse)
	// }

	coordinate := c.Param("coordinate")
	radius := c.Param("radius")
	fmt.Println("####", coordinate)

	s := strings.Split(coordinate, ",")
	latitude, longitude := s[0], s[1]
	fmt.Println("lat : ", latitude, " - long : ", longitude)

	if radius == "" {
		radius = "20"
	}

	// lat, err := strconv.ParseInt(latitude, 10, 64)
	// lng, err := strconv.ParseInt(longitude, 10, 64)
	// rad, err := strconv.ParseInt(radius, 10, 64)
	// fmt.Println("lat : ", lat, " - long : ", lng)
	// if err == nil {
	// 	fmt.Printf("%d of type %T", lat, lat)
	// }

	result, err := repositories.GetCityByCoordinate(latitude, longitude, radius)
	if err != nil {
		fmt.Println(err.Error())
		showResponseError(c, "8002", "City Not Found")
		return
	}

	if result == nil {
		showResponseError(c, "8003", "City Not Found")
		return
	}

	fmt.Println(len(result))

	var data []types.ResponseCity
	for _, row := range result {
		r := types.ResponseCity{}
		r.Name = row.Name.String
		r.Regency = row.Regency.String
		r.Province = row.Province.String
		r.Region = row.Region.String
		data = append(data, r)
	}

	gr := GeneralResponseSuccessBuild(time.Now().UTC(), true,
		"000", "Success")

	// INSERT LOG
	var log types.Log
	log.Endpoint = c.Request.URL
	log.RequestMessage = c.Request.RequestURI
	log.ResponseMessage = data
	log.RequestTime = reqTime
	log.ResponseTime = GetCurrentTimestampTimeZone("Asia/Jakarta")
	Log(&log)

	c.JSON(http.StatusOK, gin.H{
		"data":             data,
		"general_response": gr,
	})

	return
}
