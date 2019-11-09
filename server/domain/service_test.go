package domain_test

import (
	"context"
	"reflect"
	"testing"
	"todo/domain"
	"todo/domain/mock_domain"

	gomock "github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {
}

func TestLoadTodo_Process(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    []domain.Todo
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				ctx: context.Background(),
			},
			want: []domain.Todo{
				domain.Todo{ID: 1, IsDone: false, Name: "task"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_domain.NewMockTodoRepository(ctrl)
			m.EXPECT().FindAll(tt.args.ctx).Return([]domain.Todo{
				domain.Todo{ID: 1, IsDone: false, Name: "task"},
			})

			s := &domain.LoadTodo{
				TodoRepository: m,
			}
			got, err := s.Process(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTodo.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadTodo.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
