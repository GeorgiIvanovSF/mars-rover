package inputparser

import (
	"testing"
)

var file = "resources/example-commands-cmd.txt"

func TestParser_ValidateGridLine(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Test if grid line is valid",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1",
			},
			want: true,
		},
		{
			name: "Test if grid line is invalid",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1M",
			},
			want: false,
		},
		{
			name: "Test if grid line has more than 2 words",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1 1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				FilePath: tt.fields.FilePath,
			}
			if got := p.ValidateGridLine(tt.args.line); got != tt.want {
				t.Errorf("Parser.ValidateGridLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ValidatePositionLine(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Test if position line is valid",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1 N",
			},
			want: true,
		},
		{
			name: "Test if position line is invalid",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1 4",
			},
			want: false,
		},
		{
			name: "Test if position line is not 3 words",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "1 1 N M",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				FilePath: tt.fields.FilePath,
			}
			if got := p.ValidatePositionLine(tt.args.line); got != tt.want {
				t.Errorf("Parser.ValidatePositionLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ValidateInstructionsLine(t *testing.T) {
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
		want   bool
	}{
		{
			name: "Test if instructions line is valid",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "MMRMLMRMLR",
			},
			want: true,
		},
		{
			name: "Test if instructions contain forbidden characters",
			fields: fields{
				FilePath: file,
			},
			args: args{
				line: "MMARMBLMRMLR",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				FilePath: tt.fields.FilePath,
			}
			if got := p.ValidateInstructionsLine(tt.args.line); got != tt.want {
				t.Errorf("Parser.ValidateInstructionsLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
