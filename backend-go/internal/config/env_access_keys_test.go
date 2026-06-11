package config

import "testing"

func TestNewEnvConfig_ExtendAccessKeysParsesCommaSeparatedKeys(t *testing.T) {
	t.Setenv("PROXY_ACCESS_KEY", "primary-key")
	t.Setenv("EXTEND_ACCESS_KEY", "client-a, client-b,,client-a")

	envCfg := NewEnvConfig()

	want := []string{"client-a", "client-b"}
	if len(envCfg.ExtendAccessKeys) != len(want) {
		t.Fatalf("ExtendAccessKeys length = %d, want %d: %#v", len(envCfg.ExtendAccessKeys), len(want), envCfg.ExtendAccessKeys)
	}
	for i, key := range want {
		if envCfg.ExtendAccessKeys[i] != key {
			t.Fatalf("ExtendAccessKeys[%d] = %q, want %q: %#v", i, envCfg.ExtendAccessKeys[i], key, envCfg.ExtendAccessKeys)
		}
	}
}

func TestEnvConfig_IsValidProxyAccessKey(t *testing.T) {
	envCfg := &EnvConfig{
		ProxyAccessKey:   "primary-key",
		ExtendAccessKeys: []string{"client-a", "client-b"},
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
