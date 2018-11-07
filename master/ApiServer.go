package master

import (
	"net/http"
	"net"
	"time"
)

//任务的HTTP接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	//单例对象
	G_apiServer *ApiServer
)
//保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request)  {

}
//初始化服务
func InitApiServer()  {
	var (
		mux *http.ServeMux
		listener net.Listener
		err error
		httpServer *http.Server
	)
	//配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save",handleJobSave)

	//启动TCP监听
	if listener,err = net.Listen("tcp",":8070");err != nil{
		return
	}

	//创建一个HTTP服务
	httpServer = &http.Server{
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:mux,
	}

	//赋值单例
	G_apiServer = &ApiServer{
		httpServer:httpServer,
	}

	//启动了服务端
	go httpServer.Serve(listener)
	return
}