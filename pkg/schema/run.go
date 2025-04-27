package schema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/sirupsen/logrus"
)

func Run(rootPath string) {
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			logrus.Debugf("skipping dir %s", path)
		} else if filepath.Ext(path) == ".json" {
			logrus.Debugf("found json path: %s, %t, %s, %t, %t\n", path, d.IsDir(), d.Name(), d.Type().IsDir(), d.Type().IsRegular())
			bs := utils.Die(os.ReadFile(path))
			var v any
			utils.Die0(json.Unmarshal(bs, &v))
			fmt.Println(Traverse(v))
		} else {
			logrus.Debugf("skipping non-json file %s\n", path)
		}
		return nil
	})
}
