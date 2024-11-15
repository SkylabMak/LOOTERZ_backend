package utils

import "LOOTERZ_backend/models/types"

const (
	// General Errors
	// you can use and see in fiber like "fiber.StatusUnauthorized" for standard http status code
	//please add in drawn.io too if u modify
	ErrInternal           types.ErrorCode = "500"
	ErrBadReq             types.ErrorCode = "400"
	ErrDatabaseConnection types.ErrorCode = "ERR_DATABASE_CONNECTION"

	// custom error code
	ErrMissCondition types.ErrorCode = "461"
	ErrNotFound types.ErrorCode = "462"
	ErrContentType types.ErrorCode = "499"

	//room
	ErrRoomPassword types.ErrorCode = "601"
	
)
