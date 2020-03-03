package factory

import (
	"github.com/spf13/viper"
	"github.com/jukylin/esim/pkg/file-dir"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/imports"
)


func getCurDir() string {
	modelpath, err := os.Getwd()
	if err != nil {
		println(err.Error())
	}

	return modelpath
}


func delModelFile() {
	os.Remove(getCurDir() + "/plugin/model.go")
	os.Remove(getCurDir() + "/plugin/model_test.go")
}


func TestFindModel(t *testing.T) {
	modelName := "Test"
	modelPath := getCurDir() + "/example"

	info, err := FindModel(modelPath, modelName, "")
	if err != nil {
		t.Error(err.Error())
		return
	}
	if info.packName == "" {
		t.Error("error")
		return
	}
}


func TestBuildVirEnv(t *testing.T) {

	modelName := "Test"
	modelPath := getCurDir() + "/example"

	info, err := FindModel(modelPath, modelName, "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = BuildPluginEnv(info, delModelFile)
	if err != nil {
		t.Error(err.Error())
		return
	}

	e, err := file_dir.IsExistsDir("./example/plugin")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if e == false {
		t.Error("plugin 创建失败")
		return
	}
	Clear(info)
}


func TestExecPlugin(t *testing.T) {

	modelName := "Test"
	modelPath := getCurDir() + "/example"

	v := viper.New()
	v.Set("sort", false)
	v.Set("pool", false)

	info, err := FindModel(modelPath, modelName, "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = BuildPluginEnv(info, delModelFile)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = ExecPlugin(v, info)
	if err != nil {
		t.Error(err.Error())
		return
	}

	Clear(info)
}


func TestFinalContent(t *testing.T) {

	result := `package example

type Test struct {
	c int8

	i bool

	g byte

	d int16

	f float32

	a int32

	b int64

	m map[string]interface{}

	e string

	h []int

	u [3]string
}

type Tests []Test
`

	modelName := "Test"
	modelPath := getCurDir() + "/example"

	v := viper.New()
	v.Set("sort", true)
	v.Set("pool", false)
	v.Set("coverpool", false)
	v.Set("plural", false)

	info, err := FindModel(modelPath, modelName, getPluralWord(modelName))
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = BuildPluginEnv(info, delModelFile)
	if err != nil{
		t.Error(err.Error())
		return
	}

	err = ExecPlugin(v, info)
	if err != nil {
		t.Error(err.Error())
		return
	}

	BuildFrame(v, info)

	src, err := ReplaceContent(v, info)
	if err != nil {
		t.Error(err.Error())
		return
	}

	res, err := imports.Process("", []byte(src), nil)
	if err != nil {
		t.Error(err.Error())
		return
	}

	assert.Equal(t, result, string(res))

	Clear(info)
}


func TestClear(t *testing.T) {
	modelName := "Test"
	modelPath := getCurDir() + "/example"

	info, err := FindModel(modelPath, modelName, "")
	if err != nil {
		t.Error(err.Error())
	}
	Clear(info)
}


func TestNewFrame(t *testing.T)  {

	v := viper.New()
	v.Set("gen_logger_option", true)
	v.Set("gen_conf_option", true)
	v.Set("star", true)

	info := &BuildPluginInfo{}
	info.modelName = "TestFrame"

	NewVarStr(v, info)
	frame := NewFrame(v, info)
	getOptions(v, info)

	newFrame := replaceFrame(frame, info)
	assert.Empty(t, newFrame)
}


func TestGetNewImport(t *testing.T)  {

	result := `import (
        test
        test2
)
`

	imports := []string{"test", "test2"}

	newImport := getNewImport(imports)

	assert.Equal(t,  result, newImport)
}


func TestExtendField(t *testing.T)  {
	modelName := "Test"
	modelPath := getCurDir() + "/example"

	v := viper.New()
	v.Set("gen_logger_option", true)
	v.Set("gen_conf_option", true)
	v.Set("option", true)

	info, err := FindModel(modelPath, modelName, getPluralWord(modelName))
	if err != nil {
		assert.Nil(t, err)
	}

	if ExtendField(v, info) {
		err := ReWriteModelContent(info)
		assert.Nil(t, err)
	}else{
		assert.Fail(t, "not here")
	}
}