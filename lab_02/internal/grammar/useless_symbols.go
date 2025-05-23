package grammar

import (
	"reflect"
	"sort"
)

func (g *Grammar) EliminationUselessSymbols() *Grammar {
	// Определение порождающих нетерминалов
	// Инициализация N0 = пустое множество
	useful := make(map[string]bool)
	newUseful := make(map[string]bool)
	terms := make(map[string]bool)

	newUseful[g.Start] = true
	for _, t := range g.Terminals {
		terms[t] = true
	}

	// Итеративное построение Ni пока не стабилизируется
	for !reflect.DeepEqual(useful, newUseful) {
		// Обмен множеств для следующей итерации
		useful, newUseful = newUseful, useful
		for k := range newUseful {
			delete(newUseful, k)
		}

		// Обновление Ni
		for symbol, productions := range g.Rules {
			if useful[symbol] { // Если A принадлежит Vi-1
				for _, p := range productions {
					for _, elem := range p {
						// Добавляем терминалы (T остается неизменным)
						if terms[elem] {
							newUseful[elem] = true
						}
					}
				}
			}
		}

		// Объединение Vi = Vi-1 ∪ новых символов
		for k := range useful {
			newUseful[k] = true
		}
	}

	// Расширение Ni за счет нетерминалов в правых частях
	for symbol, productions := range g.Rules {
		if !useful[symbol] {
			continue
		}
		for _, p := range productions {
			for _, elem := range p {
				if terms[elem] {
					continue
				}
				useful[elem] = true // Добавляем в Ni

			}
		}

	}

	// Устранение недостижимых символов
	// Формирование новых правил P'
	newRules := make(map[string][][]string)
	for symbol, productions := range g.Rules {
		if useful[symbol] { // Фильтрация по Vi
			newRules[symbol] = productions
		}
	}

	// Формирование итоговых компонентов грамматики
	newNonTerminals, newTerminals := g.extractSymbolsFromRules(newRules)

	// Возвращаем грамматику после применения обоих алгоритмов
	return &Grammar{
		NonTerminals: newNonTerminals,
		Terminals:    newTerminals,
		Start:        g.Start,
		Rules:        newRules,
	}
}

func (g *Grammar) extractSymbolsFromRules(newRules map[string][][]string) (nonTerminals, terminals []string) {
	existsT := map[string]struct{}{}
	existsNT := map[string]struct{}{}

	for _, nt := range g.NonTerminals {
		existsNT[nt] = struct{}{}
	}

	for _, t := range g.Terminals {
		existsT[t] = struct{}{}
	}

	newNT := make(map[string]struct{})
	newT := make(map[string]struct{})

	// Собираем нетерминалы из ключей правил
	for nt := range newRules {
		newNT[nt] = struct{}{}
	}

	// Анализируем правые части правил
	for _, productions := range newRules {
		for _, production := range productions {
			for _, symbol := range production {
				if symbol == Empty {
					continue
				}

				if _, ok := existsNT[symbol]; ok {
					newNT[symbol] = struct{}{}
					continue
				}

				if _, ok := existsT[symbol]; ok {
					newT[symbol] = struct{}{}
					continue
				}
			}
		}
	}

	// Конвертируем в отсортированные слайсы
	nonTerminals = mapKeysToSlice(newNT)
	terminals = mapKeysToSlice(newT)

	return nonTerminals, terminals
}

// Вспомогательная функция для конвертации мапы в отсортированный слайс
func mapKeysToSlice(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
