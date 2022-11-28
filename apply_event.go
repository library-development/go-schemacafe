package schemacafe

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func ApplyEvent(schemasDir string, e *Event) error {
	switch e.Command {
	case "create-folder":
		in := e.Input["in"]
		name := e.Input["name"]
		dir := filepath.Join(schemasDir, in)

		if !isFolder(dir) {
			return fmt.Errorf("folder %s does not exist", in)
		}

		err := os.Mkdir(filepath.Join(dir, name), 0755)
		if err != nil {
			return err
		}
	case "create-schema":
		in := e.Input["in"]
		name := e.Input["name"]
		dir := filepath.Join(schemasDir, in)
		if !isFolder(dir) {
			return fmt.Errorf("folder %s does not exist", in)
		}
		schema := Schema{
			Fields: []Field{},
		}
		b, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return err
		}
		path := filepath.Join(dir, name)
		err = os.WriteFile(path, b, 0644)
		if err != nil {
			return err
		}
	case "move":
		from := e.Input["from"]
		to := e.Input["to"]
		fromPath := filepath.Join(schemasDir, from)
		toPath := filepath.Join(schemasDir, to)
		err := os.Rename(fromPath, toPath)
		if err != nil {
			return err
		}
	case "delete":
		path := e.Input["path"]
		fullPath := filepath.Join(schemasDir, path)
		err := os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	case "add-schema-field":
		schemaPath := e.Input["schemaPath"]
		fieldName := e.Input["fieldName"]
		fieldType := e.Input["fieldType"]
		schema, err := ReadSchema(filepath.Join(schemasDir, schemaPath))
		if err != nil {
			return err
		}
		err = schema.AddField(fieldName, fieldType)
		if err != nil {
			return err
		}
		err = WriteSchema(filepath.Join(schemasDir, schemaPath), schema)
		if err != nil {
			return err
		}
	case "remove-schema-field":
		schemaPath := e.Input["schemaPath"]
		fieldName := e.Input["fieldName"]
		schema, err := ReadSchema(filepath.Join(schemasDir, schemaPath))
		if err != nil {
			return err
		}
		err = schema.RemoveField(fieldName)
		if err != nil {
			return err
		}
		err = WriteSchema(filepath.Join(schemasDir, schemaPath), schema)
		if err != nil {
			return err
		}
	case "move-schema-field":
		schemaPath := e.Input["schemaPath"]
		fromIndex, err := strconv.Atoi(e.Input["fromIndex"])
		if err != nil {
			return fmt.Errorf("invalid fromIndex: %s", e.Input["fromIndex"])
		}
		toIndex, err := strconv.Atoi(e.Input["toIndex"])
		if err != nil {
			return fmt.Errorf("invalid toIndex: %s", e.Input["toIndex"])
		}
		schema, err := ReadSchema(filepath.Join(schemasDir, schemaPath))
		if err != nil {
			return err
		}
		err = schema.MoveField(fromIndex, toIndex)
		if err != nil {
			return err
		}
		err = WriteSchema(filepath.Join(schemasDir, schemaPath), schema)
		if err != nil {
			return err
		}
	case "change-schema-field-name":
		schemaPath := e.Input["schemaPath"]
		oldName := e.Input["oldName"]
		newName := e.Input["newName"]
		schema, err := ReadSchema(filepath.Join(schemasDir, schemaPath))
		if err != nil {
			return err
		}
		err = schema.ChangeFieldName(oldName, newName)
		if err != nil {
			return err
		}
		err = WriteSchema(filepath.Join(schemasDir, schemaPath), schema)
		if err != nil {
			return err
		}
	case "change-schema-field-type":
		schemaPath := e.Input["schemaPath"]
		fieldName := e.Input["fieldName"]
		newType := e.Input["newType"]
		schema, err := ReadSchema(filepath.Join(schemasDir, schemaPath))
		if err != nil {
			return err
		}
		err = schema.ChangeFieldType(fieldName, newType)
		if err != nil {
			return err
		}
		err = WriteSchema(filepath.Join(schemasDir, schemaPath), schema)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", e.Command)
	}
	return nil
}

func isFolder(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
