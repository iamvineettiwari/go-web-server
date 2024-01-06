package status

type HttpStatus string

const (
	HTTP_200_OK           HttpStatus = "200"
	HTTP_201_CREATED      HttpStatus = "201"
	HTTP_204_NO_CONTENT   HttpStatus = "204"
	HTTP_400_BAD_REQUEST  HttpStatus = "400"
	HTTP_404_NOT_FOUND    HttpStatus = "404"
	HTTP_500_SERVER_ERROR HttpStatus = "500"
)
