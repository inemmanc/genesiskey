package text

func DefaultWelcomeMessage() string {
	welcomeMessage := "---------------------------------------\n---------- GENESIS KEY (cli) ----------\n---------------------------------------"
	return welcomeMessage
}

func DefaultMethodText() string {
	selectText := "\nInsert the Encoding Method:\nSTD - standard base64 encoding\nRAW - unpadded base64 encoding (same as STD but omits padding characters)"
	return selectText
}

func RecallText() string {
	recallText := "\nDo you want to Generate a new Key? Type YES/Y or type any other word to exit"
	return recallText
}

func FinalScreenText() string {
	finalText := "\nPress enter to quit Genesis Key..."
	return finalText
}
