package Imagehash

import (
	"Dhash/DB"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)
// 哈希值类型为uint64，计算的时候是以10进制数字展示的，存储的时候是转换成了string，filesimhash也是一样


// 删除数据库中的目标照片
func Delete()  {
	//连接打开数据库
	DB.CinitDB()
	defer  DB.Database.Close()

	// 获取数据库中ID为1的照片，并赋值给images
	images := &[]DB.Imagesave{}
	DB.Database.Where("id<?",19).Find(&images)

	// 删除照片
	DB.Database.Delete(&images)

}

func Save(s string)  {
	//读取照片文件获得照片
	image := Readimage(s)
	// 获取照片对应的dhash值
	//imageHash, err,kind := DifferenceHash(image)
	// 获取照片对应的Ahash值
	//imageHash, err,kind := AverageHash(image)
	// 获取照片对应的phash值
	imageHash, err,kind := PerceptionHash(image)

	if err != nil{
		fmt.Println("图片哈希值计算失败：",err)
	}
	fmt.Println("新图片的iamgehash为：",imageHash)
	// 生成照片对应的结构体

	Saveimage := DB.Imagesave{Name: s,Ihash: strconv.FormatUint(imageHash.hash,10),Kind:kind,RID: 0,Distance: 9999}

	DB.CinitDB()
	if !DB.Database.HasTable(&DB.Imagesave{}){
		DB.Database.CreateTable(&DB.Imagesave{})
	}
	defer DB.Database.Close()
	//DB.Database.AutoMigrate(&Imagesave{})


	Images := []DB.Imagesave{}
	//获取对应表中的所有结构体
	DB.Database.Find(&Images)

	//判断存证照片是否是相似照片
	for k,v := range Images{
		// 只与同类哈希算法计算的哈希值比较
		if v.Kind == kind{
			hash,_ := strconv.ParseUint(v.Ihash,10,64)
			ihash := &ImageHash{hash,Kind(v.Kind)}
			if num,_ :=imageHash.Distance(ihash);num<6{
				if Saveimage.Distance > num{
					Saveimage.Distance = num
					Saveimage.RID = Images[k].ID
				}
				fmt.Println("遍历数据库的值,相似值为：",k,v)
			}
		}
	}

	//DB.Database.Create(&Saveimage)
	err = DB.Database.Create(&Saveimage).Error
	if err!=nil{
		fmt.Println("写入失败：",err)
	}else {
		fmt.Println("写入成功，要写入的值为：",Saveimage)
	}

	// 判断当前值是否具有主键
	//if DB.Database.NewRecord(Saveimage){
	//	fmt.Println("写入成功")
	//}else {
	//	fmt.Println("写入失败")
	//}
}

// 读取目标路径的照片s，
func Readimage(s string) image.Image {
	file1, err := os.Open( "D:/workspace/PHA/-/Imagehash/tupian/Example/"+s+".jpg")
	if err !=nil{
		fmt.Println("Open failed",err)
	}
	defer file1.Close()
	//读取一个JPEG文件，将其解码，返回Image.image
	img1, _ := jpeg.Decode(file1)
	return img1
}

func Comparehash(img1,img2 image.Image){
	//从系统中读取文件,并返回一个文件对象
	//file1, err := os.Open( "D:/workspace/Imagehash/tupian/"+s1+".jpg")
	//if err !=nil{
	//	fmt.Println("Open failed",err)
	//	return
	//}
	//file2, err2 := os.Open( "D:/workspace/Imagehash/tupian/"+s2+".jpg")
	//if err2 != nil{
	//	fmt.Println("获取文件2失败")
	//	return
	//}
	//
	//
	//defer file1.Close()
	//defer file2.Close()
	////读取一个JPEG文件，将其解码，返回image.Image
	//img1, _ := jpeg.Decode(file1)
	//img2, _ := jpeg.Decode(file2)

	//Dhash算法计算
	hash1, _,_ := DifferenceHash(img1)
	hash2, _,_ := DifferenceHash(img2)

	//比较两个图片哈希值的海明距离
	distance1, err1 := hash1.Distance(hash2)
	if err1 != nil{
		fmt.Println(err1)
		return
	}
	fmt.Println( "Dhash算法：Distance between images:", distance1)

	//Phash算法计算
	hash1, _,_ = PerceptionHash(img1)
	hash2, _,_ = PerceptionHash(img2)
	distance1, _ = hash1.Distance(hash2)
	fmt.Printf( "Phash算法：Distance between images: %d \n", distance1)
	//Ahash算法计算
    hash1, _,_ = AverageHash(img1)
	hash2, _,_ = AverageHash(img2)
	distance1, _ = hash2.Distance(hash1)
	fmt.Printf( "Ahash算法：Distance between images: %d \n", distance1)

}