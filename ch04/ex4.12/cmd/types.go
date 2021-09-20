package cmd

const CacheDir = "cache"
const domain = "https://xkcd.com"

type XkcdJSON struct {
	Num        int
	Title      string
	Transcript string
	Img        string
}
