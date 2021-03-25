package pingdom_ext

import (
	"reflect"
	"testing"

	"github.com/nordcloud/go-pingdom/pingdom"
)

func TestIntegrationService_Create(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		integration Integration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IntegrationStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &IntegrationService{
				client: tt.fields.client,
			}
			got, err := cs.Create(tt.args.integration)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrationService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegrationService_List(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		params []map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []IntegrationGetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &IntegrationService{
				client: tt.fields.client,
			}
			got, err := cs.List(tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrationService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationService.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegrationService_Read(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IntegrationGetResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &IntegrationService{
				client: tt.fields.client,
			}
			got, err := cs.Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrationService.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationService.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegrationService_Update(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id          int
		integration Integration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IntegrationStatus
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &IntegrationService{
				client: tt.fields.client,
			}
			got, err := cs.Update(tt.args.id, tt.args.integration)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrationService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntegrationService_Delete(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pingdom.PingdomResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &IntegrationService{
				client: tt.fields.client,
			}
			got, err := cs.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrationService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntegrationService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
