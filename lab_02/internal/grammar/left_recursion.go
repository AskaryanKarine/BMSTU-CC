package grammar

import "slices"

// EliminateLeftRecursion устраняет левую рекурсию в грамматике
func (g *Grammar) EliminateLeftRecursion() *Grammar {

	// Сохраняем исходный порядок нетерминалов для обработки
	orderedNT := make([]string, len(g.NonTerminals))
	copy(orderedNT, g.NonTerminals)

	// Проходим по всем нетерминалам согласно порядку
	for i := 0; i < len(orderedNT); i++ {
		currentNT := orderedNT[i]

		// Этап 1: Замена продукций вида Ai -> Ajγ
		for j := 0; j < i; j++ {
			prevNT := orderedNT[j]
			g.replaceProductions(currentNT, prevNT)
		}

		// Этап 2: Устранение непосредственной левой рекурсии
		g.eliminateImmediateLR(currentNT)
	}

	return g
}

// replaceProductions заменяет продукции вида Ai -> Ajγ на развернутые продукции Aj
func (g *Grammar) replaceProductions(ai, aj string) {
	productions, exists := g.Rules[ai]
	if !exists {
		return
	}

	var newProductions [][]string
	for _, prod := range productions {
		// Проверяем продукцию вида Ajγ
		if len(prod) > 0 && prod[0] == aj {
			gamma := prod[1:]
			// Получаем все продукции Aj
			ajProductions := g.Rules[aj]

			// Создаем новые продукции: σγ для каждой σ в Aj
			for _, sigma := range ajProductions {
				// Объединяем σ и γ
				newProd := append(append([]string{}, sigma...), gamma...)
				newProductions = append(newProductions, newProd)
			}
		} else {
			// Оставляем продукцию без изменений
			newProductions = append(newProductions, prod)
		}
	}
	g.Rules[ai] = newProductions
}

// eliminateImmediateLR устраняет непосредственную левую рекурсию для нетерминала
func (g *Grammar) eliminateImmediateLR(nt string) {
	productions := g.Rules[nt]
	var alphas [][]string // Продукции вида A -> Aα
	var betas [][]string  // Продукции вида A -> β

	// Разделяем продукции на альфа и бета
	for _, prod := range productions {
		if len(prod) > 0 && prod[0] == nt {
			alphas = append(alphas, prod[1:])
		} else {
			betas = append(betas, prod)
		}
	}

	// Если нет леворекурсивных продукций - выходим
	if len(alphas) == 0 {
		return
	}

	// Создаем новый нетерминал A'
	newNT := nt + "'"
	// Добавляем его в список нетерминалов
	g.NonTerminals = append(g.NonTerminals, newNT)

	// Обновляем продукции исходного нетерминала
	var newBetas [][]string
	for _, beta := range betas {
		// Добавляем к каждой бета-продукции новый нетерминал
		newBetas = append(newBetas, append(beta, newNT))
	}
	g.Rules[nt] = newBetas

	// Создаем продукции для нового нетерминала
	var newAlphaProductions [][]string
	for _, alpha := range alphas {
		// Добавляем альфа + новый нетерминал
		newAlphaProductions = append(newAlphaProductions, append(alpha, newNT))
	}
	// Добавляем эпсилон-продукцию
	newAlphaProductions = append(newAlphaProductions, []string{Empty})

	g.Rules[newNT] = newAlphaProductions
}

func (g *Grammar) RemoveCycles() *Grammar {
	// Сохраняем порядок обработки нетерминалов
	orderedNT := g.NonTerminals

	for i, Ai := range orderedNT {
		// Обрабатываем только предыдущие нетерминалы (j < i)
		for j := 0; j < i; j++ {
			Aj := orderedNT[j]

			var newCombs [][]string
			for _, comb := range g.Rules[Ai] {
				// Если первый символ комбинации - Aj
				if len(comb) > 0 && comb[0] == Aj {
					// Подставляем все комбинации Aj
					for _, jComb := range g.Rules[Aj] {
						// Клонируем комбинацию Aj и добавляем хвост
						newComb := append(slices.Clone(jComb), comb[1:]...)
						newCombs = append(newCombs, newComb)
					}
				} else {
					newCombs = append(newCombs, comb)
				}
			}
			g.Rules[Ai] = newCombs
		}
	}

	return g
}
