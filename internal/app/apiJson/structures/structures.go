package structures

type Posts struct{
	userId 		uint32
	id 			uint32
	title		string
	body	 	string
}

type Comments struct {
	postId		uint32
	id			uint32
	name		string
	email		string
	body		string
}

type Albums struct {
	userId 		uint32
	id 			uint32
	title		string
}

type Photos struct {
	albumId 	uint32
	id			uint32
	title		string
	url			string
	thumbnailUrl string
}

type Todos struct {
	userId 		uint32
	id 			uint32
	title		string
	completed	bool
}

type geoPosition struct {
	lat			string
	lng			string
}

type address struct {
	street		string
	suite		string
	city		string
	zipcode		string
	geo			geoPosition
}

type Company struct {
	name		string
	catchPhrase	string
	bs			string
}

type Users struct {
	id       uint
	name     string
	username string
	email    string
	address  address
	phone    string
	website  string
	company  Company
}
