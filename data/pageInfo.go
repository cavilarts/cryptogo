package data

type PageData struct {
	Title string
	Image string
	Made string
	Facts []string
	CurrentlyAvailable bool
}

func GetPageInfo() PageData {
	return pafeInfo
}

func SetFact(fact []string) {
	pafeInfo.Facts = append(pafeInfo.Facts, fact...)
}

var pafeInfo  = PageData{
	Title: "This is a random web page",
	Image: "https://media.tenor.com/soCyR7I18DYAAAAC/dance-dancing-dog.gif",
	Made: "made with love by Carlos",
	Facts: []string{"Handsome", "Clever", "Funny"},
	CurrentlyAvailable: true,
}