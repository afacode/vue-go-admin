package ServiceSys

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zhoukai/db"
	ModelSys "zhoukai/model/sys"
	UtilsDB "zhoukai/utils/db"
)

func DictCreate(data ModelSys.SysDict)(ModelSys.SysDict, *gorm.DB)  {
	data.ID = UtilsDB.CreateUUId()
	result := db.DB.Create(&data)
	return data, result
}
func DictList(data []ModelSys.SysDict, c *gin.Context)([]ModelSys.SysDict, *gorm.DB)  {
	result := db.DB.Scopes(UtilsDB.Paginate(c, &data)).Find(&data)
	return data, result
}
func DictUpdate(data ModelSys.SysDict)(ModelSys.SysDict, *gorm.DB)  {
	result := db.DB.Updates(&data)
	return data, result
}
func DictDel(data ModelSys.SysDict)(ModelSys.SysDict, *gorm.DB)  {
	result := db.DB.Delete(&data)
	return data, result
}