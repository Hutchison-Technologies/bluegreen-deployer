package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_DeploymentName_Returns_ValidAppName(t *testing.T) {
	assert.True(t, IsValidAppName(DeploymentName("prod", "blue", "some-api")))
}

func Test_DeploymentName_Returns_Name_PrefixedWith_TargetEnv(t *testing.T) {
	targetEnv := "prod"
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("^%s-.*", targetEnv)), DeploymentName(targetEnv, "blue", "some-api"))
}

func Test_DeploymentName_Returns_Name_Containing_AppName(t *testing.T) {
	appName := "some-api"
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf(".*-%s.*", appName)), DeploymentName("prod", "blue", appName))
}

func Test_DeploymentName_Returns_Name_Containing_Colour(t *testing.T) {
	colour := "green"
	assert.Regexp(t, regexp.MustCompile(fmt.Sprintf(".*-%s.*", colour)), DeploymentName("prod", colour, "some-api"))
}