package main
 
import (
	"io"
	"os"
	"fmt"
	"strings"
	
	"github.com/kodaimura/ddlparse"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error: The path to the DDL must be provided as the first argument.")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println(args[1])
	
	tables, _ := ddlparse.ParseForce(string(data))
	
	s := ""
	for _, table := range tables {
		s += "type " + snakeToPascal(strings.ToLower(table.Name)) + " struct {\n"
		for _, column := range table.Columns {
			cn := strings.ToLower(column.Name)
			s += "\t" + snakeToPascal(cn) + " " + 
				dataTypeToGoType(column.DataType.Name) + " " + 
				"`db:\"" + cn + "\" json:\"" + cn + "\"`\n"
		}
		s += "}"
	}
	fmt.Println(s)
}

func snakeToPascal(snake string) string {
	ls := strings.Split(strings.ToLower(snake), "_")
	for i, s := range ls {
		ls[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	return strings.Join(ls, "")
}

func dataTypeToGoType(dataType string) string {
	dataType = strings.ToUpper(dataType)

	if (strings.Contains(dataType, "INT") || strings.Contains(dataType, "BIT")) {
		return "int"
	} else if strings.Contains(dataType, "NUMERIC") ||
		strings.Contains(dataType, "DECIMAL") ||
		strings.Contains(dataType, "FLOAT") ||
		strings.Contains(dataType, "REAL") ||
		strings.Contains(dataType, "DOUBLE") {
		return "float64"
	} else {
		return "string"
	}
}