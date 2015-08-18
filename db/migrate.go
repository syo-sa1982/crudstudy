package crudstudy

import (

	"github.com/jinzhu/gorm"
	"github.com/syo-sa1982/crudstudy/models"
)

func main() {
	db, _ := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
}
