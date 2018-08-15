package web

import (
	"github.com/MrChang0/idls/db"
	"github.com/MrChang0/idls/device"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router = gin.Default()

func init() {
	router.Static("/static", "web/static")
	router.LoadHTMLGlob("web/template/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/manager", func(context *gin.Context) {
		context.HTML(http.StatusOK, "manager.html", device.GetAllOnlineDevice())
	})
	router.GET("/manager/:uuid/namechange", func(context *gin.Context) {
		uuid := context.Param("uuid")
		newname := context.Query("newname")

		d, ok := device.FindOnlineDevice(uuid)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
			})
			return
		}
		device.ChangeName(d.Name, newname)
		d.Name = newname
		db.UpdateDevice(d)
		context.JSON(http.StatusOK, gin.H{
			"statue": "success",
		})
	})
	router.GET("/manager/:uuid/vm", func(context *gin.Context) {
		uuid := context.Param("uuid")
		d, ok := device.FindOnlineDevice(uuid)
		if !ok {
			context.Redirect(http.StatusOK, "/manager")
		}
		code, err := d.GetCode()
		type Data struct {
			UUID string
			Code string
			Err  string
		}
		data := Data{UUID: uuid}
		if err != nil {
			data.Err = err.Error()
			data.Code = ""
		} else {
			data.Code = code
			data.Err = ""
		}
		context.HTML(http.StatusOK, "edit.html", data)
	})
	router.GET("/manger/:uuid/error", func(context *gin.Context) {
		uuid := context.Param("uuid")
		d, ok := device.FindOnlineDevice(uuid)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
				"err":    "uuid is not found",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"statue": "success",
			"data":   d.Error(),
		})
	})
	router.GET("/manager/:uuid/event", func(context *gin.Context) {
		uuid := context.Param("uuid")
		d, ok := device.FindOnlineDevice(uuid)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
				"err":    "uuid is not found",
			})
			return
		}
		str, err := d.EventType()
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
				"err":    err.Error(),
			})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{
				"statue": "success",
				"data":   str,
			})
		}
	})
	router.POST("/manger/:uuid/code", func(context *gin.Context) {
		uuid := context.Param("uuid")
		d, ok := device.FindOnlineDevice(uuid)
		if !ok {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
				"err":    "uuid is not found",
			})
			return
		}
		code := context.PostForm("code")
		err := d.NewCode(code)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"statue": "fail",
				"err":    err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"statue": "success",
			})
		}
	})
}

func Run() {
	router.Run("0.0.0.0:8080")
}
