/////////////////////////////////////////////////////////////
//
//					Home Page View
//
/////////////////////////////////////////////////////////////

var HomePageView = Backbone.View.extend({
	events: {},
	render: function(){
		var template = _.template($('#homePageTemplate').html());

		this.$el.html(template);

		return this;
	}
});

var homePageView = new HomePageView({el: $('#viewport')});

/////////////////////////////////////////////////////////////
//
//					Create User View
//
/////////////////////////////////////////////////////////////
var CreateUserView = Backbone.View.extend({
	events: {
		'click .btn-register': 'createUser'
	},
	render: function(){
		var template = _.template($('#createUserTemplate').html());

		this.$el.html(template);

		return this;
	},
	createUser: function(){
		var user = $('#newUserForm').serializeObject();
		user.token = hex_md5(user.studentid);
		user.password = hex_md5(user.password);
		console.log(user);

		delete user["conf-password"];
		var newUser = new UserModel(user);
		newUser.save(null, {success : function(model, response){
			//appRouter.navigate("/", {trigger: true});
			listUsersView.render();
			appRouter.navigate("/", {trigger: true});
		}});

	}
});

var createUserView = new CreateUserView({el: $('#viewport')});

/////////////////////////////////////////////////////////////
//
//					List Users View
//
/////////////////////////////////////////////////////////////
var ListUsersView = Backbone.View.extend({
	events: {
		'click .editBtn': 'editUser',
		'click .delBtn': 'delUser'
	},
	render: function(){
		var users = new UserCollection();
		var that = this;
		users.fetch({success: function(model, res){
			//var users = model;

			var template = _.template($('#listUsersTemplate').html(), {users: model.models});

			that.$el.html(template);

		}});

		return this;
	},
	editUser: function(e){

	},
	delUser: function(e){
		var that = $(e.target);
		var id = that.parent();
		console.log(id);
	}
});

var listUsersView = new ListUsersView({el: $('#viewport')});

/////////////////////////////////////////////////////////////
//
//					Profile View
//
/////////////////////////////////////////////////////////////

var ProfileView = Backbone.View.extend({
	events:{},
	render: function(){

    	var user =  {
    	UUID: "123456789",
    	ProfilePicture: "images/default.jpg",
    	FirstName: "Matt",
    	LastName: "Hintzke",
    	Resume: "Path/To/Resume", // For now, resume will be a single uploaded file
    	Skills: [{name:"C/C++/C#", value:"Advanced"}, {name:"HTML/CSS3", value:"Intermediate"}],
    	Interests: ["Guitar", "Piano"],
    	Email: "matt.hintze@email.wsu.edu",
    	Major: "Computer Science",
    	Minors: ["Math", "Computer Engineering"],
    	Projects: ["134624978", "98726542", "719541376"]
    	};

		var template = _.template($('#profileviewTemplate').html(), {user: user});

		this.$el.html(template);
	}
});

var profileView = new ProfileView({el: $('#viewport')});
