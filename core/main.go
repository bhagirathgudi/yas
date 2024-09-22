package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

const (
	DefaultTemplateType = "GITLAB"
)
func main() {
	//Create tmp directory
	dir := CreateAndChngeDir()
	fmt.Println("%s", dir)
	//Perform git clone
	repo := PerformCloneAndInit("https://github.com/bhagirathgudi/dashcon.git")
	//Create a branch
	branchPrefixes := []string{"scaffold", CreateRandomString(5)}
	CreateBranch(strings.Join(branchPrefixes, "-"), repo)
	//Load the plugin
	pluginName := os.Getenv("TEMPLATE_SRC_TYPE") ?: DefaultTemplateType
	//Download template/templates
	//Replace place holders
	//Perform git commit and push
	log.Println(os.Getwd())
	GitPush(repo)
	//Clean up directory
	CleanUpDir(dir)
}