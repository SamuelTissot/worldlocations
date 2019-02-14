package actions

import (
	"testing"
	"worldlocations/models"

	"github.com/gobuffalo/buffalo"
)

func TestCities_List(t *testing.T) {
	type fields struct {
		Count int
		Data  *models.Cities
	}
	type args struct {
		c buffalo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := Cities{
				Count: tt.fields.Count,
				Data:  tt.fields.Data,
			}
			if err := cs.List(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Cities.List() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCities_CountryCities(t *testing.T) {
	type fields struct {
		Count int
		Data  *models.Cities
	}
	type args struct {
		c buffalo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := Cities{
				Count: tt.fields.Count,
				Data:  tt.fields.Data,
			}
			if err := cs.CountryCities(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Cities.CountryCities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCities_SubdivisionCities(t *testing.T) {
	type fields struct {
		Count int
		Data  *models.Cities
	}
	type args struct {
		c buffalo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := Cities{
				Count: tt.fields.Count,
				Data:  tt.fields.Data,
			}
			if err := cs.SubdivisionCities(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Cities.SubdivisionCities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
