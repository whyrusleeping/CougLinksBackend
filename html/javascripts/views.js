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