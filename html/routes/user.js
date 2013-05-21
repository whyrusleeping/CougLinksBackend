var mongoose = require('mongoose');

mongoose.connect('mongodb://localhost/wsu');

var db = mongoose.connection;

db.on('error', console.error.bind(console, 'connection error:'));

db.once('open', function callback () {
  // yay!
  console.log("Connected successfully to WSU database");

  //Create some test data

});

var userSchema = mongoose.Schema({
    FirstName: String,
    LastName: String,
    UUID: String,
    Email: String,
    password: String,
    token: String
});

var User = mongoose.model('User', userSchema, "users");
/*
 * GET users listing.
 */

exports.list = function(req, res){
  User.find(function(err, users){
  	if(err){
  		console.log("Cannot find users");
  		res.send(err);
  	}else{
  		console.log("Found users");
  		res.send(users);
  	}
  });
};

exports.create = function(req, res){
	var newUser = new User(req.body);
	newUser.save(function(err, user){
		if(err){
			console.log("Error creating user: " + err);
			res.send(err);
		}else{
			console.log("Created new user");
			res.send(user);
		}
	});
}
