package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/", index)
	r.GET("/upload", upload)
	r.POST("/upload", doUpload)
	r.GET("/downfile", fileDownload)
	r.GET("/delfile", delFilename)
	r.Run(":8080")
}

func index(c *gin.Context){
	fileList := make([]string, 0)
	files, _ := ioutil.ReadDir("upload/")
	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"content":fileList})
}

func upload(c *gin.Context){
	c.HTML(http.StatusOK, "upload.html", gin.H{})
}

func doUpload(c *gin.Context){
	//获取表单数据 参数为name值
	f, err := c.FormFile("f1")
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		//将文件保存至本项目根目录中
		c.SaveUploadedFile(f, "upload/" + f.Filename)
		//保存成功返回正确的Json数据
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}

func fileDownload(c *gin.Context){
	filename := c.Query("filename")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./upload/" + filename)
}

func delFilename(c *gin.Context){
	filename := c.Query("filename")
	err := os.Remove("upload/" + filename)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Failed",
		})
	} else {
		// 删除成功
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}