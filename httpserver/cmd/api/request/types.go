// Package request common define
// Created by chenguolin 2018-11-17
package request

// CommonParams 请求通用参数
type CommonParams struct {
	Gid        string `form:"gid"`
	GidStatus  string `form:"gid_status"`
	SigTime    int    `form:"sigTime"`
	SigVersion string `form:"sigVersion"`
	OsType     string `form:"os_type"`
	OsVersion  string `form:"os_version"`
	AppVersion string `form:"app_version"`
	Idfa       string `form:"idfa"`
	Idfv       string `form:"idfv"`
	AndroidID  string `form:"android_id"`
	Imei       string `form:"imei"`
	Iccid      string `form:"iccid"`
	Model      string `form:"model"`
	Channel    string `form:"channel"`
	Carrier    string `form:"carrier"`
	Network    string `form:"network"`
	Language   string `form:"language"`
	MacAddr    string `form:"mac_addr"`
	BundleID   string `form:"bundleid"`
}
