package model

import "testing"

func Test_parseModel(t *testing.T) {
	tests := []struct {
		name string

		entity    any
		wantModel *model
		wantErr   error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := parseModel(tt.entity)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantModel, m)
		})
	}
}
