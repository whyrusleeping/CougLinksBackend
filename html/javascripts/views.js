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
    	Bio: "Founder of the Windows Phone Development Group and ex-Vice President of Delta Chi Fraternity.  I am the co-creator of the Windows 8 Store Application, TweetMaps, and am heading into my final year of my computer science career path.",
    	ResumeID: "resume-number", // For now, resume will be a single uploaded file
    	Skills: [{name:"C/C++/C#", value:"Advanced"}, {name:"HTML/CSS3", value:"Intermediate"}],
    	Interests: ["Guitar", "Piano"],
    	Email: "matt.hintze@email.wsu.edu",
    	Major: "Computer Science",
    	Minors: ["Math", "Computer Engineering"],
    	Projects: ["134624978", "98726542", "719541376"],
    	Resume: {}
    	};

    	var resume =  {
        ResumeID: "123456789",
        ownerID: "987654321",
        objective: "I have a knack for the web and everything about it. I love to create web applications and develop web sites for friends and family just for fun. My plan is to find a small start-up that will allow me to use my skills as a C programmer as well as web designer to create content rich applications for clients",
        education: [
            {
                school:"Washington State University",
                degree:"Bachelor of Science",
                gpa:"3.5",
                classes: [{abbr: "CptS", number:"121", name:"Intro to Computer Science"},
                		{abbr: "CptS", number:"122", name:"Data Structures"},
                		{abbr: "CptS", number:"223", name:"Advanced Data Structures"}]
            }
        ],
        skills: [{name: "C/C++/C#", value: "Advanced"},
        		{name: "HTML/CSS", value: "Intermediate"},
        		{name: "Linux Programming", value: "Intermediate"}],
        employment: [
            {
                startDate:"May 2012",
                endDate: "August 2012",
                jobTitle: "Web Application Developer",
                company: "IDD Aerospace",
                city: "Redmond, Wa",
                description: "Working at IDD, I personally designed and implemented a system for storing and viewing data from a centralized MySQL database. Originally, IDD ran lighting tests using radiometers and spectrometers to ensure the requirements of each part are met. However, this data was never saved in any way once the test was done. This led to my design of a PHP, Javascript, and MySQL system that allowed data stored in an Excel spreadsheet to be translated to a database and then viewed using a PHP client on the browser. The client also allowed the user to fully manage the database tables in order to handle the need of any future data specifications."
            },
            {
            	startDate:"May 2011",
            	endDate:"August 2011",
            	jobTitle:"Mechanical Engineering Intern",
            	company:"Coffman Engineering",
            	city: "Seattle, Wa",
            	description: "As an intern at Coffman Engineers I worked along side colleagues to complete a contracting job at The Boeing Company in Everett where I was put to the task of mapping out the extensive steam pipe system that powers much of the equipment used by the assemblers. My daily tasks included drawing out steam lines and their corresponding loads throughout the Boeing plant, reviewing previously created CAD drawings for accuracy and consistency, and modeling HVAC load scenarios using computer software. This internship provided me the opportunity to work with a team to deliver quality work and consultation in a professional environment."
            }
        ],
        hobbies: "By now, it is clear that I have a passion for the web. I spend much of my time learning new web application developing standards and design architectures. I have basic to intermediate experience using MVC frameworks such as NodeJS, BackboneJS, ExpressJS, and Meteor. Apart from programming, I love to play sports such as soccer and basketball and spend time fulfilling my passion for music by playing guitar and piano.", //Or should we make this a string?
        links: [{name:"Google", href: "https://www.google.com", description: "This is google's website" },
        		{name:"DamonDevelopment", href: "https://www.google.com", description: "My Personal Website" }],
        projects: ["02384209385", "00234583452", "0234247234702"]
    }

    	user.Resume = resume;
    	//var user = new UserModel(obj);

		var template = _.template($('#profileviewTemplate').html(), {user : user});

		this.$el.html(template);
	}
});

var profileView = new ProfileView({el: $('#viewport')});
