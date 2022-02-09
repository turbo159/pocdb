package main

//type containerList struct {
//	containers []container `json:"containers"`
//}

type container struct {
	SysId      int    `json:"sysid"`
	Id         string `json:"id"`
	Value      string `json:"value"`
	CreateUser string `json:"createuser"`
	CreateDate string `json:"createdate"`
	UpdateUser string `json:"updateuser"`
	UpdateDate string `json:"updatedate"`
}
