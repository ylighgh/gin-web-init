package utils

import "testing"

func TestCpuResource(t *testing.T) {
	type args struct {
		resource string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "example_cpu_resource_test1",
			args:    args{resource: "3000m"},
			want:    3000,
			wantErr: false,
		},
		{
			name:    "example_cpu_resource_test2",
			args:    args{resource: "3000"},
			want:    3000,
			wantErr: false,
		},
		{
			name:    "example_cpu_resource_test3",
			args:    args{resource: "300a0m"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CpuResource(tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("CpuResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CpuResource() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryResource(t *testing.T) {
	type args struct {
		resource string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "example_memory_resource_test1",
			args:    args{resource: "1024Mi"},
			want:    1024 * Legitimate["Mi"],
			wantErr: false,
		},
		{
			name:    "example_memory_resource_test2",
			args:    args{resource: "1024Gi"},
			want:    1024 * Legitimate["Gi"],
			wantErr: false,
		},
		{
			name:    "example_memory_resource_test3",
			args:    args{resource: "1024Ki"},
			want:    1024 * Legitimate["Ki"],
			wantErr: false,
		},
		{
			name:    "example_memory_resource_test4",
			args:    args{resource: "1024T"},
			want:    1024 * Legitimate["T"],
			wantErr: false,
		},
		{
			name:    "example_memory_resource_test5",
			args:    args{resource: "1024TB"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MemoryResource(tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MemoryResource() got = %v, want %v", got, tt.want)
			}
		})
	}
}
