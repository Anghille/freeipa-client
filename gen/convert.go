package main

import (
	"strings"
	"unicode"

)

var reservedWords = []string{
	"continue",
	"break",
	"return",
	"type",
	"map",
}

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string {
        var result string
        var words []string
        var lastPos int
        rs := []rune(s)

        for i := 0; i < len(rs); i++ {
                if i > 0 && unicode.IsUpper(rs[i]) {
                        if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
                                words = append(words, initialism)

                                i += len(initialism) - 1
                                lastPos = i
                                continue
                        }

                        words = append(words, s[lastPos:i])
                        lastPos = i
                }
        }

        // append the last word
        if s[lastPos:] != "" {
                words = append(words, s[lastPos:])
        }

        for k, word := range words {
                if k > 0 {
                        result += "_"
                }

                result += strings.ToLower(word)
        }

        return result
}

func snakeToCamel(s string, upperCase bool) string {
        var result string

        words := strings.Split(s, "_")

        for i, word := range words {
                if exception := snakeToCamelExceptions[word]; len(exception) > 0 {
                        result += exception
                        continue
                }

                if upperCase || i > 0 {
                        if upper := strings.ToUpper(word); commonInitialisms[upper] {
                                result += upper
                                continue
                        }
                }

                if (upperCase || i > 0) && len(word) > 0 {
                        w := []rune(word)
                        w[0] = unicode.ToUpper(w[0])
                        result += string(w)
                } else {
                        result += word
                }
        }package main

		import (
			"strings"
			"unicode"
		
		)
		
		var reservedWords = []string{
			"continue",
			"break",
			"return",
			"type",
			"map",
		}
		
		// CamelToSnake converts a given string to snake case
		func CamelToSnake(s string) string {
				var result string
				var words []string
				var lastPos int
				rs := []rune(s)
		
				for i := 0; i < len(rs); i++ {
						if i > 0 && unicode.IsUpper(rs[i]) {
								if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
										words = append(words, initialism)
		
										i += len(initialism) - 1
										lastPos = i
										continue
								}
		
								words = append(words, s[lastPos:i])
								lastPos = i
						}
				}
		
				// append the last word
				if s[lastPos:] != "" {
						words = append(words, s[lastPos:])
				}
		
				for k, word := range words {
						if k > 0 {
								result += "_"
						}
		
						result += strings.ToLower(word)
				}
		
				return result
		}
		
		func snakeToCamel(s string, upperCase bool) string {
				var result string
		
				words := strings.Split(s, "_")
		
				for i, word := range words {
						if exception := snakeToCamelExceptions[word]; len(exception) > 0 {
								result += exception
								continue
						}
		
						if upperCase || i > 0 {
								if upper := strings.ToUpper(word); commonInitialisms[upper] {
										result += upper
										continue
								}
						}
		
						if (upperCase || i > 0) && len(word) > 0 {
								w := []rune(word)
								w[0] = unicode.ToUpper(w[0])
								result += string(w)
						} else {
								result += word
						}
				}
		
				return result
		}
		
		// SnakeToCamel returns a string converted from snake case to uppercase
		func SnakeToCamel(s string) string {
				return snakeToCamel(s, true)
		}
		
		// SnakeToCamelLower returns a string converted from snake case to lowercase
		func SnakeToCamelLower(s string) string {
				return snakeToCamel(s, false)
		}
		
		// startsWithInitialism returns the initialism if the given string begins with it
		func startsWithInitialism(s string) string {
				var initialism string
				// the longest initialism is 5 char, the shortest 2
				for i := 1; i <= 5; i++ {
						if len(s) > i-1 && commonInitialisms[s[:i]] {
								initialism = s[:i]
						}
				}
				return initialism
		}
		
		// commonInitialisms, taken from
		// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
		var commonInitialisms = map[string]bool{
				"ACL":   true,
				"API":   true,
				"ASCII": true,
				"CPU":   true,
				"CSS":   true,
				"DNS":   true,
				"EOF":   true,
				"ETA":   true,
				"GPU":   true,
				"GUID":  true,
				"HTML":  true,
				"HTTP":  true,
				"HTTPS": true,
				"ID":    true,
				"IP":    true,
				"JSON":  true,
				"LHS":   true,
				"OS":    true,
				"QPS":   true,
				"RAM":   true,
				"RHS":   true,
				"RPC":   true,
				"SLA":   true,
				"SMTP":  true,
				"SQL":   true,
				"SSH":   true,
				"TCP":   true,
				"TLS":   true,
				"TTL":   true,
				"UDP":   true,
				"UI":    true,
				"UID":   true,
				"UUID":  true,
				"URI":   true,
				"URL":   true,
				"UTF8":  true,
				"VM":    true,
				"XML":   true,
				"XMPP":  true,
				"XSRF":  true,
				"XSS":   true,
				"OAuth": true,
		}
		
		// add exceptions here for things that are not automatically convertable
		var snakeToCamelExceptions = map[string]string{
				"oauth": "OAuth",
		}
		
		
		func safeName(s string) string {
			for _, r := range reservedWords {
				if r == s {
					return "_" + s
				}
			}
			return s
		}
		
		func toGoType(ipaType string) string {
			switch ipaType {
			case "":
				return "interface{}"
			case "dict":
				return "interface{}"
			case "object":
				return "interface{}"
			case "NoneType":
				return "interface{}"
			case "unicode":
				return "string"
			case "str":
				return "string"
			case "bytes":
				return "string"
			case "datetime":
				return "time.Time"
			case "DN":
				return "string"
			case "Principal":
				return "string"
			case "DNSName":
				return "string"
			case "Decimal":
				return "float64"
			case "Certificate":
				return "interface{}"
			case "CertificateSigningRequest":
				return "string"
			default:
				return ipaType
			}
		}
		
		func (t *CommandOutput) GoType(parent *Command) string {
			if t.Type == "dict_guess_receiver" {
				cls := strings.Split(parent.ObjClass, "/")[0]
				if cls != "" {
					return upperName(cls)
				}
			}
			return toGoType(t.Type)
		}
		
		func lowerName(s string) string {
			return safeName(SnakeToCamelLower(s))
		}
		
		func upperName(s string) string {
			return safeName(SnakeToCamel(s))
		}
		
		func (t *Topic) LowerName() string {
			return lowerName(t.Name)
		}
		
		func (t *Topic) UpperName() string {
			return upperName(t.Name)
		}
		
		func (t *Class) LowerName() string {
			return lowerName(t.Name)
		}
		
		func (t *Class) UpperName() string {
			return upperName(t.Name)
		}
		
		func (t *Command) LowerName() string {
			return lowerName(t.Name)
		}
		
		func (t *Command) UpperName() string {
			return upperName(t.Name)
		}
		
		func (t *Param) LowerName() string {
			return lowerName(t.Name)
		}
		
		func (t *Param) UpperName() string {
			return upperName(t.Name)
		}
		
		func (t *CommandOutput) LowerName() string {
			return lowerName(t.Name)
		}
		
		func (t *CommandOutput) UpperName() string {
			return upperName(t.Name)
		}
		

        return result
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string) string {
        return snakeToCamel(s, true)
}

