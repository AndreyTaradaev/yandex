package contacts

import "fmt"

// Email 
const Email = "support@example.com" // глобальная экспортируемая константа

var support string // глобальная неэкспортируемая переменная

// SetSupport устанавливает значение переменной support.
func SetSupport(s string) { 
    support = s
}

// GetContact возвращает имя и email.
func GetContact() string {
    return fmt.Sprintf("%s <%s>", support, Email)
} 