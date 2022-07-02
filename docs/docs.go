// Package classification Attendance Platform API.
//
// This is a API for employee to to check in attendance
//
//   Schemes: https
//   Host: attendance-platform.herokuapp.com
//   BasePath: /
//   Version: 1.0.0
//   Contact: muhammadarash1997@gmail.com
//
//   Consumes:
//   - application/json
//
//   Produces:
//   - application/json
//
//   SecurityDefinitions:
//   Bearer:
//    description: Type 'Bearer' before you enter your token. ex = Bearer tokentokentoken
//    type: apiKey
//    name: Authorization
//    in: header
//
// swagger:meta
package docs

import "attendance-platform/dto"

// Success testing API
// swagger:response testAPI
type testAPI struct {
	// in: Body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:response errorResponse
type errorResponse struct {
	// in: Body
	Body struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   string `json:"data"`
	}
}

// Success registering an account
// swagger:response registerUser
type registerUser struct {
	// in: Body
	Body struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   string `json:"data"`
	}
}

// Success logging in an account
// swagger:response loginUser
type loginUser struct {
	// in: Body
	Body struct {
		Code   int               `json:"code"`
		Status string            `json:"status"`
		Data   dto.LoginResponse `json:"data"`
	}
}

// Success logging out an account
// swagger:response logoutUser
type logoutUser struct {
	// in: Body
	Body struct {
		Code   int                `json:"code"`
		Status string             `json:"status"`
		Data   dto.LogoutResponse `json:"data"`
	}
}

// Success checking in for attendance
// swagger:response checkIn
type checkIn struct {
	// in: Body
	Body struct {
		Code   int                 `json:"code"`
		Status string              `json:"status"`
		Data   dto.CheckInResponse `json:"data"`
	}
}

// Success checking out for attendance
// swagger:response checkOut
type checkOut struct {
	// in: Body
	Body struct {
		Code   int                  `json:"code"`
		Status string               `json:"status"`
		Data   dto.CheckOutResponse `json:"data"`
	}
}

// Success getting all attendance of user
// swagger:response getAllUserAttendance
type getAllUserAttendance struct {
	// in: Body
	Body struct {
		Code   int                              `json:"code"`
		Status string                           `json:"status"`
		Data   dto.GetAllUserAttendanceResponse `json:"data"`
	}
}

// Success creating an activity of attendance
// swagger:response createActivity
type createActivity struct {
	// in: Body
	Body struct {
		Code   int                        `json:"code"`
		Status string                     `json:"status"`
		Data   dto.CreateActivityResponse `json:"data"`
	}
}

// Success updating an activity of attendance
// swagger:response updateActivity
type updateActivity struct {
	// in: Body
	Body struct {
		Code   int                        `json:"code"`
		Status string                     `json:"status"`
		Data   dto.UpdateActivityResponse `json:"data"`
	}
}

// Success deleting an activity of attendance
// swagger:response deleteActivity
type deleteActivity struct {
	// in: Body
	Body struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   string `json:"data"`
	}
}

// Success getting user activities by date
// swagger:response getUserActivitiesByDate
type getUserActivitiesByDate struct {
	// in: Body
	Body struct {
		Code   int                                 `json:"code"`
		Status string                              `json:"status"`
		Data   dto.GetUserActivitiesByDateResponse `json:"data"`
	}
}

// swagger:parameters registerUser
type registerUserParams struct {
	// User object that needs to be registered
	// in: body
	// required: true
	Body dto.RegisterRequest
}

// swagger:parameters loginUser
type loginUserParams struct {
	// User login object that needs to be logged in
	// in: body
	// required: true
	Body dto.LoginRequest
}

// swagger:parameters checkOut
type checkOutParams struct {
	// The attendance id that needs to be checked out
	// in: path
	// required: true
	ID string `json:"attendance_id"`
}

// swagger:parameters createActivity
type createActivityParams struct {
	// Activity object that needs to be created
	// in: body
	// required: true
	Body dto.CreateActivityRequest
}

// swagger:parameters UpdateActivity
type updateActivityParams struct {
	// Activity object that needs to be updated
	// in: body
	// required: true
	Body dto.UpdateActivityRequest
}

// swagger:parameters deleteActivity
type deleteActivityParams struct {
	// The activity id that needs to be deleted
	// in: path
	// required: true
	ID string `json:"activity_id"`
}

// swagger:parameters getUserActivitiesByDate
type getUserActivitiesByDateParams struct {
	// The date that want to be used as parameter for getting activities
	// in: path
	// required: true
	ID string `json:"date"`
}
