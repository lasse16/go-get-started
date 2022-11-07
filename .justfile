# https://just.systems/man/en/
run:
	go run .

update-packages:
	go mod tidy

# Intended way
update-packages-alt:
	go get .
	# for specific packages
	# go get example.com/theirmodule
