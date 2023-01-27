package benchmark

import (
	"fmt"
	"testing"
)

//func BenchmarkUnmarshalGoEncoding(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		UnmarshalEncoding()
//	}
//
//}
//
//func BenchmarkUnmarshalGoJson(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		UnmarshalGoJson()
//	}
//
//}
//
//func BenchmarkUnmarshalAndMarshalGoEncoding(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		UnmarshalAndMarshalEncoding()
//	}
//
//}
//
//func BenchmarkUnmarshalAndMarshalGoJson(b *testing.B) {
//
//	for i := 0; i < b.N; i++ {
//		UnmarshalAndMarshalGoJson()
//	}
//
//}

func Benchmark(b *testing.B) {
	fileName := []string{"1000.json", "2000.json", "4000.json"}
	for _, l := range fileName {
		b.Run(fmt.Sprintf("UnmarshalEncoding-%s", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UnmarshalEncoding(l)
			}
		})

		b.Run(fmt.Sprintf("UnmarshalGoJson-%s", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UnmarshalGoJson(l)
			}
		})
		b.Run(fmt.Sprintf("NewDecoderEncoding-%s", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewDecoderEncoding(l)
			}
		})
		b.Run(fmt.Sprintf("NewDecoderGoJson-%s", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewDecoderGoJson(l)
			}
		})
		//b.Run(fmt.Sprintf("UnmarshalAndMarshalEncoding-%s", l), func(b *testing.B) {
		//	for i := 0; i < b.N; i++ {
		//		UnmarshalAndMarshalEncoding(l)
		//	}
		//})
		//b.Run(fmt.Sprintf("UnmarshalAndMarshalGoJson-%s", l), func(b *testing.B) {
		//	for i := 0; i < b.N; i++ {
		//		UnmarshalAndMarshalGoJson(l)
		//	}
		//})
	}
}
