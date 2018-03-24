package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

var defaults = Configuration{
	DbUser:          "db_user",
	DbPassword:      "db_pw",
	DbName:          "bd_name",
	PkgName:         "DbStructs",
	TagLabel:        "db",
	Xorm:            false,
	OnlyBaseTables:  false,
	IgnoreNullables: false,
}

var config Configuration

type Configuration struct {
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
	// PkgName gives name of the package using the stucts
	PkgName string `json:"pkg_name"`
	// TagLabel produces tags commonly used to match database field names with Go struct members
	TagLabel string `json:"tag_label"`
	// Adds the tablename return for the xorm ORM
	Xorm            bool `json:"xorm"`
	OnlyBaseTables  bool `json:"only_base_tables"`
	IgnoreNullables bool `json:"ignore_nullables"`
}

type ColumnSchema struct {
	TableName              string
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	ColumnKey              string
}

func writeStructs(schemas []ColumnSchema) (int, error) {

	file, err := os.Create("db_structs.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currentTable := ""

	neededImports := make(map[string]bool)

	// First, get body text into var out
	out := ""
	for _, cs := range schemas {

		if cs.TableName != currentTable {
			if currentTable != "" {
				out = out + "}\n\n"
				if config.Xorm {
					out = out + "func (t " + formatName(currentTable) + ") TableName() string {\n" +
						"\t return \"" + currentTable + "\"\n" +
						"}\n\n"

					out = out + "func (t " + formatName(currentTable) + ") SetId(id int64) {\n" +
						"\tt.Id = id\n" +
						"}\n\n"

					out = out + "func (t " + formatName(currentTable) + ") GetId() int64 {\n" +
						"\treturn t.Id\n" +
						"}\n\n"
				}
			}
			out = out + "type " + formatName(cs.TableName) + " struct{\n"
		}

		goType, requiredImport, err := goType(&cs)
		if requiredImport != "" {
			neededImports[requiredImport] = true
		}

		if err != nil {
			log.Fatal(err)
		}
		out = out + "\t" + formatName(cs.ColumnName) + " " + goType
		if len(config.TagLabel) > 0 {
			if config.Xorm {
				out = out + "\t`" + config.TagLabel + ":"
				if cs.ColumnName == "id" {
					out = out + "\"'" + cs.ColumnName + "' pk autoincr"
				} else {
					out = out + "\"" + cs.ColumnName
				}
				out = out + "\" json:\"" + cs.ColumnName + "\""
			} else {
				out = out + "\t`" + config.TagLabel + ":\"" + cs.ColumnName + "\""
			}

			// Need to make this an option at some point
			if true {
				out = out + " schema:\"" + cs.ColumnName + "\""
				if goType == "bool" {
					out = out + " sql:\"default: false\""
				}
			}

			out = out + "`"
		}
		out = out + "\n"
		currentTable = cs.TableName

	}
	out = out + "}"

	// Now add the header section
	header := "package " + config.PkgName + "\n\n"
	if len(neededImports) > 0 {
		header = header + "import (\n"
		for imp := range neededImports {
			header = header + "\t\"" + imp + "\"\n"
		}
		header = header + ")\n\n"
	}

	totalBytes, err := fmt.Fprint(file, header+out)
	if err != nil {
		log.Fatal(err)
	}
	return totalBytes, nil
}

func getSchema() []ColumnSchema {
	conn, err := sql.Open("mysql", config.DbUser+":"+config.DbPassword+"@/information_schema")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	q := "SELECT COLUMNS.TABLE_NAME, COLUMNS.COLUMN_NAME, COLUMNS.IS_NULLABLE, COLUMNS.DATA_TYPE, " +
		"COLUMNS.CHARACTER_MAXIMUM_LENGTH, COLUMNS.NUMERIC_PRECISION, COLUMNS.NUMERIC_SCALE, COLUMNS.COLUMN_TYPE, " +
		"COLUMNS.COLUMN_KEY FROM COLUMNS "

	if config.OnlyBaseTables {
		q = q + "LEFT JOIN TABLES ON TABLES.TABLE_NAME = COLUMNS.TABLE_NAME AND TABLES.TABLE_SCHEMA = COLUMNS.TABLE_SCHEMA "
	}

	q = q + "WHERE COLUMNS.TABLE_SCHEMA = ? "

	if config.OnlyBaseTables {
		q = q + "AND TABLES.TABLE_TYPE = \"BASE TABLE\" "
	}

	q = q + "ORDER BY COLUMNS.TABLE_NAME, COLUMNS.ORDINAL_POSITION"
	rows, err := conn.Query(q, config.DbName)
	if err != nil {
		log.Fatal(err)
	}
	columns := []ColumnSchema{}
	for rows.Next() {
		cs := ColumnSchema{}
		err := rows.Scan(&cs.TableName, &cs.ColumnName, &cs.IsNullable, &cs.DataType,
			&cs.CharacterMaximumLength, &cs.NumericPrecision, &cs.NumericScale,
			&cs.ColumnType, &cs.ColumnKey)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, cs)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return columns
}

func formatName(name string) string {
	parts := strings.Split(name, "_")
	newName := ""
	for _, p := range parts {
		if len(p) < 1 {
			continue
		}
		newName = newName + strings.Replace(p, string(p[0]), strings.ToUpper(string(p[0])), 1)
	}
	return newName
}

func goType(col *ColumnSchema) (string, string, error) {
	requiredImport := ""
	if col.IsNullable == "YES" && !config.IgnoreNullables {
		requiredImport = "database/sql"
	}
	var gt string = ""
	switch col.DataType {
	case "char", "varchar", "enum", "set", "text", "longtext", "mediumtext", "tinytext":
		if col.IsNullable == "YES" && !config.IgnoreNullables {
			gt = "sql.NullString"
		} else {
			gt = "string"
		}
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		gt = "[]byte"
	case "date", "time", "datetime", "timestamp":
		gt, requiredImport = "time.Time", "time"
	case "bit", "tinyint", "smallint", "int", "mediumint", "bigint":
		if col.ColumnType == "tinyint(1) unsigned" {
			gt = "bool"
		} else {
			if col.IsNullable == "YES" && !config.IgnoreNullables {
				gt = "sql.NullInt64"
			} else {
				gt = "int64"
			}
		}
	case "float", "decimal", "double":
		if col.IsNullable == "YES" && !config.IgnoreNullables {
			gt = "sql.NullFloat64"
		} else {
			gt = "float64"
		}
	}
	if gt == "" {
		n := col.TableName + "." + col.ColumnName
		return "", "", errors.New("No compatible datatype (" + col.DataType + ") for " + n + " found")
	}
	return gt, requiredImport, nil
}

var configFile = flag.String("json", "", "Config file")

func main() {
	flag.Parse()

	if len(*configFile) > 0 {
		f, err := os.Open(*configFile)
		if err != nil {
			log.Fatal(err)
		}
		err = json.NewDecoder(f).Decode(&config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		config = defaults
	}

	columns := getSchema()
	bytes, err := writeStructs(columns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ok %d\n", bytes)
}
