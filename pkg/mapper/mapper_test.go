package mapper

import (
	"reflect"
	"sort"
	"testing"
)

func Test_bisectSlice(t *testing.T) {
	type args struct {
		src []string
		sep string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "Normal",
			args: args{
				src: []string{"a", "b", "c", ";", "d", "e", "f"},
				sep: ";",
			},
			want:  []string{"a", "b", "c"},
			want1: []string{"d", "e", "f"},
		},
		{
			name: "No Match",
			args: args{
				src: []string{"a", "b", "c", ";", "d", "e", "f"},
				sep: "not there",
			},
			want:  nil,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bisectSlice(tt.args.src, tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bisectSlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("bisectSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseMappings(t *testing.T) {
	type args struct {
		src []string
		sep string
	}
	tests := []struct {
		name    string
		args    args
		want    []envMapping
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				src: []string{"A:B", "B:A"},
				sep: ":",
			},
			want: []envMapping{
				{"A", "B"},
				{"B", "A"},
			},
			wantErr: false,
		},
		{
			name: "Missing Separator",
			args: args{
				src: []string{"A:B", "BA"},
				sep: ":",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty",
			args: args{
				src: []string{""},
				sep: ":",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Duplicate Target",
			args: args{
				src: []string{"A:B", "B:A", "A:B"},
				sep: ":",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMappings(tt.args.src, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMappings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMappings() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveMappings(t *testing.T) {
	type args struct {
		mappings []envMapping
		resolver func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Normal",
			args: args{
				mappings: []envMapping{
					{"A", "B"},
					{"B", "C"},
				},
				resolver: func(s string) string {
					return "ValueOf" + s
				},
			},
			want: []string{
				"A=ValueOfB",
				"B=ValueOfC",
			},
		},
		{
			name: "Normal Empty Value",
			args: args{
				mappings: []envMapping{
					{"A", "B"},
					{"B", "C"},
				},
				resolver: func(s string) string {
					return ""
				},
			},
			want: []string{
				"A=",
				"B=",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveMappings(tt.args.mappings, tt.args.resolver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveMappings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommandWithOverrides(t *testing.T) {
	type args struct {
		inputArgs []string
		inputEnv  []string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantArgs []string
		wantEnv  []string
		wantErr  bool
	}{
		{
			name: "Normal",
			args: args{
				inputArgs: []string{"A:B", "B:C", "--", "/usr/bin/env", "-v"},
				inputEnv:  []string{"PATH=/bin:/usr/bin"},
			},
			wantPath: "/usr/bin/env",
			wantArgs: []string{"/usr/bin/env", "-v"}, //argv[0] is always path to executable
			wantEnv:  []string{"A=", "B=", "PATH=/bin:/usr/bin"},
			wantErr:  false,
		},
		{
			name: "Missing command",
			args: args{
				inputArgs: []string{"A:B", "B:C", "--"},
				inputEnv:  []string{"PATH=/bin:/usr/bin"},
			},
			wantPath: "/usr/bin/env",
			wantArgs: []string{"/usr/bin/env", "-v"}, //argv[0] is always path to executable
			wantEnv:  []string{"PATH=/bin:/usr/bin", "A=", "B="},
			wantErr:  true,
		},
		{
			name: "Broken envMapping",
			args: args{
				inputArgs: []string{"NOT_A_MAPPING", "B:C", "--", "/usr/bin/env", "-v"},
				inputEnv:  []string{"PATH=/bin:/usr/bin"},
			},
			wantPath: "/usr/bin/env",
			wantArgs: []string{"/usr/bin/env", "-v"}, //argv[0] is always path to executable
			wantEnv:  []string{"PATH=/bin:/usr/bin", "A=", "B="},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommandWithEnvOverrides(tt.args.inputArgs, tt.args.inputEnv)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommandWithEnvOverrides() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil { // do not test expected values if there is an error
				return
			}
			sort.Strings(got.Env) // using map results in non-deterministic order, but we do not care
			if !reflect.DeepEqual(got.Path, tt.wantPath) {
				t.Errorf("CommandWithEnvOverrides().Path got = %v, want %v", got.Path, tt.wantPath)
			}
			if !reflect.DeepEqual(got.Args, tt.wantArgs) {
				t.Errorf("CommandWithEnvOverrides().Path got = %v, want %v", got.Args, tt.wantArgs)
			}
			if !reflect.DeepEqual(got.Env, tt.wantEnv) {
				t.Errorf("CommandWithEnvOverrides().Path got = %v, want %v", got.Env, tt.wantEnv)
			}
		})
	}
}
