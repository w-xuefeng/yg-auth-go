package services

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/w-xuefeng/yg-auth-go/web/app/config"
)

func Login(stuid string, password string) (any, error) {
	// header('Access-Control-Allow-Origin: *');
	// header('Access-Control-Allow-Methods: POST, GET, PUT, PATCH, OPTIONS, DELETE');
	// header('Access-Control-Allow-Headers: Origin, X-Requested-With, Content-Type, Accept, X-URL-PATH, X-Access-Token, Authorization');
	// header('Content-type: application/json');
	// $infor = [
	//     'url' => $GLOBALS['BASEURL'] . '/session',
	//     'data' => [
	//         'stuid' => base64_decode($data['stuid']),
	//         'password' => base64_decode($data['password']),
	//     ],
	//     'refererUrl' => 'https://auth.youngon.work',
	//     'contentType' => 'application/json',
	// ];
	// $resdata = json_decode(httpPost($infor), true);
	// $a = $resdata['status'] ? ['status' => true, 'resdata' => ['token' => $resdata['dToken']]] : $resdata;

	decodeStuid, err := base64.StdEncoding.DecodeString(stuid)
	if err != nil {
		return nil, err
	}

	decodePassword, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, err
	}

	requestBody := strings.NewReader("stuid=" + string(decodeStuid) + "&password=" + string(decodePassword))

	req, err := http.NewRequest("POST", config.UrlLogin, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RefererUrl", config.AuthIndex)
	response, err := (&http.Client{}).Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, err
	}
	reader := response.Body
	return reader, nil
}
func GetUserByToken(token string) {}
func GetRegCode(regCode string)   {}
