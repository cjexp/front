package main

import (
	"os"
	"path/filepath"

	"github.com/cjexp/base/task"
	"github.com/cjtoolkit/taskforce"
)

func initTask() *taskforce.TaskForce {
	var (
		tf             = taskforce.InitTaskForce()
		yarnRun        = task.YarnRun(tf)
		devEnvironment = true
	)

	tf.Register("yarn", func() {
		tf.ExecCmd("gnode", "yarn")
		//tf.ExecCmd("./createLink")
	})

	tf.Register("clean", func() {
		os.RemoveAll("live")

		os.Mkdir("live", 0755)
	})

	{
		type sass struct {
			dest string
			src  string
		}

		destSrc := []sass{
			{
				dest: "live/stylesheets/styles.css",
				src:  "dev/sass/styles.scss",
			},
		}

		tf.Register("sass", func() {
			args := []string{"--source-map", "--embed-sources", "--style=compressed"}
			for _, v := range destSrc {
				args := append(args, v.src, v.dest)
				yarnRun("sass", args...)
			}
		})
	}

	tf.Register("rollup", func() {
		env := "BUILD:production"
		if devEnvironment {
			env = "BUILD:development"
		}
		yarnRun("rollup", "-c", "-m", "--environment", env)
	})

	tf.Register("copy", func() {
		copyFolder := task.CopyFolder(tf)

		copyFolder("live/fonts", "module/fontawesome/webfonts")
	})

	tf.Register("zip", func() {
		zipUtil := task.Zip(tf)
		zipUtil("../asset.zip", "live")
	})

	tf.Register("dev", func() {
		tf.Run("yarn", "clean", "sass", "rollup", "copy")
	})

	tf.Register("prod", func() {
		devEnvironment = false
		tf.Run("yarn", "clean", "sass", "rollup", "copy", "zip")
	})

	tf.Register("quick:sass", func() {
		os.RemoveAll(filepath.FromSlash("live/stylesheets"))
		tf.Run("sass")
	})

	tf.Register("quick:js", func() {
		os.RemoveAll(filepath.FromSlash("live/javascript"))
		tf.Run("rollup")
	})

	return tf
}

func main() {
	initTask().Run(os.Args[1:]...)
}
