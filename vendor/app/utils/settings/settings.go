package settings

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Templates ...
type Templates struct {
	Root        string
	Ext         string
	Folder      string
	TemplateMap map[string]string
}

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// LoadConfig ...
func LoadConfig(configName string, p Parser) {
	// We need to parse config json file into Config struct
	configFile, err := ioutil.ReadFile(configName)
	if err = p.ParseJSON(configFile); err != nil {
		log.Fatalf("Couldn`t parse %s: %v", configName, configFile)
	}
}

// LoadTemplates ...
func LoadTemplates(t *Templates) {
	// Fill template struct from our templates dir
	templ := "%s/%s.%s"
	templateDir := fmt.Sprintf("./%s", t.Folder)
	// list of files from templates dir
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatalln("Read template Directory", err)
	}
	// Set template Root file, all templates will Parse with it
	t.Root = fmt.Sprintf(templ, t.Folder, t.Root, t.Ext)
	// Create map with template names
	templates := make(map[string]string)
	for _, f := range files {
		ext := fmt.Sprintf(".%s", t.Ext)
		fullName := f.Name()
		index := strings.Index(fullName, ext)
		// name without extension
		name := fullName[:index]
		templates[name] = fmt.Sprintf("./%s/%s", t.Folder, fullName)
	}
	t.TemplateMap = templates
}
