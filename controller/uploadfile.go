package controller

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/donnie4w/go-logger/logger"
)

//上传图片
func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var guid string = r.URL.Query().Get("guid")
	//获取当前url
	logger.Debug("url:", guid)
	logger.Debug("url:", r.URL.EscapedPath())
	logger.Debug("url:", r.URL.Query())
	logger.Debug("url:", r.URL.String())
	logger.Debug("url:", r.Host)

	t := time.Now().Unix()
	logger.Debug("时间戳：", t)
	sname := strconv.FormatInt(t, 10)
	logger.Debug("时间戳：", sname)

	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	if r.Method == "GET" {
		logger.Debug("GET")
		//t, err := template.ParseFiles("upload.gptl")
		//t.Execute(w, nil)
		buffer.WriteString("{\"success\":0,\"url\":\"\",\"message\":\"上传失败！\"}")
	} else {
		logger.Debug("POST")
		file, handle, err := r.FormFile("editormd-image-file")
		if err != nil {
			logger.Fatal(err.Error())
			return
		}

		fileSuffix := path.Ext(handle.Filename)
		logger.Debug("后缀名：", fileSuffix)
		pathfile := "./static/upload/" + sname + fileSuffix

		f, err := os.OpenFile(pathfile, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)
		if err != nil {
			logger.Fatal(err.Error())
			return
		}
		defer f.Close()
		defer file.Close()
		buffer.WriteString("{\"success\":1,\"url\":\"")
		buffer.WriteString("http://")
		buffer.WriteString(r.Host)
		buffer.WriteString("/upload/")
		buffer.WriteString(sname)
		buffer.WriteString(fileSuffix)
		buffer.WriteString("\",\"message\":\"上传成功！\"}")
	}
	w.Write(buffer.Bytes())
}

func Upload(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("static/admin/upload.html")
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	logger.Debug("文件上传页面加载完成")
}
