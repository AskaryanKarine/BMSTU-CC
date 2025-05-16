package dfa

func (dfa *DFA) buildReverseTransitions() map[int]map[rune][]int {
	reverse := make(map[int]map[rune][]int)
	// Инициализируем для всех состояний
	for id := range dfa.States {
		reverse[id] = make(map[rune][]int)
	}
	// Заполняем обратные переходы
	for fromID, state := range dfa.States {
		for symbol, toID := range state.Transitions {
			reverse[toID][symbol] = append(reverse[toID][symbol], fromID)
		}
	}
	return reverse
}

func (dfa *DFA) findReachable() map[int]bool {
	reachable := make(map[int]bool)
	queue := []int{dfa.Start}
	reachable[dfa.Start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range dfa.States[current].Transitions {
			if !reachable[next] {
				reachable[next] = true
				queue = append(queue, next)
			}
		}
	}
	return reachable
}

func (dfa *DFA) buildMarkedTable(reverseTransitions map[int]map[rune][]int) [][]bool {
	maxID := 0
	for id := range dfa.States {
		if id > maxID {
			maxID = id
		}
	}
	marked := make([][]bool, maxID+1)
	for i := range marked {
		marked[i] = make([]bool, maxID+1)
	}

	// Шаг 3: Пометить пары с разной терминальностью
	var queue [][2]int
	for i := range maxID + 1 {
		for j := range maxID + 1 {
			if !marked[i][j] && dfa.States[i].IsFinal != dfa.States[j].IsFinal {
				marked[i][j] = true
				marked[j][i] = true
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	//for i := range marked {
	//	for j := range marked[i] {
	//		fmt.Print(marked[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	//
	//fmt.Println(queue)

	// Шаг 4: Обработка очереди
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		u, v := pair[0], pair[1]

		for _, symbol := range dfa.Alphabet {
			// Получаем все r: r → u по symbol
			for _, r := range reverseTransitions[u][symbol] {
				// Получаем все s: s → v по symbol
				for _, s := range reverseTransitions[v][symbol] {
					if !marked[r][s] {
						marked[r][s] = true
						marked[s][r] = true
						queue = append(queue, [2]int{r, s})
					}
				}
			}
		}
	}

	return marked
}

func (dfa *DFA) buildComponents(marked [][]bool, reachable map[int]bool) []int {
	component := make([]int, len(dfa.States))
	currentComp := 0

	// Заполняем массив -1
	for id := range dfa.States {
		component[id] = -1
	}

	for i := range len(dfa.States) {
		if !marked[0][i] {
			component[i] = currentComp
		}
	}

	// Назначение компонент для достижимых состояний
	for id := 1; id < len(component); id++ {
		if !reachable[id] {
			continue
		}
		if component[id] == -1 {
			currentComp++
			component[id] = currentComp
			for j := id + 1; j < len(component); j++ {
				if !marked[id][j] {
					component[j] = currentComp
				}
			}
		}
	}

	return component
}

func (dfa *DFA) buildMinimizedDFA(component []int) *DFA {
	minimized := &DFA{
		States:   make(map[int]*State),
		Alphabet: dfa.Alphabet,
	}

	// Сопоставление компонент с новыми состояниями
	compToState := make(map[int]*State)
	for _, comp := range component {
		if comp == -1 {
			continue // Игнорируем недостижимые
		}
		if _, exists := compToState[comp]; !exists {
			isFinal := false
			// Проверяем, есть ли в классе терминальное состояние
			for stateID, c := range component {
				if c == comp && dfa.States[stateID].IsFinal {
					isFinal = true
					break
				}
			}
			newState := &State{
				ID:          comp,
				IsFinal:     isFinal,
				Transitions: make(map[rune]int),
			}
			compToState[comp] = newState
			minimized.States[comp] = newState
		}
	}

	// Заполняем переходы
	for comp, state := range compToState {
		// Берем первое состояние из класса для примера
		var sampleStateID int
		for id, c := range component {
			if c == comp {
				sampleStateID = id
				break
			}
		}
		// Копируем переходы из любого состояния класса
		for symbol, targetID := range dfa.States[sampleStateID].Transitions {
			targetComp := component[targetID]
			if targetComp != -1 {
				state.Transitions[symbol] = targetComp
			}
		}
	}

	// Назначаем стартовое состояние
	minimized.Start = component[dfa.Start]
	return minimized
}

func (dfa *DFA) Minimize() *DFA {
	// Добавляем состояние 0 (ловушку) как требует алгоритм
	extendedDFA := dfa.addZeroState()

	// Шаг 1: Построить обратные переходы
	reverseTransitions := extendedDFA.buildReverseTransitions()

	// Шаг 2: Найти достижимые состояния
	reachable := extendedDFA.findReachable()

	// Шаг 3-4: Построить таблицу различимых пар
	marked := extendedDFA.buildMarkedTable(reverseTransitions)

	//fmt.Println()
	//for i := range marked {
	//	for j := range marked[i] {
	//		fmt.Print(marked[i][j], " ")
	//	}
	//	fmt.Println()
	//}
	//
	//fmt.Println()
	//fmt.Println()

	// Шаг 5: Разбить на классы эквивалентности
	component := extendedDFA.buildComponents(marked, reachable)

	//fmt.Println(component)

	// Шаг 6: Построить минимизированный ДКА
	minimized := extendedDFA.buildMinimizedDFA(component)

	// Удаляем состояние-ловушку 0, если оно не достижимо
	return minimized.removeZeroState()
}

// Добавляет состояние 0 и корректирует переходы
func (dfa *DFA) addZeroState() *DFA {
	newDFA := &DFA{
		Start:    dfa.Start + 1, // Стартовое состояние теперь ID+1
		Alphabet: dfa.Alphabet,
		States:   make(map[int]*State),
	}

	// 1. Сдвигаем все ID состояний на +1
	for oldID, state := range dfa.States {
		newState := &State{
			ID:          oldID + 1,
			Transitions: make(map[rune]int),
			IsFinal:     state.IsFinal,
		}
		for symbol, target := range state.Transitions {
			newState.Transitions[symbol] = target + 1
		}
		newDFA.States[newState.ID] = newState
	}

	// 2. Добавляем состояние 0
	trap := &State{
		ID:          0,
		Transitions: make(map[rune]int),
		IsFinal:     false,
	}
	for _, symbol := range newDFA.Alphabet {
		trap.Transitions[symbol] = 0 // Все переходы ведут в себя
	}
	newDFA.States[0] = trap

	// 3. Добавляем недостающие переходы в 0
	for _, state := range newDFA.States {
		if state.ID == 0 {
			continue
		}
		for _, symbol := range newDFA.Alphabet {
			if _, exists := state.Transitions[symbol]; !exists {
				state.Transitions[symbol] = 0
			}
		}
	}

	return newDFA
}

// Удаляет недостижимые состояния (включая 0, если нужно)
func (dfa *DFA) removeZeroState() *DFA {
	cleaned := &DFA{
		Start:    dfa.Start,
		Alphabet: dfa.Alphabet,
		States:   make(map[int]*State),
	}

	for id, state := range dfa.States {
		if id == 0 { // Удаляем 0
			continue
		}
		newState := &State{
			ID:          state.ID,
			IsFinal:     state.IsFinal,
			Transitions: make(map[rune]int),
		}

		// Копируем переходы, исключая те, что ведут в 0
		for symbol, target := range state.Transitions {
			if target != 0 {
				newState.Transitions[symbol] = target
			}
		}

		cleaned.States[id] = newState
	}

	return cleaned
}
