module github.com/duhnnie/soccerclub-scoring

go 1.21.6

replace github.com/duhnnie/jexp => ../jexp

replace github.com/duhnnie/valuebox => ../valuebox

replace github.com/duhnnie/godash => ../godash

require (
	github.com/duhnnie/godash v0.3.0
	github.com/duhnnie/jexp v0.2.0
	github.com/duhnnie/valuebox v0.1.2
)

require github.com/json-e/json-e/v4 v4.8.0
