package schemacafe

type Schema struct {
	Fields []Field
}

func (s *Schema) AddField()        {}
func (s *Schema) RemoveField()     {}
func (s *Schema) MoveField()       {}
func (s *Schema) ChangeFieldName() {}
func (s *Schema) ChangeFieldType() {}
