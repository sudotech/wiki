package controller

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/donnie4w/go-logger/logger"
)

//编辑页面绑定事件监听
func Index(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	path := v.Get("file")
	if path == "" {
		return
	}
	logger.Debug("接收参数为：", path)
	//尝试打开原有文件
	f, err := os.OpenFile("static/md/"+path+".md", os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在md目录
	defer f.Close()
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	t, err := template.ParseFiles("static/admin/index.html")
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	err = t.Execute(w, &map[string]interface{}{
		"mdfile": path + ".md",
		"mdname": path,
	})
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	logger.Debug("编辑页面加载完成")
}

//full页面demo事件监听
func Full(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("static/admin/full.html")
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	logger.Debug("编辑页面加载完成")
}

//保存markdown文件事件监听
func SaveMarkdown(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	mdstr := r.Form.Get("mdstr")
	htmlstr := r.Form.Get("htmlstr")
	filename := r.Form.Get("filename")
	if mdstr == "" {
		return
	}
	logger.Debug("文档名称：", filename, "md文档：", mdstr, "html:", htmlstr)

	//保存md
	var mdarray = []byte(mdstr)
	err := ioutil.WriteFile("static/md/"+filename+".md", mdarray, 0666) //写入文件(字节数组)
	if err != nil {
		logger.Fatal("保存md文件异常：", err.Error())
		return
	}

	//保存静态页开始
	//获取模板
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	//添加头部
	buffer = GetMasterTop(filename, filename)
	buffer.WriteString(mdstr)
	buffer.WriteString(GetMasterBottom())

	//打印日志
	logger.Debug(buffer.String())
	err1 := ioutil.WriteFile("static/html/"+filename+".html", buffer.Bytes(), 0666) //写入文件(字节数组)
	if err1 != nil {
		logger.Fatal("保存html文件异常：", err1.Error())
		return
	}
	//保存静态页结束

	//logger.Debug("接收到markdown信息为：", mdstr, "；html信息为：", htmlstr, "；文件名称：", filename)
}

//页面拼接
func GetMasterTop(tilte string, filename string) bytes.Buffer {
	var buffer bytes.Buffer
	/*
		func (b *Buffer) WriteString(s string) (n int, err error)
		Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
	*/
	buffer.WriteString(`<!DOCTYPE html>
<html lang="zh">
    <head>
        <meta charset="utf-8" />
        <title>`)
	buffer.WriteString(tilte)
	buffer.WriteString(`</title>
		<link rel="stylesheet" href="../resources/css/ncEditormdView.css" />
		<link rel="stylesheet" href="../resources/css/editormd.preview.css" />

    </head>
    <body>
	<a href="../" target="_blank" >欢迎新同学</a> | <a href="index.html" target="_blank" >首页</a> | <a href="/admin/?file=`)
	buffer.WriteString(filename)
	buffer.WriteString(`" target="_blank" >编辑</a> | <a href="help.html" target="_blank" >帮助</a>
        <div id="layout">
            <div id="sidebar">
                <h1>目录</h1>
                <div class="markdown-body editormd-preview-container" id="custom-toc-container">#custom-toc-container</div>
            </div>
            <div id="txtMdSourceView">
                <textarea id="txtMdSource" style="display:none;">`)
	return buffer
}

func GetMasterBottom() string {
	return `</textarea>          
            </div>
        </div>
		<script src="../resources/js/jquery.min.js"></script>
		<script src="../resources/js/lib/marked.min.js"></script>
		<script src="../resources/js/lib/prettify.min.js"></script>
		<script src="../resources/js/lib/raphael.min.js"></script>
		<script src="../resources/js/lib/underscore.min.js"></script>
		<script src="../resources/js/lib/sequence-diagram.min.js"></script>
		<script src="../resources/js/lib/flowchart.min.js"></script>
		<script src="../resources/js/lib/jquery.flowchart.min.js"></script>
		<script src="../resources/js/core/editormd.min.js"></script>
		<script type="text/javascript">
           $(function() {
                var testEditormdView, testEditormdView2;
                 
	testEditormdView = editormd.markdownToHTML("txtMdSourceView", {
                        markdown        : $("#txtMdSource").text(),
                        htmlDecode      : "style,script,iframe",  // you can filter tags decode
                        tocm            : true,    // Using [TOCM]
                        tocContainer    : "#custom-toc-container", // 自定义 ToC 容器层
                        emoji           : true,
                        taskList        : true,
                        tex             : true,  // 默认不解析
                        flowChart       : true,  // 默认不解析
                        sequenceDiagram : true,  // 默认不解析
                    }); 
					console.log(testEditormdView);
            });
        </script>
    </body>
</html>`
}
