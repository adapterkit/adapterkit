package generate

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"moul.io/climan"
)

func Generate() *climan.Command {
	var moduleRepo string
	var outputDir string
	var templateDir string

	return &climan.Command{
		Name:       "generate",
		ShortHelp:  "Generate code, you must specify the path to your protoc bin in the environment variable PROTOC_BIN",
		ShortUsage: "adapterkit generate [global flags] [flags] [args]",
		FlagSetBuilder: func(fs *flag.FlagSet) {
			fs.StringVar(&moduleRepo, "mod", "", "github repo where the module is located")
			fs.StringVar(&outputDir, "out", ".", "output directory")
			fs.StringVar(&templateDir, "tpl", "template", "template directory")
		},
		Exec: func(_ context.Context, strings []string) error {
			protocBin := os.Getenv("PROTOC_BIN")
			cmd := exec.Command(protocBin, "-I.", fmt.Sprintf("--gotemplate_out=template_dir=%s,debug=true:%s", templateDir, outputDir), strings[0])
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Println("error: ", err)
			}

			currentModule := getModule(outputDir)

			err = fillImport(outputDir, moduleRepo, currentModule)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func getModule(path string) string {
	dir, err := os.ReadDir("./" + path)
	if err != nil {
		return ""
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if file.Name() == "go.mod" {
			content, err := os.ReadFile(path + "/" + file.Name())
			if err != nil {
				return ""
			}

			reg := regexp.MustCompile(`module (.*)`)

			return reg.FindStringSubmatch(string(content))[1]
		}
	}

	return ""
}

func fillImport(path, logicPackage, currentModule string) error {
	dir, err := os.ReadDir("./" + path)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if file.IsDir() {
			err := fillImport(path+"/"+file.Name(), logicPackage, currentModule)
			if err != nil {
				return err
			}
			continue
		}

		content, err := os.ReadFile(path + "/" + file.Name())
		if err != nil {
			return err
		}

		reg := regexp.MustCompile(`\$\[ADAPTERKIT_GOMOD]`)
		content = reg.ReplaceAll(content, []byte(currentModule))

		reg = regexp.MustCompile(`\$\[ADAPTERKIT_LOGIC_PACKAGE]`)
		content = reg.ReplaceAll(content, []byte(logicPackage))

		err = os.WriteFile(path+"/"+file.Name(), content, 0o600)
		if err != nil {
			return err
		}
	}

	return nil
}
