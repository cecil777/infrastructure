/*
生成api相关的注册信息，每个项目需要在main函数里添加api.GenerateRegistry()方法。
第一次运行，生成相应的api metadata.go文件，第二次运行会让metadata.go文件生效。

在api.GenerateRegistry()方法时，其项目的目录结构必须以gopath目录结构进行组织，否则api metadata数据将无效生成。
*/
package api

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Metadata 模板数据结构
type metadata struct {
	Endpoint string //端点结构
	Name     string //api 名字
	RegName  string //api 处理struct的名
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
func GenerateRegistry() error {

	rootFile, err := os.Open("api")
	if err != nil {
		return err
	}

	dirList, err := rootFile.ReadDir(0)
	if err != nil {
		return err
	}

	list := make([]metadata, 0)
	for index := range dirList {

		fileInfo, err := dirList[index].Info()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			goFileList, err := os.ReadDir(rootFile.Name() + "/" + fileInfo.Name())
			if err != nil {
				return err
			}

			for goIndex := range goFileList {

				if goFileList[goIndex].IsDir() == false {

					fSet := token.NewFileSet()
					goFileName := goFileList[goIndex].Name()
					if strings.ContainsAny(goFileName, ".go") {
						f, err := parser.ParseFile(fSet, rootFile.Name()+"/"+fileInfo.Name()+"/"+goFileName, nil, 0)
						if err != nil {
							return err
						}
						infos := parseGoFile(fileInfo.Name(), strings.Split(goFileName, ".")[0], f.Scope.Objects)
						list = append(list, infos...)
					}

				}
			}
		}

	}

	return writeGoFile(list)
}

// 写入到指定目录下的metadata.go文件
func writeGoFile(list []metadata) error {

	gopath := os.Getenv("GOPATH")

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	t, err := template.New("").Parse(metadataTemplate)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(nil)
	err = t.Execute(buffer, map[string]interface{}{"List": list, "ProjectPath": projectApiPath(wd, gopath)})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("api/metadata.go", buffer.Bytes(), fs.ModePerm)
	if err != nil {
		return err
	}
	return nil
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