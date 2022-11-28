package schemacafe

import "fmt"

type Schema struct {
	Fields []Field `json:"fields"`
}

func (s *Schema) AddField(name, t string) error {
	for _, f := range s.Fields {
		if f.Name == name {
			return fmt.Errorf("field %s already exists", name)
		}
	}
	s.Fields = append(s.Fields, Field{
		Name: name,
		Type: t,
	})
	return nil
}

func (s *Schema) RemoveField(name string) error {
	for i, f := range s.Fields {
		if f.Name == name {
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

func (s *Schema) ChangeFieldName(fieldIndex int, newName string) error {
	for _, f := range s.Fields {
		if f.Name == newName {
			return fmt.Errorf("field %s already exists", newName)
		}
	}
	s.Fields[fieldIndex].Name = newName
	return nil
}

func (s *Schema) ChangeFieldType(fieldIndex int, newType string) {
	s.Fields[fieldIndex].Type = newType
}
