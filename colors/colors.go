package colors


// def getWhite(text):
//     return "\033[1;37m" + text + "\033[0m"

// def getYellow(text):
//     return "\033[1;33m" + text + "\033[0m"

// def getBrightCyan(text):
//     return "\033[1;36;40m" + text + "\033[0m"

// def getBlue(text):
//     return "\033[1;34m" + text + "\033[0m"

// def getRed(text):
//     return "\033[1;31m" + text + "\033[0m"

// def getGreen(text):
//     return "\033[1;32;40m" + text + "\033[0m"

// def showInfo(text):
//     print(getBlue("[#]"), text)

// def showWarning(text):
//     print(getYellow("[!]"), text)

// def showSuccess(text):
//     print(getGreen("[+]"), text)

// def showError(text):
//     print(getRed("[-]"), text)

func info(text string) string {
	return "\033[1;36;40m " + text + " \033[0m"
}