var UserModel = Backbone.Model.extend({
	idAttribute: "_id",
	urlRoot: "/users"
});

var UserCollection = Backbone.Collection.extend({
	model: UserModel,
	url: "/users"
});