package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/w-xuefeng/yg-auth-go/web/app/interfaces"

	"github.com/w-xuefeng/yg-auth-go/web/app/config"
)

func Login(stuid string, password string) (interfaces.ResLogin, error) {
	decodeStuid, err := base64.StdEncoding.DecodeString(stuid)
	res := interfaces.ResLogin{}
	if err != nil {
		return res, err
	}

	decodePassword, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return res, err
	}

	requestBody := strings.NewReader("stuid=" + string(decodeStuid) + "&password=" + string(decodePassword))

	req, err := http.NewRequest("POST", config.UrlLogin, requestBody)
	if err != nil {
		return res, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RefererUrl", config.AuthIndex)
	response, err := (&http.Client{}).Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		return res, err
	}

	rs, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	type RsType struct {
		Status bool   `json:"status"`
		Title  string `json:"title"`
		DToken string `json:"dToken"`
	}

	var rsJson RsType
	err = json.Unmarshal(rs, &rsJson)
	if err != nil {
		return res, err
	}

	if rsJson.Status {
		res.Token = rsJson.DToken
		return res, nil
	}

	errMessage := "request error"

	if rsJson.Title != "" {
		errMessage = rsJson.Title
	}

	return res, errors.New(errMessage)

}

func GetUserByToken(token string) (interfaces.AuthUser, error) {
	queries := "?dToken=" + token
	res := interfaces.AuthUser{}
	req, err := http.NewRequest("GET", config.UrlUsers+queries, nil)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RefererUrl", config.AuthIndex)
	response, err := (&http.Client{}).Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		return res, err
	}

	rs, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	type RsType struct {
		Status  bool                `json:"status"`
		Title   string              `json:"title"`
		Resdata interfaces.AuthUser `json:"resdata"`
	}

	var rsJson RsType
	err = json.Unmarshal(rs, &rsJson)
	if err != nil {
		return res, err
	}

	if rsJson.Status {
		return rsJson.Resdata, nil
	}

	errMessage := "request error"

	if rsJson.Title != "" {
		errMessage = rsJson.Title
	}

	return res, errors.New(errMessage)
}

func CheckRegCode(regCode string) (bool, error) {
	decodeRegCode, err := base64.StdEncoding.DecodeString(regCode)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", config.UrlRegCode, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RefererUrl", config.AuthIndex)
	response, err := (&http.Client{}).Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		return false, err
	}

	rs, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	type RsType struct {
		Rcode string `json:"Rcode"`
	}

	var rsJson RsType
	err = json.Unmarshal(rs, &rsJson)
	if err != nil {
		return false, err
	}
	return rsJson.Rcode == string(decodeRegCode), nil
}
