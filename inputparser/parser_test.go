package inputparser

import (
	"reflect"
	"testing"

	"github.com/GeorgiIvanovSF/mars-rover/navigation"
)

var filepath = "resources/example-commands-cmd.txt"

func TestParser_GetGridDefinition(t *testing.T) {
	type fields struct {
		FilePath string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   navigation.Grid
	}{
		{
			name: "Test if grid definition outputs a correct grid",
			fields: fields{
				FilePath: filepath,
			},
			args: args{
				line: "1 1",
			},
			want: navigation.Grid{Width: 1, Height: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				FilePath: tt.fields.FilePath,
			}
			if got := p.GetGridDefinition(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.GetGridDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetPositionDefinition(t *testing.T) {
	type fields struct {
		FilePath string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   navigation.Position
	}{
		{
			name: "Generate correct position",
			fields: fields{
				FilePath: filepath,
			},
			args: args{
				line: "1 1 N",
			},
			want: navigation.Position{X: 1, Y: 1, Direction: navigation.North},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				FilePath: tt.fields.FilePath,
			}
			if got := p.GetPositionDefinition(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.GetPositionDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
