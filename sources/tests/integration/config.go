package integration

import "time"

// Service URLs
const (
	ContentServiceURL = "http://localhost:8081"
	ScoringServiceURL = "http://localhost:8082"
	LearnerModelURL   = "http://localhost:8080"
	AdaptiveEngineURL = "http://localhost:8084"
)

// Timeouts and intervals
const (
	MasteryUpdateTimeout = 10 * time.Second
	PollingInterval      = 200 * time.Millisecond
	HTTPTimeout          = 5 * time.Second
)

// Test users
const (
	TestUser1 = "integration-test-user-1"
	TestUser2 = "integration-test-user-2"
	TestUser3 = "integration-test-user-3"
)

// Skills
const (
	SkillMath    = "math"
	SkillScience = "science"
)
