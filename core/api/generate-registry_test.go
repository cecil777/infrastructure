package api

import (
	"go/ast"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile() {
	err := os.MkdirAll("api/open", os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile("api/open/deploy.go", []byte(`package open

type Deploy struct {
}

func (m Deploy) Call() (interface{}, error) {
	return nil, nil
}`), os.ModePerm)
	if err != nil {
		log.Panic(err)
	}

}
func TestGenerateRegistry(t *testing.T) {
	createTempFile()

	t.Run("TestGenerateRegistry", func(t *testing.T) {
		err := GenerateRegistry()

		assert.NoError(t, err, "执行GenerateRegistry()出错")

		b, err := ioutil.ReadFile("api/metadata.go")

		assert.NoError(t, err, "没有找到api/metadata.go文件，GenerateRegistry执行失败")

		actual := `package api
import (
    "gitlab.dev.daaokeji.com/infrastructure/core/api"
    
    "gitlab.dev.daaokeji.com/infrastructure/core/api/api/open"
)

func init() {
	api.Register("open", "deploy", open.Deploy{})
    
}
`

		assert.Equal(t, string(b), actual)

	})

	assert.NoError(t, os.RemoveAll("api"))

}

func Test_projectApiPath(t *testing.T) {

	type args struct {
		wd     string
		gopath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test_projectApiPath_win", args: args{
			wd:     "C:\\Users\\daao-010\\go\\src\\gitlab.dev.daaokeji.com\\infrastructure\\core",
			gopath: "C:\\Users\\daao-010\\go",
		}, want: "gitlab.dev.daaokeji.com/infrastructure/core/api"},

		{name: "Test_projectApiPath_linux", args: args{
			wd:     "/home/user/go/src/gitlab.dev.daaokeji.com/infrastructure/core",
			gopath: "/home/user/go",
		}, want: "gitlab.dev.daaokeji.com/infrastructure/core/api"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := projectApiPath(tt.args.wd, tt.args.gopath)
			assert.Equal(t, got, tt.want)

		})
	}
}

func Test_parseGoFile(t *testing.T) {

	type args struct {
		endpoint   string
		goFileName string
		objs       map[string]*ast.Object
	}
	tests := []struct {
		name string
		args args
		want []metadata
	}{
		{name: "Test_parseGoFile", args: args{
			endpoint:   "open",
			goFileName: "deploy",
			objs:       map[string]*ast.Object{"DeployStruct": &ast.Object{Name: "DeployStruct"}},
		}, want: []metadata{{
			Endpoint: "open",
			Name:     "deploy",
			RegName:  "DeployStruct",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseGoFile(tt.args.endpoint, tt.args.goFileName, tt.args.objs)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_writeGoFile(t *testing.T) {
	createTempFile()

	type args struct {
		list   []metadata
		actual string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test_writeGoFile", args: args{list: []metadata{
			{
				Endpoint: "open",
				Name:     "deploy",
				RegName:  "DeployStruct",
			},
		},
			actual: `package api
import (
    "gitlab.dev.daaokeji.com/infrastructure/core/api"
    
    "gitlab.dev.daaokeji.com/infrastructure/core/api/api/open"
)

func init() {
	api.Register("open", "deploy", open.DeployStruct{})
    
}
`,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := writeGoFile(tt.args.list)
			assert.NoErrorf(t, err, "写入api/metadata.go出错")

			b, err := ioutil.ReadFile("api/metadata.go")

			assert.NoErrorf(t, err, "没有找到api/metadata.go文件，writeGoFile执行失败")

			assert.Equal(t, string(b), tt.args.actual)
		})
	}

	assert.NoError(t, os.RemoveAll("api"))
}