package schema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattfenwick/algorithms/pkg/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Run(rootPath string) {
	t := NewTraverser()
	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			logrus.Debugf("skipping dir %s", path)
		} else if filepath.Ext(path) == ".json" {
			logrus.Debugf("found json path: %s, %t, %s, %t, %t\n", path, d.IsDir(), d.Name(), d.Type().IsDir(), d.Type().IsRegular())
			bs := utils.Die(os.ReadFile(path))
			var v any
			utils.Die0(json.Unmarshal(bs, &v))
			t.Add(v)
		} else if filepath.Ext(path) == ".jsonl" {
			for _, line := range strings.Split(string(utils.Die(os.ReadFile(path))), "\n") {
				if strings.TrimSpace(line) == "" {
					continue
				}
				var v any
				utils.Die0(errors.Wrapf(json.Unmarshal([]byte(line), &v), "unable to read jsonl from %s", path))
				t.Add(v)
			}
		} else {
			logrus.Debugf("skipping non-json file %s\n", path)
		}
		return nil
	})
	fmt.Println(strings.Join(t.Lines(), "\n"))
}
