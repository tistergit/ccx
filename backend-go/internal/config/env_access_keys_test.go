package config

import "testing"

func TestNewEnvConfig_ProxyAccessKeysIncludesPrimaryAndAdditionalKeys(t *testing.T) {
	t.Setenv("PROXY_ACCESS_KEY", "primary-key")
	t.Setenv("PROXY_ACCESS_KEYS", "client-a, client-b\nclient-c,,client-a")

	envCfg := NewEnvConfig()

	want := []string{"primary-key", "client-a", "client-b", "client-c"}
	if len(envCfg.ProxyAccessKeys) != len(want) {
		t.Fatalf("ProxyAccessKeys length = %d, want %d: %#v", len(envCfg.ProxyAccessKeys), len(want), envCfg.ProxyAccessKeys)
	}
	for i, key := range want {
		if envCfg.ProxyAccessKeys[i] != key {
			t.Fatalf("ProxyAccessKeys[%d] = %q, want %q: %#v", i, envCfg.ProxyAccessKeys[i], key, envCfg.ProxyAccessKeys)
		}
	}
}

func TestEnvConfig_IsValidProxyAccessKey(t *testing.T) {
	envCfg := &EnvConfig{
		ProxyAccessKey:  "primary-key",
		ProxyAccessKeys: []string{"primary-key", "client-a", "client-b"},
	}

	tests := []struct {
		name string
		key  string
		want bool
	}{
		{name: "primary key", key: "primary-key", want: true},
		{name: "additional key", key: "client-a", want: true},
		{name: "unknown key", key: "missing", want: false},
		{name: "empty key", key: "", want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := envCfg.IsValidProxyAccessKey(tt.key); got != tt.want {
				t.Fatalf("IsValidProxyAccessKey(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestEnvConfig_IsValidProxyAccessKeyHonorsUpdatedPrimaryKey(t *testing.T) {
	envCfg := NewEnvConfig()
	envCfg.ProxyAccessKey = "updated-key"

	if !envCfg.IsValidProxyAccessKey("updated-key") {
		t.Fatal("IsValidProxyAccessKey should accept ProxyAccessKey when it is updated after NewEnvConfig")
	}
}
