module github.com/BielPinto/curso_go/5-packaging/3-go-mod/system

go 1.22.5

replace github.com/BielPinto/curso_go/5-packaging/3-go-mod/math => ../math

require github.com/BielPinto/curso_go/5-packaging/3-go-mod/math v0.0.0-00010101000000-000000000000
