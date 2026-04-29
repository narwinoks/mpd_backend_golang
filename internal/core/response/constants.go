package response

import "net/http"

type Status struct {
	HttpCode int
	RC       string
	Message  string
}

var (
	Success       = Status{http.StatusOK, "200", "Successfully"}
	SuccessCreate = Status{http.StatusCreated, "201", "Successfully Create"}
	SuccessUpdate = Status{http.StatusCreated, "201", "Successfully Update"}
	SuccessDelete = Status{http.StatusCreated, "201", "Successfully Delete"}
)

var (
	//400 Bad Request
	BadRequest   = Status{http.StatusBadRequest, "400", "Bad Request"}
	Validation   = Status{http.StatusBadRequest, "400", "Validation Error"}
	AlreadyLogin = Status{http.StatusBadRequest, "400", "Sudah Login di Device Lain"}

	// 401 & 403 Auth
	Unauthorized = Status{http.StatusUnauthorized, "401", "Unauthenticated"}
	Forbidden    = Status{http.StatusForbidden, "403", "You are not authorized to access this resource"}

	// 404 Not Found
	NotFound      = Status{http.StatusNotFound, "404", "Not Found"}
	DataNotFound  = Status{http.StatusNotFound, "404", "Data Not Found"}
	PathNotFound  = Status{http.StatusNotFound, "404", "Path Not Found"}
	UserNotFound  = Status{http.StatusNotFound, "404", "Data User Tidak Ditemukan"}
	PasswordSalah = Status{http.StatusNotFound, "404", "Password Yang Dimasukkan Tidak Sesuai"}

	// 405 Method
	MethodNotAllowed = Status{http.StatusMethodNotAllowed, "405", "Method Not Allowed"}

	// 409 Conflict
	Conflict            = Status{http.StatusConflict, "409", "Data Already Exist"}
	VerificationAlready = Status{http.StatusConflict, "409", "Data Has Been Verified"}

	// 422 Unprocessable Entity
	UnprocessableEntity = Status{http.StatusUnprocessableEntity, "422", "Unprocessable Entity"}
	DataAlreadyUsed     = Status{http.StatusUnprocessableEntity, "422", "Data sudah digunakan, tidak bisa dihapus atau dinonaktifkan"}

	// 429 Too Many Requests
	TooManyRequests = Status{http.StatusTooManyRequests, "429", "Too Many Requests"}

	// 500 Server Error
	InternalServerError = Status{http.StatusInternalServerError, "500", "Internal Server Error."}
	DataBridgingNotSend = Status{http.StatusInternalServerError, "500", "Send Data Bridging Failed"}
)
