package dto

import (
	"time"

	"github.com/fauziagustian/delos-aqua/internal/models"
)

type UserAgentResponse struct {
	UserAgentId uint      `json:"user_agent_id"`
	MethodUrl   string    `json:"method_url"`
	UserId      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAgentResponseCount struct {
	MethodUrl       string `json:"method_url"`
	Count           int    `json:"count"`
	UniqueUserAgent int    `json:"unique_user_agent"`
}

func FormatUserAgent(ua *models.UserAgents) UserAgentResponse {
	return UserAgentResponse{
		UserAgentId: uint(ua.UserAgentId),
		MethodUrl:   ua.MethodUrl,
		UserId:      uint(ua.UserId),
		CreatedAt:   ua.CreatedAt,
		UpdatedAt:   ua.UpdatedAt,
	}
}

func FormatUserAgents(ua []*models.UserAgents) []UserAgentResponse {
	formattedUserAgent := []UserAgentResponse{}
	for _, userAgent := range ua {
		formattedBook := FormatUserAgent(userAgent)
		formattedUserAgent = append(formattedUserAgent, formattedBook)
	}
	return formattedUserAgent
}

func FormatUserAgentCount(ua *UserAgentResponseCount) UserAgentResponseCount {
	return UserAgentResponseCount{
		MethodUrl:       ua.MethodUrl,
		Count:           ua.Count,
		UniqueUserAgent: ua.UniqueUserAgent,
	}
}

func FormatUserAgentsCount(ua []*UserAgentResponseCount) []UserAgentResponseCount {
	formattedUserAgent := []UserAgentResponseCount{}
	for _, ua := range ua {
		formattedBook := FormatUserAgentCount(ua)
		formattedUserAgent = append(formattedUserAgent, formattedBook)
	}
	return formattedUserAgent
}
