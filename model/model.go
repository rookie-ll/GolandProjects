package model

type model struct {
	tableName string
	fields    map[string]*field
}

type field struct {
	colName string
}

func parseModel(entity any) (*model, error) {
	return &model{
		tableName: "",
		fields:    make(map[string]*field, 8),
	}, nil
}
