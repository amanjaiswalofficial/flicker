module demo

go 1.20

replace job_submitter => ../job_submitter

require job_submitter v0.0.0-00010101000000-000000000000

require (
	github.com/go-gota/gota v0.12.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	gonum.org/v1/gonum v0.13.0 // indirect
)
