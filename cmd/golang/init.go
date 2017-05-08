package golang

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Template(outputDir string, loc string, name string) error {
	fmt.Println("what is your import package? Example github.com/maleck13/test")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tData := map[string]interface{}{}
	tData["root_package"] = scanner.Text()
	tData["app"] = name
	dirs := []string{}
	templates := []string{}
	files := []string{}
	//Walk the filepath of the template to use to setup our new app. We sort into dirs,templates and regular files
	filepath.Walk(loc, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			relative := strings.Replace(path, loc, "", 1)
			dirs = append(dirs, filepath.Join(outputDir, relative))
			return nil
		}
		name := info.Name()
		if name[len(name)-4:] == ".tpl" {
			templates = append(templates, path)
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err := mkdirs(dirs); err != nil {
		return err
	}
	if err := writeTemplatedFiles(tData, loc, outputDir, templates); err != nil {
		return err
	}
	if err := writeRegularFiles(loc, outputDir, files); err != nil {
		return err
	}
	return nil
}

func mkdirs(dirs []string) error {
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}
	return nil
}

func writeTemplatedFiles(tdata map[string]interface{}, loc, outpath string, ts []string) error {
	for _, tloc := range ts {
		t, err := template.ParseFiles(tloc)
		if err != nil {
			return err
		}

		outFileName := tloc[0 : len(tloc)-4]
		relative := strings.Replace(outFileName, loc, "", 1)
		outFileName = filepath.Join(outpath, relative)
		f, err := os.Create(outFileName)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := t.Execute(f, tdata); err != nil {
			return err
		}
	}
	return nil
}

func writeRegularFiles(loc, outpath string, fs []string) error {
	for _, floc := range fs {
		//read the file data and write it out
		data, err := ioutil.ReadFile(floc)
		if err != nil {
			return err
		}
		relative := strings.Replace(floc, loc, "", 1)
		outFileName := filepath.Join(outpath, relative)
		if err := ioutil.WriteFile(outFileName, data, 0755); err != nil {
			return err
		}
	}
	return nil
}
