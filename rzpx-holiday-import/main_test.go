package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseDate(t *testing.T) {
	tTime, _ := time.Parse("Mon 2 Jan 2006", "Fri 18 Mar 2022")
	type args struct {
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "Parse Fri, 18th Mar",
			args: args{
				dateStr: "Fri, 18th Mar",
			},
			want:    &tTime,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDate(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
