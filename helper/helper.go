package helper

/*
jadi perbedaannya Pascal dan camelCase adalah bisa atau tidaknya diakses ketika di import
pada camelCase atau huruf kecil hanya bisa di akses di dalam satu file ini saja,
mirip private pada PHP
dan untuk yang PascalCase mirip public di PHP
*/
var version = "1.0.0"      // ini tidak bisa di akses keluar (import)
var Application = "golang" // ini bisa di akses di luar (import)

func SayHello(name string) string {
	return "Hello " + name
}
