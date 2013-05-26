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
		"home": "goHome",
		"profile": "viewProfile"
	}
});

var appRouter = new AppRouter;

appRouter.on('route:goHome', function(){
	console.log("RENDER: goHome");
	homePageView.render();
});

appRouter.on('route:listAll', function(){
	console.log("RENDER: listAll");
	listUsersView.render();

});

appRouter.on('route:registerUser', function(){
	console.log("RENDER: registerUser");
	createUserView.render();
});

appRouter.on('route:viewProfile', function(){
	console.log("RENDER: viewProfile");
	profileView.render();
});
  // Start Backbone history a necessary step for bookmarkable URL's
Backbone.history.start();