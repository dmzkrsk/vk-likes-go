package vk

type VkApiError struct {
	Code int `json:"error_code"`
	Message string `json:"error_msg"`
}

func (e VkApiError) Error() string {
	return e.Message
}

type VkError struct {
	message string
}

func (e VkError) Error() string {
	return e.message
}

func newError(m string) VkError {
	return VkError{message: m}
}

type VkAuthError struct {
	message string
}

func (e VkAuthError) Error() string {
	return e.message
}

func newAuthError(m string) VkAuthError {
	return VkAuthError{message: m}
}
