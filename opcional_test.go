package opcional

import (
	"testing"
)

func TestOpcionalTieneValor(t *testing.T) {
	tests := []struct {
		name     string
		opcional Opcional[string]
		want     bool
	}{
		{
			name:     "opcional vacío debe retornar falso",
			opcional: Opcional[string]{vacio: true},
			want:     false,
		},
		{
			name:     "opcional con valor debe retornar verdadero",
			opcional: Opcional[string]{valor: "test", vacio: false},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opcional.TieneValor(); got != tt.want {
				t.Errorf("TieneValor() = %v, quiere %v", got, tt.want)
			}
		})
	}
}

func TestOpcionalConsumir(t *testing.T) {
	tests := []struct {
		name           string
		opcional       Opcional[int]
		wantValor      int
		wantTieneValor bool
	}{
		{
			name:           "opcional vacío debe retornar valor cero y falso",
			opcional:       Opcional[int]{vacio: true},
			wantValor:      0,
			wantTieneValor: false,
		},
		{
			name:           "opcional con valor debe retornar valor y verdadero",
			opcional:       Opcional[int]{valor: 42, vacio: false},
			wantValor:      42,
			wantTieneValor: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valor, tieneValor := tt.opcional.Consumir()
			if valor != tt.wantValor {
				t.Errorf("Consumir() valor = %v, quiere %v", valor, tt.wantValor)
			}
			if tieneValor != tt.wantTieneValor {
				t.Errorf("Consumir() tieneValor = %v, quiere %v", tieneValor, tt.wantTieneValor)
			}
		})
	}
}

func TestOpcionalValorOCero(t *testing.T) {
	tests := []struct {
		name     string
		opcional Opcional[string]
		want     string
	}{
		{
			name:     "opcional vacío debe retornar valor cero",
			opcional: Opcional[string]{vacio: true},
			want:     "",
		},
		{
			name:     "opcional con valor debe retornar el valor",
			opcional: Opcional[string]{valor: "test", vacio: false},
			want:     "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opcional.ValorOCero(); got != tt.want {
				t.Errorf("ValorOCero() = %v, quiere %v", got, tt.want)
			}
		})
	}
}

func TestOpcionalValorO(t *testing.T) {
	tests := []struct {
		name       string
		opcional   Opcional[int]
		porDefecto int
		want       int
	}{
		{
			name:       "opcional vacío debe retornar valor por defecto",
			opcional:   Opcional[int]{vacio: true},
			porDefecto: 99,
			want:       99,
		},
		{
			name:       "opcional con valor debe retornar el valor",
			opcional:   Opcional[int]{valor: 42, vacio: false},
			porDefecto: 99,
			want:       42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opcional.ValorO(tt.porDefecto); got != tt.want {
				t.Errorf("ValorO() = %v, quiere %v", got, tt.want)
			}
		})
	}
}

func TestOpcionalValorOFunc(t *testing.T) {
	funcionPorDefecto := func(args []any) string {
		return "porDefecto"
	}

	tests := []struct {
		name     string
		opcional Opcional[string]
		f        func([]any) string
		args     []any
		want     string
	}{
		{
			name:     "opcional vacío debe ejecutar la función",
			opcional: Opcional[string]{vacio: true},
			f:        funcionPorDefecto,
			args:     []any{},
			want:     "porDefecto",
		},
		{
			name:     "opcional con valor debe retornar el valor",
			opcional: Opcional[string]{valor: "test", vacio: false},
			f:        funcionPorDefecto,
			args:     []any{},
			want:     "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opcional.ValorOFunc(tt.f, tt.args); got != tt.want {
				t.Errorf("ValorOFunc() = %v, quiere %v", got, tt.want)
			}
		})
	}
}
