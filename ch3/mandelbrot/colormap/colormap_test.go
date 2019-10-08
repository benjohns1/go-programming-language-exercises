package colormap

import (
	"image/color"
	"testing"
)

func TestGradient_FloatToColor(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		g    *Gradient
		args args
		want color.Color
	}{
		{
			name: "single gradient should return white",
			g:    &Gradient{[]color.Color{color.White}, Range{0, 1}},
			args: args{0},
			want: color.White,
		},
		{
			name: "two gradient min should return white",
			g:    &Gradient{[]color.Color{color.White, color.Black}, Range{0, 1}},
			args: args{0},
			want: color.White,
		},
		{
			name: "two gradient max should return black",
			g:    &Gradient{[]color.Color{color.White, color.Black}, Range{0, 1}},
			args: args{1},
			want: color.Black,
		},
		{
			name: "three gradient middle should return blue",
			g:    &Gradient{[]color.Color{color.White, color.RGBA{0x00, 0x00, 0xFF, 0xFF}, color.Black}, Range{0, 1}},
			args: args{0.5},
			want: color.RGBA{0x00, 0x00, 0xFF, 0xFF},
		},
		{
			name: "two gradient middle should return purple",
			g:    &Gradient{[]color.Color{color.RGBA{0xFF, 0x00, 0x00, 0xFF}, color.RGBA{0x00, 0x00, 0xFF, 0xFF}}, Range{0, 1}},
			args: args{0.5},
			want: color.RGBA{0x7F, 0x00, 0x7F, 0xFF},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.FloatToColor(tt.args.n)
			// Check for nil
			if tt.want == nil && got != nil {
				t.Errorf("Gradient.FloatToColor() = %v, want nil", got)
			}
			// Check for color equality
			wr, wg, wb, wa := tt.want.RGBA()
			r, g, b, a := got.RGBA()
			if wr != r || wg != g || wb != b || wa != a {
				t.Errorf("Gradient.FloatToColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
