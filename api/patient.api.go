package api

import (
	"main/db"
	"main/model"

	"github.com/gin-gonic/gin"
)

func setupPatientAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{

	//	authenAPI.POST("/edituser", edituser)
		//authenAPI.POST("/getuser", interceptor.JwtVerify, getuser)
		authenAPI.POST("/getpatient/:cid", getPatient)
	}
}

func getPatient(c *gin.Context) {
	var patient []model.Patient
	if c.ShouldBind(&patient)== nil{

		//db.GetDB().Find(&patient)
if err:=	db.GetDB().Where("cid = ?", c.Param("cid")).First(&patient).Error; err == nil{
	c.JSON(200, gin.H{"data": patient,"msg":"found"})
}else{
	c.JSON(200,gin.H{"msg":"notfound"})
}
	
	}else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
	
}


