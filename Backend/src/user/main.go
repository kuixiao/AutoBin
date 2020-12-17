package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kuixiao/AutoBin/Backend/src/user/conf"
	"github.com/kuixiao/AutoBin/Backend/src/user/handlers"
	pb "github.com/kuixiao/AutoBin/Backend/src/user/protos"
	"github.com/rs/cors"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
)

func startGrpcServer() error {
	var config = conf.Config
	/// 开启Grpc端口监听
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", config.GrpcEndpoint))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()	// 获取grpcServer
	pb.RegisterUserServiceServer(grpcServer, handlers.NewService())	//
	log.Println("Grpc Listen(config.GrpcEndpoint) : "+config.GrpcEndpoint)
	return grpcServer.Serve(lis)
}

func startHttpServer() error {
	// fileServer
	var config = conf.Config
	verify := config.Verify
	fileServer := http.StripPrefix("/autobin/user/statics/", http.FileServer(http.Dir(verify)))

	ctx := context.Background()
	ctx, cancel :=context.WithCancel(ctx)
	defer  cancel()

	mux := runtime.NewServeMux()	// 路由处理
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.GrpcEndpoint, opts)
	if err != nil {
		fmt.Println(48,err.Error())
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods:[]string{
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPost,
			http.MethodGet,
			http.MethodHead,
			http.MethodPatch,
			http.MethodPut,
		},
		AllowedHeaders:[]string{"*"},
		AllowCredentials:false,
	})
	handler := c.Handler(mux)
	return http.ListenAndServe(fmt.Sprintf("%s",config.HttpHost+":"+config.HttpPort),
		setFileServer(fileServer, wsproxy.WebsocketProxy(handler)))
}

func setFileServer(fileServer, other http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := conf.Config
		prefix := config.Prefix
		if len(prefix) > 1 {
			prefix = prefix[:len(prefix) - 1]
		}
		if strings.HasPrefix(r.RequestURI, prefix+"/statics/"){
			fileServer.ServeHTTP(w,r)
		} else {
			other.ServeHTTP(w,r)
		}
	})
}

func run() error {
	var config = conf.Config
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, config.GrpcEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.HttpHost + ":" + config.HttpPort, mux)

}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

//func main(){
//	wg := new (sync.WaitGroup)
//	wg.Add(2)
//
//	go func() {
//		err := startGrpcServer()
//		if err != nil {
//			log.Fatal("startGrpcServer:", err)
//		}
//		wg.Done()
//	}()
//
//	go func() {
//		err := startHttpServer()
//		if err != nil {
//			log.Fatal("startHttpServer:",err)
//		}
//		wg.Done()
//	}()
//	wg.Wait()
//}
