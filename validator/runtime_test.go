package validator

import (
	"testing"
)

func Test_imageRepositoryValidator(t *testing.T) {
	type args struct {
		urn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "image_repository_validator_1",
			args: args{urn: "harbor.axao.cn/pro/apollo:latest"},
			want: true,
		},
		{
			name: "image_repository_validator_2",
			args: args{urn: "https://harbor.axzo.cn/pro/apollo:latest"},
			want: false,
		},
		{
			name: "image_repository_validator_3",
			args: args{urn: "harbor.axzo.cn/pro/custom/apollo:latest"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageRepositoryValidator(tt.args.urn); got != tt.want {
				t.Errorf("imageRepositoryValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
