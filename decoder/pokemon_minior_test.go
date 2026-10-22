package decoder

import (
	"testing"

	"golbat/pogo"

	"github.com/guregu/null/v6"
)

// Minior's encounter values are per player, so a wild Minior gets no derived
// stats; the same sighting of another species does.
func TestDeriveIvIfAbsentSkipsMinior(t *testing.T) {
	for _, tc := range []struct {
		species pogo.HoloPokemonId
		derived bool
	}{{pogo.HoloPokemonId_MINIOR, false}, {pogo.HoloPokemonId_PIDGEY, true}} {
		p := &Pokemon{}
		p.Id = Uint64Str(8429844873227918432)
		p.PokemonId = int16(tc.species)
		p.SetWeather(null.IntFrom(int64(pogo.GameplayWeatherProto_RAINY)))
		p.SetSeenType(SeenTypeCodeWild)

		p.deriveIvIfAbsent()
		p.deriveSizeAndMovesIfAbsent()

		if p.AtkIv.Valid != tc.derived || p.Level.Valid != tc.derived || p.Size.Valid != tc.derived {
			t.Errorf("%s: atk=%v level=%v size=%v, want derived=%v", tc.species, p.AtkIv.Valid, p.Level.Valid, p.Size.Valid, tc.derived)
		}
	}
}
