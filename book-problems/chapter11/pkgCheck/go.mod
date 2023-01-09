module example.com/chapter11/pkgCheck

go 1.18

replace example.com/chapter11/math => ../math

require example.com/chapter11/math v0.0.0-00010101000000-000000000000
