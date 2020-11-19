package cpf

import (
	"strconv"
	"testing"
)

func TestCpfComDV(t *testing.T) {
	t.Run("Calcula CPF com DV", func(t *testing.T) {
		cpfsATestar := map[string]string{"296539278": "29653927850", "069319878": "06931987887", "453178287": "45317828791"}
		for cpfATestar, resultado := range cpfsATestar {
			cpffinal, err := CpfComDV(cpfATestar)
			if cpffinal != resultado {
				t.Errorf("O CPF %s não bate com o resultado esperado %s", cpffinal, resultado)
			}
			if err != nil {
				t.Errorf("Não deveria ter dado erro '%s' no CPF: %s", err, cpfATestar)
			}
		}
	})
}

func TestCalculaDVCPF(t *testing.T) {
	t.Run("Calcula DV para CPF com menos de 9 números deve retornar erro", func(t *testing.T) {
		for _, cpfATestar := range [3]string{"02110202", "12345678901", "1234567890"} {
			_, err := CalculaDVCPF(cpfATestar)
			if err == nil {
				t.Errorf("Deveria ter dado erro no CPF: %s", cpfATestar)
			}
		}
	})
	t.Run("Calcula DV para CPF correto", func(t *testing.T) {
		cpfsATestar := map[string]string{"987654321": "00", "123456789": "09", "021913247": "03", "121913247": "03"}
		for cpfATestar, resultado := range cpfsATestar {
			dvfinal, err := CalculaDVCPF(cpfATestar)
			if dvfinal != resultado {
				t.Errorf("O DV do CPF %s não bate com o resultado esperado %s", dvfinal, resultado)
			}
			if err != nil {
				t.Errorf("Não deveria ter dado erro '%s' no CPF: %s", err, cpfATestar)
			}
		}
	})
}

func TestSomaCPF(t *testing.T) {
	t.Run("Soma CPF com numero diferente de 10 ou 9 deve retornar erro", func(t *testing.T) {
		for _, cpfATestar := range [2]string{"02110202", "12345678901"} {
			_, err := SomaCPF(cpfATestar)
			if err == nil {
				t.Errorf("Deveria ter dado erro no CPF: %s", cpfATestar)
			}
		}
	})

	t.Run("Soma CPF com multiplicação do cálculo DV", func(t *testing.T) {
		cpfsATestar := map[string]int{"987654321": 330, "123456789": 210, "0219132470": 173, "1219132470": 184}
		for cpfATestar, resultado := range cpfsATestar {
			soma, err := SomaCPF(cpfATestar)
			if soma != resultado {
				t.Errorf("A soma %d não bate com o resultado esperado %d", soma, resultado)
			}
			if err != nil {
				t.Errorf("Não deveria ter dado erro no CPF: %s", cpfATestar)
			}
		}
	})
}

func TestRestoDV11(t *testing.T) {
	t.Run("DV11 zerado", func(t *testing.T) {
		r12 := RestoDV11(12)
		r11 := RestoDV11(11)
		if r12 != 0 || r11 != 0 {
			t.Errorf("Esperado 0 para todos. Resultados: %d e %d", r11, r12)
		}
	})

	t.Run("DV11 2", func(t *testing.T) {
		r9 := RestoDV11(9)
		if r9 != 2 {
			t.Errorf("Esperado 2 para DV. Resultados: %d", r9)
		}
	})
	t.Run("DV11 9", func(t *testing.T) {
		r13 := RestoDV11(13)
		if r13 != 9 {
			t.Errorf("Esperado 9 para DV. Resultados: %d", r13)
		}
	})
	t.Run("DV11 6", func(t *testing.T) {
		r16 := RestoDV11(16)
		if r16 != 6 {
			t.Errorf("Esperado 6 para DV. Resultados: %d", r16)
		}
	})
}

func BenchmarkAdiciona(b *testing.B) {
	var cpfATestar string
	for i := 0; i < b.N; i++ {
		c := 123456789 + i
		cpfATestar = strconv.Itoa(c)
		CpfComDV(cpfATestar)
	}
}
