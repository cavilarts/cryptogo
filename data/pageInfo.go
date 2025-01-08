package data

type PageData struct {
	Title string
	Image string
	Made string
	Facts [3]string
	CurrentlyAvailable bool
}

func GetPageInfo() PageData {
	return pafeInfo
}

var pafeInfo  = PageData{
	Title: "This is a random web page",
	Image: "https://media.tenor.com/soCyR7I18DYAAAAC/dance-dancing-dog.gif",
	Made: "made with love by Carlos",
	Facts: [3]string{"Handsome", "Clever", "Funny"},
	CurrentlyAvailable: true,
}