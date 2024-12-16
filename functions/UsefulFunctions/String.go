package usefulfunctions

// Сложный шаблон
/*
	в начале в (int) пишется длина строки
	(...) - символа может и не быть, но если он есть, то в количестве 1 штуки
	((int) ...) - символов может и не быть, но если он есть, то в количестве int штук
	{...} - непроизвольное кол. символов
	{(int) ...} - до int символов
	{(int1:int2) ...} - от int1 до int2 символов
	[...] - 1 символ
	[(int) ...] - int символов
	... (без скобок) обозначает определённую последовательность символов/символ
*/
type ComplexPattern string

// Шаблон
/*
	% - непроизвольное кол. символов
	_ - 1 символ
*/
type Pattern string

// Проверка подходит ли строка под шаблон
func (p Pattern) Check(s string) bool {
	im := 0
	for i, e := range s {
		switch p[i-im] {
		case byte('_'):
			continue
		case byte('%'):
			continue
		default:
			if p[i-im] != byte(e) {
				return false
			}
		}
	}
	return true
}
