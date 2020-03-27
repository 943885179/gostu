# 打包全部平台(会重命名对应的生成文件)
# 安装插件`go get github.com/mitchellh/gox`
# 打包全部`gox`
# 打包在不同文件夹下`gox -output "{{.Dir}}_{{.OS}}_{{.Arch}}/swbatch`
# 只打包Linux`gox -os="linux"` 只打包Linux Bit64` gox -osarch="linux/amd64"`
# Mac 下打包 Linux 环境下运行的程序 `env GOOS=linux GOARCH=386 go build main.go`

# windows打包
>cmd到目录下执行`go build`后会生成对应的exe;
