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
		"profile": "viewProfile",
		"login": "loginUser",
		"support": "support"
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
	 $('.section-pulldown').on('click', function(){
			        var el = $(this).next();
			        var par = $(this).parent();
			        var arrow = $(this).children('img');

			        if(el.css('display') == 'none')
			        {
			        	el.css('padding-bottom', '50px');
			        	arrow.attr('src', 'images/downarrow.png');
			        	par.css('box-shadow', '0px 0px 15px 3px #A8A8A8 inset');
			        	$(this).next().slideToggle(700);
				        
			        }else{
			      		$(this).next().slideToggle(700, function(){
			         	el.css('padding-bottom', '0px');
			      		par.css('box-shadow', '2px 2px 15px 2px #CCCCCC inset');
			      		arrow.attr('src', 'images/uparrow.png');
			         });
			           
			        }

			    });
});

appRouter.on('route:loginUser', function(){
	console.log("RENDER: loginUser");
	loginView.render();
});

appRouter.on('route:support', function(){
	console.log("RENDER: support");
	supportView.render();
});
  // Start Backbone history a necessary step for bookmarkable URL's
Backbone.history.start();