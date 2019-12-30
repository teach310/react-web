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

func TestLoadTodo_Run(t *testing.T) {
	type args struct {
		ctx     context.Context
		ownerID string
	}

	type recorders struct {
		taskRepository *mock_domain.MockTaskRepositoryMockRecorder
	}

	tests := []struct {
		name    string
		args    args
		record  func(recorders recorders, args args)
		want    []domain.Task
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				ctx:     context.Background(),
				ownerID: "dummy_owner01",
			},
			record: func(recorders recorders, args args) {
				recorders.taskRepository.FindAllByOwnerID(args.ctx, args.ownerID).Return([]domain.Task{
					domain.Task{ID: 1, IsDone: false, Name: "task", OwnerID: "dummy_owner01"},
				})
			},
			want: []domain.Task{
				domain.Task{ID: 1, IsDone: false, Name: "task", OwnerID: "dummy_owner01"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_domain.NewMockTaskRepository(ctrl)
			if tt.record != nil {
				recorders := recorders{
					taskRepository: m.EXPECT(),
				}
				tt.record(recorders, tt.args)
			}

			s := &domain.LoadTodo{
				OwnerID:        tt.args.ownerID,
				TaskRepository: m,
			}
			got, err := s.Run(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTodo.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadTodo.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
