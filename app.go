package schemacafe

import "github.com/library-development/go-web"

var app = web.App{
	Models: []web.Model{
		{
			Name: "schema",
			Fields: []web.Field{
				{
					Name: "fields",
					Type: web.Type{
						IsArray: true,
						BaseType: web.ID{
							Path: "field",
						},
					},
				},
			},
		},
		{
			Name: "field",
			Fields: []web.Field{
				{
					Name: "name",
					Type: web.Type{
						BaseType: web.ID{
							Path: "name",
						},
					},
				},
				{
					Name: "type",
					Type: web.Type{
						BaseType: web.ID{
							Path: "type",
						},
					},
				},
			},
		},
		{
			Name: "type",
			Fields: []web.Field{
				{
					Name: "is_pointer",
					Type: web.Type{
						BaseType: web.ID{
							Path: "bool",
						},
					},
				},
				{
					Name: "is_array",
					Type: web.Type{
						BaseType: web.ID{
							Path: "bool",
						},
					},
				},
				{
					Name: "is_map",
					Type: web.Type{
						BaseType: web.ID{
							Path: "bool",
						},
					},
				},
				{
					Name: "base_type",
					Type: web.Type{
						BaseType: web.ID{
							Path: "id",
						},
					},
				},
			},
		},
		{
			Name: "id",
			Fields: []web.Field{
				{
					Name: "path",
					Type: web.Type{
						BaseType: web.ID{
							Path: "path",
						},
					},
				},
			},
		},
		{
			Name: "path",
			Fields: []web.Field{
				{
					Name: "parts",
					Type: web.Type{
						IsArray: true,
						BaseType: web.ID{
							Path: "name",
						},
					},
				},
			},
		},
		{
			Name: "name",
			Fields: []web.Field{
				{
					Name: "words",
					Type: web.Type{
						IsArray: true,
						BaseType: web.ID{
							Path: "string",
						},
					},
				},
			},
		},
	},
}
