// tar包实现了文件的打包功能,可以将多个文件或者目录存储到单一的.tar压缩文件中
// tar本身不具有压缩功能,只能打包文件或目录
package main

import (
	//"fmt" // 输入输出包
	"archive/tar" //tar包
	"io"          //数据流操作
	"log"         //日志
	"os"          //文件操作
)

func main() {
	TarTest()
	UnTarTest()
}

// 测试压缩
func TarTest() {
	dst := "G:\\Go学习\\assets\\test.tar" //压缩路径
	files := []string{"G:\\Go学习\\assets\\test.txt", "G:\\Go学习\\assets\\ee.txt"}
	if err := Tar(files, dst); err != nil {
		log.Println(err)
	}
}

// 压缩为tar
func Tar(src []string, dst string) error {
	//创建tar文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()
	//通过fw创建一个tar.Writer
	tw := tar.NewWriter(fw)
	//如果失败关闭造成tar包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			log.Println(err)
		}
	}()
	//变量读取文件
	for _, fileName := range src {
		fi, err := os.Stat(fileName)
		if err != nil {
			log.Println(err)
			continue
			//return err
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		//将tar的文件信息hdr写入到tw
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
		//写入文件数据
		fs, err := os.Open(fileName)
		if err != nil {
			return err
		}
		if _, err = io.Copy(tw, fs); err != nil {
			return err
		}
		defer fs.Close()
	}
	return nil
}
func UnTarTest() {
	UnTar("G:\\Go学习\\assets\\test.tar", "G:\\Go学习\\assets\\")
}

//解压tar文件
func UnTar(tarFileName string, tarToFile string) {
	//打开tar
	fr, err := os.Open(tarFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fr.Close()
	//读取tar文件
	tr := tar.NewReader(fr)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			log.Println(err)
			continue
		}
		// 读取文件信息
		fi := hdr.FileInfo()
		//创建空白文件，解包写入数据
		fw, err := os.Create(tarToFile + fi.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		if _, err := io.Copy(fw, tr); err != nil {
			log.Println(err)
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		defer fw.Close()
	}
}
