package analyzer

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type NodeJSDetector struct{}

func NewNodeJSDetector() *NodeJSDetector {
	return &NodeJSDetector{}
}

type packageJSON struct {
	Scripts map[string]string `json:"scripts"`
}

func (d *NodeJSDetector) Detect(directory string) (*ProjectInfo, bool) {
	packageJSONPath := filepath.Join(directory, "package.json")

	if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
		return nil, false
	}

	file, err := ioutil.ReadFile(packageJSONPath)
	if err != nil {
		return nil, false
	}

	var pkg packageJSON
	if err := json.Unmarshal(file, &pkg); err != nil {
		return nil, false
	}

	info := &ProjectInfo{
		Language:       "Node.js",
		InstallCommand: "npm install",
		DockerImage:    "node:18-alpine",
	}

	if _, ok := pkg.Scripts["build"]; ok {
		info.BuildCommand = "npm run build"
	}
	if _, ok := pkg.Scripts["test"]; ok {
		info.TestCommand = "npm run test"
	}

	yarnLockPath := filepath.Join(directory, "yarn.lock")
	if _, err := os.Stat(yarnLockPath); err == nil {
		info.InstallCommand = "yarn install"
		if info.BuildCommand != "" {
			info.BuildCommand = "yarn build"
		}
		if info.TestCommand != "" {
			info.TestCommand = "yarn test"
		}
	}

	return info, true
}
