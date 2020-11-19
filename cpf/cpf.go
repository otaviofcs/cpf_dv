package cpf

import (
	"errors"
	"strconv"
)

func CpfComDV(cpf string) (string, error) {
	dvcpf, err := CalculaDVCPF(cpf)
	if err != nil {
		return "", err
	}
	return cpf + dvcpf, nil
}

func CalculaDVCPF(cpf string) (string, error) {
	if len(cpf) != 9 {
		return "", errors.New("Tamanho de CPF sem DV deve ser 9")
	}
	primSoma, err := SomaCPF(cpf)
	if err != nil {
		return "", err
	}
	primDv := RestoDV11(primSoma)
	segSoma, err2 := SomaCPF(cpf + strconv.Itoa(primDv))
	if err2 != nil {
		return "", err2
	}
	segDv := RestoDV11(segSoma)
	return strconv.Itoa(primDv) + strconv.Itoa(segDv), nil
}

func SomaCPF(cpf string) (int, error) {
	if len(cpf) != 9 && len(cpf) != 10 {
		return 0, errors.New("Tamanho de CPF sem DV completo deve ser 10 ou 9 " + cpf)
	}
	multiplo := 1
	total := 0
	for i := len(cpf) - 1; i >= 0; i-- {
		multiplo++
		total += (int(cpf[i]) - int('0')) * multiplo
	}
	return total, nil
}

func RestoDV11(valorTotal int) int {
	resto11 := valorTotal % 11
	resto := 11 - resto11
	if resto >= 10 {
		return 0
	}
	return resto
}
