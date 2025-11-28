// ============================================================
// @file: password.go
// @author: Yosemar Andrade
// @date: 2025-11-26
// @lastModified: 2025-11-26
// @description: Define reglas de validación para la contraseña.
// ============================================================

package rules

import (
	errorsDomain "api-user-go/internal/domain/user/error"
	"strings"
)

// ValidatePasswordNotEmpty verifica que la contraseña no esté vacía.
//
// Parámetros:
//   - pass: la contraseña a validar.
//
// Retorna:
//   - error: retorna error si la contraseña está vacía o solo contiene espacios.
//
// Errores:
//   - Retorna `user.ErrInvalidPassword` si la validación falla.
func ValidatePasswordNotEmpty(pass string) error {
	if strings.TrimSpace(pass) == "" {
		return errorsDomain.ErrInvalidPassword
	}
	return nil
}
