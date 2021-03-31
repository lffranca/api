package export

import (
	"reflect"
	"testing"
)

func TestNewHTML(t *testing.T) {
	type args struct {
		input *NewHTMLInput
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "TestNewHTML 01",
			args: args{
				input: &NewHTMLInput{
					Template: []byte(`<!DOCTYPE html>
					<html lang="en">
						<head>
							<meta charset="UTF-8" />
							<meta http-equiv="X-UA-Compatible" content="IE=edge" />
							<meta name="viewport" content="width=device-width, initial-scale=1.0" />
							<title>Document</title>
						</head>
						<body>
							<table>
								<thead>
									<tr>
										<th>CONTA</th>
										<th>TRANSAÇÃO</th>
										<th>VALOR</th>
									</tr>
								</thead>
								<tbody>
									{{range .Transactions}}
										<tr>
											<td>{{.ID}}</td>
											<td>{{.Description}}</td>
											<td>{{.Value}}</td>
										</tr>
									{{end}}
								</tbody>
							</table>
						</body>
					</html>`),
					Input: map[string]interface{}{"Transactions": []struct {
						ID          string
						Description string
						Value       float64
					}{{ID: "alksjd", Description: "Netflix", Value: 123.12}}},
				},
			},
			want:    650,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHTML(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			value := len(got.Bytes())
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("NewHTML() = %v, want %v", value, tt.want)
			}
		})
	}
}
