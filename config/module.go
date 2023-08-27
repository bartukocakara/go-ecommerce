package config

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"path/filepath"

)

type StubData struct {
	ModuleVar        string
	ModuleTitle      string
	ModuleCamelTitle string
}

func loadTemplateContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func loadStubTemplates() (*stubTemplates, error) {
	templates := &stubTemplates{}

	repositoryContent, err := loadTemplateContent(repositoryStubPath)
	if err != nil {
		return nil, err
	}
	templates.Repository = repositoryContent

	serviceContent, err := loadTemplateContent(serviceStubPath)
	if err != nil {
		return nil, err
	}
	templates.Service = serviceContent

	routeContent, err := loadTemplateContent(routeStubPath)
	if err != nil {
		return nil, err
	}
	templates.Route = routeContent

	dtoContent, err := loadTemplateContent(dtoStubPath)
	if err != nil {
		return nil, err
	}
	templates.Dto = dtoContent

	entityContent, err := loadTemplateContent(entityStubPath)
	if err != nil {
		return nil, err
	}
	templates.Entity = entityContent

	containerContent, err := loadTemplateContent(containerStubPath)
	if err != nil {
		return nil, err
	}
	templates.Container = containerContent

	handlerContent, err := loadTemplateContent(handlerStubPath)
	if err != nil {
		return nil, err
	}
	templates.Handler = handlerContent

	migrationContent, err := loadTemplateContent(migrationStubPath)
	if err != nil {
		return nil, err
	}
	templates.Migration = migrationContent

	seederContent, err := loadTemplateContent(seederStubPath)
	if err != nil {
		return nil, err
	}
	templates.Seeder = seederContent

	return templates, nil
}

func createFile(filePath, content string) error {
	// Create the directory path if it doesn't exist
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// Generate files for the specified module using the provided stub templates
func generateFiles(moduleName string, templates *stubTemplates) {
	fmt.Println("Generating files for", moduleName)

	// Convert the module name to title case
	titleCase, camelCase := toTitleCaseWithSpecialChars(moduleName)

	// Define the file names and content for the module
	files := map[string]string{
		fmt.Sprintf("internal/handler/%s.handler.go", moduleName):       generateContent(templates.Handler, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("internal/repository/%s.repository.go", moduleName): generateContent(templates.Repository, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("internal/service/%s.service.go", moduleName):       generateContent(templates.Service, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("internal/route/%s.route.go", moduleName):           generateContent(templates.Route, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("internal/entity/%s.entity.go", moduleName):         generateContent(templates.Entity, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("internal/dto/%s.dto.go", moduleName):               generateContent(templates.Dto, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("container/%s.container.go", moduleName):            generateContent(templates.Container, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("database/seeder/%s.seeder.go", moduleName):         generateContent(templates.Seeder, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
		fmt.Sprintf("database/migration/%s.migration.go", moduleName):   generateContent(templates.Migration, StubData{ModuleVar: moduleName, ModuleTitle: titleCase, ModuleCamelTitle: camelCase}),
	}

	// Generate the files
	for filePath, content := range files {
		err := createFile(filePath, content)
		if err != nil {
			fmt.Println("Failed to generate file:", err)
		} else {
			fmt.Println("Generated file:", filePath)
		}
	}
}

func generateContent(templateContent string, data StubData) string {
	tmpl, err := template.New("stub").Parse(templateContent)
	if err != nil {
		fmt.Println("Failed to parse template:", err)
		return ""
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Failed to execute template:", err)
		return ""
	}

	return buf.String()
}

func toTitleCaseWithSpecialChars(input string) (string, string) {
	var titleCase string
	var camelCase string

	words := strings.Fields(strings.ReplaceAll(input, "-", " "))
	for index, word := range words {
		if index == 0 {
			camelCase += strings.ToLower(word)
		} else {
			camelCase += strings.Title(word)
		}
		titleCase += strings.Title(word)
	}

	return titleCase, camelCase
}