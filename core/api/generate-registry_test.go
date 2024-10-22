package api

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"os"
	"testing"
)

func TestGenerateRegistry(t *testing.T) {
	t.Run("TestGenerateRegistry", func(t *testing.T) {
		os.Mkdir("api", os.ModePerm)
		GenerateRegistry()
		_, err := os.Stat("api/metadata.go")
		if err != nil {
			assert.Error(t, err, "没有找到api/metadata.go文件，GenerateRegistry执行失败")
		}
		os.RemoveAll("api")

	})
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
	type args struct {
		list []metadata
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
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Mkdir("api", os.ModePerm)
			writeGoFile(tt.args.list)
			os.RemoveAll("api")
		})
	}
}
