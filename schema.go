package schemacafe

import (
	"fmt"

	"lib.dev/nameconv"
)

type Schema struct {
	Fields []Field `json:"fields"`
}

func (s *Schema) AddField(name nameconv.Name, t Type) error {
	for _, f := range s.Fields {
		if f.Name.String() == name.String() {
			return fmt.Errorf("field %s already exists", name)
		}
	}
	s.Fields = append(s.Fields, Field{
		Name: name,
		Type: t,
	})
	return nil
}

func (s *Schema) RemoveField(name nameconv.Name) error {
	for i, f := range s.Fields {
		if f.Name.String() == name.String() {
			s.Fields = append(s.Fields[:i], s.Fields[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("field %s does not exist", name)
}

func (s *Schema) MoveField(fromIndex, toIndex int) error {
	if fromIndex < 0 || fromIndex >= len(s.Fields) {
		return fmt.Errorf("invalid fromIndex %d", fromIndex)
	}
	if toIndex < 0 || toIndex >= len(s.Fields) {
		return fmt.Errorf("invalid toIndex %d", toIndex)
	}
	field := s.Fields[fromIndex]
	s.Fields = append(s.Fields[:fromIndex], s.Fields[fromIndex+1:]...)
	s.Fields = append(s.Fields[:toIndex], append([]Field{field}, s.Fields[toIndex:]...)...)
	return nil
}

func (s *Schema) ChangeFieldName(oldName, newName nameconv.Name) error {
	for _, f := range s.Fields {
		if f.Name.String() == newName.String() {
			return fmt.Errorf("field %s already exists", newName)
		}
	}
	for i, f := range s.Fields {
		if f.Name.String() == oldName.String() {
			s.Fields[i].Name = newName
			return nil
		}
	}
	return fmt.Errorf("field %s does not exist", oldName)
}

func (s *Schema) ChangeFieldType(fieldName nameconv.Name, newType Type) error {
	for i, f := range s.Fields {
		if f.Name.String() == fieldName.String() {
			s.Fields[i].Type = newType
			return nil
		}
	}
	return fmt.Errorf("field %s does not exist", fieldName)
}
