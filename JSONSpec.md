#JSON Object Specification

##Basic JSON Models


    //////////////////////////////
    //
    //		User JSON Model
    //
    //////////////////////////////
    {
    	UUID: "123456789", // This is going to be the hash of studentID
        ProfilePicture: "images/default.jpg",
        FirstName: "Matt",
        LastName: "Hintzke",
        StudentID: "11086824",
        Bio: "Founder of the Windows Phone Development Group and ex-Vice President of Delta Chi Fraternity.  I am the co-creator of the Windows 8 Store Application, TweetMaps, and am heading into my final year of my computer science career path.",
        Email: "matt.hintze@email.wsu.edu",
        Major: "Computer Science",
        Minors: ["Math", "Computer Engineering"],
        ResumeID: "resume-number", // Resume number is the ID of a resume Object
        Resume: {}, // Resume-User relationship is going to be 1-1. This object is empty at initial registration
        ProjectIDs: ["134624978", "98726542", "719541376"], // These are the ids of the project Objects that this student contributes to
        Projects: [List of project objects will be created on front]
    }

    //////////////////////////////
    //
    //		Project JSON Model
    //
    //////////////////////////////
    {
    	UUID: "987654321", // Just set up some counter for this id
    	Name: "CougLinks",
    	Description: "Blah blah blah",
    	OwnerID: "123456789",
    	Contributors: [OwnerID, "654321987", "46793158"], // This will be initialized as just [OwnerID]
    	Images: ["Path/to/image1", "Path/to/image2"], // Each user will have their own storage spot
    }

    //////////////////////////////
    //
    //      Resume JSON Model
    //
    //////////////////////////////
    {
        resumeID: "123456789",
        ownerID: "987654321",
        objective: "Description of objective",
        education: [
            {
                school:"Washington State University",
                degree:"Bachelor of Science",
                gpa:"3.5",
                classes: [{abbr: "CptS", number:"121", name:"Intro to Computer Science"}],
            }
        ],
        skills: [{name: "C/C++/C#", value: "Advanced"}],
        education: [
            {
                startDate:"May 2012",
                endDate: "August 2012",
                jobTitle: "Web Application Developer",
                company: "IDD Aerospace",
                city: "Redmond, Wa",
                description: "Here is a description of the job"
            }
        ],
        hobbies: ["Guitar", "Linux", "Snowboarding"], //Or should we make this a string?
        links: [{name:"Google", href: "www.google.com", description: "This is google's website" }]
    }

## Actions

###These do not require authentication

Creating a user # POST /users

Viewing all users # GET /users

Viewing a single user # GET /users/:id



###These DO require authentication

Editing Profile # PUT /users/:id

Deleting Profile # DELETE /users/:id

Creating a project # POST /projects

Creating a resume # POST /resume


###These require authentication AND ownership of project or resume

Editing a project # PUT /projects/:id

Deleting a project # DELETE /projects/:id

Editing a resume # PUT /resume/:id

Deleting resume # DELETE /resume/:id


ALL REQUESTS THAT NEED AUTHENTICATION WILL NEED TO SEND AN OBJECT AND HAVE THE SERVER DO A CHECK FOR THE Token ATTRIBUTE.  WE CAN STORE A SESSION OBJECT ON THE SERVER THAT WILL BE USED TO VALIDATE A REQUEST. 


    /////////////////////////////
    //
    //	Authentication Object
    //
    /////////////////////////////
    {
    	Token: "haldkfjoiewuvalksdjfoiujlkavj",
    	Value: {} // This object will either be a user, resume or a project
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
- Token : Authentication token (string)
- Value : A User Object
