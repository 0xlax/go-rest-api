package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)
func RequestFollowUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// funciton to request a follow to a user
}
func AcceptFollowRequest(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// funciton to Accept Follow Requests
}
func FollowRequests(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// function to deploy follow request
}