var AppRouter = Backbone.Router.extend({
	routes: {
		"": "listAll",
		"register": "registerUser"
	}
});

var appRouter = new AppRouter;

appRouter.on('route:listAll', function(){
	console.log("RENDER: listAll");
	listUsersView.render();

});

appRouter.on('route:registerUser', function(){
	console.log("RENDER: registerUser");
	createUserView.render();
})

  // Start Backbone history a necessary step for bookmarkable URL's
Backbone.history.start();