# https://just.systems/man/en/
run:
	go run .

clean-dependencies:
	# Add references to go.mod and clean unused dependencies
	go mod tidy

# Intended way
update-packages-alt:
	go get .
	# for specific packages
	# go get example.com/theirmodule
	#

list-dependencies:
	# include -u to check for available updates
	go list -m all
