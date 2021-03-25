package pingdom_ext

import (
	"reflect"
	"testing"
)

func TestWebHookIntegration_PostParams(t *testing.T) {
	type fields struct {
		Active     bool
		ProviderId int
		UserData   *WebHookData
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wi := &WebHookIntegration{
				Active:     tt.fields.Active,
				ProviderId: tt.fields.ProviderId,
				UserData:   tt.fields.UserData,
			}
			if got := wi.PostParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebHookIntegration.PostParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebHookIntegration_Valid(t *testing.T) {
	type fields struct {
		Active     bool
		ProviderId int
		UserData   *WebHookData
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wi := &WebHookIntegration{
				Active:     tt.fields.Active,
				ProviderId: tt.fields.ProviderId,
				UserData:   tt.fields.UserData,
			}
			if err := wi.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("WebHookIntegration.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLibratoIntegration_PostParams(t *testing.T) {
	type fields struct {
		Active     bool
		ProviderId int
		UserData   *LibratoData
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := &LibratoIntegration{
				Active:     tt.fields.Active,
				ProviderId: tt.fields.ProviderId,
				UserData:   tt.fields.UserData,
			}
			if got := li.PostParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LibratoIntegration.PostParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLibratoIntegration_Valid(t *testing.T) {
	type fields struct {
		Active     bool
		ProviderId int
		UserData   *LibratoData
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := &LibratoIntegration{
				Active:     tt.fields.Active,
				ProviderId: tt.fields.ProviderId,
				UserData:   tt.fields.UserData,
			}
			if err := li.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("LibratoIntegration.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
