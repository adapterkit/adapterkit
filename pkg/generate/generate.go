package generate

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"moul.io/climan"
	"moul.io/u"
)

func Cmd() *climan.Command {
	var moduleRepo string
	var outputDir string
	var templateDir string
	var protocBin string

	return &climan.Command{
		Name:       "generate",
		ShortHelp:  "Generate code",
		ShortUsage: "adapterkit generate [global flags] [flags] [args]",
		FlagSetBuilder: func(fs *flag.FlagSet) {
			fs.StringVar(&moduleRepo, "mod", "", "github repo where the module is located")
			fs.StringVar(&outputDir, "out", ".", "output directory")
			fs.StringVar(&templateDir, "tpl", "template", "template directory")
			fs.StringVar(&protocBin, "protoc-bin", "protoc", "path to the 'protoc' binary")
		},
		Exec: func(_ context.Context, args []string) error {
			if len(args) != 1 {
				return flag.ErrHelp
			}
			if !u.CommandExists(protocBin) {
				return fmt.Errorf("protoc binary not found: %s", protocBin) //nolint:goerr113
			}

			gotemplateOpts := fmt.Sprintf("--gotemplate_out=template_dir=%s,debug=true:%s", templateDir, outputDir)
			protoPath := args[0]
			cmd := exec.Command(protocBin, "-I.", gotemplateOpts, protoPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Println("error: ", err)
			}

			currentModule, err := getModule(outputDir)
			if err != nil {
				return err
			}

			templateDirName, err := getTemplateDirName(templateDir)
			if err != nil {
				return err
			}

			err = fillImport(outputDir, moduleRepo, currentModule, templateDirName)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func getTemplateDirName(path string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("empty path") //nolint:goerr113
	}
	tabPath := strings.Split(path, "/")
	templateDirName := tabPath[len(tabPath)-1]

	return templateDirName, nil
}

var errGetModule = errors.New("can't get module")

func getModule(path string) (string, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if file.Name() == "go.mod" {
			content, err := os.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				return "", err
			}

			reg := regexp.MustCompile(`module (.*)`)

			match := reg.FindStringSubmatch(string(content))
			if len(match) < 2 { //nolint:gomnd
				return "", fmt.Errorf("%w: can't find module name", errGetModule)
			}

			return match[1], nil
		}
	}

	return "", fmt.Errorf("%w: can't find go.mod file", errGetModule)
}

func fillImport(path, logicPackage, currentModule, templateDirName string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range dir {
		filePath := filepath.Join(path, file.Name())
		if file.Name()[0] == '.' || file.Name() == templateDirName {
			continue
		}

		if file.IsDir() {
			err := fillImport(filePath, logicPackage, currentModule, templateDirName)
			if err != nil {
				return err
			}
			continue
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("can't read file: %w", err)
		}

		reg := regexp.MustCompile(`\$\[ADAPTERKIT_GOMOD]`)
		content = reg.ReplaceAll(content, []byte(currentModule))

		reg = regexp.MustCompile(`\$\[ADAPTERKIT_LOGIC_PACKAGE]`)
		content = reg.ReplaceAll(content, []byte(logicPackage))

		err = os.WriteFile(filePath, content, 0o600) //nolint:gomnd
		if err != nil {
			return fmt.Errorf("can't write in file: %w", err)
		}
	}

	return nil
}
