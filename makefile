.PHONY: gorm
gorm:
	go run ./gorm

.PHONY: restAPI
restAPI:
	go run ./restAPI

auth-basic:
	go run ./authentication/basic/
auth-jwt:
	go run ./authentication/jwt/

quicky:
	go run ./quicktest/