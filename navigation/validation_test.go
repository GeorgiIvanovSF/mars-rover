package navigation

import (
	"testing"
)

func TestGrid_IsValid(t *testing.T) {
	type fields struct {
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test if grid is valid",
			fields: fields{
				Width:  1,
				Height: 1,
			},
			want: true,
		},
		{
			name: "Test if grid is invalid",
			fields: fields{
				Width:  0,
				Height: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := g.IsValid(); got != tt.want {
				t.Errorf("Grid.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_IsValid(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test if position is valid",
			fields: fields{
				X:         1,
				Y:         1,
				Direction: North,
			},
			want: true,
		},
		{
			name: "Test if position is invalid",
			fields: fields{
				X:         -1,
				Y:         1,
				Direction: North,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("Position.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_IsWithinGrid(t *testing.T) {
	type fields struct {
		X         int
		Y         int
		Direction Direction
	}
	type args struct {
		g Grid
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Test if Position is within the Grid",
			fields: fields{
				X:         1,
				Y:         1,
				Direction: North,
			},
			args: args{
				g: Grid{Width: 2, Height: 2},
			},
			want: true,
		},
		{
			name: "Test if Position is outside the Grid",
			fields: fields{
				X:         10,
				Y:         10,
				Direction: North,
			},
			args: args{
				g: Grid{Width: 2, Height: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				X:         tt.fields.X,
				Y:         tt.fields.Y,
				Direction: tt.fields.Direction,
			}
			if got := p.IsWithinGrid(tt.args.g); got != tt.want {
				t.Errorf("Position.IsWithinGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstruction_IsValid(t *testing.T) {
	tests := []struct {
		name string
		i    Instruction
		want bool
	}{
		{
			name: "Test if Instruction is Valid moving",
			i:    Move,
			want: true,
		},
		{
			name: "Test if Instruction is Valid rotating left",
			i:    RotateLeft,
			want: true,
		},
		{
			name: "Test if Instruction is Valid rotating right",
			i:    RotateRight,
			want: true,
		},

		{
			name: "Test if Instruction is invalid",
			i:    Instruction('A'),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsValid(); got != tt.want {
				t.Errorf("Instruction.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstruction_IsRotation(t *testing.T) {
	tests := []struct {
		name string
		i    Instruction
		want bool
	}{
		{
			name: "Test if Instruction is Valid Rotation left",
			i:    RotateLeft,
			want: true,
		},
		{
			name: "Test if Instruction is Valid Rotation right",
			i:    RotateRight,
			want: true,
		},
		{
			name: "Test if Instruction is Not a rotation",
			i:    Move,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IsRotation(); got != tt.want {
				t.Errorf("Instruction.IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirection_IsValid(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want bool
	}{
		{
			name: "test if Direction is valid",
			d:    Direction('E'),
			want: true,
		},
		{
			name: "test if Direction is invalid",
			d:    Direction('B'),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsValid(); got != tt.want {
				t.Errorf("Direction.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
