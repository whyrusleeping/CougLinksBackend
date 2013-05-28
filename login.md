#Login Scheme

TODO: Registration...

##Initial login request.

Post the following JSON Object to /login:

	{
		"ID":"users id for now, i can do email later",
		"Password":"users hashed password"
	}

You will be returned the following similar object:

	{
		"ID":"The same user ID you send",
		"Token":"the session token for this login"
	}

Any further requests that require authentication can be handed this token (which is a base64 encoded md5 hash of a few different things)

In order to log out, send a DELETE request to /login with the users ID and Token.
