package services

func LegacyLogin() {
	Login()
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
}

func LegacyGetUserByToken() {
	GetUserByToken()
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

func LegacyGetRegCode() {
  GetRegCode()
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
