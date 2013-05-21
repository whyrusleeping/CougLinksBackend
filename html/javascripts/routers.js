/////////////////////////////////////////////////////////////
//
//					Main view routes
//
/////////////////////////////////////////////////////////////

var AppRouter = Backbone.Router.extend({
	routes: {
		"": "goHome",
		"register": "registerUser",
		"students": "listAll",
		"home": "goHome"
	}
});

var appRouter = new AppRouter;

appRouter.on('route:goHome', function(){
	console.log("REDNER: goHome");
	homePageView.render();
});

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