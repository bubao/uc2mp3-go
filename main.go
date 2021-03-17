/*
 * @description:
 * @author: bubao
 * @date: 2021-03-14 17:53:48
 * @last author: bubao
 * @last edit time: 2021-03-17 18:30:05
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func uc2mp3(ucFilename string, mp3Filename string) {

	buf, err := ioutil.ReadFile(ucFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	arr := make([]uint8, len(buf))
	for i, v := range buf {
		arr[i] = v ^ 0xA3
	}
	err = ioutil.WriteFile(mp3Filename, arr, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	var dirName string
	var outputDir string
	var inputFilename string
	var rename string
	flag.StringVar(&dirName, "d", "", "输入文件夹")
	flag.StringVar(&outputDir, "o", "", "输出文件夹")
	flag.StringVar(&inputFilename, "f", "", "uc!文件")
	flag.StringVar(&rename, "r", "", "mp3 重命名")
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		return
	}
	var ext string = ".uc!"

	if inputFilename == "" && dirName != "" {
		stat, err := os.Stat(dirName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
			panic(err.Error())
		}
		if stat.IsDir() {
			dir, err := ioutil.ReadDir(dirName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
				panic(err.Error())
			}

			for _, info := range dir {
				filename := info.Name()

				if !info.IsDir() && path.Ext(filename) == ext {
					var outputFile string
					outputFile = strings.TrimSuffix(path.Base(filename), ext)
					uc2mp3(path.Join(dirName, filename), path.Join(outputDir, outputFile))
				}
			}
		}
	} else if inputFilename != "" {
		if path.Ext(inputFilename) == ext {
			var outputFile string
			if rename != "" {
				outputFile = rename
			} else {
				outputFile = strings.TrimSuffix(path.Base(inputFilename), ext)
			}
			uc2mp3(inputFilename, path.Join(outputFile))
		}
	}
}
