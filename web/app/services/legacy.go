package services

func LegacyLogin(stuid string, password string) {
	Login(stuid, password);
}

func LegacyGetUserByToken(token string) {
	GetUserByToken(token)
	// get user information with token
	// return user information
	// $infor = [
	//     'url' => $GLOBALS['BASEURL'] . '/users',
	//     'data' => [
	//         'dToken' => $data['token'],
	//     ],
	//     'refererUrl' => 'https://auth.youngon.work',
	//     'contentType' => 'application/json',
	// ];
	// $resdata = json_decode(httpGet($infor), true);
	// return json($resdata);
}

func LegacyGetRegCode(regCode string) {
  GetRegCode(regCode)
	// check register code
	// return a boolean
	// $infor = [
	//     'url' => $GLOBALS['BASEURL'] . '/commonset/index/getrcode',
	//     'data' => [],
	//     'refererUrl' => 'http://auth.youngon.work',
	//     'contentType' => 'application/json',
	// ];
	// $resdata = json_decode(httpGet($infor), true);
	// return json(['status' => $resdata['Rcode'] == base64_decode($data['regcode'])]);
}
