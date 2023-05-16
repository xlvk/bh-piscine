package piscine

func DivMod(a *int, b *int, div int, mod int) {
	div = *a / *b
	mod = *a % *b
	*a = div
	*b = mod
}
