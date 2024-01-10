# prorobuf管理中心
## 使用方式
在api目录下新建各个server的目录，然后版本目录，下面放protobuf文件，可参考lotterysvr和shortUrlXsvr
每次新增或者更新完protobuf文件后，在项目根目录执行


```make api```

更新完pb之后，打tag，tag以v1.0.0类似的形势命名，打完tag推送到远程仓库即可

```git tag ${版本号}```

```git push --tags```

## 其他仓库引用
在其他仓库引用proto_center

1. 初次使用
```go get github.com/BitofferHub/pkg```
2. 更新完pb之后，更新版本，只需要在go.mod文件中修改引用的
```github.com/BitofferHub/pkg```的版本号为最新打的tag，然后执行
```go mod tidy```即可
