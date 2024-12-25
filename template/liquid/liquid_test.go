package liquid

import (
	"github.com/osteele/liquid"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBasicUsage(t *testing.T) {
	engine := liquid.NewEngine()
	template := `{{ page.title }}`
	bindings := map[string]any{
		"page": map[string]string{
			"title": "hello world",
		},
	}
	out, err := engine.ParseAndRenderString(template, bindings)
	assert.Equal(t, nil, err)
	assert.Equal(t, "hello world", out)
}

func TestIfStmt(t *testing.T) {
	template := `{% if user != null %}
		hello {{ user.name }}
	{% endif %}`
	bindings := map[string]any{
		"user": map[string]string{
			"name": "Peter",
		},
	}
	engine := liquid.NewEngine()
	out, err := engine.ParseAndRenderString(template, bindings)
	assert.Equal(t, nil, err)
	assert.Equal(t, "hello Peter", strings.TrimSpace(out))
}
