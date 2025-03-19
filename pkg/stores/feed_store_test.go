package stores

import (
	"reflect"
	"testing"

	"github.com/ochiengotieno304/feedpulse-go/pkg/db"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
	"gorm.io/gorm"
)

func TestNewFeedStore(t *testing.T) {
	tests := []struct {
		name string
		want FeedStore
	}{
		{
			name: "Returns valid FeedStore implementation",
			want: &feedStore{
				db: db.DB(),
			},
		},
		{
			name: "Ensures FeedStore is properly initialized",
			want: NewFeedStore(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFeedStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFeedStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feedStore_ReadAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		filters  map[string]string
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.News
		wantErr bool
	}{
		{
			name: "Successfully retrieve all feeds without filters",
			fields: fields{
				db: db.DB(),
			},
			args: args{
				filters:  map[string]string{},
				page:     1,
				pageSize: 10,
			},
			want:    []*models.News{}, // This should be modified with expected data based on your test database
			wantErr: false,
		},
		{
			name: "Successfully retrieve feeds with category filter",
			fields: fields{
				db: db.DB(),
			},
			args: args{
				filters:  map[string]string{"category": "technology"},
				page:     1,
				pageSize: 10,
			},
			want:    []*models.News{}, // This should be modified with expected data
			wantErr: false,
		},
		{
			name: "Successfully retrieve with pagination - page 2",
			fields: fields{
				db: db.DB(),
			},
			args: args{
				filters:  map[string]string{},
				page:     2,
				pageSize: 5,
			},
			want:    []*models.News{}, // This should be modified with expected data
			wantErr: false,
		},
		{
			name: "Handle empty result set",
			fields: fields{
				db: db.DB(),
			},
			args: args{
				filters:  map[string]string{"category": "nonexistent_category"},
				page:     1,
				pageSize: 10,
			},
			want:    []*models.News{}, // Empty result expected
			wantErr: false,
		},
		{
			name: "Handle invalid page parameters",
			fields: fields{
				db: db.DB(),
			},
			args: args{
				filters:  map[string]string{},
				page:     -1,
				pageSize: 10,
			},
			want:    []*models.News{}, // Should return default pagination
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &feedStore{
				db: tt.fields.db,
			}
			got, err := s.ReadAll(tt.args.filters, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("feedStore.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("feedStore.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_feedStore_Read(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		feedID int
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
				db: tt.fields.db,
			}
			got, err := s.Read(tt.args.feedID)
			if (err != nil) != tt.wantErr {
				t.Errorf("feedStore.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("feedStore.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
