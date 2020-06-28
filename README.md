# go-micro-foobar
花了一天时间 终于把文档中的例子搭了出来 ，一步一个坑 ，不容易。
记录一下才过的坑 ：
1.依赖问题
仁者见仁智者见智 不论是科学上网 还是其他渠道 都可以 。
2.go.mod问题 相当于maven的pom文件
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
这个地方的意思是项目中引用到replace google.golang.org/grpc  这个包的内容的话 会去r => 后面的地方拉取依赖
不是科学上网的话 拉到死 也拉不下来 ，项目就会标红。
解决办法是：
找到引用包爆红的地方 把包名替换成本地，例如：
replace google.golang.org/protobuf => C:/Users/admin/go/src/google.golang.org/protobuf


require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.8.0

)
##go.mod文件中这个地方会自动识别本地，和maven pom文件一样的

##go mod 的适用方法，个人实践，最好的适用方法是新建一个项目，然后 go mod init  然后 go mod tidy 
##go mod 会把依赖包下载下来存到一个目录，新项目的时候 会加载本地存好的依赖，完全可以理解是pom 文件，加载本地jar包


#一步一坑 加油！
