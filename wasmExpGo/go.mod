module github.com/hollowaykeanho/ExperimentingGoWASM/wasmExpGo

go 1.18

replace (
	github.com/hollowaykeanho/ExperimentingGoWASM/wasmExpGo => ./
	hestiaGo => ./hestiaGo
	wasmExpGo => ./
)

require hestiaGo v0.0.0-00010101000000-000000000000
