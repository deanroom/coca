package unused

import (
	. "../../adapter/models"
	. "../../utils"
	. "../../utils/models"
	. "../base/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var parsedChange []RefactorChangeRelate
var nodes []JMoveStruct

type RemoveMethodApp struct {
}

var depsFile string
var configPath string
var conf string
var parsedDeps []JClassNode

func RenameMethodApp(dep string, p string) *RemoveMethodApp {
	nodes = nil
	depsFile = dep
	configPath = p
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Start() {
	file := ReadFile(depsFile)
	if file == nil {
		return
	}

	_ = json.Unmarshal(file, &parsedDeps)

	configBytes := ReadFile(configPath)
	if configBytes == nil {
		return
	}

	conf = string(configBytes)

	parsedChange = ParseRelates(conf)

	startParse(parsedDeps, parsedChange)
}

func startParse(nodes []JClassNode, relates []RefactorChangeRelate) {
	for _, related := range relates {
		oldInfo := BuildMethodPackageInfo(related.OldObj)
		//newInfo := BuildMethodPackageInfo(related.NewObj)

		for _, pkgNode := range nodes {
			fmt.Println(pkgNode.Package+pkgNode.Class, oldInfo.Package+oldInfo.Class)
			if pkgNode.Package+pkgNode.Class == oldInfo.Package+oldInfo.Class {
				for _, method := range pkgNode.Methods {
					updateSelfRefs(pkgNode, method)
				}
			}
			//
			//for methodCall := range pkgNode.MethodCalls {
			//
			//}
		}
	}
}

func updateSelfRefs(node JClassNode, method JMethod) {
	path := node.Path
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if i == method.StartLine-1 {
			fmt.Println(line[method.StartLinePosition:])
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func updateDepsRefs(node JClassNode, info *PackageClassInfo, info2 *PackageClassInfo) {

}
