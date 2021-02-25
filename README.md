# GoStudy
我的Go语言学习仓库


# GOPATH与工作空间
1、只要进入对应的应用包目录，然后执行go install，就可以安装了
2、在任意的目录执行如下代码go install mymath (生成.a文件)

安装完之后，我们可以进入如下目录

cd $GOPATH/pkg/${GOOS}_${GOARCH}
//可以看到如下文件
mymath.a  （.a文件是应用包）

bin目录下面存的是编译之后可执行的文件，pkg下面存放的是应用包，src下面保存的是应用源代码

# 
|——main.go         入口文件
|——conf            配置文件和处理模块
|——controllers     控制器入口
|——models          数据库处理模块
|——utils           辅助函数库
|——static          静态文件目录
|——views           视图库
