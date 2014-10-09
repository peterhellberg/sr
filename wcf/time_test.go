package wcf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeUnmarshalJSON(t *testing.T) {

	tt := Time{}
	in := []byte(`\/Date(1411653960000)\/`)

	tt.UnmarshalJSON(in)

	assert.Equal(t, "2014-09-25 14:06:00 +0000 UTC", tt.UTC().String())
}
