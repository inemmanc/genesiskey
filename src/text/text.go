package text

func DefaultWelcomeMessage() string {
	welcomeMessage := "---------------------------------------\n---------- GENESIS KEY (cli) ----------\n---------------------------------------"
	return welcomeMessage
}

func DefaultMethodText() string {
	selectText := "Insert the Encoding Method:\nSTD - standard base64 encoding\nURL - used in URLs and file names"
	return selectText
}
