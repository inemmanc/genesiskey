package text

func DefaultWelcomeMessage() string {
	welcomeMessage := "---------------------------------------\n---------- GENESIS KEY (cli) ----------\n---------------------------------------"
	return welcomeMessage
}

func DefaultMethodText() string {
	selectText := "Insert the Encoding Method:\nSTD - standard base64 encoding\nRAW - unpadded base64 encoding (same as STD but omits padding characters)"
	return selectText
}
