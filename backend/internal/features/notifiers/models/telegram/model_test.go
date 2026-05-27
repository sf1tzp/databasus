package telegram_notifier

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type passthroughEncryptor struct{}

func (p passthroughEncryptor) Encrypt(plaintext string) (string, error) {
	return plaintext, nil
}

func (p passthroughEncryptor) Decrypt(ciphertext string) (string, error) {
	return ciphertext, nil
}

func Test_Validate_WhenHTTPProxyEnabledWithoutURL_ReturnsError(t *testing.T) {
	notifier := &TelegramNotifier{
		BotToken:           "token",
		TargetChatID:       "123456",
		IsHTTPProxyEnabled: true,
	}

	err := notifier.Validate(passthroughEncryptor{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP proxy URL is required")
}

func Test_Validate_WhenHTTPProxyURLIsNotHTTP_ReturnsError(t *testing.T) {
	notifier := &TelegramNotifier{
		BotToken:           "token",
		TargetChatID:       "123456",
		IsHTTPProxyEnabled: true,
		HTTPProxyURL:       "https://proxy.example.com:3128",
	}

	err := notifier.Validate(passthroughEncryptor{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "valid http:// URL")
}

func Test_BuildHTTPClient_WhenHTTPProxyEnabled_UsesConfiguredProxy(t *testing.T) {
	notifier := &TelegramNotifier{
		IsHTTPProxyEnabled: true,
		HTTPProxyURL:       "http://user:password@proxy.example.com:3128",
	}

	client, err := notifier.buildHTTPClient(passthroughEncryptor{})
	require.NoError(t, err)

	transport, ok := client.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, transport.Proxy)

	req, err := http.NewRequest(http.MethodGet, "https://api.telegram.org", nil)
	require.NoError(t, err)

	proxyURL, err := transport.Proxy(req)
	require.NoError(t, err)
	require.NotNil(t, proxyURL)
	assert.Equal(t, "http://user:password@proxy.example.com:3128", proxyURL.String())
}

func Test_BuildHTTPClient_WhenHTTPProxyDisabled_UsesDefaultTransport(t *testing.T) {
	notifier := &TelegramNotifier{}

	client, err := notifier.buildHTTPClient(passthroughEncryptor{})

	require.NoError(t, err)
	assert.Nil(t, client.Transport)
}
