#JSON Object Specification

##User Object (all strings unless otherwise noted)
- FirstName
- LastName
- Resume
- Skills
- Interests
- Email
- UUID
- Major
- Minor

##Request Object
###For adding, deleting, and updating users
- Action : one of NEW, UPDATE, DELETE (string)
- Token : Authentication token (string)
- Value : A User Object
