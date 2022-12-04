package main

import (
	"testing"
)

func Test_countBox(t *testing.T) {
	type args struct {
		cakes  int
		apples int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "cakes: 20 apple: 25",
			args: args{
				cakes:  20,
				apples: 25,
			},
			want: 5,
		},
		{
			name: "cakes: 12 apple: 18",
			args: args{
				cakes:  12,
				apples: 18,
			},
			want: 6,
		},
		{
			name: "cakes: 96 apple: 48",
			args: args{
				cakes:  96,
				apples: 48,
			},
			want: 48,
		},
		{
			name: "cakes: 10 apple: 8",
			args: args{
				cakes:  10,
				apples: 8,
			},
			want: 2,
		},
		{
			name: "cakes: 10 apple: 7",
			args: args{
				cakes:  10,
				apples: 7,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBox(tt.args.cakes, tt.args.apples); got != tt.want {
				t.Errorf("countBox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getItemsPerBox(t *testing.T) {
	type args struct {
		cakes  int
		apples int
	}
	tests := []struct {
		name           string
		args           args
		wantBoxCount   int
		wantCakeCount  int
		wantAppleCount int
	}{
		{
			name: "cake: 20 apple: 25",
			args: args{
				cakes:  20,
				apples: 25,
			},
			wantBoxCount:   5,
			wantCakeCount:  4,
			wantAppleCount: 5,
		},
		{
			name: "cake: 12 apple: 18",
			args: args{
				cakes:  12,
				apples: 18,
			},
			wantBoxCount:   6,
			wantCakeCount:  2,
			wantAppleCount: 3,
		},
		{
			name: "cakes: 96 apple: 48",
			args: args{
				cakes:  96,
				apples: 48,
			},
			wantBoxCount:   48,
			wantCakeCount:  2,
			wantAppleCount: 1,
		},
		{
			name: "cakes: 10 apple: 8",
			args: args{
				cakes:  10,
				apples: 8,
			},
			wantBoxCount:   2,
			wantCakeCount:  5,
			wantAppleCount: 4,
		},
		{
			name: "cakes: 10 apple: 7",
			args: args{
				cakes:  10,
				apples: 7,
			},
			wantBoxCount:   1,
			wantCakeCount:  10,
			wantAppleCount: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoxCount, gotCakeCount, gotAppleCount := GetItemsPerBox(tt.args.cakes, tt.args.apples)
			if gotBoxCount != tt.wantBoxCount {
				t.Errorf("getItemsPerBox() gotBoxCount = %v, want %v", gotBoxCount, tt.wantBoxCount)
			}
			if gotCakeCount != tt.wantCakeCount {
				t.Errorf("getItemsPerBox() gotCakeCount = %v, want %v", gotCakeCount, tt.wantCakeCount)
			}
			if gotAppleCount != tt.wantAppleCount {
				t.Errorf("getItemsPerBox() gotAppleCount = %v, want %v", gotAppleCount, tt.wantAppleCount)
			}
		})
	}
}
