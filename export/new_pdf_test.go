package export

import (
	"reflect"
	"testing"
)

func TestNewPDF(t *testing.T) {
	type args struct {
		input *NewPDFInput
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "TestNewPDF 01",
			args: args{
				input: &NewPDFInput{
					Template: []byte(`
					<!DOCTYPE html>
					<html lang="en">
						<head>
							<meta charset="UTF-8" />
							<meta http-equiv="X-UA-Compatible" content="IE=edge" />
							<meta name="viewport" content="width=device-width, initial-scale=1.0" />
							<title>Document</title>
						</head>
						<body>
							<h1>TESTE</h1>
						</body>
					</html>
					`),
				},
			},
			want:    6596,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPDF(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(len(got.Bytes()), tt.want) {
				t.Errorf("NewPDF() = %v, want %v", nil, tt.want)
			}
		})
	}
}
