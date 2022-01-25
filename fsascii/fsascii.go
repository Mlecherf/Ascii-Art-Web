package ascii

import (
	"bufio"
	"os"
)

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func Display(text, choice string) []string {
	rns := []rune(text)

	AsciiSentence := make([]string, 8)
	line := ""
	for j := 0; j < 8; j++ {
		for i := 0; i < len(rns); i++ {

			line += AsciiCarac(rns[i], j, choice)
		}
		AsciiSentence[j] = line

		line = ""
	}

	return AsciiSentence

}

func AsciiCarac(r rune, index int, style string) string {

	ToReturn := ""

	lines := LinesInFile("")
	if style == "shadow" {
		lines = LinesInFile(`Shadow.txt`)
	} else if style == "thinkertoy" {
		lines = LinesInFile(`thinkertoy.txt`)
	} else {
		lines = LinesInFile(`Ascii_Art_Font.txt`)
	}

	TabSpace := []int{1, 2, 3, 4, 5, 6, 7, 8}
	TabExcla := []int{10, 11, 12, 13, 14, 15, 16, 17}
	TabDoubleQuote := []int{19, 20, 21, 22, 23, 24, 25, 26}
	TabHashtag := []int{28, 29, 30, 31, 32, 33, 34, 35}
	TabDollard := []int{37, 38, 39, 40, 41, 42, 43, 44}
	TabPourCentage := []int{46, 47, 48, 49, 50, 51, 52, 53}
	TabEsperluette := []int{55, 56, 57, 58, 59, 60, 61, 62}
	TabSmpleQuote := []int{64, 65, 66, 67, 68, 69, 70, 71}
	TabParentheseGauche := []int{73, 74, 75, 76, 77, 78, 79, 80}
	TabParentheseDroite := []int{82, 83, 84, 85, 86, 87, 88, 89}
	TabAsterisque := []int{91, 92, 93, 94, 95, 96, 97, 98}
	TabPlus := []int{100, 101, 102, 103, 104, 105, 106, 107}
	TabVirgule := []int{109, 110, 111, 112, 113, 114, 115, 116}
	TabMoins := []int{118, 119, 120, 121, 122, 123, 124, 125}
	TabPoint := []int{127, 128, 129, 130, 131, 132, 133, 134}
	TabSlache := []int{136, 137, 138, 139, 140, 141, 142, 143}
	TabZero := []int{145, 146, 147, 148, 149, 150, 151, 152}
	TabOne := []int{154, 155, 156, 157, 158, 159, 160, 161}
	TabTwo := []int{163, 164, 165, 166, 167, 168, 169, 170}
	TabThree := []int{172, 173, 174, 175, 176, 177, 178, 179}
	TabFour := []int{181, 182, 183, 184, 185, 186, 187, 188}
	TabFive := []int{190, 191, 192, 193, 194, 195, 196, 197}
	TabSix := []int{199, 200, 201, 202, 203, 204, 205, 206}
	TabSeven := []int{208, 209, 210, 211, 212, 213, 214, 215}
	TabEight := []int{217, 218, 219, 220, 221, 222, 223, 224}
	TabNine := []int{226, 227, 228, 229, 230, 231, 232, 233}
	TabDoublePoint := []int{235, 236, 237, 238, 239, 240, 241, 242}
	TabpointVirgule := []int{244, 245, 246, 247, 248, 249, 250, 251}
	TabLeftArrow := []int{253, 254, 255, 256, 257, 258, 259, 260}
	TabEqual := []int{262, 263, 264, 265, 266, 267, 268, 269}
	TabRightArrow := []int{271, 272, 273, 274, 275, 276, 277, 278}
	Tabinterrogation := []int{280, 281, 282, 283, 284, 285, 286, 287}
	TabArrowBase := []int{289, 290, 291, 292, 293, 294, 295, 296}
	TabA := []int{298, 299, 300, 301, 302, 303, 304, 305}
	TabB := []int{307, 308, 309, 310, 311, 312, 313, 314}
	TabC := []int{316, 317, 318, 319, 320, 321, 322, 323}
	TabD := []int{325, 326, 327, 328, 329, 330, 331, 332}
	TabE := []int{334, 335, 336, 337, 338, 339, 340, 341}
	TabF := []int{343, 344, 345, 346, 347, 348, 349, 350}
	TabG := []int{352, 353, 354, 355, 356, 357, 358, 359}
	TabH := []int{361, 362, 363, 364, 365, 366, 367, 368}
	TabI := []int{370, 371, 372, 373, 374, 375, 376, 377}
	TabJ := []int{379, 380, 381, 382, 383, 384, 385, 386}
	TabK := []int{388, 389, 390, 391, 392, 393, 394, 395}
	TabL := []int{397, 398, 399, 400, 401, 402, 403, 404}
	TabM := []int{406, 407, 408, 409, 410, 411, 412, 413}
	TabN := []int{415, 416, 417, 418, 419, 420, 421, 422}
	TabO := []int{424, 425, 426, 427, 428, 429, 430, 431}
	TabP := []int{433, 434, 435, 436, 437, 438, 439, 440}
	TabQ := []int{442, 443, 444, 445, 446, 447, 448, 449}
	TabR := []int{451, 452, 453, 454, 455, 456, 457, 458}
	TabS := []int{460, 461, 462, 463, 464, 465, 466, 467}
	TabT := []int{469, 470, 471, 472, 473, 474, 475, 476}
	TabU := []int{478, 479, 480, 481, 482, 483, 484, 485}
	TabV := []int{487, 488, 489, 490, 491, 492, 493, 494}
	TabW := []int{496, 497, 498, 499, 500, 501, 502, 503}
	TabX := []int{505, 506, 507, 508, 509, 510, 511, 512}
	TabY := []int{514, 515, 516, 517, 518, 519, 520, 521}
	TabZ := []int{523, 524, 525, 526, 527, 528, 529, 530}
	TabCrochetGauche := []int{532, 533, 534, 535, 536, 537, 538, 539}
	TabAntiSlache := []int{541, 542, 543, 544, 545, 546, 547, 548}
	TabCrochetDroit := []int{550, 551, 552, 553, 554, 555, 556, 557}
	TabChapeau := []int{559, 560, 561, 562, 563, 564, 565, 566}
	TabTiretDuBas := []int{568, 569, 570, 571, 572, 573, 574, 575}
	TabBackticke := []int{577, 578, 579, 580, 581, 582, 583, 584}
	Taba := []int{586, 587, 588, 589, 590, 591, 592, 593}
	Tabb := []int{307 + 288, 308 + 288, 309 + 288, 310 + 288, 311 + 288, 312 + 288, 313 + 288, 314 + 288}
	Tabc := []int{316 + 288, 317 + 288, 318 + 288, 319 + 288, 320 + 288, 321 + 288, 322 + 288, 323 + 288}
	Tabd := []int{325 + 288, 326 + 288, 327 + 288, 328 + 288, 329 + 288, 330 + 288, 331 + 288, 332 + 288}
	Tabe := []int{334 + 288, 335 + 288, 336 + 288, 337 + 288, 338 + 288, 339 + 288, 340 + 288, 341 + 288}
	Tabf := []int{343 + 288, 344 + 288, 345 + 288, 346 + 288, 347 + 288, 348 + 288, 349 + 288, 350 + 288}
	Tabg := []int{352 + 288, 352 + 288, 353 + 288, 354 + 288, 355 + 288, 356 + 288, 357 + 288, 358 + 288}
	Tabh := []int{361 + 288, 362 + 288, 363 + 288, 364 + 288, 365 + 288, 366 + 288, 367 + 288, 368 + 288}
	Tabi := []int{370 + 288, 371 + 288, 372 + 288, 373 + 288, 374 + 288, 375 + 288, 376 + 288, 377 + 288}
	Tabj := []int{379 + 288, 380 + 288, 381 + 288, 382 + 288, 383 + 288, 384 + 288, 385 + 288, 386 + 288}
	Tabk := []int{388 + 288, 389 + 288, 390 + 288, 391 + 288, 392 + 288, 393 + 288, 394 + 288, 395 + 288}
	Tabl := []int{397 + 288, 398 + 288, 399 + 288, 400 + 288, 401 + 288, 402 + 288, 403 + 288, 404 + 288}
	Tabm := []int{406 + 288, 407 + 288, 408 + 288, 409 + 288, 410 + 288, 411 + 288, 412 + 288, 413 + 288}
	Tabn := []int{415 + 288, 416 + 288, 417 + 288, 418 + 288, 419 + 288, 420 + 288, 421 + 288, 422 + 288}
	Tabo := []int{424 + 288, 425 + 288, 426 + 288, 427 + 288, 428 + 288, 429 + 288, 430 + 288, 431 + 288}
	Tabp := []int{433 + 288, 434 + 288, 435 + 288, 436 + 288, 437 + 288, 438 + 288, 439 + 288, 440 + 288}
	Tabq := []int{442 + 288, 443 + 288, 444 + 288, 445 + 288, 446 + 288, 447 + 288, 448 + 288, 449 + 288}
	Tabr := []int{451 + 288, 452 + 288, 453 + 288, 454 + 288, 455 + 288, 456 + 288, 457 + 288, 458 + 288}
	Tabs := []int{460 + 288, 461 + 288, 462 + 288, 463 + 288, 464 + 288, 465 + 288, 466 + 288, 467 + 288}
	Tabt := []int{469 + 288, 470 + 288, 471 + 288, 472 + 288, 473 + 288, 474 + 288, 475 + 288, 476 + 288}
	Tabu := []int{478 + 288, 479 + 288, 480 + 288, 481 + 288, 482 + 288, 483 + 288, 484 + 288, 485 + 288}
	Tabv := []int{487 + 288, 488 + 288, 489 + 288, 490 + 288, 491 + 288, 492 + 288, 493 + 288, 494 + 288}
	Tabw := []int{496 + 288, 497 + 288, 498 + 288, 499 + 288, 500 + 288, 501 + 288, 502 + 288, 503 + 288}
	Tabx := []int{505 + 288, 506 + 288, 507 + 288, 508 + 288, 509 + 288, 510 + 288, 511 + 288, 512 + 288}
	Taby := []int{514 + 288, 515 + 288, 516 + 288, 517 + 288, 518 + 288, 519 + 288, 520 + 288, 521 + 288}
	Tabz := []int{523 + 288, 524 + 288, 525 + 288, 526 + 288, 527 + 288, 528 + 288, 529 + 288, 530 + 288}
	TabBracketGauche := []int{820, 821, 822, 823, 824, 825, 826, 827}
	TabPipe := []int{829, 830, 831, 832, 833, 834, 835, 836}
	TabBracketDroit := []int{838, 839, 840, 841, 842, 843, 844, 845}
	TabVague := []int{847, 848, 849, 850, 851, 852, 853, 853}
	if r == ' ' {
		ToReturn = (lines[TabSpace[index]])
	} else if r == '!' {
		ToReturn = (lines[TabExcla[index]])
	} else if r == '"' {
		ToReturn = (lines[TabDoubleQuote[index]])
	} else if r == '#' {
		ToReturn = (lines[TabHashtag[index]])
	} else if r == '$' {
		ToReturn = (lines[TabDollard[index]])
	} else if r == '%' {
		ToReturn = (lines[TabPourCentage[index]])
	} else if r == '&' {
		ToReturn = (lines[TabEsperluette[index]])
	} else if r == '\'' {
		ToReturn = (lines[TabSmpleQuote[index]])
	} else if r == '(' {
		ToReturn = (lines[TabParentheseGauche[index]])
	} else if r == ')' {
		ToReturn = (lines[TabParentheseDroite[index]])
	} else if r == '*' {
		ToReturn = (lines[TabAsterisque[index]])
	} else if r == '+' {
		ToReturn = (lines[TabPlus[index]])
	} else if r == ',' {
		ToReturn = (lines[TabVirgule[index]])
	} else if r == '-' {
		ToReturn = (lines[TabMoins[index]])
	} else if r == '.' {
		ToReturn = (lines[TabPoint[index]])
	} else if r == '/' {
		ToReturn = (lines[TabSlache[index]])
	} else if r == '0' {
		ToReturn = (lines[TabZero[index]])
	} else if r == '1' {
		ToReturn = (lines[TabOne[index]])
	} else if r == '2' {
		ToReturn = (lines[TabTwo[index]])
	} else if r == '3' {
		ToReturn = (lines[TabThree[index]])
	} else if r == '4' {
		ToReturn = (lines[TabFour[index]])
	} else if r == '5' {
		ToReturn = (lines[TabFive[index]])
	} else if r == '6' {
		ToReturn = (lines[TabSix[index]])
	} else if r == '7' {
		ToReturn = (lines[TabSeven[index]])
	} else if r == '8' {
		ToReturn = (lines[TabEight[index]])
	} else if r == '9' {
		ToReturn = (lines[TabNine[index]])
	} else if r == ':' {
		ToReturn = (lines[TabDoublePoint[index]])
	} else if r == ';' {
		ToReturn = (lines[TabpointVirgule[index]])
	} else if r == '<' {
		ToReturn = (lines[TabLeftArrow[index]])
	} else if r == '=' {
		ToReturn = (lines[TabEqual[index]])
	} else if r == '>' {
		ToReturn = (lines[TabRightArrow[index]])
	} else if r == '?' {
		ToReturn = (lines[Tabinterrogation[index]])
	} else if r == '@' {
		ToReturn = (lines[TabArrowBase[index]])
	} else if r == 'A' {
		ToReturn = (lines[TabA[index]])
	} else if r == 'B' {
		ToReturn = (lines[TabB[index]])
	} else if r == 'C' {
		ToReturn = (lines[TabC[index]])
	} else if r == 'D' {
		ToReturn = (lines[TabD[index]])
	} else if r == 'E' {
		ToReturn = (lines[TabE[index]])
	} else if r == 'F' {
		ToReturn = (lines[TabF[index]])
	} else if r == 'G' {
		ToReturn = (lines[TabG[index]])
	} else if r == 'H' {
		ToReturn = (lines[TabH[index]])
	} else if r == 'I' {
		ToReturn = (lines[TabI[index]])
	} else if r == 'J' {
		ToReturn = (lines[TabJ[index]])
	} else if r == 'K' {
		ToReturn = (lines[TabK[index]])
	} else if r == 'L' {
		ToReturn = (lines[TabL[index]])
	} else if r == 'M' {
		ToReturn = (lines[TabM[index]])
	} else if r == 'N' {
		ToReturn = (lines[TabN[index]])
	} else if r == 'O' {
		ToReturn = (lines[TabO[index]])
	} else if r == 'P' {
		ToReturn = (lines[TabP[index]])
	} else if r == 'Q' {
		ToReturn = (lines[TabQ[index]])
	} else if r == 'R' {
		ToReturn = (lines[TabR[index]])
	} else if r == 'S' {
		ToReturn = (lines[TabS[index]])
	} else if r == 'T' {
		ToReturn = (lines[TabT[index]])
	} else if r == 'U' {
		ToReturn = (lines[TabU[index]])
	} else if r == 'V' {
		ToReturn = (lines[TabV[index]])
	} else if r == 'W' {
		ToReturn = (lines[TabW[index]])
	} else if r == 'X' {
		ToReturn = (lines[TabX[index]])
	} else if r == 'Y' {
		ToReturn = (lines[TabY[index]])
	} else if r == 'Z' {
		ToReturn = (lines[TabZ[index]])
	} else if r == '[' {
		ToReturn = (lines[TabCrochetGauche[index]])
	} else if r == '\\' {
		ToReturn = (lines[TabAntiSlache[index]])
	} else if r == ']' {
		ToReturn = (lines[TabCrochetDroit[index]])
	} else if r == '^' {
		ToReturn = (lines[TabChapeau[index]])
	} else if r == '_' {
		ToReturn = (lines[TabTiretDuBas[index]])
	} else if r == '`' {
		ToReturn = (lines[TabBackticke[index]])
	} else if r == 'a' {
		ToReturn = (lines[Taba[index]])
	} else if r == 'b' {
		ToReturn = (lines[Tabb[index]])
	} else if r == 'c' {
		ToReturn = (lines[Tabc[index]])
	} else if r == 'd' {
		ToReturn = (lines[Tabd[index]])
	} else if r == 'e' {
		ToReturn = (lines[Tabe[index]])
	} else if r == 'f' {
		ToReturn = (lines[Tabf[index]])
	} else if r == 'g' {
		ToReturn = (lines[Tabg[index]])
	} else if r == 'h' {
		ToReturn = (lines[Tabh[index]])
	} else if r == 'i' {
		ToReturn = (lines[Tabi[index]])
	} else if r == 'j' {
		ToReturn = (lines[Tabj[index]])
	} else if r == 'k' {
		ToReturn = (lines[Tabk[index]])
	} else if r == 'l' {
		ToReturn = (lines[Tabl[index]])
	} else if r == 'm' {
		ToReturn = (lines[Tabm[index]])
	} else if r == 'n' {
		ToReturn = (lines[Tabn[index]])
	} else if r == 'o' {
		ToReturn = (lines[Tabo[index]])
	} else if r == 'p' {
		ToReturn = (lines[Tabp[index]])
	} else if r == 'q' {
		ToReturn = (lines[Tabq[index]])
	} else if r == 'r' {
		ToReturn = (lines[Tabr[index]])
	} else if r == 's' {
		ToReturn = (lines[Tabs[index]])
	} else if r == 't' {
		ToReturn = (lines[Tabt[index]])
	} else if r == 'u' {
		ToReturn = (lines[Tabu[index]])
	} else if r == 'v' {
		ToReturn = (lines[Tabv[index]])
	} else if r == 'w' {
		ToReturn = (lines[Tabw[index]])
	} else if r == 'x' {
		ToReturn = (lines[Tabx[index]])
	} else if r == 'y' {
		ToReturn = (lines[Taby[index]])
	} else if r == 'z' {
		ToReturn = (lines[Tabz[index]])
	} else if r == '{' {
		ToReturn = (lines[TabBracketGauche[index]])
	} else if r == '|' {
		ToReturn = (lines[TabPipe[index]])
	} else if r == '}' {
		ToReturn = (lines[TabBracketDroit[index]])
	} else if r == '~' {
		ToReturn = (lines[TabVague[index]])
	}

	return ToReturn
}
