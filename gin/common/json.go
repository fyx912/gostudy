package common

import()

func Json(data interface{})map[string]interface{}{
	jsonMap := make(map[string]interface{})
	jsonMap["code"] = 0
	jsonMap["msg"] = "success"
	jsonMap["data"] = data 
	return jsonMap
}

func Jsons(code int,msg string,data interface{})map[string]interface{}{
	jsonMap := make(map[string]interface{})
	jsonMap["code"] = code
	jsonMap["msg"] = msg
	jsonMap["data"] = data 
	return jsonMap
}