// SnakeToCamelLower returns a string converted from snake case to lowercase
func SnakeToCamelLower(s string) string {
        return snakeToCamel(s, false)
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
        var initialism string
        // the longest initialism is 5 char, the shortest 2
        for i := 1; i <= 5; i++ {
                if len(s) > i-1 && commonInitialisms[s[:i]] {
                        initialism = s[:i]
                }
        }
        return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/206c0f020eba0f7fbcfbc467a5eb808037df2ed6/lint.go#L731
var commonInitialisms = map[string]bool{
        "ACL":   true,
        "API":   true,
        "ASCII": true,
        "CPU":   true,
        "CSS":   true,
        "DNS":   true,
        "EOF":   true,
        "ETA":   true,
        "GPU":   true,
        "GUID":  true,
        "HTML":  true,
        "HTTP":  true,
        "HTTPS": true,
        "ID":    true,
        "IP":    true,
        "JSON":  true,
        "LHS":   true,
        "OS":    true,
        "QPS":   true,
        "RAM":   true,
        "RHS":   true,
        "RPC":   true,
        "SLA":   true,
        "SMTP":  true,
        "SQL":   true,
        "SSH":   true,
        "TCP":   true,
        "TLS":   true,
        "TTL":   true,
        "UDP":   true,
        "UI":    true,
        "UID":   true,
        "UUID":  true,
        "URI":   true,
        "URL":   true,
        "UTF8":  true,
        "VM":    true,
        "XML":   true,
        "XMPP":  true,
        "XSRF":  true,
        "XSS":   true,
        "OAuth": true,
}

// add exceptions here for things that are not automatically convertable
var snakeToCamelExceptions = map[string]string{
        "oauth": "OAuth",
}


func safeName(s string) string {
	for _, r := range reservedWords {
		if r == s {
			return "_" + s
		}
	}
	return s
}

func toGoType(ipaType string) string {
	switch ipaType {
	case "":
		return "interface{}"
	case "dict":
		return "interface{}"
	case "object":
		return "interface{}"
	case "NoneType":
		return "interface{}"
	case "unicode":
		return "string"
	case "str":
		return "string"
	case "bytes":
		return "string"
	case "datetime":
		return "time.Time"
	case "DN":
		return "string"
	case "Principal":
		return "string"
	case "DNSName":
		return "string"
	case "Decimal":
		return "float64"
	case "Certificate":
		return "interface{}"
	case "CertificateSigningRequest":
		return "string"
	default:
		return ipaType
	}
}

func (t *CommandOutput) GoType(parent *Command) string {
	if t.Type == "dict_guess_receiver" {
		cls := strings.Split(parent.ObjClass, "/")[0]
		if cls != "" {
			return upperName(cls)
		}
	}
	return toGoType(t.Type)
}

func lowerName(s string) string {
	return safeName(SnakeToCamelLower(s))
}

func upperName(s string) string {
	return safeName(SnakeToCamel(s))
}

func (t *Topic) LowerName() string {
	return lowerName(t.Name)
}

func (t *Topic) UpperName() string {
	return upperName(t.Name)
}

func (t *Class) LowerName() string {
	return lowerName(t.Name)
}

func (t *Class) UpperName() string {
	return upperName(t.Name)
}

func (t *Command) LowerName() string {
	return lowerName(t.Name)
}

func (t *Command) UpperName() string {
	return upperName(t.Name)
}

func (t *Param) LowerName() string {
	return lowerName(t.Name)
}

func (t *Param) UpperName() string {
	return upperName(t.Name)
}

func (t *CommandOutput) LowerName() string {
	return lowerName(t.Name)
}

func (t *CommandOutput) UpperName() string {
	return upperName(t.Name)
}
