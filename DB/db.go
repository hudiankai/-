package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	gorm.Model
	Age int
	Name string
}

//图片作品往数据库中存储的结构体
type Imagesave struct {
	gorm.Model
	Name string
	Ihash  string
	Kind int
	RID uint
	Distance int
}

//simhash算法对应的文件内容的存储结构体
type Filesave struct {
	gorm.Model
	Name string
	Hash string
	RID int
}

// 赋值数据库
var Database *gorm.DB

//连接mysql数据库，前提是启动本地的mysql
func CreateDB()  {
	db,err := gorm.Open("mysql","root:rootroot@/pha.db?charset=utf8&parseTime=True&loc=Local")
	if err != nil && db != nil{
		fmt.Println("数据库创建失败",err)
	}
	defer db.Close()
	db.SingularTable(true)
}

//连接sqlite3数据库（文本性数据库，比mysql更小、更轻）
func CinitDB()  {
	if db, err := gorm.Open("sqlite3", "pha.db?cache=shared"); err != nil {
		fmt.Println("数据库创建失败：",err)
		return
	} else {
		Database =db
	}
}

type A struct {
	Na string
}
func ZSGCDB()  {
	//任何操作数据库都应该先Open数据库，否则会在之后操作数据库是报错：无效的内存地址或空指针
	CinitDB()
	//Open()数据库之后要对应数据库的关闭
	defer Database.Close()

	// 判断并添加目标表格（结构体对应的是表格）
	if !Database.HasTable(&A{}){
		Database.CreateTable(&A{})
	}
	fmt.Println(Database.HasTable(&Student{}))
	if !Database.HasTable(&Student{}){
		Database.CreateTable(&Student{})
	}

	//随意插入数据记录
	for i := 0;i<5;i++{
		s := Student{ Age: i,Name: "gfhaslk"}
		Database.Create(&s)
		//判断插入数据是够缺少主键
		if Database.NewRecord(&s){
			fmt.Println("包含主键")
		}else {
			fmt.Println("缺少主键")
		}
		//若要判断插入数据是否成功,注意下面语句包含了插入操作
		if err := Database.Create(&s).Error;err != nil{
			fmt.Println("插入失败",err)
		}
	}

	//查询
	s := []Student{}
	//第一个参数必须是地址，第二个参数为查询条件，此处意思为查询age=17的用户信息，因为数据库中不止一条，所以返回多个，通过ID可以区分
	Database.Find(&s,"age=?",17)
	//与上一行的代码相等，查询age=17的所有记录
	Database.Where("age=?",17).Find(&s)
	//查询获得id为1,3,4,7的记录
	Database.Find(&s,"id in (?)",[]int{1,3,4,7})
	//查询获取table中的所有记录
	Database.Find(&s)


	//更新,当Model中的变量是单个变量时，就是更新单条记录，
	//若是多个结构体的列表或者是空的原结构体（&Student），则更新所有记录中的某个数据(此处为将所有记录的age改为16)
	Database.Model(&s).Update("age",16)
	//保存目标对象中的值
	Database.Save(&s)
	//依次修改目标查询列表记录中的数值
	for _,v := range s{
		Database.Model(v).Update("age",88)
	}

	//删除，此处因为有gorm.Model,删除后在对应的数据库中还会有记录，但是删除时间已不再是null，查询也无法再次查询到
	// 若参数为多个结构体的列表或原结构体,则删除全部数据
	Database.Delete(s)
	// 依据条件删除目标记录
	Database.Where("id>?",20).Delete(&Student{})

}