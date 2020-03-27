package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	// ZipTest()
	UnZipTest()
}

//测试文件压缩
func ZipTest() {
	err := Zip("G:\\Go学习\\assets\\test.zip", []string{"G:\\Go学习\\assets\\test.txt", "G:\\Go学习\\assets\\ee.txt"})
	if err != nil {
		log.Println(err)
	}
}

//文件压缩
func Zip(zipFile string, fileList []string) error {
	//创建zip文件
	fw, err := os.Create(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()
	//实例化新的Zip.Writer
	zw := zip.NewWriter(fw)
	defer func() {
		//检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()
	for _, fileName := range fileList {
		fr, err := os.Open(fileName)
		if err != nil {
			log.Println(err)
			return err
		}
		fi, err := fr.Stat()
		if err != nil {
			return err
		}
		//写入文件头信息
		fh, err := zip.FileInfoHeader(fi)
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return nil
		}
		//写入内容
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
		defer fr.Close()
	}
	return nil
}
func UnZipTest() {
	err := UnZip("G:\\Go学习\\assets\\test.zip", "G:\\Go学习\\assets\\")
	if err != nil {
		log.Fatal(err)
	}
}

//文件解压
func UnZip(zipFile string, zipOutPath string) error {
	zr, err := zip.OpenReader(zipFile)
	defer zr.Close()
	if err != nil {
		return err
	}
	for _, file := range zr.File {
		//如果是目录则创建目录
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(file.Name, file.Mode()); err != nil {
				return err
			}
			continue
		}
		//获取Reader
		fr, err := file.Open()
		if err != nil {
			return nil
		}
		fw, err := os.OpenFile(zipOutPath+file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		if _, err = io.Copy(fw, fr); err != nil {
			return err
		}
		defer fw.Close()
		defer fr.Close()
	}
	return nil
}
