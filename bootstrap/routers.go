package bootstrap

import (
	"net/http"

	"gome.com.cn/httpserver/controller"
)

//初始化路由
func InitRouter() {
	//静态资源地址
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	//设置md编辑界面
	http.HandleFunc("/admin/", controller.Index)

	//设置full编辑界面
	http.HandleFunc("/admin/full/", controller.Full)

	//设置md保存功能
	http.HandleFunc("/admin/savefile/", controller.SaveMarkdown)

	http.HandleFunc("/admin/upload/", controller.Upload)
	//图片上传功能
	http.HandleFunc("/admin/uploadfile/", controller.UploadFile)
}
