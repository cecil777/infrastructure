package api

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type metadata struct {
	Endpoint string
	Name     string
	RegName  string
}

var metadataTemplate = `package api
{{ $ProjectPath:=.ProjectPath -}}

import (
    "gitlab.dev.daaokeji.com/infrastructure/core/api"
    {{range $k,$v:=.List}}
    "{{$ProjectPath}}/{{$v.Endpoint}}"
    {{- end}}
)

func init() {
	{{range $k,$v:=.List -}}
    api.Register("{{$v.Endpoint}}", "{{$v.Name}}", {{$v.Endpoint}}.{{$v.RegName}}{})
    {{end}}
}
`

// 描述dir下的api定义，使用tpl模板，写入到wDir目录下
func GenerateRegistry() {

	rootFile, err := os.Open("api")
	if err != nil {
		log.Panic(err)
	}

	dirList, err := rootFile.ReadDir(0)
	if err != nil {
		log.Panic(err)
	}

	list := make([]metadata, 0)
	for index := range dirList {

		fileInfo, err := dirList[index].Info()
		if err != nil {
			log.Panic(err)
		}

		if fileInfo.IsDir() {
			goFileList, err := os.ReadDir(rootFile.Name() + "/" + fileInfo.Name())
			if err != nil {
				log.Panic(err)
			}

			for goIndex := range goFileList {

				if goFileList[goIndex].IsDir() == false {

					fSet := token.NewFileSet()
					goFileName := goFileList[goIndex].Name()
					if strings.ContainsAny(goFileName, ".go") {
						f, err := parser.ParseFile(fSet, rootFile.Name()+"/"+fileInfo.Name()+"/"+goFileName, nil, 0)
						if err != nil {
							log.Panic(err)
						}
						infos := parseGoFile(fileInfo.Name(), strings.Split(goFileName, ".")[0], f.Scope.Objects)
						list = append(list, infos...)
					}

				}
			}
		}

	}
	writeGoFile(list)
}

// 写入到指定目录下的metadata.go文件
func writeGoFile(list []metadata) {

	gopath := os.Getenv("GOPATH")

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	t, err := template.New("").Parse(metadataTemplate)
	if err != nil {
		log.Panic(err)
	}
	buffer := bytes.NewBuffer(nil)
	err = t.Execute(buffer, map[string]interface{}{"List": list, "ProjectPath": projectApiPath(wd, gopath)})
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile("api/metadata.go", buffer.Bytes(), fs.ModePerm)
	if err != nil {
		log.Panic(err)
	}

}

// 分析go文件
func parseGoFile(endpoint string, goFileName string, objs map[string]*ast.Object) []metadata {
	l := len(objs)
	outList := make([]metadata, 0, l)
	for key := range objs {
		outList = append(outList, metadata{
			Endpoint: endpoint,
			Name:     goFileName,
			RegName:  objs[key].Name,
		})
	}
	return outList
}

// 获取项目路径
func projectApiPath(wd, gopath string) string {
	dir := strings.ReplaceAll(wd, gopath, "")
	dir = strings.ReplaceAll(filepath.ToSlash(dir), "/src/", "") + "/api"
	return dir
}
