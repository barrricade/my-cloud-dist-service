package services

type RepoBase struct {
	// nothing
}

// type curd interface {
// 	add() map[string]string
// 	update() map[string]string
// 	delete() map[string]string
// 	query() map[string]string
// }

// 公共的增删改查服务
// func add(ctx gin.Context, model map[string]interface) {
// 	var m model
// 	ctx.BindJSON(&m)
// 	// err,_ := DB().Create(&model)
// 	// c.BindJSON(&person)
// 	DB().Create(&m)
// }

// func GetProjects(c *gin.Context) {
// 	var people []Person
// 	if err := db.Find(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}
// }

// func GetPerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, person)
// 	}
// }

// func CreatePerson(c *gin.Context) {
// 	var person Person
// 	c.BindJSON(&person)
// 	db.Create(&person)
// }

// func UpdatePerson(c *gin.Context) {
// 	var person Person
// 	id := c.Params.ByName("id")
// 	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	}
// 	c.BindJSON(&person)
// 	db.Save(&person)
// 	c.JSON(200, person)
// }

// func DeletePerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	d := db.Where("id = ?", id).Delete(&person)
// 	fmt.Println(d)
// 	c.JSON(200, gin.H{"id #" + id: "deleted"})
// }
