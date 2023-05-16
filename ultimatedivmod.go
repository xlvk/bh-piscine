package piscine

func UltimateDivMod(a *int, b *int, div int, mod int) {
	div = *a / *b
	mod = *a % *b
	*a = div
	*b = mod
}
