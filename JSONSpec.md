#JSON Object Specification

##User Object (all strings unless otherwise noted)
- FirstName
- LastName
- Resume
- Skills (list of strings)
- Interests (list of strings)
- Email
- UUID
- Major
- Minors (list of strings)

//////////////////////////////
//
//		User JSON Model
//
//////////////////////////////
    {
    	UUID: "123456789",
    	FirstName: "Matt",
    	LastName: "Hintzke",
    	Resume: "Path/To/Resume", // For now, resume will be a single uploaded file
    	Skills: [{name:"C/C++/C#", value:"Advanced"}, {name:"HTML/CSS3", value:"Intermediate"}],
    	Interests: ["Guitar", "Piano"],
    	Email: "matt.hintze@email.wsu.edu",
    	Major: "Computer Science",
    	Minors: ["Math", "Computer Engineering"],
    	Projects: ["134624978", "98726542", "719541376"]
    }

//////////////////////////////
//
//		Project JSON Model
//
//////////////////////////////
    {
    	UUID: "987654321",
    	Name: "CougLinks",
    	Description: "Blah blah blah",
    	OwnerID: "123456789",
    	Contributors: [OwnerID, "654321987", "46793158"],
    	Images: ["Path/to/image1", "Path/to/image2"],
    }

########### Actions ###########

// These do not require authentication
Creating a user # POST /users
Viewing all users # GET /users
Viewing a single user # GET /users/:id

// These DO require authentication
Editing Profile # PUT /users/:id
Deleting Profile # DELETE /users/:id
Creating a project # POST /projects

//These require authentication AND ownership of project
Editing a project # PUT /projects/:id
Deleting a project # DELETE /projects/:id

ALL REQUESTS THAT NEED AUTHENTICATION WILL NEED TO SEND AN OBJECT AND HAVE THE SERVER DO A CHECK FOR THE Token ATTRIBUTE.  WE CAN STORE A SESSION OBJECT ON THE SERVER THAT WILL BE USED TO VALIDATE A REQUEST. 

    /////////////////////////////
    //
    //	Authentication Object
    //
    /////////////////////////////
    {
    	Token: "haldkfjoiewuvalksdjfoiujlkavj",
    	Body: {} // This object will either be a user or a project
    }

    /////////////////////////////
    //
    //		Session Object
    //
    /////////////////////////////
    {
        // Not sure how to format this one yet
	}


##Request Object
###For adding, deleting, and updating users
- Action : one of NEW, UPDATE, DELETE (string)
- Token : Authentication token (string)
- Value : A User Object
