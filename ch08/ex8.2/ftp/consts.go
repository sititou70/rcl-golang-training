package ftp

type FTPSessionState struct {
	User          string
	LatestAddress string // e.g. "127.0.0.1:12345"
	CWD           string
	Type          string // e.g. "A", "I"
}

var FTPReplies = map[int]string{
	150: "150 File status okay; about to open data connection.\r\n",
	200: "200 %s command okay.\r\n",
	215: "215 UNIX system type.\r\n",
	220: "220 My FTP server ready for new user.\r\n",
	226: "226 Closing data connection.\r\n",
	230: "230 User %s logged in, proceed.\r\n",
	250: "250 Requested file action okay, completed.\r\n",
	257: "257 \"%s\" created.\r\n",
	331: "331 Password required for %s.\r\n",
	500: "500 Syntax error, command unrecognized.\r\n",
	502: "502 Command not implemented.\r\n",
	550: "550 Requested action not taken.\r\n",
}
