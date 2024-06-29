package stores

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
)

func TestNewFeedStore(t *testing.T) {
	tests := []struct {
		name string
		want FeedStore
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFeedStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFeedStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feedStore_GetAll(t *testing.T) {
	type fields struct {
		db  *pgx.Conn
		ctx context.Context
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.News
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &feedStore{
				db:  tt.fields.db,
				ctx: tt.fields.ctx,
			}
			got, err := s.GetAll(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("feedStore.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("feedStore.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feedStore_GetSingle(t *testing.T) {
	type fields struct {
		db  *pgx.Conn
		ctx context.Context
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.News
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &feedStore{
				db:  tt.fields.db,
				ctx: tt.fields.ctx,
			}
			got, err := s.GetSingle(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("feedStore.GetSingle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("feedStore.GetSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}
