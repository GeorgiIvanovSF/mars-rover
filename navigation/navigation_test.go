package navigation

import (
	"reflect"
	"testing"
)

func TestPosition_String(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Position string output",
			fields: fields{
				X:         1,
				Y:         1,
				Direction: North,
			},
			want: "1 1 N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildInstructions(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []Instruction
	}{
		{name: "Position string output",

			args: args{
				line: "MMRMMRMMR",
			},

			want: []Instruction{'M', 'M', 'R', 'M', 'M', 'R', 'M', 'M', 'R'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildInstructions(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_ChangeCoordinates(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   Position
	}{
		{name: "Test Move Forward",

			fields: fields{
				X:         1,
				Y:         1,
				Direction: North,
			},

			want: Position{
				X:         1,
				Y:         2,
				Direction: North,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			p.ChangeCoordinates()
			if !reflect.DeepEqual(*p, tt.want) {
				t.Errorf("ChangeCoordinates() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPosition_ChangeDirection(t *testing.T) {
	type fields struct {
		X           int
		Y           int
		Direction   Direction
		Instruction Instruction
	}
	type args struct {
		i Instruction
	}
	tests := []struct {
		name   string
		fields fields
		want   Position
	}{
		{name: "Change Direction Left",

			fields: fields{
				X:           1,
				Y:           1,
				Direction:   North,
				Instruction: RotateLeft,
			},

			want: Position{
				X:         1,
				Y:         1,
				Direction: West,
			},
		},
		{name: "Change Direction Right",

			fields: fields{
				X:           1,
				Y:           1,
				Direction:   North,
				Instruction: RotateRight,
			},

			want: Position{
				X:         1,
				Y:         1,
				Direction: East,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			p.ChangeDirection(tt.fields.Instruction)
			if !reflect.DeepEqual(*p, tt.want) {
				t.Errorf("ChangeDirection() = %v, want %v", p, tt.want)
			}
		})
	}
}

func Test_compass_Use(t *testing.T) {
	tests := []struct {
		name string
		want map[Direction]map[Instruction]Direction
	}{
		{
			name: "Test if Compass was changed in the library",

			want: map[Direction]map[Instruction]Direction{
				North: {RotateLeft: West,
					RotateRight: East},
				East: {RotateLeft: North,
					RotateRight: South},
				South: {RotateLeft: East,
					RotateRight: West},
				West: {RotateLeft: South,
					RotateRight: North},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewCompass().Use(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compass.UseCompass() = %v, want %v", got, tt.want)
			}
		})
	}
}
