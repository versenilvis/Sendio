package worker

import (
	"fmt"
	"time"

	"github.com/versenilvis/sendio/internal/crypto"
	"github.com/versenilvis/sendio/internal/logger"
	"github.com/versenilvis/sendio/internal/models"
	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
)

// sendEmail sends an email via SMTP using the provided config
func (p *Processor) sendEmail(config *models.SMTPConfig, msg *mail.Msg) error {
	opts := []mail.Option{
		mail.WithPort(config.Port),
		mail.WithTimeout(5 * time.Second),
	}

	// configure TLS policy
	if config.UseSSL {
		opts = append(opts, mail.WithSSL())
	} else if config.UseTLS {
		opts = append(opts, mail.WithTLSPolicy(mail.TLSMandatory))
	} else {
		opts = append(opts, mail.WithTLSPolicy(mail.NoTLS))
	}

	// add auth if credentials exist
	if config.Username != nil && config.PasswordEnc != nil {
		password, err := crypto.Decrypt(*config.PasswordEnc, p.encryptionKey)
		if err != nil {
			return fmt.Errorf("failed to decrypt smtp password: %w", err)
		}
		opts = append(opts, mail.WithSMTPAuth(mail.SMTPAuthPlain))
		opts = append(opts, mail.WithUsername(*config.Username))
		opts = append(opts, mail.WithPassword(password))
	}

	client, err := mail.NewClient(config.Host, opts...)
	if err != nil {
		return fmt.Errorf("failed to create smtp client: %w", err)
	}

	logger.Log.Info("sending email via SMTP",
		zap.String("host", config.Host),
		zap.Int("port", config.Port),
	)

	return client.DialAndSend(msg)
}
