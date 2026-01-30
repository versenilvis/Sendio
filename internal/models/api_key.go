// Package models - api key
package models

import (
	"time"

	"github.com/google/uuid"
)

type APIKey struct {
	ID        uuid.UUID  `json:"id"`
	AppName   string     `json:"app_name"`
	KeyHash   string     `json:"-"`
	KeyPrefix string     `json:"key_prefix"`
	IsActive  bool       `json:"is_active"`
	RateLimit int        `json:"rate_limit"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type APIKeyUsage struct {
	ID          uuid.UUID `json:"id"`
	APIKeyID    uuid.UUID `json:"api_key_id"`
	UsageDate   time.Time `json:"usage_date"` // date in DB
	SentCount   int       `json:"sent_count"`
	FailedCount int       `json:"failed_count"`
}

type SubscriptionStatus string

const (
	SubscriptionActive    SubscriptionStatus = "active"
	SubscriptionPastDue   SubscriptionStatus = "past_due"
	SubscriptionCancelled SubscriptionStatus = "cancelled"
	SubscriptionExpired   SubscriptionStatus = "expired"
	SubscriptionTrialing  SubscriptionStatus = "trialing"
)

type Subscription struct {
	ID                 uuid.UUID          `json:"id"`
	APIKeyID           uuid.UUID          `json:"api_key_id"`
	PlanID             uuid.UUID          `json:"plan_id"`
	Status             SubscriptionStatus `json:"status"`
	BillingPeriod      string             `json:"billing_period"`
	CurrentPeriodStart time.Time          `json:"current_period_start"`
	CurrentPeriodEnd   time.Time          `json:"current_period_end"`
	QuotaLimit         int                `json:"quota_limit"`
	QuotaUsed          int                `json:"quota_used"`
	QuotaResetAt       time.Time          `json:"quota_reset_at"`
	PaymentProvider    string             `json:"payment_provider"`
	ProviderSubID      *string            `json:"provider_sub_id,omitempty"`
	PaymentMetadata    map[string]any     `json:"payment_metadata"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	CancelledAt        *time.Time         `json:"cancelled_at,omitempty"`
}

type Plan struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	MaxQuota  int       `json:"max_quota"`
	RateLimit int       `json:"rate_limit"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
