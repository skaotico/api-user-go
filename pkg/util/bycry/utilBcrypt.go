package bycry

import (
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = bcrypt.DefaultCost // 10

// Opcional: pepper recomendado (clave secreta global).
// Puedes comentarlo si no lo usarás.
var pepper = os.Getenv("PEPPER_SECRET")

// HashPassword genera un hash bcrypt seguro.
// Si no pasas cost, usará DefaultCost.
func HashPassword(password string, cost int) (string, error) {
	if cost <= 0 {
		cost = DefaultCost
	}

	// Agregar pepper si está definido
	if pepper != "" {
		password += pepper
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashBytes), nil
}

// CheckPassword verifica si el password coincide con el hash almacenado.
func CheckPassword(hash string, password string) error {
	if pepper != "" {
		password += pepper
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// NeedsRehash permite saber si un hash viejo
// debe ser re-hasheado con un cost más alto.
func NeedsRehash(hash string, desiredCost int) bool {
	currentCost, err := bcrypt.Cost([]byte(hash))
	if err != nil {
		return true
	}
	return currentCost < desiredCost
}

// BenchmarkCost mide cuánto se demora un cost
// para elegir el ideal para producción.
func BenchmarkCost(cost int) (time.Duration, error) {
	start := time.Now()
	_, err := bcrypt.GenerateFromPassword([]byte("testPassword123!"), cost)
	if err != nil {
		return 0, err
	}
	return time.Since(start), nil
}
