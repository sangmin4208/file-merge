package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)


func main(){
	path := "./output/202101"
	// InPutPath(&path)
	files,err := ioutil.ReadDir(path)
	 if err != nil {
		 log.Fatal(err, "\n경로를 찾을 수 없습니다.")
	 }
	 contents := [][]byte{}
	 for _,file := range files {
		fileName := file.Name()
		content := readFile(path, fileName)
		contents = append(contents,content)
	 }
	 writeFile(path, merge(contents))
}

func InPutPath(path *string){
	fmt.Println("경로 입력")
	fmt.Scanln(path)
}

func readFile(path string,fileName string) []byte{
	f,_:= os.ReadFile(path + "/" + fileName)
	return f
}

func merge(contents [][]byte) []byte{
	merged := []byte{}

	for _,content := range contents{
		fmt.Println(string(content))
		content = []byte(string(content) + "\n")
		merged = append(merged,content...)
	}
	return merged
}
func writeFile(path string, content []byte){
	os.WriteFile(path+"/merged.txt", content,fs.ModeAppend)
}