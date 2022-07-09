## To Do

- Make user owner to update his userprofile
    > Write post request to update userprofile
    * [ ] should be updated fields is Email, first name, last name, personal value, secret
- Make authenticated users view their user profile
    > Need to provide signed users to view their user profile. Need to add some extra fields: First Name, Last Name,   Personal    value. Personal value should be visible only to this user.Username should be email with appropriate validation. Also need to provide Secret field which could be visible only for signed users
    * [ ] Add First Name field with validation not empty
    * [ ] Add Last Name field with validation not empty
    * [ ] Add Personal Value field with validation not empty and visible only to owner user
    * [ ] Add Secret Field with validation not empty and visibility only for signed user.
    * [x] All fields except Personal value and secret should be visible for anonymous user
    * [ ] add update User endpoint
    * [ ] 
- write more tests

## Doing

- use of mediator to connect commands and handlers

## Done

- ref
    > generated with go-swagger package
    * [x] generate client from spec
    * [x] generate tests from spec
- Add validation
    * [x] Add payload validation to sign-in
    * [x] Add payload validation to sign-up
