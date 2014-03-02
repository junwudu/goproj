package errors

import (
	"strconv"
	"io/ioutil"
	"encoding/json"
	"io"
	"fmt"
)

type BaiduError struct {
	StatusCode  int
	ErrorCode   int
	Description string
}


func (baiduError *BaiduError) Parse(reader io.Reader) {

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		baiduError.Description = err.Error()
		return
	}

	var v interface {}
	err = json.Unmarshal(data, &v)
	if err != nil {
		baiduError.Description = err.Error()
		return
	}

	ve := v.(map[string]interface {})

	eObj := ve["Error"]

	eoe := eObj.(map[string]interface {})

	errCode := eoe["code"].(string)
	c, err := strconv.ParseInt(errCode, 10, 32)
	if err == nil {
		baiduError.ErrorCode = int(c)
	}

	baiduError.Description = eoe["Message"].(string)
}


func (baiduError *BaiduError) Error() string {
	return fmt.Sprintf("%d->%d:%s", baiduError.ErrorCode, baiduError.StatusCode, baiduError.Description)
}


func GetBaiduError(code int) BaiduError {
	codeMap := BaiduErrorCodeMap()
	return codeMap[string(code)]
}

var BaiduErrorCodeMap = func() func() map[string]BaiduError {
	list := BaiduErrorCodeList()
	mapList := make(map[string]BaiduError, len(list))


	for _, be := range list {
		mapList[strconv.FormatInt(int64(be.ErrorCode), 10)] = be
	}

	return func() map[string]BaiduError {
		return mapList
	}
}()

var BaiduErrorCodeList = func() func() []BaiduError {

	codeList := []BaiduError{
		{400, 1, "签名失败"},         //{"Error":{"code":"1","Message":"signature errors","LogId":"11111111"}}
		{400, 3, "ACL上传错误"},      //{"Error":{"code":"3","Message":"acl put errors","LogId":"11111111"}}
		{400, 4, "ACL查询出错"},      //{"Error":{"code":"4","Message":"acl query errors","LogId":"11111111"}}
		{400, 5, "获取ACL出错"},      //{"Error":{"code":"5","Message":"acl get errors","LogId":"11111111"}}
		{400, 7, "Bucket已存在"},    //{"Error":{"code":"7","Message":"bucket already exists","LogId":"11111111"}}
		{400, 8, "请求错误"},         //{"Error":{"code":"8","Message":"bad request","LogId":"11111111"}}
		{403, 11, "拒绝访问"},        //{"Error":{"code":"11","Message":"access denied","LogId":"11111111"}}
		{403, 18, "超出容量限额"},      //{"Error":{"code":"18","Message":"storage exceed limit","LogId":"11111111"}}
		{403, 19, "请求数超出限额"},     //{"Error":{"code":"19","Message":"request exceed limit","LogId":"11111111"}}
		{403, 20, "流量超出限额"},      //{"Error":{"code":"20","Message":"transfer exceed limit","LogId":"11111111"}}
		{404, 2, "文件不存在"},        //{"Error":{"code":"2","Message":"object not exists","LogId":"11111111"}}
		{404, 6, "获取ACL出错"},      //{"Error":{"code":"6","Message":"acl get errors","LogId":"11111111"}}
		{413, 24, "请求文件过大"},      //{"Error":{"code":"24","Message":"request entity too large","LogId":"11111111”}}
		{416, 21, "请求范围不符合要求"},   //{"Error":{"code":"21","Message":"requested range not satisfiable","LogId":"11111111"}}
		{500, 9, "服务内部错误"},       //{"Error":{"code":"9","Message":"baidubs internal errors","LogId":"11111111"}}
		{501, 10, "不支持该请求"},      //{"Error":{"code":"9","Message":"no implement","LogId":"11111111"}}
		{503, 12, "系统繁忙，请稍后再试"},  //{"Error":{"code":"12","Message":"service unavailable","LogId":"11111111"}}
		{503, 13, "系统繁忙，请稍后再试"},  //{"Error":{"code":"13","Message":"service unavailable","LogId":"11111111"}}
		{503, 14, "上传文件数据信息失败"},  //{"Error":{"code":"14","Message":"put object data errors","LogId":"11111111"}}
		{503, 15, "上传文件元信息失败"},   //{"Error":{"code":"15","Message":"put object meta errors","LogId":"11111111"}}
		{503, 16, "无法获取文件的数据信息"}, //{"Error":{"code":"16","Message":"get object data errors","LogId":"11111111"}}
		{503, 17, "无法获取文件元信息"},   //{"Error":{"code":"17","Message":"get object meta errors","LogId":"11111111"}}
		{504, 22, "后端连接超时"},      //{"Error":{"code":"22","Message":"connect backend timeout","LogId":"11111111"}}
	}
	return func() []BaiduError {
		return codeList
	}
}()

