/////////////////////////////////////////////////////////////
//
//					User Model & Collection
//
/////////////////////////////////////////////////////////////


/*

User Model Format

{
	"token":"11229324",
	"Name":"Jeromy",
	"Resume":"None Yet",
	"Skills":["Golang", "C", "Genius"],
	"Major":"Computer Science",
	"Minors":["Math", "CompE"],
	"Interests":"Females",
	"Email":"me@jero.my"
}
*/
var UserModel = Backbone.Model.extend({
	idAttribute: "UUID",
	urlRoot: "/users"
});

var UserCollection = Backbone.Collection.extend({
	model: UserModel,
	url: "/users"
});