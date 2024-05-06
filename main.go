package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	projectName := getProjectName()
	entities := getEntities()

	createProjectStructure(projectName, entities)

	// Abrir Visual Studio Code
	err := exec.Command("code", projectName).Run()
	if err != nil {
		fmt.Println("Error al abrir Visual Studio Code:", err)
	}
}

func getProjectName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el nombre del proyecto: ")
	projectName, _ := reader.ReadString('\n')
	return strings.TrimSpace(projectName)
}

func getEntities() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese todas las entidades (dominios) separadas por coma: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	entities := strings.Split(input, ",")
	for i, entity := range entities {
		entities[i] = strings.TrimSpace(entity)
	}
	return entities
}

func createProjectStructure(projectName string, entities []string) {
	fmt.Printf("Creando estructura para el proyecto '%s'...\n", projectName)

	// Crear directorio del proyecto
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Printf("Error al crear el directorio '%s': %v\n", projectName, err)
		return
	}
	fmt.Printf("Directorio '%s' creado correctamente.\n", projectName)

	// Crear directorio cmd
	cmdDir := filepath.Join(projectName, "cmd")
	err = os.Mkdir(cmdDir, 0755)
	if err != nil {
		fmt.Printf("Error al crear el directorio 'cmd' dentro de '%s': %v\n", projectName, err)
		return
	}
	fmt.Printf("Directorio 'cmd' dentro de '%s' creado correctamente.\n", projectName)

	// Crear archivo main.go dentro de cmd
	mainFilePath := filepath.Join(cmdDir, "main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		fmt.Printf("Error al crear el archivo 'main.go': %v\n", err)
		return
	}
	defer mainFile.Close()
	fmt.Fprintln(mainFile, "package main")
	fmt.Fprintln(mainFile, "\nfunc main() {")
	fmt.Fprintln(mainFile, "\t// Código principal aquí")
	fmt.Fprintln(mainFile, "}")

	fmt.Printf("Archivo 'main.go' creado correctamente en '%s'.\n", cmdDir)

	// Crear directorio pkg
	pkgDir := filepath.Join(projectName, "pkg")
	err = os.Mkdir(pkgDir, 0755)
	if err != nil {
		fmt.Printf("Error al crear el directorio 'pkg' dentro de '%s': %v\n", projectName, err)
		return
	}
	fmt.Printf("Directorio 'pkg' dentro de '%s' creado correctamente.\n", projectName)

	// Crear directorio internal
	internalDir := filepath.Join(projectName, "internal")
	err = os.Mkdir(internalDir, 0755)
	if err != nil {
		fmt.Printf("Error al crear el directorio 'internal' dentro de '%s': %v\n", projectName, err)
		return
	}
	fmt.Printf("Directorio 'internal' dentro de '%s' creado correctamente.\n", projectName)

	// Crear directorio domain dentro de internal
	domainDir := filepath.Join(internalDir, "domain")
	err = os.Mkdir(domainDir, 0755)
	if err != nil {
		fmt.Printf("Error al crear el directorio 'domain' dentro de '%s': %v\n", internalDir, err)
		return
	}
	fmt.Printf("Directorio 'domain' dentro de '%s' creado correctamente.\n", internalDir)

	// Crear archivos Go para cada entidad dentro de domain
	for _, entity := range entities {
		entityFile := filepath.Join(domainDir, entity+".go")
		_, err = os.Create(entityFile)
		if err != nil {
			fmt.Printf("Error al crear el archivo '%s': %v\n", entityFile, err)
			return
		}
		fmt.Printf("Archivo '%s' creado correctamente.\n", entityFile)

		// Crear directorio para cada entidad dentro de internal
		entityDir := filepath.Join(internalDir, entity)
		err := os.Mkdir(entityDir, 0755)
		if err != nil {
			fmt.Printf("Error al crear el directorio '%s': %v\n", entityDir, err)
			return
		}
		fmt.Printf("Directorio '%s' creado correctamente.\n", entityDir)

		// Crear archivos controller, service y repository dentro de cada entidad
		files := []string{"controller.go", "repository.go", "service.go"}
		for _, file := range files {
			filePath := filepath.Join(entityDir, file)
			_, err = os.Create(filePath)
			if err != nil {
				fmt.Printf("Error al crear el archivo '%s': %v\n", filePath, err)
				return
			}
			fmt.Printf("Archivo '%s' creado correctamente.\n", filePath)
		}
	}

	fmt.Printf("Estructura para el proyecto '%s' creada exitosamente.\n", projectName)
}
