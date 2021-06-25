package simhash

import (
	"Dhash/DB"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//读取文件中的内容,返回二进制文件
func Readfile(s string) []byte {
	f, err := os.OpenFile("D:/workspace/PHA/simhash/wenben/"+s,os.O_RDONLY,0600)
	if err!=nil{
		fmt.Println("作品读取失败：",err)
	}else {
		contentbyte,err := ioutil.ReadAll(f)
		if err != nil{
			fmt.Println("作品读取失败！：",err)
		}
		return contentbyte
	}
	defer f.Close()
	return []byte("")
}

//将二进制转换成simhash值
func ContTOSimhash(c []byte) uint64 {
	var hash uint64
	hash = Simhash(NewWordFeatureSet(c))
	return hash
}

// 比较两个哈希值的大小
func Filecomp(ha1 uint64,ha2 string) uint8 {
	//h1,_ := strconv.ParseUint(ha1,10,64)
	h2,_ := strconv.ParseUint(ha2,10,64)
	xiangsi := Compare(ha1,h2)
	return xiangsi
}

func Savefile(s string)  {
	DB.CinitDB()
	defer DB.Database.Close()
	//判断并创建表
	if !DB.Database.HasTable(&DB.Filesave{}){
		DB.Database.CreateTable(&DB.Filesave{})
	}
	//获取目标文件中内容的二进制
	contentbyte := Readfile(s)
	//依据二进制内容转换成simhash值
	hash := ContTOSimhash(contentbyte)
	//存储目标
	filesace := DB.Filesave{Name: s,Hash: strconv.FormatUint(hash,10),RID: 0}

	//判断是否还有相似的文档记录
	fhashs := []DB.Filesave{}
	DB.Database.Find(&fhashs)
	for k,v := range fhashs {
		if Filecomp(hash,v.Hash)<5{
			filesace.RID = k+1
		}
	}

	fmt.Println("新的文档记录为：",filesace)
	//将目标值写入到数据库中
	err := DB.Database.Create(&filesace).Error
	if err!= nil{
		fmt.Println("新文档对应的哈希写入失败",err)
	}


	// 判断当前值是否具有主键
	//if DB.Database.NewRecord(&filesace){
	//	fmt.Println("新文档记录写入成功！")
	//}else {
	//	fmt.Println("新文档记录写入失败！")
	//}


}
