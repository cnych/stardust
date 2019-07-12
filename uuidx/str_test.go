package uuidx

import (
	"github.com/cnych/stardust/stringsx"
	"github.com/satori/go.uuid"
	"testing"
)

const uuidSymbol = "-"

func TestHex(t *testing.T) {
	type args struct {
		o uuid.UUID
	}
	uuid0 := uuid.NewV2(uuid.DomainPerson)
	uuid1 := uuid.NewV2(uuid.DomainGroup)
	uuid2 := uuid.NewV2(uuid.DomainOrg)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "uuid0",
			args: args{o: uuid0},
			want: stringsx.RemoveSymbol(uuid0.String(), uuidSymbol),
		}, {
			name: "uuid1",
			args: args{o: uuid1},
			want: stringsx.RemoveSymbol(uuid1.String(), uuidSymbol),
		}, {
			name: "uuid2",
			args: args{o: uuid2},
			want: stringsx.RemoveSymbol(uuid2.String(), uuidSymbol),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hex(tt.args.o); got != tt.want {
				t.Errorf("Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexV1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexV1(); got != tt.want {
				t.Errorf("HexV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexV4(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexV4(); got != tt.want {
				t.Errorf("HexV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64(t *testing.T) {
	type args struct {
		o uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64(tt.args.o); got != tt.want {
				t.Errorf("Base64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64V1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64V1(); got != tt.want {
				t.Errorf("Base64V1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64V4(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64V4(); got != tt.want {
				t.Errorf("Base64V4() = %v, want %v", got, tt.want)
			}
		})
	}
}
