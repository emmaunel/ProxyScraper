package colors

var red = "\033[1;31m"
var green = "\033[8;32m"
var yellow = "\033[8;33m"
var defcol = "\033[0m"
var pink = "\033[8;35m" // pink. Why??? COuldn't find a better color

func info(text string) string {
	return red + "[" + defcol + getGreen(text) + red + "]" + defcol 
}

func getGreen(text string) string {
	return green + text + defcol
}

func Error(msg string) string {
	return red + "[ERROR] - " + msg + defcol
}

func ShowWarning(msg string) string {
	return yellow + "[WARNING] - " + msg + defcol
}

func ShowInfo(msg string) string {
	return pink + "[INFO] - " + msg + defcol
}

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

// func info(text string) string {
// 	return "\033[1;36;40m " + text + " \033[0m"
// }


func PrintProxy(ip string, port string, protocolType string) string {
	return info(protocolType) + " " + ip + ":" + port
}